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

// TestAccessHandler tests the AccessHandler function
func TestAccessHandler(t *testing.T) {

	tests := []struct {
		body      string
		wantError bool
	}{
		{`{"url":"https://nb.no", "token":"3"}`, false},
		{`{"url": "https://nb.no/paywall/", "token"}`, true}, // invalid json
	}
	for _, tt := range tests {
		t.Run(tt.body, func(t *testing.T) {
			e := echo.New()

			req := httptest.NewRequest(http.MethodPost, "http://example.test/auth", strings.NewReader(tt.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err := AccessHandler(c)
			if err != nil {
				if !tt.wantError {
					t.Errorf("Unexpexted error = %v", err)
				}
				return
			}

			var got = pdp.AccessResponse{}

			err = json.Unmarshal(rec.Body.Bytes(), &got)
			if err != nil {
				t.Errorf("Error unmarshaling AccessResponse %v, rec.body.Bytes %v", err, rec.Body.Bytes())
			}
			if got.Permission != pdp.Deny {
				t.Errorf("Expected permission %v, got %v", pdp.Deny, got.Permission)
			}
		})
	}
}
