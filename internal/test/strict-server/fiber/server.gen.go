// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version (devel) DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/gerhardwagner/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gofiber/fiber/v2"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /json)
	JSONExample(c *fiber.Ctx) error

	// (POST /multipart)
	MultipartExample(c *fiber.Ctx) error

	// (POST /multiple)
	MultipleRequestAndResponseTypes(c *fiber.Ctx) error

	// (GET /reserved-go-keyword-parameters/{type})
	ReservedGoKeywordParameters(c *fiber.Ctx, pType string) error

	// (POST /reusable-responses)
	ReusableResponses(c *fiber.Ctx) error

	// (POST /text)
	TextExample(c *fiber.Ctx) error

	// (POST /unknown)
	UnknownExample(c *fiber.Ctx) error

	// (POST /unspecified-content-type)
	UnspecifiedContentType(c *fiber.Ctx) error

	// (POST /urlencoded)
	URLEncodedExample(c *fiber.Ctx) error

	// (POST /with-headers)
	HeadersExample(c *fiber.Ctx, params HeadersExampleParams) error

	// (POST /with-union)
	UnionExample(c *fiber.Ctx) error
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

type MiddlewareFunc fiber.Handler

// JSONExample operation middleware
func (siw *ServerInterfaceWrapper) JSONExample(c *fiber.Ctx) error {

	return siw.Handler.JSONExample(c)
}

// MultipartExample operation middleware
func (siw *ServerInterfaceWrapper) MultipartExample(c *fiber.Ctx) error {

	return siw.Handler.MultipartExample(c)
}

// MultipleRequestAndResponseTypes operation middleware
func (siw *ServerInterfaceWrapper) MultipleRequestAndResponseTypes(c *fiber.Ctx) error {

	return siw.Handler.MultipleRequestAndResponseTypes(c)
}

// ReservedGoKeywordParameters operation middleware
func (siw *ServerInterfaceWrapper) ReservedGoKeywordParameters(c *fiber.Ctx) error {

	var err error

	// ------------- Path parameter "type" -------------
	var pType string

	err = runtime.BindStyledParameter("simple", false, "type", c.Params("type"), &pType)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter type: %w", err).Error())
	}

	return siw.Handler.ReservedGoKeywordParameters(c, pType)
}

// ReusableResponses operation middleware
func (siw *ServerInterfaceWrapper) ReusableResponses(c *fiber.Ctx) error {

	return siw.Handler.ReusableResponses(c)
}

// TextExample operation middleware
func (siw *ServerInterfaceWrapper) TextExample(c *fiber.Ctx) error {

	return siw.Handler.TextExample(c)
}

// UnknownExample operation middleware
func (siw *ServerInterfaceWrapper) UnknownExample(c *fiber.Ctx) error {

	return siw.Handler.UnknownExample(c)
}

// UnspecifiedContentType operation middleware
func (siw *ServerInterfaceWrapper) UnspecifiedContentType(c *fiber.Ctx) error {

	return siw.Handler.UnspecifiedContentType(c)
}

// URLEncodedExample operation middleware
func (siw *ServerInterfaceWrapper) URLEncodedExample(c *fiber.Ctx) error {

	return siw.Handler.URLEncodedExample(c)
}

// HeadersExample operation middleware
func (siw *ServerInterfaceWrapper) HeadersExample(c *fiber.Ctx) error {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params HeadersExampleParams

	headers := c.GetReqHeaders()

	// ------------- Required header parameter "header1" -------------
	if value, found := headers[http.CanonicalHeaderKey("header1")]; found {
		var Header1 string

		err = runtime.BindStyledParameterWithLocation("simple", false, "header1", runtime.ParamLocationHeader, value, &Header1)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter header1: %w", err).Error())
		}

		params.Header1 = Header1

	} else {
		err = fmt.Errorf("Header parameter header1 is required, but not found: %w", err)
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// ------------- Optional header parameter "header2" -------------
	if value, found := headers[http.CanonicalHeaderKey("header2")]; found {
		var Header2 int

		err = runtime.BindStyledParameterWithLocation("simple", false, "header2", runtime.ParamLocationHeader, value, &Header2)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter header2: %w", err).Error())
		}

		params.Header2 = &Header2

	}

	return siw.Handler.HeadersExample(c, params)
}

