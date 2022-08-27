package test

import (
	"net/http"
	"rvnx_doener_service/internal/api"
	"rvnx_doener_service/internal/api/twitch"
	"strconv"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type APITestEnvironment struct {
	BaseTestEnvironment
	Expect *httpexpect.Expect
}

func NewAPITestEnvironment(t *testing.T, base BaseTestEnvironment) *APITestEnvironment {
	engine := api.BuildEngine()
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
	testCase func(t *testing.T, env *APITestEnvironment)) {
	t.Helper()

	base := NewBaseTestEnvironment(t)
	env := NewAPITestEnvironment(t, *base)

	t.Run(name, func(t *testing.T) {
		t.Helper()
		defer env.Cleanup()
		testCase(t, env)
	})
}
