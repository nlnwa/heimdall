package handlers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/nlnwa/heimdall/pdp"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_pdp(t *testing.T) {

	tests := []struct {
		url  string
		body string
		want pdp.AccessResponse
	}{
		{"http://localhost:8080/auth", `{"url":"http://nb.no", "token":"3"}`, pdp.AccessResponse{Permission: pdp.Allow}},                // default policy should allow access
		{"http://localhost:8080/auth", `{"url": "http://nb.no/paywall/", "token": "3"}`, pdp.AccessResponse{Permission: pdp.Deny}},      // default policy should deny access
		{"http://localhost:8080/auth", `{"url": "http://nb.no/confidential/", "token": "3"}`, pdp.AccessResponse{Permission: pdp.Deny}}, // default policy should deny access
		{"http://localhost:8080/auth", `{"url": "http://nb.no", "token": "2"}`, pdp.AccessResponse{Permission: pdp.Allow}},              // curator policy should allow access
		{"http://localhost:8080/auth", `{"url": "http://nb.no/paywall", "token": "2"}`, pdp.AccessResponse{Permission: pdp.Allow}},      // curator policy should allow access
		{"http://localhost:8080/auth", `{"url": "http://nb.no/confidential", "token": "2"}`, pdp.AccessResponse{Permission: pdp.Deny}},  // curator policy should deny access
		{"http://localhost:8080/auth", `{"url": "http://nb.no", "token": "1"}`, pdp.AccessResponse{Permission: pdp.Allow}},              // admin policy should allow access
		{"http://localhost:8080/auth", `{"url": "http://nb.no/paywall", "token": "1"}`, pdp.AccessResponse{Permission: pdp.Allow}},      // admin policy should allow access
		{"http://localhost:8080/auth", `{"url": "http://nb.no/confidential", "token": "1"}`, pdp.AccessResponse{Permission: pdp.Allow}}, // admin policy should allow access
	}
	for _, tt := range tests {
		t.Run(tt.url, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, tt.url, strings.NewReader(tt.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if err := AccessHandler(c); err != nil {
				t.Errorf("AccessHandler() error = %v", err)
			}

			var got = &pdp.AccessResponse{}

			err := json.Unmarshal([]byte(rec.Body.Bytes()), got)
			if err != nil {
				t.Errorf("Error unmarshaling AccessResponse %v, rec.body.Bytes %v", err, rec.Body.Bytes())
			}
			if got.Permission != tt.want.Permission {
				t.Errorf("CanAcces(): got = %v, want %v", got, tt.want)
			}
		})
	}
}