// UnionExample operation middleware
func (siw *ServerInterfaceWrapper) UnionExample(c *fiber.Ctx) error {

	return siw.Handler.UnionExample(c)
}

// FiberServerOptions provides options for the Fiber server.
type FiberServerOptions struct {
	BaseURL     string
	Middlewares []MiddlewareFunc
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router fiber.Router, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, FiberServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router fiber.Router, si ServerInterface, options FiberServerOptions) {
	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	for _, m := range options.Middlewares {
		router.Use(m)
	}

	router.Post(options.BaseURL+"/json", wrapper.JSONExample)

	router.Post(options.BaseURL+"/multipart", wrapper.MultipartExample)

	router.Post(options.BaseURL+"/multiple", wrapper.MultipleRequestAndResponseTypes)

	router.Get(options.BaseURL+"/reserved-go-keyword-parameters/:type", wrapper.ReservedGoKeywordParameters)

	router.Post(options.BaseURL+"/reusable-responses", wrapper.ReusableResponses)

	router.Post(options.BaseURL+"/text", wrapper.TextExample)

	router.Post(options.BaseURL+"/unknown", wrapper.UnknownExample)

	router.Post(options.BaseURL+"/unspecified-content-type", wrapper.UnspecifiedContentType)

	router.Post(options.BaseURL+"/urlencoded", wrapper.URLEncodedExample)

	router.Post(options.BaseURL+"/with-headers", wrapper.HeadersExample)

	router.Post(options.BaseURL+"/with-union", wrapper.UnionExample)

}

type BadrequestResponse struct {
}

type ReusableresponseResponseHeaders struct {
	Header1 string
	Header2 int
}
type ReusableresponseJSONResponse struct {
	Body Example

	Headers ReusableresponseResponseHeaders
}

type JSONExampleRequestObject struct {
	Body *JSONExampleJSONRequestBody
}

type JSONExampleResponseObject interface {
	VisitJSONExampleResponse(ctx *fiber.Ctx) error
}

type JSONExample200JSONResponse Example

func (response JSONExample200JSONResponse) VisitJSONExampleResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(200)

	return ctx.JSON(&response)
}

type JSONExample400Response = BadrequestResponse

func (response JSONExample400Response) VisitJSONExampleResponse(ctx *fiber.Ctx) error {
	ctx.Status(400)
	return nil
}

type JSONExampledefaultResponse struct {
	StatusCode int
}

func (response JSONExampledefaultResponse) VisitJSONExampleResponse(ctx *fiber.Ctx) error {
	ctx.Status(response.StatusCode)
	return nil
}

type MultipartExampleRequestObject struct {
	Body *multipart.Reader
}

type MultipartExampleResponseObject interface {
	VisitMultipartExampleResponse(ctx *fiber.Ctx) error
}

type MultipartExample200MultipartResponse func(writer *multipart.Writer) error

func (response MultipartExample200MultipartResponse) VisitMultipartExampleResponse(ctx *fiber.Ctx) error {
	writer := multipart.NewWriter(ctx.Response().BodyWriter())
	ctx.Response().Header.Set("Content-Type", writer.FormDataContentType())
	ctx.Status(200)

	defer writer.Close()
	return response(writer)
}

type MultipartExample400Response = BadrequestResponse

func (response MultipartExample400Response) VisitMultipartExampleResponse(ctx *fiber.Ctx) error {
	ctx.Status(400)
	return nil
}

type MultipartExampledefaultResponse struct {
	StatusCode int
}

func (response MultipartExampledefaultResponse) VisitMultipartExampleResponse(ctx *fiber.Ctx) error {
	ctx.Status(response.StatusCode)
	return nil
}

type MultipleRequestAndResponseTypesRequestObject struct {
	JSONBody      *MultipleRequestAndResponseTypesJSONRequestBody
	FormdataBody  *MultipleRequestAndResponseTypesFormdataRequestBody
	Body          io.Reader
	MultipartBody *multipart.Reader
	TextBody      *MultipleRequestAndResponseTypesTextRequestBody
}

type MultipleRequestAndResponseTypesResponseObject interface {
	VisitMultipleRequestAndResponseTypesResponse(ctx *fiber.Ctx) error
}

type MultipleRequestAndResponseTypes200JSONResponse Example

