package zstripe

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestRequest(t *testing.T) {
	n := 0
	tests := []struct {
		name    string
		handler http.HandlerFunc
		want    string
		wantErr string
	}{
		{
			"regular",
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				fmt.Fprintln(w, `{"x": "y"}`)
			},
			"y",
			"",
		},
		{
			"error",
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(400)
				fmt.Fprintln(w, `{"error": {"code": "xx_yy", "message": "oh noes"}}`)
			},
			"",
			"(xx_yy: oh noes)",
		},
		{
			"retry once",
			func(w http.ResponseWriter, r *http.Request) {
				if n == 0 {
					n++
					w.Header().Set("Stripe-Should-Retry", "true")
					w.WriteHeader(400)
					return
				}

				w.WriteHeader(http.StatusOK)
				fmt.Fprintln(w, `{"x": "y"}`)
			},
			"y",
			"",
		},
		{
			"retry indefinitely",
			func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Stripe-Should-Retry", "true")
				w.WriteHeader(400)
			},
			"",
			"retried longer than MaxRetry",
		},
	}

	MaxRetry = 1 * time.Second

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n = 0

			api := httptest.NewServer(http.HandlerFunc(tt.handler))
			defer api.Close()
			API = api.URL

			SecretKey = "sk_test_xxx"
			var scan struct {
				X string `json:"x"`
			}
			_, err := Request(&scan, "GET", "/", nil)
			if !errorContains(err, tt.wantErr) {
				t.Fatalf("wrong error:\nwant: %s\ngot:  %s", tt.wantErr, err)
			}
			if scan.X != tt.want {
				t.Fatalf("scan.X wrong\nwant: %q\ngot:  %q", tt.want, scan.X)
			}
		})
	}
}

func errorContains(out error, want string) bool {
	if out == nil {
		return want == ""
	}
	if want == "" {
		return false
	}
	return strings.Contains(out.Error(), want)
}
