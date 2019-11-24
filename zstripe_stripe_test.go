// +build teststripe

// Tests against the live Stripe API; set STRIPE_SECRET_KEY to your secret
// Stripe testing key.

package zstripe

import (
	"fmt"
	"net/url"
	"os"
	"strings"
	"testing"
)

func init() {
	DebugURL = true
	DebugBody = true

	SecretKey = os.Getenv("STRIPE_SECRET_KEY")
	if SecretKey == "" {
		panic("must set STRIPE_SECRET_KEY")
	}
	if !strings.HasPrefix(SecretKey, "sk_test_") {
		panic("STRIPE_SECRET_KEY doesn't start with sk_test_; don't use a live key for testing!")
	}
}

func TestCreateCustomer(t *testing.T) {
	type Customer struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	c := Customer{}
	body := make(url.Values)
	body.Set("name", "Martin Tournoij")
	body.Set("email", "martin@arp242.net")
	_, err := Request(&c, "POST", "/v1/customers", strings.NewReader(body.Encode()))
	if err != nil {
		t.Fatal(err)
	}

	if !strings.HasPrefix(c.ID, "cus_") {
		t.Fatalf("wrong ID: %q", c.ID)
	}

	fmt.Printf("%#v\n", c)
}
