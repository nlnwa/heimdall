package main

import (
	"flag"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/nlnwa/heimdall/docs"
	"github.com/nlnwa/heimdall/handlers"
	"github.com/nlnwa/heimdall/pdp"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
	"os"
)

// @title Heimdall API
// @version 0.1.0-alpha
// @description This is the Heimdall API documentation.

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func main() {

	policyFile := flag.String("policy", "", "Path to policy file")
	flag.Parse()

	if *policyFile == "" {
		fmt.Println("Missing required flag -policy")
		os.Exit(1)
	}

	err := pdp.SetPolicies(*policyFile)
	if err != nil {
		fmt.Printf("Failed to initialize pdp: %v\n", err)
		os.Exit(1)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello from server")
	})

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	e.POST("/auth", handlers.AccessHandler)

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
