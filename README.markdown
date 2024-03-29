zstripe is a set of utility functions for working with the Stripe API.

It's not a full "client library"; but just a few functions that make it easy
to call api.stripe.com.

API docs: https://godocs.io/zgo.at/zstripe

---

Personally, I prefer working with the API "directly", instead of using an
elaborate wrapper/client library. YMMV. Check out
[`stripe-go`](https://github.com/stripe/stripe-go) if you want a full client
library.

See [`zstripe_stripe_test.go`](zstripe_stripe_test.go) for usage examples.
