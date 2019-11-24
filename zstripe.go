// Package zstripe is a set of utility functions for working with the Stripe
// API.
//
// It's not a full "client library"; but just a few functions that make it easy
// to call api.stripe.com.
package zstripe

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	SecretKey = ""                       // Your Stripe secret key (sk_*).
	API       = "https://api.stripe.com" // API base URL.
	DebugURL  = false                    // Show URLs as they're requested.
	DebugBody = false                    // Show body of requests.
	MaxRetry  = 30 * time.Second         // Max time to retry requests.
)

// ErrRetry is used when we've retried longer than MaxRetry.
var ErrRetry = errors.New("retried longer than MaxRetry")

type (
	// Error is used when the status code is not 200 OK.
	Error struct {
		Method, URL string
		Status      string
		StatusCode  int
		StripeError StripeError `json:"error"`
	}

	// StripeError is Stripe's response on errors.
	// https://stripe.com/docs/api/errors
	StripeError struct {
		// Error type; always set and one of: "api_connection_error, api_error,
		// authentication_error, card_error, idempotency_error,
		// invalid_request_error, or rate_limit_error".
		Type string `json:"type"`

		// Parameter related to the error to display a message near the form
		// field.
		Param string `json:"param"`

		Message      string `json:"message"`       // Human-readable message.
		Code         string `json:"code"`          // Error code; may be blank.
		DocURL       string `json:"doc_url"`       // URL to more information.
		Charge       string `json:"charge"`        // ID of the failed charge for card errors.
		DeclinedCode string `json:"declined_code"` // Card issuer's reason for declining a card, if provided.
	}
)

func (e Error) Error() string {
	return fmt.Sprintf("code %s for %s %s (%s: %s)",
		e.Status, e.Method, e.URL, e.StripeError.Code, e.StripeError.Message)
}

var client = http.Client{Timeout: 10 * time.Second}

// Request something from the Stripe API.
//
// The response body is unmarshaled to scan as JSON.
//
// Responses with the Stripe-Should-Retry header set will be retried every two
// seconds. ErrRetry is returned if it still fails after MaxRetry.
//
// A response code higher than 399 will return an Error, but won't affect the
// behaviour of this function.
//
// The request body is an URL-encoded form (Stripe doesn't accept JSON), usually
// you will want to do something like this:
//
//   f := make(url.Values)
//   f.Set("name", "Martin Tournoij")
//   body := strings.NewReader(body.Encode())
//
// There are many libraries to convert a struct or map to an encoded form, but
// for many simpler application it's not really needed, which is why it's not
// done automatically.
//
// The Body on the returned http.Response is closed.
//
// This will use the global SecretKey, which must be set.
func Request(scan interface{}, method, url string, body io.Reader) (*http.Response, error) {
	if SecretKey == "" {
		panic("zstripe: must set SecretKey")
	}

	start := time.Now()

	if !strings.HasPrefix(url, "https://") {
		url = API + url
	}

	r, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("zstripe: http.NewRequest: %s", err)
	}

	r.Header.Add("Authorization", "Bearer "+SecretKey)
	r.Header.Add("Idempotency-Key", rnd())
	// TODO: /v1/files needs multipart/form-data
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Stripe-Version", "2019-11-05")
	r.Header.Add("User-Agent", "Go-http-client/1.1; client=zstripe")

doreq:
	if DebugURL {
		fmt.Printf("%v %v\n", method, url)
	}

	resp, err := client.Do(r)
	if err != nil {
		return resp, fmt.Errorf("zstripe: client.Do: %s", err)
	}
	defer resp.Body.Close()

	// 202 Accepted: retry the request after a short delay.
	if resp.Header.Get("Stripe-Should-Retry") == "true" {
		resp.Body.Close()
		if time.Now().Sub(start) > MaxRetry {
			return resp, ErrRetry
		}
		time.Sleep(2 * time.Second)
		goto doreq
	}

	rbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp, fmt.Errorf("zstripe: read body: %s", err)
	}

	if DebugBody {
		fmt.Println(string(rbody))
	}

	err = json.Unmarshal(rbody, scan)
	if resp.StatusCode >= 400 {
		serr := Error{
			Status:     resp.Status,
			StatusCode: resp.StatusCode,
			Method:     method,
			URL:        url,
		}
		_ = json.Unmarshal(rbody, &serr)

		// Intentionally override the JSON status error; chances are this is the
		// root cause.
		err = serr
	}

	return resp, err
}

var max = big.NewInt(0).SetUint64(18446744073709551615)

func rnd() string {
	var key strings.Builder
	for i := 0; i < 4; i++ {
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			panic(fmt.Errorf("zstripe.rnd: %s", err))
		}
		_, _ = key.WriteString(strconv.FormatUint(n.Uint64(), 36))
	}
	return key.String()
}
