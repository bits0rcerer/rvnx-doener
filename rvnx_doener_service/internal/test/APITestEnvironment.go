package test

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"rvnx_doener_service/internal/api"
	"rvnx_doener_service/internal/api/twitch"

	"github.com/gavv/httpexpect"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers"
	"github.com/getkin/kin-openapi/routers/gorillamux"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type APITestEnvironment struct {
	BaseTestEnvironment
	Expect *httpexpect.Expect
}

type bodyLogWriter struct {
    gin.ResponseWriter
    body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
    w.body.Write(b)
    return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
    w.body.WriteString(s)
    return w.ResponseWriter.WriteString(s)
}

func apiSpecsValidationMiddleware(t *testing.T, routeResolver routers.Router) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ignore special test requests
		if strings.HasPrefix(c.Request.URL.Path, "/api/_test") {
			return
		}

		// get request body
		var buf bytes.Buffer
		tee := io.TeeReader(c.Request.Body, &buf)
		bodyBytes, _ := ioutil.ReadAll(tee)

		// restore body for api spec check
		c.Request.Body = ioutil.NopCloser(bytes.NewReader(bodyBytes))

		// Inject custom writer
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
        c.Writer = blw

		// pretend this is real
		c.Request.URL.Scheme = "https"
		c.Request.URL.Host = "doener.rvnx.community"

		// Find route
		route, pathParams, err := routeResolver.FindRoute(c.Request)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		
		// Validate request
		requestValidationInput := &openapi3filter.RequestValidationInput{
			Request:     c.Request,
			PathParams:  pathParams,
			Route:       route,
			QueryParams: c.Request.URL.Query(),
		}
		if !assert.NoError(t, openapi3filter.ValidateRequest(c, requestValidationInput)) {
			t.FailNow()
		}

		// restore body for gin
		c.Request.Body = ioutil.NopCloser(bytes.NewReader(bodyBytes))
		
		c.Next()

		// Validate response
		responseValidationInput := &openapi3filter.ResponseValidationInput{
			RequestValidationInput: requestValidationInput,
			Status:                 c.Writer.Status(),
			Header:                 c.Writer.Header().Clone(),
		}
		responseValidationInput.SetBodyBytes(blw.body.Bytes())
		if !assert.NoError(t, openapi3filter.ValidateResponse(c, responseValidationInput)) {
			t.FailNow()
		}
	}
}

func NewAPITestEnvironment(t *testing.T, specsFile []byte, base BaseTestEnvironment) *APITestEnvironment {
	ctx := context.Background()
	loader := &openapi3.Loader{Context: ctx, IsExternalRefsAllowed: true}
	specs, _ := loader.LoadFromData(specsFile)
	if !assert.NoError(t, specs.Validate(ctx)) {
		t.FailNow()
	}
	routeResolver, _ := gorillamux.NewRouter(specs)

	engine := api.BuildEngine()
	engine.Use(apiSpecsValidationMiddleware(t, routeResolver))
	apiRouter := engine.Group("/api")
	api.RouteAPI(apiRouter, base.Services)

	apiRouter.POST("/_test/setSession", func(c *gin.Context) {
		s := sessions.Default(c)

		var payload gin.H
		err := c.BindJSON(&payload)
		if err != nil {
			assert.FailNow(t, "unable to set session values")
		}

		for k, v := range payload {
			if k == twitch.UserIDSessionKey {
				id, _ := strconv.ParseInt(v.(string), 10, 64)
				s.Set(k, id)
			} else {
				s.Set(k, v)
			}
		}

		err = s.Save()
		if err != nil {
			assert.FailNow(t, "unable to set session values")
		}

		c.Status(http.StatusOK)
	})

	expect := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(engine),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})

	return &APITestEnvironment{
		BaseTestEnvironment: base,
		Expect:              expect,
	}
}

func DoAPITest(
	t *testing.T,
	name string,
	specsFile []byte,
	testCase func(t *testing.T, env *APITestEnvironment),
) {
	t.Helper()

	base := NewBaseTestEnvironment(t)
	env := NewAPITestEnvironment(t, specsFile, *base)

	t.Run(name, func(t *testing.T) {
		t.Helper()
		defer env.Cleanup()
		testCase(t, env)
	})
}
