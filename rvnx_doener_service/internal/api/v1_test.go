package api_test

import (
	"net/http"
	"rvnx_doener_service/internal/test"
	"strconv"
	"testing"
)

func TestV1KebabShops_Box(t *testing.T) {
	test.DoAPITest(t, "Request kebab shops within a box",
		func(t *testing.T, env *test.APITestEnvironment) {
			env.LoadOSMTestData(t)

			resp := env.Expect.GET("/api/v1/kebabshops/box").
				WithQuery("ltm", 50).
				WithQuery("ltx", 60).
				WithQuery("lnm", 10).
				WithQuery("lnx", 20).
				Expect().Status(http.StatusOK).JSON()

			cords := resp.Path("$.cords").Array()
			cords.NotEmpty()
			for _, v := range cords.Iter() {
				v.Schema(`{
					"type": "object",
					"properties": {
					   "id": {
						   "type": "string"
					   },
					   "lat": {
						   "type": "number"
					   },
					   "lng": {
						   "type": "number"
					   }
				   },
				   "require": ["id", "lat", "lng"]
				 }`)
			}
		})

	test.DoAPITest(t, "Invalid request",
		func(t *testing.T, env *test.APITestEnvironment) {
			env.LoadOSMTestData(t)

			env.Expect.GET("/api/v1/kebabshops/box").
				Expect().Status(http.StatusBadRequest)

			env.Expect.GET("/api/v1/kebabshops/box").
				WithQuery("ltx", 60).
				WithQuery("lnm", 10).
				WithQuery("lnx", 20).
				Expect().Status(http.StatusBadRequest)

			env.Expect.GET("/api/v1/kebabshops/box").
				WithQuery("ltm", 50).
				WithQuery("lnm", 10).
				WithQuery("lnx", 20).
				Expect().Status(http.StatusBadRequest)

			env.Expect.GET("/api/v1/kebabshops/box").
				WithQuery("ltm", 50).
				WithQuery("ltx", 60).
				WithQuery("lnx", 20).
				Expect().Status(http.StatusBadRequest)

			env.Expect.GET("/api/v1/kebabshops/box").
				WithQuery("ltm", 50).
				WithQuery("ltx", 60).
				WithQuery("lnm", 10).
				Expect().Status(http.StatusBadRequest)

			env.Expect.GET("/api/v1/kebabshops/box").
				WithQuery("ltm", "not a number").
				WithQuery("ltx", "not a number").
				WithQuery("lnm", "not a number").
				WithQuery("lnx", "not a number").
				Expect().Status(http.StatusBadRequest)
		})
}

func TestV1KebabShops_Cluster(t *testing.T) {
	// TODO: add tests
	t.Skip("TODO: add tests")
}

func TestV1KebabShops_Auto(t *testing.T) {
	// TODO: add tests
	t.Skip("TODO: add tests")
}

func TestV1KebabShops_ShopByID(t *testing.T) {
	test.DoAPITest(t, "Request kebab shop by its id",
		func(t *testing.T, env *test.APITestEnvironment) {
			s1 := env.CreateKebabShop(t, "Shop1", 13, 37)
			_ = env.CreateKebabShop(t, "Shop2", -4, 20)

			resp := env.Expect.GET("/api/v1/kebabshops/{shop_id}", s1.ID).
				Expect().Status(http.StatusOK).JSON()

			shop := resp.Path("$.shop").Object()
			shop.Schema(`{
					"type": "object",
					"properties": {
					   "id": {
						   "type": "string"
					   },
					   "name": {
						   "type": "string"
					   },
					   "lat": {
						   "type": "number"
					   },
					   "lng": {
						   "type": "number"
					   }
				   },
				   "require": ["id", "name", "lat", "lng"]
				 }`)
			shop.Path("$.id").String().Equal(strconv.Itoa(s1.ID))
			shop.Path("$.name").String().Equal(s1.Name)
			shop.Path("$.lat").Number().Equal(s1.Lat)
			shop.Path("$.lng").Number().Equal(s1.Lng)
		})

	test.DoAPITest(t, "ID not present",
		func(t *testing.T, env *test.APITestEnvironment) {
			env.Expect.GET("/api/v1/kebabshops/{shop_id}", 1337).
				Expect().Status(http.StatusNotFound)
		})

	test.DoAPITest(t, "ID invalid",
		func(t *testing.T, env *test.APITestEnvironment) {
			env.Expect.GET("/api/v1/kebabshops/{shop_id}", "an invalid id").
				Expect().Status(http.StatusBadRequest)
		})
}
