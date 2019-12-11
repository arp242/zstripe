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
	"io/ioutil"
	"math/big"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var (
	SecretKey     = ""                       // Your Stripe secret key (sk_*).
	API           = "https://api.stripe.com" // API base URL.
	DebugURL      = false                    // Show URLs as they're requested.
	DebugReqBody  = false                    // Show body of request.
	DebugRespBody = false                    // Show body of response
	MaxRetry      = 30 * time.Second         // Max time to retry requests.
)

// ErrRetry is used when we've retried longer than MaxRetry.
var ErrRetry = errors.New("retried longer than MaxRetry")

type (
	// ID if you're interested in just retrieving the ID from a response.
	ID struct {
		ID string `json:"id"`
	}

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
	sc := ""
	if e.StripeError.Code != "" {
		sc = e.StripeError.Code + ": "
	}
	return fmt.Sprintf("code %s for %s %s (%s%s)",
		e.Status, e.Method, e.URL, sc, e.StripeError.Message)
}

// Body for requests.
type Body map[string]string

// Encode the values with url.Values.Encode().
func (b Body) Encode() string {
	body := make(url.Values)
	for k, v := range b {
		body.Set(k, v)
	}
	return body.Encode()
}

// Client to use for all API requests.
var Client = http.Client{Timeout: 10 * time.Second}

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
func Request(scan interface{}, method, url string, body string) (*http.Response, error) {
	if SecretKey == "" {
		panic("zstripe: must set SecretKey")
	}

	start := time.Now()

	if !strings.HasPrefix(url, "https://") {
		url = API + url
	}

	r, err := http.NewRequest(method, url, strings.NewReader(body))
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
	if DebugReqBody {
		fmt.Println(body)
	}

	resp, err := Client.Do(r)
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

	if DebugRespBody {
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
