package apimanager

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"apimanager/internal/app/apimanager/actions"
	"apimanager/internal/app/apimanager/gateways"
	"net/http"
)

func StartServer() {

	// Set up logging using zerolog
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Print("Starting application")

	swagger, err := gateways.GetSwagger()
	if err != nil {
		log.Print("Error loading swagger spec", err)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	//Echo instance
	router := echo.New()

	// Middleware

	// log all requests
	router.Use(middleware.Logger())

	// add recovery
	//router.Use(echoMiddleware.Recover())

	// Use codegen validation middleware to check all requests against the OpenAPI schema.
	//router.Use(codegenMiddleware.OapiRequestValidator(swagger))

	router.Static("/swaggerui", "swaggerui/dist")
	router.Static("/swagger", "api/openapi-spec")

	// Create an instance of our handler which satisfies the generated interface
	gatewayInstance := gateways.NewGateway()
	gateways.RegisterHandlers(router, gatewayInstance)

	actionInstance := actions.NewAction()
	actions.RegisterHandlers(router, actionInstance)

	router.GET("/v1/ping", ping)
	router.GET("/v2/ping", pingV2)

	// Start Server
	router.Logger.Fatal(router.Start(":8080"))
}

func ping(c echo.Context) error {
	return c.String(http.StatusOK, "ping")
}

func pingV2(c echo.Context) error {
	return c.String(http.StatusOK, "pingV2")
}