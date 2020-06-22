// Package system provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package system

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

	// (GET /system/export)
	ExportData1(ctx echo.Context, params ExportData1Params) error

	// (POST /system/import)
	ImportData1(ctx echo.Context) error

	// (GET /system/status)
	GetStatus1(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// ExportData1 converts echo context to params.
func (w *ServerInterfaceWrapper) ExportData1(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params ExportData1Params
	// ------------- Optional query parameter "download" -------------

	err = runtime.BindQueryParameter("form", true, false, "download", ctx.QueryParams(), &params.Download)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter download: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ExportData1(ctx, params)
	return err
}

// ImportData1 converts echo context to params.
func (w *ServerInterfaceWrapper) ImportData1(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ImportData1(ctx)
	return err
}

// GetStatus1 converts echo context to params.
func (w *ServerInterfaceWrapper) GetStatus1(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetStatus1(ctx)
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

	router.GET("/system/export", wrapper.ExportData1)
	router.POST("/system/import", wrapper.ImportData1)
	router.GET("/system/status", wrapper.GetStatus1)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/7RYzXLbNhB+FQ3aIyMoyU2nprEn45m29sSTU8eHNbSikOKvAChF5ejdOwBBCZJoh5Sp",
	"iwTsLj7sAosF+NWEaWm0QuUdmdfEsRVKiM3HrfMoHz34yv2OoILMWG3Qeo7R4rniwt9Hhd8aJHPivOWq",
	"JLuCLNAxy43nulvPF51iqS3eqaXuVCqQ2KmoTCZ+1loEd3cFWaN13Q7silain78j82QXRDzNzLTywHxo",
	"ogQuyJwsKim3v8XfKdOStN40ClKQygazlfdmTmkUNsZcny0HuTeoPj3cTT5OZxNnkPElZxB0E72cBMWf",
	"oKBESwoiOEPlYthpwk8G2AonH6azo1ndnNLNZjOFqJ5qW9I01tE/7j7f/vV4++7DdDZdeSmCRx6tdPfL",
	"R7RrzvDgOhguQdGopqQgnnsRpz1ya7+y5P10Np0FQG1QgeFkTj5GUUEM+FVMFAosRBfau4LQhd4ooWHh",
	"aN027xa7pCzB4wa27qRL69Q6t8xVFNXCaK58stG2BMX/g3z6Ixmt8+4B/HWjGNCa+21Pa8Ndf0tag+H9",
	"PclGXODWfmjaUfeGobROrd0oIGNEk6GFQ22B+XEipAtccsVTfRsB77XEHY4m0VvOHGWCo/Jf0RmtHIZi",
	"7q4xwTcHJY4LbASoK/od4K/gtb2ex66SEuz2imtSjbceYXndiFC0Dn/BwmjB2fbzCvhIBy8CNi+aEdFo",
	"3Th6WR3vALao7QLtw6jeuvi664fVnPRhxrRuGr1X4XzcwFugA2DYzfYKwND7rRfUePHlmIb/g2MjWiy5",
	"83ZLv7u+l95w7B9SjAk98NIfCErrtvmWBH+lIoPhAy6pftjD6t0wzMFVrx/8RbVPonxGO8yY1pVD29v7",
	"Abfc8S12yZiBheJk8LAi+MLgoQXwpzDjxHRxcvfHG5zYP4e+KKmtFoNMaR3+Qvcks42oSr5Ph9SjsAYu",
	"4FngQ6e6bhpnKJkivc9ucOn62LTreoPLoFhqK9thZzhdg9In2MGjfIGOVyAJHYJlq1BZP4MHoUuKytvD",
	"BpzrFUh0BliHyYnk+ImUhF0URFLlziZR2KW9KDJwFH8YbSMZVWL80wZt2mYyJ7dRfQMe3kfOxYJEH0H+",
	"rkl4JJN/K7TbjLFKnAspEt3XxY89FaT9lokUzgKXUAnfNHMuy1WMoXPLSkz2fjX8GpTBh8QjkqddFhKX",
	"bUhGu46Y7mQe0xUd2T+Au9f2C/qGAD1z48Ns1jKFGD/aCRgjEo3XvpCy9f3V4pLMyS/0wLfSRLbSM6Y1",
	"kpFviS5mEWWVtah86NBEbR6UJ/XgWHhal0+1hmtbuheVL2ia4/HySFxwH4rPu+NzdGLVdZ5OTAxayZ17",
	"xUCAetmPNcfNqR+7gji06/ZcHbG9YEzIaEHCqUn7caBrWbNVxV5y056/TPalYTBz0X0WZy5vSvORJFbD",
	"yaEc5sqvWmDef4x15kjSZE4m+ebQkt3T7v8AAAD//4U5dsIeGAAA",
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

