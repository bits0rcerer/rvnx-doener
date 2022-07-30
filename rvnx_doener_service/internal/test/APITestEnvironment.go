package test

import (
	"github.com/gavv/httpexpect"
	"net/http"
	"rvnx_doener_service/internal/api"
	"testing"
)

type APITestEnvironment struct {
	BaseTestEnvironment
	Expect *httpexpect.Expect
}

func NewAPITestEnvironment(t *testing.T, base BaseTestEnvironment) *APITestEnvironment {
	engine := api.BuildEngine()
	api.RouteAPI(engine.Group("/api"), base.Services)

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
		t.Parallel()

		defer env.Cleanup()
		testCase(t, env)
	})
}