func (response MultipleRequestAndResponseTypes200JSONResponse) VisitMultipleRequestAndResponseTypesResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(200)

	return ctx.JSON(&response)
}

type MultipleRequestAndResponseTypes200FormdataResponse Example

func (response MultipleRequestAndResponseTypes200FormdataResponse) VisitMultipleRequestAndResponseTypesResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/x-www-form-urlencoded")
	ctx.Status(200)

	if form, err := runtime.MarshalForm(response, nil); err != nil {
		return err
	} else {
		_, err := ctx.WriteString(form.Encode())
		return err
	}
}

type MultipleRequestAndResponseTypes200ImagepngResponse struct {
	Body          io.Reader
	ContentLength int64
}

func (response MultipleRequestAndResponseTypes200ImagepngResponse) VisitMultipleRequestAndResponseTypesResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "image/png")
	if response.ContentLength != 0 {
		ctx.Response().Header.Set("Content-Length", fmt.Sprint(response.ContentLength))
	}
	ctx.Status(200)

	if closer, ok := response.Body.(io.ReadCloser); ok {
		defer closer.Close()
	}
	_, err := io.Copy(ctx.Response().BodyWriter(), response.Body)
	return err
}

type MultipleRequestAndResponseTypes200MultipartResponse func(writer *multipart.Writer) error

func (response MultipleRequestAndResponseTypes200MultipartResponse) VisitMultipleRequestAndResponseTypesResponse(ctx *fiber.Ctx) error {
	writer := multipart.NewWriter(ctx.Response().BodyWriter())
	ctx.Response().Header.Set("Content-Type", writer.FormDataContentType())
	ctx.Status(200)

	defer writer.Close()
	return response(writer)
}

type MultipleRequestAndResponseTypes200TextResponse string

func (response MultipleRequestAndResponseTypes200TextResponse) VisitMultipleRequestAndResponseTypesResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "text/plain")
	ctx.Status(200)

	_, err := ctx.WriteString(string(response))
	return err
}

type MultipleRequestAndResponseTypes400Response = BadrequestResponse

func (response MultipleRequestAndResponseTypes400Response) VisitMultipleRequestAndResponseTypesResponse(ctx *fiber.Ctx) error {
	ctx.Status(400)
	return nil
}

type ReservedGoKeywordParametersRequestObject struct {
	Type string `json:"type"`
}

type ReservedGoKeywordParametersResponseObject interface {
	VisitReservedGoKeywordParametersResponse(ctx *fiber.Ctx) error
}

type ReservedGoKeywordParameters200TextResponse string

func (response ReservedGoKeywordParameters200TextResponse) VisitReservedGoKeywordParametersResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "text/plain")
	ctx.Status(200)

	_, err := ctx.WriteString(string(response))
	return err
}

type ReusableResponsesRequestObject struct {
	Body *ReusableResponsesJSONRequestBody
}

type ReusableResponsesResponseObject interface {
	VisitReusableResponsesResponse(ctx *fiber.Ctx) error
}

type ReusableResponses200JSONResponse struct{ ReusableresponseJSONResponse }

func (response ReusableResponses200JSONResponse) VisitReusableResponsesResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("header1", fmt.Sprint(response.Headers.Header1))
	ctx.Response().Header.Set("header2", fmt.Sprint(response.Headers.Header2))
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(200)

	return ctx.JSON(&response.Body)
}

type ReusableResponses400Response = BadrequestResponse

func (response ReusableResponses400Response) VisitReusableResponsesResponse(ctx *fiber.Ctx) error {
	ctx.Status(400)
	return nil
}

type ReusableResponsesdefaultResponse struct {
	StatusCode int
}

func (response ReusableResponsesdefaultResponse) VisitReusableResponsesResponse(ctx *fiber.Ctx) error {
	ctx.Status(response.StatusCode)
	return nil
}

type TextExampleRequestObject struct {
	Body *TextExampleTextRequestBody
}

type TextExampleResponseObject interface {
	VisitTextExampleResponse(ctx *fiber.Ctx) error
}

type TextExample200TextResponse string

func (response TextExample200TextResponse) VisitTextExampleResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "text/plain")
	ctx.Status(200)

	_, err := ctx.WriteString(string(response))
	return err
}

type TextExample400Response = BadrequestResponse

