// Package plugins provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package plugins

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

	// (GET /plugins)
	List(ctx echo.Context) error

	// (POST /plugins)
	Create1(ctx echo.Context) error

	// (GET /plugins/availablePlugins)
	GetAvailablePlugins1(ctx echo.Context) error

	// (DELETE /plugins/{pluginId})
	Delete1(ctx echo.Context, pluginId int64) error

	// (GET /plugins/{pluginId})
	Get1(ctx echo.Context, pluginId int64) error

	// (GET /plugins/{pluginId}/policyDefs)
	GetPolicyDefs1(ctx echo.Context, pluginId int64) error

	// (GET /plugins/{pluginId}/policyDefs/{policyDefId}/form)
	GetPolicyForm1(ctx echo.Context, pluginId int64, policyDefId string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// List converts echo context to params.
func (w *ServerInterfaceWrapper) List(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.List(ctx)
	return err
}

// Create1 converts echo context to params.
func (w *ServerInterfaceWrapper) Create1(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.Create1(ctx)
	return err
}

// GetAvailablePlugins1 converts echo context to params.
func (w *ServerInterfaceWrapper) GetAvailablePlugins1(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetAvailablePlugins1(ctx)
	return err
}

// Delete1 converts echo context to params.
func (w *ServerInterfaceWrapper) Delete1(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "pluginId" -------------
	var pluginId int64

	err = runtime.BindStyledParameter("simple", false, "pluginId", ctx.Param("pluginId"), &pluginId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pluginId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.Delete1(ctx, pluginId)
	return err
}

// Get1 converts echo context to params.
func (w *ServerInterfaceWrapper) Get1(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "pluginId" -------------
	var pluginId int64

	err = runtime.BindStyledParameter("simple", false, "pluginId", ctx.Param("pluginId"), &pluginId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pluginId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.Get1(ctx, pluginId)
	return err
}

// GetPolicyDefs1 converts echo context to params.
func (w *ServerInterfaceWrapper) GetPolicyDefs1(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "pluginId" -------------
	var pluginId int64

	err = runtime.BindStyledParameter("simple", false, "pluginId", ctx.Param("pluginId"), &pluginId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pluginId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPolicyDefs1(ctx, pluginId)
	return err
}

// GetPolicyForm1 converts echo context to params.
func (w *ServerInterfaceWrapper) GetPolicyForm1(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "pluginId" -------------
	var pluginId int64

	err = runtime.BindStyledParameter("simple", false, "pluginId", ctx.Param("pluginId"), &pluginId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pluginId: %s", err))
	}

	// ------------- Path parameter "policyDefId" -------------
	var policyDefId string

	err = runtime.BindStyledParameter("simple", false, "policyDefId", ctx.Param("policyDefId"), &policyDefId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter policyDefId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPolicyForm1(ctx, pluginId, policyDefId)
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

	router.GET("/plugins", wrapper.List)
	router.POST("/plugins", wrapper.Create1)
	router.GET("/plugins/availablePlugins", wrapper.GetAvailablePlugins1)
	router.DELETE("/plugins/:pluginId", wrapper.Delete1)
	router.GET("/plugins/:pluginId", wrapper.Get1)
	router.GET("/plugins/:pluginId/policyDefs", wrapper.GetPolicyDefs1)
	router.GET("/plugins/:pluginId/policyDefs/:policyDefId/form", wrapper.GetPolicyForm1)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9xZTXPbNhP+Kx6871ER5KTTg05N7DTjThp74uaU8WENLimkxEcB0IrK0X/vAKTED1ES",
	"KdFpk4sE7i4e7i4WH3yQE6aEVhKls2SeE8sWKCA0P+DyLs0SLt8gSC/QRmk0jmNQg3E8BuZuIv/kVhrJ",
	"nFhnuEzIekJYCtbymKPpVEdomeHacSU79YlRmd4DLUFgp6IQdCgynRiI6rpHpVIf1npCntDYbje2kEQ9",
	"fkHmvPXzpYQZBIfRm9Uh7W14bayMAEfmJAKHLxwXSCZdOU7RYdQd9TkDwKOGE1y6n3+qHODSYYLmtJE6",
	"ZTDuMyHArL6TMfkR0q5SzlbXGHPJfSQHB+BYxD6eP0q/UGaCzD+Ta4whSx2ZkN+skvdhTSIPHenkbA8q",
	"H7h06FBJN30zrEMGboROe6XMi7iMlTdmSjpgLoQrgKe+YjIhVr+E3ylTgmz8LBRkQjLjzRbO6TmlQVgY",
	"c7VTU+RWo3x9d3Pxajq7sBoZjzkDr7tQ8YVX/A4SfBgTknKG0oaElC98rYEt8OLldNZ4q51TulwupxDU",
	"U2USWva19P3N1dsP929fvJzOpgsn0lBnaIS9je/RPHGGleuguQBJg5r6xHKXhtc23NoWI7mczqYzD6g0",
	"StCczMmrIJoQDW4R6osC89H59npCaKSWMlUQWZpvmjfRulQm4HAJK9t6pHnZ2rWsqyjKSCsuXWmjTAKS",
	"/w311zdkNK8/VuCHjUJAT9ytelprbvtb0hw07+9JrccJbm27liNqz+hK87K1HgVkjGhqaH5SG2BunAhp",
	"tF1ax8E7VLjD0QQ6w5mlLOUo3Ue0WkmL9w7GCr/5gk8WEhwXWKcgn9FvD/8MXpvn89gWO/gz5iQbLx8+",
	"vXZEKJr7P29RbOtXC+AjTbwAWByERkSjeXn+iEZaDA0qE6G5G9Vb68BlPbGKmT7MmOZFo3cWdvsN3AU6",
	"AIbtbAcAhu5vvaDGi6+OqfmfODaiwYRbZ1b0i+276Q3H/irSMaEHbvoDQWm+aZ5T4AdWZNB8wCbVD3vY",
	"ejcMc/Cq1w/+pLVPoHhEM8yY5plF09v7Abtccxc7pc/AhaLVedgiuKfz0AXwKMw4MZ1c3P3xBhf2ceiT",
	"itqodJApzf2ff2xVdkGmhK/zBAPToTSasjOZk/fcOjIhm/NsMHw5m23YEQwfKgS0TkvqYrMrlPS0b3GH",
	"InT8v8GYzMn/aEVk05LFprv8YMXNgDGwKqiZJoFiM8bQ2jhLL7ZuF6QOJJbMP5esoyUPgQiyHQFeBWLw",
	"MsT4V4bWvVHRalB4h6JqEvPrEMJZuTyews2bzshVrTAoPAFP4THFuyOV8g7d65bt5Q9ROfVs5BvycV3w",
	"pSk63E3FdZBfBurLgEAX9p/POfHfKoEOq3jDLZtZFCA3GJG5MxlOank4SnOuH1qZjkpats3qDpsy+wb6",
	"Xw3tPz9lqiIpv1KvMT44b7Ykvf2uEttvdh64gPgW87Q2BJvN+xpjr/AJOj4qvyojvumoTLrBK88P4rfv",
	"N84d4zbe2QNUnxBNQW14ylqpjin1807zQFMKLYJhC/+hdAUOUpVQlM5U56ldvU+s1cA6TFqSJuNRCrtu",
	"FEpV3dlS5A9dW9HKOhQUv2plXFPGxa6swc0EHMoyY1A6/0DLu6pK2TrgNYXtg3Zbq7kyid2r3KMpErS/",
	"J0bc+WPBi2YmW1ZdGW2ZaDSCW3vAIAW5348njsu2H+sJsWieNvO6cX0HWvshSYmfRmU9V/dvrCj1yVZy",
	"XV5i1WXviiupuui2FmddXkyUhiTMh4tqQtSVH1WK9ef7UGkNSSiguuST9UvMw/qfAAAA//8lQ4J1SyIA",
	"AA==",
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

