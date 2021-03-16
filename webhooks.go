package zstripe

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Possible errors when validating a webhook.
var (
	ErrWebhookTooOld           = errors.New("webhook too old")
	ErrWebhookInvalidHeader    = errors.New("invalid Stripe-Signature header")
	ErrWebhookInvalidSignature = errors.New("invalid signature")
)

var (
	// Stripe signing secret (whsec_*).
	SignSecret string

	// Reject signatures older than this, to prevent replay attacks.
	MaxAge = 300 * time.Second
)

// https://stripe.com/docs/api#events.
type Event struct {
	ID              string `json:"id"`
	Type            string `json:"type"`
	Livemode        bool   `json:"livemode"`
	Created         int64  `json:"created"`
	Account         string `json:"account"`          // Account that originated the event (Connect only).
	PendingWebhooks int64  `json:"pending_webhooks"` // Number of webhooks that still need to be delivered.

	Data struct {
		Raw json.RawMessage `json:"object"`

		// Relevant resource, e.g. "invoice.created" will have the full invoice
		// object.
		Object map[string]interface{}

		// Names of changed attributes with their previous values for *.updated
		// events.
		PreviousAttributes map[string]interface{} `json:"previous_attributes"`
	} `json:"data"`

	// Details about the request that created the event; may be empty as not all
	// events are created by a request.
	Request struct {
		ID             string `json:"id"`
		IdempotencyKey string `json:"idempotency_key"`
	} `json:"request"`
}

// Read the event from the request body and validate the signature.
func (e *Event) Read(r *http.Request) error {
	if SignSecret == "" {
		panic("zstripe.Event.Read: must set zstripe.SignSecret")
	}

	ts, sigs, err := parseHeader(r.Header.Get("Stripe-Signature"))
	if err != nil {
		return err
	}
	if time.Since(ts) > MaxAge {
		return ErrWebhookTooOld
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	mac := hmac.New(sha256.New, []byte(SignSecret))
	mac.Write([]byte(fmt.Sprintf("%d", ts.Unix())))
	mac.Write([]byte("."))
	mac.Write(b)
	sig := mac.Sum(nil)
	found := false
	for _, s := range sigs {
		if hmac.Equal(sig, s) {
			found = true
			break
		}
	}
	if !found {
		return ErrWebhookInvalidSignature
	}

	return json.Unmarshal(b, &e)
}

// The Stripe-Signature header contains a timestamp and one or more signatures.
// The timestamp is prefixed by t=, and each signature is prefixed by a scheme.
// Schemes start with v, followed by an integer. Currently, the only valid
// signature scheme is v1. To aid with testing, Stripe sends an additional
// signature with a fake v0 scheme, for test-mode events.
//
// Stripe-Signature: t=1492774577,
//     v1=5257a869e7ecebeda32affa62cdca3fa51cad7e77a0e56ff536d0ce8e108d8bd,
//     v0=6ffbb59b2300aae63f272406069a9788598b792a944a07aba816edb039989a39
//
// Note that newlines have been added in the example above for clarity, but a
// real Stripe-Signature header will be all one line.
//
// Stripe generates signatures using a hash-based message authentication code
// (HMAC) with SHA-256. To prevent downgrade attacks, you should ignore all
// schemes that are not v1.
//
// It is possible to have multiple signatures with the same scheme/secret pair.
// This can happen when you roll an endpoint’s secret from the Dashboard, and
// choose to keep the previous secret active for up to 24 hours. During this
// time, your endpoint has multiple active secrets and Stripe generates one
// signature for each secret.
func parseHeader(header string) (time.Time, [][]byte, error) {
	var (
		ts   time.Time
		sigs [][]byte
	)

	if header == "" {
		return ts, nil, ErrWebhookInvalidSignature
	}

	for _, item := range strings.Split(header, ",") {
		parts := strings.Split(item, "=")
		if len(parts) != 2 {
			return ts, nil, ErrWebhookInvalidSignature
		}

		switch parts[0] {
		case "t":
			timestamp, err := strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				return ts, nil, ErrWebhookInvalidSignature
			}
			ts = time.Unix(timestamp, 0)

		case "v1":
			sig, err := hex.DecodeString(parts[1])
			if err != nil {
				continue // Ignore invalid signatures.
			}
			sigs = append(sigs, sig)

		default:
			continue // Ignore unknown parts of the header.
		}
	}

	if len(sigs) == 0 {
		return ts, nil, ErrWebhookInvalidSignature
	}
	return ts, sigs, nil
}
