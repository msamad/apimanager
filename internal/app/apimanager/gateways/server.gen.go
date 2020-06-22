// Package gateways provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package gateways

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /gateways)
	ListGateways(ctx echo.Context) error

	// (POST /gateways)
	PostGateways(ctx echo.Context) error

	// (PUT /gateways)
	PutGateways(ctx echo.Context) error

	// (DELETE /gateways/{gatewayId})
	DeleteGatewayById(ctx echo.Context, gatewayId string) error

	// (GET /gateways/{gatewayId})
	GetGatewayById(ctx echo.Context, gatewayId string) error

	// (PUT /gateways/{gatewayId})
	PutGatewayById(ctx echo.Context, gatewayId string) error

	// (GET /gateways/{gatewayId}/endpoint)
	GetGatewayEndpoint1(ctx echo.Context, gatewayId string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// ListGateways converts echo context to params.
func (w *ServerInterfaceWrapper) ListGateways(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ListGateways(ctx)
	return err
}

// PostGateways converts echo context to params.
func (w *ServerInterfaceWrapper) PostGateways(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostGateways(ctx)
	return err
}

// PutGateways converts echo context to params.
func (w *ServerInterfaceWrapper) PutGateways(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PutGateways(ctx)
	return err
}

// DeleteGatewayById converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteGatewayById(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "gatewayId" -------------
	var gatewayId string

	err = runtime.BindStyledParameter("simple", false, "gatewayId", ctx.Param("gatewayId"), &gatewayId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter gatewayId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteGatewayById(ctx, gatewayId)
	return err
}

// GetGatewayById converts echo context to params.
func (w *ServerInterfaceWrapper) GetGatewayById(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "gatewayId" -------------
	var gatewayId string

	err = runtime.BindStyledParameter("simple", false, "gatewayId", ctx.Param("gatewayId"), &gatewayId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter gatewayId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetGatewayById(ctx, gatewayId)
	return err
}

// PutGatewayById converts echo context to params.
func (w *ServerInterfaceWrapper) PutGatewayById(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "gatewayId" -------------
	var gatewayId string

	err = runtime.BindStyledParameter("simple", false, "gatewayId", ctx.Param("gatewayId"), &gatewayId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter gatewayId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PutGatewayById(ctx, gatewayId)
	return err
}

// GetGatewayEndpoint1 converts echo context to params.
func (w *ServerInterfaceWrapper) GetGatewayEndpoint1(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "gatewayId" -------------
	var gatewayId string

	err = runtime.BindStyledParameter("simple", false, "gatewayId", ctx.Param("gatewayId"), &gatewayId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter gatewayId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetGatewayEndpoint1(ctx, gatewayId)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET("/gateways", wrapper.ListGateways)
	router.POST("/gateways", wrapper.PostGateways)
	router.PUT("/gateways", wrapper.PutGateways)
	router.DELETE("/gateways/:gatewayId", wrapper.DeleteGatewayById)
	router.GET("/gateways/:gatewayId", wrapper.GetGatewayById)
	router.PUT("/gateways/:gatewayId", wrapper.PutGatewayById)
	router.GET("/gateways/:gatewayId/endpoint", wrapper.GetGatewayEndpoint1)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9RZTXOjOBP+Ky6975GxPDM3n3YmSaVSNTtJxZPTVg6KaLBm0cdKwh4v5f++JSFswNgB",
	"B+9WcohFfzxqPXS3BBSISq6kAGENmhdIw185GPtVxgy84Dusb4mFNdl8BSIed+qNU1IpLAjrhkSpjFFi",
	"mRT4p5HCyQxdAidu9H8NCZqj/+H9bLjUGtycAW3dXxR8fQh17bxASksF2ob4qBQJS3PtZ3YCu1GA5shY",
	"zUSKthGiGoiF+OvmlPbe+yZSc2LRHMXEwgfLOKDo0CUGQzVTRydkcaeYy5gl7GgglXpIJIJw6AQrBQUC",
	"kXM0/wM93ix+oOcDhJ0lki8/gVrnGti+EbGSTNhFzjnRR8iHYNQRwwnok5BnkntBJn6AsY9g8swei9gS",
	"lnVObnJKwZia7kXKrMzyjglblXBGrr/G3iVoelIuRy8c+NnxORETiazaFaE+W4H7e4binPPNb/7/lEqO",
	"KopKBYpQrp3Z0lo1x9gLS2MmD8JG9wrEl4e7yefpbGIUUJaEjjiRycQpfieCpKBRhDJGQRi/pDDhF0Xo",
	"EiafprPGrGaO8Xq9nhKvnkqd4uBr8Le7q5vvi5sPn6az6dLyzBMFmpv7ZAF6xSjsQyeKcSKwV2PXS5jN",
	"/LSNsFagTbmWj9PZdOYApQJBFENz9NmLIqSIXfrbigl1q3PjbYRwLNcikyQ2uKiGd/E2KNMyP7xfCv4e",
	"uAzx9NzFaI6+MWNvK6MIaTBKulU6y0+z2aDthlng5rV9p6Mb7dOHaE02ZfY073Eo6STPJrv4y7wjqXGZ",
	"uVvD8zZCSpqOpT7I1lIbm2pXzI1tGR/fk7dvJK4HX9Um/UZi8i5e8vdLS2uXeCNB9YrBRRiVxYRiyMDC",
	"IXvXXl4xsLmLfalqwsGCdhMUiLk4XPnu+9wOOzDONMRobnUOUY2Adpt9bhEaQ0LyzB5s4AOzorMx3IL9",
	"bxb1Xovn36Bp/MeAw2NE+SAwfp4dKy5cP86+korV8fjj+87HrlP+CK1L6pQI9jepnw4aMlzUL/enhNNG",
	"/ryxYnbT01ox098SF0Sx/pHUPM4Ia+caDlzmDa64CKPtKCBjrKaG5rJVE2rHWSGOIWGChceEEfBqJT8C",
	"GgerGTWYZgyEOw744l1YMtbymxM8GZLCuMAqcyeni8Xt4C8Qtb5cxKbsjBfkJB+PD0evGREKF+7HWSiZ",
	"Mbq5WhI2UuF5wPL1wIhouCgDPa+PdwBrkDoG/TBqtMYSm/fEKit9mDEuykFvFg79Bu4CHQDDdrYTAEP3",
	"t15Q462vjqnYnzA2ooaUGas31QHvIti/eDYm9MBNfyAoLqrhWxL8REcmig3YpPphD+t3wzAHd71+8Gf1",
	"Pg78xT8JDTDGRW5A945+wC7X3MXO8RnYKFrOw5rgEeehDfBVmHHWdHZy98cbnNivQ5+V1Fpmg0xx4X7c",
	"ZSuzVZanbJcO4QqTFWEZecngoVNdlIMDlJoinM+uITF9bCperyFxikRqXrkd4HQ5hUewfUR1gpoMBKEB",
	"ounSddYrYkkmUwzC6v0NONQLwsEoQjtMWpLmESkIu15BBFU92CByd2kn2hgLHMMvJbVtyhg/lDUOcx4H",
	"01xrENZd4PDtaa9sZURT2K7MtlYxqVNzVHlEUxJ03BNiZl36fWgy2bLqYrRlokBzZswJg4yI43GsGKzb",
	"cWwjZECvqpdrjc9xRCl3SzK0fd69itp/T6PlW6poJ7kOH6XqsvDiqi66r62zLi+LsyHx9TDZF0Rd+Sgz",
	"qF8vfKY1JD6B6pInAxptn7f/BAAA//9Ot/QPliEAAA==",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}

