[![GoDoc](https://godoc.org/zgo.at/zstripe?status.svg)](https://pkg.go.dev/zgo.at/zstripe)
[![Build Status](https://travis-ci.org/zgoat/zstripe.svg?branch=master)](https://travis-ci.org/zgoat/zstripe)

zstripe is a set of utility functions for working with the Stripe API.

It's not a full "client library"; but just a few functions that make it easy
to call api.stripe.com.

Personally, I prefer working with the API "directly", instead of using an
elaborate wrapper/client library. YMMV. Check out
[`stripe-go`](https://github.com/stripe/stripe-go) if you want a full client
library.

See [`zstripe_stripe_test.go`](zstripe_stripe_test.go) for usage examples.