func (response TextExample400Response) VisitTextExampleResponse(ctx *fiber.Ctx) error {
	ctx.Status(400)
	return nil
}

type TextExampledefaultResponse struct {
	StatusCode int
}

func (response TextExampledefaultResponse) VisitTextExampleResponse(ctx *fiber.Ctx) error {
	ctx.Status(response.StatusCode)
	return nil
}

type UnknownExampleRequestObject struct {
	Body io.Reader
}

type UnknownExampleResponseObject interface {
	VisitUnknownExampleResponse(ctx *fiber.Ctx) error
}

type UnknownExample200Videomp4Response struct {
	Body          io.Reader
	ContentLength int64
}

func (response UnknownExample200Videomp4Response) VisitUnknownExampleResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "video/mp4")
	if response.ContentLength != 0 {
		ctx.Response().Header.Set("Content-Length", fmt.Sprint(response.ContentLength))
	}
	ctx.Status(200)

	if closer, ok := response.Body.(io.ReadCloser); ok {
		defer closer.Close()
	}
	_, err := io.Copy(ctx.Response().BodyWriter(), response.Body)
	return err
}

type UnknownExample400Response = BadrequestResponse

func (response UnknownExample400Response) VisitUnknownExampleResponse(ctx *fiber.Ctx) error {
	ctx.Status(400)
	return nil
}

type UnknownExampledefaultResponse struct {
	StatusCode int
}

func (response UnknownExampledefaultResponse) VisitUnknownExampleResponse(ctx *fiber.Ctx) error {
	ctx.Status(response.StatusCode)
	return nil
}

type UnspecifiedContentTypeRequestObject struct {
	ContentType string
	Body        io.Reader
}

type UnspecifiedContentTypeResponseObject interface {
	VisitUnspecifiedContentTypeResponse(ctx *fiber.Ctx) error
}

type UnspecifiedContentType200VideoResponse struct {
	Body          io.Reader
	ContentType   string
	ContentLength int64
}

func (response UnspecifiedContentType200VideoResponse) VisitUnspecifiedContentTypeResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", response.ContentType)
	if response.ContentLength != 0 {
		ctx.Response().Header.Set("Content-Length", fmt.Sprint(response.ContentLength))
	}
	ctx.Status(200)

	if closer, ok := response.Body.(io.ReadCloser); ok {
		defer closer.Close()
	}
	_, err := io.Copy(ctx.Response().BodyWriter(), response.Body)
	return err
}

type UnspecifiedContentType400Response = BadrequestResponse

func (response UnspecifiedContentType400Response) VisitUnspecifiedContentTypeResponse(ctx *fiber.Ctx) error {
	ctx.Status(400)
	return nil
}

type UnspecifiedContentType401Response struct {
}

func (response UnspecifiedContentType401Response) VisitUnspecifiedContentTypeResponse(ctx *fiber.Ctx) error {
	ctx.Status(401)
	return nil
}

type UnspecifiedContentType403Response struct {
}

func (response UnspecifiedContentType403Response) VisitUnspecifiedContentTypeResponse(ctx *fiber.Ctx) error {
	ctx.Status(403)
	return nil
}

type UnspecifiedContentTypedefaultResponse struct {
	StatusCode int
}

func (response UnspecifiedContentTypedefaultResponse) VisitUnspecifiedContentTypeResponse(ctx *fiber.Ctx) error {
	ctx.Status(response.StatusCode)
	return nil
}

type URLEncodedExampleRequestObject struct {
	Body *URLEncodedExampleFormdataRequestBody
}

type URLEncodedExampleResponseObject interface {
	VisitURLEncodedExampleResponse(ctx *fiber.Ctx) error
}

type URLEncodedExample200FormdataResponse Example

func (response URLEncodedExample200FormdataResponse) VisitURLEncodedExampleResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/x-www-form-urlencoded")
	ctx.Status(200)

	if form, err := runtime.MarshalForm(response, nil); err != nil {
		return err
	} else {
		_, err := ctx.WriteString(form.Encode())
		return err
	}
}

type URLEncodedExample400Response = BadrequestResponse

func (response URLEncodedExample400Response) VisitURLEncodedExampleResponse(ctx *fiber.Ctx) error {
	ctx.Status(400)
	return nil
}

type URLEncodedExampledefaultResponse struct {
	StatusCode int
}

