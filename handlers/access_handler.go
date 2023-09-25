package handlers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/nlnwa/heimdall/pdp"
	"log"
	"net/http"
)

// AccessHandler godoc
// @Summary Check if user with a certain role has access to an url
// @Description Check if user with a certain role has access to an url
// @ID access-handler
// @Accept  json
// @Produce  json
// @Param request body pdp.AccessRequest true "query params"
// @Success 200 {object} pdp.AccessResponse
// @Router /auth [post]
func AccessHandler(c echo.Context) error {
	accRec := pdp.AccessRequest{}
	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&accRec)
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
	}

	accRes := pdp.CanAccess(accRec)

	return c.JSON(http.StatusOK, accRes)
}
