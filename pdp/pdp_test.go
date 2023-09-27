package pdp

import (
	"fmt"
	"testing"
)

// TestCanAccess tests the CanAccess function
func TestCanAccess(t *testing.T) {
	err := Init("../testdata/policy_example.yaml")
	if err != nil {
		t.Fatalf("Failed to initialize policy: %v", err)
	}

	tests := []struct {
		req  AccessRequest
		want AccessResponse
	}{
		{AccessRequest{Url: "http://nb.no", Token: "3"}, AccessResponse{Permission: Allow}},              // default policy should allow access
		{AccessRequest{Url: "http://nb.no/paywall/", Token: "3"}, AccessResponse{Permission: Deny}},      // default policy should deny access
		{AccessRequest{Url: "http://nb.no/confidential/", Token: "3"}, AccessResponse{Permission: Deny}}, // default policy should deny access
		{AccessRequest{Url: "http://nb.no", Token: "2"}, AccessResponse{Permission: Allow}},              // curator policy should allow access
		{AccessRequest{Url: "http://nb.no/paywall", Token: "2"}, AccessResponse{Permission: Allow}},      // curator policy should allow access
		{AccessRequest{Url: "http://nb.no/confidential", Token: "2"}, AccessResponse{Permission: Deny}},  // curator policy should deny access
		{AccessRequest{Url: "http://nb.no", Token: "1"}, AccessResponse{Permission: Allow}},              // admin policy should allow access
		{AccessRequest{Url: "http://nb.no/paywall", Token: "1"}, AccessResponse{Permission: Allow}},      // admin policy should allow access
		{AccessRequest{Url: "http://nb.no/confidential", Token: "1"}, AccessResponse{Permission: Allow}}, // admin policy should allow access
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.req), func(t *testing.T) {
			t.Parallel()

			var got = CanAccess(tt.req)

			if got.Permission != tt.want.Permission {
				t.Errorf("CanAcces(): got = %v, want %v", got, tt.want)
			}
		})
	}
}