func (response URLEncodedExampledefaultResponse) VisitURLEncodedExampleResponse(ctx *fiber.Ctx) error {
	ctx.Status(response.StatusCode)
	return nil
}

type HeadersExampleRequestObject struct {
	Params HeadersExampleParams
	Body   *HeadersExampleJSONRequestBody
}

type HeadersExampleResponseObject interface {
	VisitHeadersExampleResponse(ctx *fiber.Ctx) error
}

type HeadersExample200ResponseHeaders struct {
	Header1 string
	Header2 int
}

type HeadersExample200JSONResponse struct {
	Body    Example
	Headers HeadersExample200ResponseHeaders
}

func (response HeadersExample200JSONResponse) VisitHeadersExampleResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("header1", fmt.Sprint(response.Headers.Header1))
	ctx.Response().Header.Set("header2", fmt.Sprint(response.Headers.Header2))
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(200)

	return ctx.JSON(&response.Body)
}

type HeadersExample400Response = BadrequestResponse

func (response HeadersExample400Response) VisitHeadersExampleResponse(ctx *fiber.Ctx) error {
	ctx.Status(400)
	return nil
}

type HeadersExampledefaultResponse struct {
	StatusCode int
}

func (response HeadersExampledefaultResponse) VisitHeadersExampleResponse(ctx *fiber.Ctx) error {
	ctx.Status(response.StatusCode)
	return nil
}

type UnionExampleRequestObject struct {
	Body *UnionExampleJSONRequestBody
}

type UnionExampleResponseObject interface {
	VisitUnionExampleResponse(ctx *fiber.Ctx) error
}

type UnionExample200ResponseHeaders struct {
	Header1 string
	Header2 int
}

type UnionExample200JSONResponse struct {
	Body struct {
		union json.RawMessage
	}
	Headers UnionExample200ResponseHeaders
}

func (response UnionExample200JSONResponse) VisitUnionExampleResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("header1", fmt.Sprint(response.Headers.Header1))
	ctx.Response().Header.Set("header2", fmt.Sprint(response.Headers.Header2))
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(200)

	return ctx.JSON(&response.Body.union)
}

type UnionExample400Response = BadrequestResponse

func (response UnionExample400Response) VisitUnionExampleResponse(ctx *fiber.Ctx) error {
	ctx.Status(400)
	return nil
}

type UnionExampledefaultResponse struct {
	StatusCode int
}

func (response UnionExampledefaultResponse) VisitUnionExampleResponse(ctx *fiber.Ctx) error {
	ctx.Status(response.StatusCode)
	return nil
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (POST /json)
	JSONExample(ctx context.Context, request JSONExampleRequestObject) (JSONExampleResponseObject, error)

	// (POST /multipart)
	MultipartExample(ctx context.Context, request MultipartExampleRequestObject) (MultipartExampleResponseObject, error)

	// (POST /multiple)
	MultipleRequestAndResponseTypes(ctx context.Context, request MultipleRequestAndResponseTypesRequestObject) (MultipleRequestAndResponseTypesResponseObject, error)

	// (GET /reserved-go-keyword-parameters/{type})
	ReservedGoKeywordParameters(ctx context.Context, request ReservedGoKeywordParametersRequestObject) (ReservedGoKeywordParametersResponseObject, error)

	// (POST /reusable-responses)
	ReusableResponses(ctx context.Context, request ReusableResponsesRequestObject) (ReusableResponsesResponseObject, error)

	// (POST /text)
	TextExample(ctx context.Context, request TextExampleRequestObject) (TextExampleResponseObject, error)

	// (POST /unknown)
	UnknownExample(ctx context.Context, request UnknownExampleRequestObject) (UnknownExampleResponseObject, error)

	// (POST /unspecified-content-type)
	UnspecifiedContentType(ctx context.Context, request UnspecifiedContentTypeRequestObject) (UnspecifiedContentTypeResponseObject, error)

	// (POST /urlencoded)
	URLEncodedExample(ctx context.Context, request URLEncodedExampleRequestObject) (URLEncodedExampleResponseObject, error)

	// (POST /with-headers)
	HeadersExample(ctx context.Context, request HeadersExampleRequestObject) (HeadersExampleResponseObject, error)

	// (POST /with-union)
	UnionExample(ctx context.Context, request UnionExampleRequestObject) (UnionExampleResponseObject, error)
}

type StrictHandlerFunc func(ctx *fiber.Ctx, args interface{}) (interface{}, error)

type StrictMiddlewareFunc func(f StrictHandlerFunc, operationID string) StrictHandlerFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// JSONExample operation middleware
func (sh *strictHandler) JSONExample(ctx *fiber.Ctx) error {
	var request JSONExampleRequestObject

	var body JSONExampleJSONRequestBody
	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	request.Body = &body

	handler := func(ctx *fiber.Ctx, request interface{}) (interface{}, error) {
		return sh.ssi.JSONExample(ctx.UserContext(), request.(JSONExampleRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "JSONExample")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else if validResponse, ok := response.(JSONExampleResponseObject); ok {
		if err := validResponse.VisitJSONExampleResponse(ctx); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	} else if response != nil {
		return fmt.Errorf("Unexpected response type: %T", response)
	}
	return nil
}

// MultipartExample operation middleware
func (sh *strictHandler) MultipartExample(ctx *fiber.Ctx) error {
	var request MultipartExampleRequestObject

	request.Body = multipart.NewReader(bytes.NewReader(ctx.Request().Body()), string(ctx.Request().Header.MultipartFormBoundary()))

	handler := func(ctx *fiber.Ctx, request interface{}) (interface{}, error) {
		return sh.ssi.MultipartExample(ctx.UserContext(), request.(MultipartExampleRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "MultipartExample")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else if validResponse, ok := response.(MultipartExampleResponseObject); ok {
		if err := validResponse.VisitMultipartExampleResponse(ctx); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	} else if response != nil {
		return fmt.Errorf("Unexpected response type: %T", response)
	}
	return nil
}

// MultipleRequestAndResponseTypes operation middleware
func (sh *strictHandler) MultipleRequestAndResponseTypes(ctx *fiber.Ctx) error {
	var request MultipleRequestAndResponseTypesRequestObject

	if strings.HasPrefix(string(ctx.Request().Header.ContentType()), "application/json") {
		var body MultipleRequestAndResponseTypesJSONRequestBody
		if err := ctx.BodyParser(&body); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		request.JSONBody = &body
	}
	if strings.HasPrefix(string(ctx.Request().Header.ContentType()), "application/x-www-form-urlencoded") {
		var body MultipleRequestAndResponseTypesFormdataRequestBody
		if err := ctx.BodyParser(&body); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		request.FormdataBody = &body
	}
	if strings.HasPrefix(string(ctx.Request().Header.ContentType()), "image/png") {
		request.Body = bytes.NewReader(ctx.Request().Body())
	}
	if strings.HasPrefix(string(ctx.Request().Header.ContentType()), "multipart/form-data") {
		request.MultipartBody = multipart.NewReader(bytes.NewReader(ctx.Request().Body()), string(ctx.Request().Header.MultipartFormBoundary()))
	}
	if strings.HasPrefix(string(ctx.Request().Header.ContentType()), "text/plain") {
		data := ctx.Request().Body()
		body := MultipleRequestAndResponseTypesTextRequestBody(data)
		request.TextBody = &body
	}

	handler := func(ctx *fiber.Ctx, request interface{}) (interface{}, error) {
		return sh.ssi.MultipleRequestAndResponseTypes(ctx.UserContext(), request.(MultipleRequestAndResponseTypesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "MultipleRequestAndResponseTypes")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else if validResponse, ok := response.(MultipleRequestAndResponseTypesResponseObject); ok {
		if err := validResponse.VisitMultipleRequestAndResponseTypesResponse(ctx); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	} else if response != nil {
		return fmt.Errorf("Unexpected response type: %T", response)
	}
	return nil
}

// ReservedGoKeywordParameters operation middleware
func (sh *strictHandler) ReservedGoKeywordParameters(ctx *fiber.Ctx, pType string) error {
	var request ReservedGoKeywordParametersRequestObject

	request.Type = pType

	handler := func(ctx *fiber.Ctx, request interface{}) (interface{}, error) {
		return sh.ssi.ReservedGoKeywordParameters(ctx.UserContext(), request.(ReservedGoKeywordParametersRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ReservedGoKeywordParameters")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else if validResponse, ok := response.(ReservedGoKeywordParametersResponseObject); ok {
		if err := validResponse.VisitReservedGoKeywordParametersResponse(ctx); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	} else if response != nil {
		return fmt.Errorf("Unexpected response type: %T", response)
	}
	return nil
}

// ReusableResponses operation middleware
func (sh *strictHandler) ReusableResponses(ctx *fiber.Ctx) error {
	var request ReusableResponsesRequestObject

	var body ReusableResponsesJSONRequestBody
	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	request.Body = &body

	handler := func(ctx *fiber.Ctx, request interface{}) (interface{}, error) {
		return sh.ssi.ReusableResponses(ctx.UserContext(), request.(ReusableResponsesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ReusableResponses")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else if validResponse, ok := response.(ReusableResponsesResponseObject); ok {
		if err := validResponse.VisitReusableResponsesResponse(ctx); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	} else if response != nil {
		return fmt.Errorf("Unexpected response type: %T", response)
	}
	return nil
}

// TextExample operation middleware
func (sh *strictHandler) TextExample(ctx *fiber.Ctx) error {
	var request TextExampleRequestObject

	data := ctx.Request().Body()
	body := TextExampleTextRequestBody(data)
	request.Body = &body

	handler := func(ctx *fiber.Ctx, request interface{}) (interface{}, error) {
		return sh.ssi.TextExample(ctx.UserContext(), request.(TextExampleRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "TextExample")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else if validResponse, ok := response.(TextExampleResponseObject); ok {
		if err := validResponse.VisitTextExampleResponse(ctx); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	} else if response != nil {
		return fmt.Errorf("Unexpected response type: %T", response)
	}
	return nil
}

// UnknownExample operation middleware
func (sh *strictHandler) UnknownExample(ctx *fiber.Ctx) error {
	var request UnknownExampleRequestObject

	request.Body = bytes.NewReader(ctx.Request().Body())

	handler := func(ctx *fiber.Ctx, request interface{}) (interface{}, error) {
		return sh.ssi.UnknownExample(ctx.UserContext(), request.(UnknownExampleRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UnknownExample")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else if validResponse, ok := response.(UnknownExampleResponseObject); ok {
		if err := validResponse.VisitUnknownExampleResponse(ctx); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	} else if response != nil {
		return fmt.Errorf("Unexpected response type: %T", response)
	}
	return nil
}

// UnspecifiedContentType operation middleware
func (sh *strictHandler) UnspecifiedContentType(ctx *fiber.Ctx) error {
	var request UnspecifiedContentTypeRequestObject

	request.ContentType = string(ctx.Request().Header.ContentType())

	request.Body = bytes.NewReader(ctx.Request().Body())

	handler := func(ctx *fiber.Ctx, request interface{}) (interface{}, error) {
		return sh.ssi.UnspecifiedContentType(ctx.UserContext(), request.(UnspecifiedContentTypeRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UnspecifiedContentType")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else if validResponse, ok := response.(UnspecifiedContentTypeResponseObject); ok {
		if err := validResponse.VisitUnspecifiedContentTypeResponse(ctx); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	} else if response != nil {
		return fmt.Errorf("Unexpected response type: %T", response)
	}
	return nil
}

// URLEncodedExample operation middleware
func (sh *strictHandler) URLEncodedExample(ctx *fiber.Ctx) error {
	var request URLEncodedExampleRequestObject

	var body URLEncodedExampleFormdataRequestBody
	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	request.Body = &body

	handler := func(ctx *fiber.Ctx, request interface{}) (interface{}, error) {
		return sh.ssi.URLEncodedExample(ctx.UserContext(), request.(URLEncodedExampleRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "URLEncodedExample")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else if validResponse, ok := response.(URLEncodedExampleResponseObject); ok {
		if err := validResponse.VisitURLEncodedExampleResponse(ctx); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	} else if response != nil {
		return fmt.Errorf("Unexpected response type: %T", response)
	}
	return nil
}

// HeadersExample operation middleware
func (sh *strictHandler) HeadersExample(ctx *fiber.Ctx, params HeadersExampleParams) error {
	var request HeadersExampleRequestObject

	request.Params = params

	var body HeadersExampleJSONRequestBody
	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	request.Body = &body

	handler := func(ctx *fiber.Ctx, request interface{}) (interface{}, error) {
		return sh.ssi.HeadersExample(ctx.UserContext(), request.(HeadersExampleRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "HeadersExample")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else if validResponse, ok := response.(HeadersExampleResponseObject); ok {
		if err := validResponse.VisitHeadersExampleResponse(ctx); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	} else if response != nil {
		return fmt.Errorf("Unexpected response type: %T", response)
	}
	return nil
}

// UnionExample operation middleware
func (sh *strictHandler) UnionExample(ctx *fiber.Ctx) error {
	var request UnionExampleRequestObject

	var body UnionExampleJSONRequestBody
	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	request.Body = &body

	handler := func(ctx *fiber.Ctx, request interface{}) (interface{}, error) {
		return sh.ssi.UnionExample(ctx.UserContext(), request.(UnionExampleRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UnionExample")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else if validResponse, ok := response.(UnionExampleResponseObject); ok {
		if err := validResponse.VisitUnionExampleResponse(ctx); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	} else if response != nil {
		return fmt.Errorf("Unexpected response type: %T", response)
	}
	return nil
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xYS2/jNhD+KwO2p4VkOdmcdOsGi227bVM4yanIgRZHNnclkiVHVgzD/72gKL9ixbW3",
	"fqDB3vSYF795cmYs06XRChU5ls6YRWe0cti8DLmw+HeFjvybQJdZaUhqxVL2gYtB+28eMYuV48MCF+ye",
	"PtOKUDWs3JhCZtyzJl+c558xl42x5P7pR4s5S9kPycqUJPx1CT7z0hTI5vN59MKCu88sYmPkAm1jbXi8",
	"2pRNU4MsZY6sVCPmhQSy604yqQhHaL02T9oa4QkWdqQzZqw2aEkGjCa8qLBbU/tFD79gRuEEUuV6G8tb",
	"rYhL5UDIPEeLiqAFD7wMB64yRltCAcMpeA0ZgUM7QcsiRpK8Yex+/Tu0BjsWsQlaFxRd9fq9vveXNqi4",
	"kSxl75tPETOcxs2Blg4yusvvv97f/QHSAa9Il5xkxotiCiW3bswLFCAVaW9ilZHrsUaTbRz/i2i5P7ZQ",
	"+qhpAuiDFtNTBEwTl2vhfN3vnyku5xG7Ccq6ZCyNStYSrBGT86rowPxRfVW6VoDWatueLCmrgqThltZ9",
	"tYn27wuSfSBfyktybctYcOInQv1Ymi4KfFsLOnPkfqxrB2NdA2kQyAuoJY1hwfgiuaUCDk6qUYGwMCrq",
	"9GSBbcn9SYlBe5YHL+PkuRRtSHmO67qOG+dVtkCVaYHi28TKko8wMWq0ye5lc2IpG07Jh+12cT1SEEWM",
	"8JkSU3CpdneOM5WT70gfLbFDulpsOqKIRzr+itNaWxEbbnmJhNYlM6997gWPsCOV/1xSQsYVDBEUL1EA",
	"zwktfNLQinRbKTto9X7SnwPJSlTTbpcv6V8z5iFpWjCLmFfA0oBKyGtpvdPJVhjtgO3pX+PzPzlggWYY",
	"9OINVd1lcFGiltBZzJ0viV2e68AvaBqsUVxmYNgdcVuj7zl6kPfk633/AZ/3avlHLH3nzu1DAavCx9cx",
	"a7n2ge0bK+keKE6kQJ2U5uZAyRcD1RnMZC5RxO0p4mDbayXhVqvMIm2OQP46oTTBUpi/5dAYISAQgdNQ",
	"I5SVIzDcOZDUVJFChpuSwK3i8biy7DZoeliV011efXcin767lEdv+leHs7w/cdxsjDKv5OPgt4+B5tD7",
	"4tFmpgMnvuPpvVA6+0tKvLZQ6U7hnwPBqqdnKCd+IlICLFJlFQqYSL5YAmzlZitg5dauWSiYsZqGFsud",
	"QwaiaKesaxbtWgA9veH1xCnXZueK00rJXWuqR/8b2hn6ZW+QWv1vllBa4V3e5MULn0R7mvD09mJgHrGw",
	"5QwFo7KFz2oikyZJ2I72XM1HI7Q9qRNupEfhnwAAAP//4wU1quoWAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
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

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
