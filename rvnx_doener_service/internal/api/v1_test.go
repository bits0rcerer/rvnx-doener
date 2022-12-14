package api_test

import (
	"context"
	"net/http"
	"strconv"
	"testing"
	"time"

	"rvnx_doener_service/doc/apispecs"
	"rvnx_doener_service/ent"
	"rvnx_doener_service/ent/event"
	"rvnx_doener_service/ent/kebabshop"
	"rvnx_doener_service/ent/scorerating"
	"rvnx_doener_service/ent/shopprice"
	"rvnx_doener_service/ent/twitchuser"
	"rvnx_doener_service/internal/api/session"
	"rvnx_doener_service/internal/api/twitch"
	"rvnx_doener_service/internal/test"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	test.CommonTestMain(m)
}

func TestV1KebabShops_Box(t *testing.T) {
	test.DoAPITest(t, "Request kebab shops within a box", apispecs.API_V1_SpecsFile,
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

	test.DoAPITest(t, "Request only kebab shops within a box that are rated by users", apispecs.API_V1_SpecsFile,
		func(t *testing.T, env *test.APITestEnvironment) {
			env.LoadOSMTestData(t)

			u := env.CreateUser(t, "u1")
			shop, err := env.Client.KebabShop.Query().First(context.Background())
			require.NoError(t, err)

			nf, err := env.Services.KebabShopService.AddUserScore(shop.ID, u.ID, false, 5.0)
			require.NoError(t, err)
			require.False(t, nf)

			resp := env.Expect.GET("/api/v1/kebabshops/box").
				WithQuery("ltm", -180).
				WithQuery("ltx", 180).
				WithQuery("lnm", -180).
				WithQuery("lnx", 180).
				WithQuery("rvnx_only", true).
				Expect().Status(http.StatusOK).JSON()

			cords := resp.Path("$.cords").Array()
			cords.Length().Equal(1)
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

	test.DoAPITest(t, "Request only kebab shops within a box that are submitted by users", apispecs.API_V1_SpecsFile,
		func(t *testing.T, env *test.APITestEnvironment) {
			env.LoadOSMTestData(t)

			u := env.CreateUser(t, "u1")
			shop, err := env.Client.KebabShop.Create().SetName("Mega D??ner").SetLat(13).SetLng(37).AddSubmittedBy(u).Save(context.Background())
			require.NoError(t, err)

			resp := env.Expect.GET("/api/v1/kebabshops/box").
				WithQuery("ltm", -180).
				WithQuery("ltx", 180).
				WithQuery("lnm", -180).
				WithQuery("lnx", 180).
				WithQuery("rvnx_only", true).
				Expect().Status(http.StatusOK).JSON()

			cords := resp.Path("$.cords").Array()
			cords.Length().Equal(1)
			cords.First().Path("$.id").Equal(strconv.Itoa(int(shop.ID)))
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

	test.DoAPITest(t, "Invalid request", nil,
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
	test.DoAPITest(t, "Request kebab shop by its id", apispecs.API_V1_SpecsFile,
		func(t *testing.T, env *test.APITestEnvironment) {
			s1 := env.CreateKebabShop(t, "Shop1", 13, 37)
			_ = env.CreateKebabShop(t, "Shop2", -4, 20)

			resp := env.Expect.GET("/api/v1/kebabshops/{shop_id}", s1.ID).
				Expect().Status(http.StatusOK).JSON()

			shop := resp.Path("$.shop").Object()
			shop.Path("$.id").String().Equal(strconv.Itoa(int(s1.ID)))
			shop.Path("$.name").String().Equal(s1.Name)
			shop.Path("$.lat").Number().Equal(s1.Lat)
			shop.Path("$.lng").Number().Equal(s1.Lng)
			rating := shop.Path("$.rating").Object()
			rating.Path("$.score")
			rating.Path("$.prices")
			rating.Path("$.reviews")
		})

	test.DoAPITest(t, "ID not present", apispecs.API_V1_SpecsFile,
		func(t *testing.T, env *test.APITestEnvironment) {
			env.Expect.GET("/api/v1/kebabshops/{shop_id}", 1337).
				Expect().Status(http.StatusNotFound)
		})

	test.DoAPITest(t, "ID invalid", nil,
		func(t *testing.T, env *test.APITestEnvironment) {
			env.Expect.GET("/api/v1/kebabshops/{shop_id}", "an invalid id").
				Expect().Status(http.StatusBadRequest)
		})
}

func TestV1KebabShops_Rating(t *testing.T) {
	test.DoAPITest(t, "a user rates a kebab shop and it is logged as an event", apispecs.API_V1_SpecsFile,
		func(t *testing.T, env *test.APITestEnvironment) {
			shop := env.CreateKebabShop(t, "Test Shop", 13, 37)
			user := env.CreateUser(t, "Test User")

			// set session cookie
			cookie := env.Expect.POST("/api/_test/setSession").WithJSON(
				gin.H{
					twitch.UserDisplaySessionKey:   user.DisplayName,
					twitch.UserIDSessionKey:        strconv.Itoa(int(user.ID)),
					twitch.UserActivatedSessionKey: true,
				}).Expect().Cookie(session.SessionCookieName)

			env.Expect.POST("/api/v1/kebabshops/{shop_id}/rate", shop.ID).
				WithCookie(cookie.Name().Raw(), cookie.Value().Raw()).
				WithJSON(gin.H{
					"rating": gin.H{
						"anonymous": false,
						"prices": gin.H{
							"normalKebab": gin.H{
								"price":    4.50,
								"currency": "EUR",
							},
							"vegiKebab": gin.H{
								"price":    5.50,
								"currency": "EUR",
							},
						},
						"opinion":   "Schmeckt ziemlich gut",
						"userScore": 3,
					},
				}).Expect().Status(http.StatusOK)

			rating := shop.QueryUserScores().Order(ent.Desc(scorerating.FieldID)).FirstX(context.Background())
			require.NotNil(t, rating)
			assert.Equal(t, rating.Anonymous, false)
			assert.Equal(t, rating.Score, 3.0)
			assert.Equal(t, rating.QueryAuthor().FirstX(context.Background()).ID, user.ID)

			opinion := shop.QueryUserOpinions().Order(ent.Desc(scorerating.FieldID)).FirstX(context.Background())
			assert.Equal(t, opinion.Anonymous, false)
			assert.Equal(t, opinion.Opinion, "Schmeckt ziemlich gut")
			assert.Equal(t, opinion.QueryAuthor().FirstX(context.Background()).ID, user.ID)

			var latestPriceIDs []struct {
				ID        uint64              `json:"max"`
				PriceType shopprice.PriceType `json:"price_type"`
			}
			shop.QueryUserPrices().
				GroupBy(shopprice.FieldPriceType).
				Aggregate(ent.Max(shopprice.FieldID)).
				ScanX(context.Background(), &latestPriceIDs)

			var id []uint64
			for _, d := range latestPriceIDs {
				id = append(id, d.ID)
			}

			prices := shop.QueryUserPrices().Where(shopprice.IDIn(id...)).AllX(context.Background())
			for _, p := range prices {
				if p.PriceType == shopprice.PriceTypeNormalKebab {
					assert.Equal(t, p.Currency, shopprice.CurrencyEuro)
					require.NotNil(t, p.Price.Int)
					assert.Equal(t, p.Price.Int.Int64(), int64(45))
					assert.Equal(t, p.Price.Exp, int32(-1))
				}
				if p.PriceType == shopprice.PriceTypeVegetarianKebab {
					assert.Equal(t, p.Currency, shopprice.CurrencyEuro)
					require.NotNil(t, p.Price.Int)
					assert.Equal(t, p.Price.Int.Int64(), int64(55))
					assert.Equal(t, p.Price.Exp, int32(-1))
				}
			}

			env.Log.WaitUntil(event.EventTypeUserSubmittedARating, time.Second, func(t *testing.T, event ent.Event) {
				if !(assert.NotNil(t, event.Info) &&
					assert.Equal(t, shop.ID, event.Info["shopID"]) &&
					assert.Equal(t, user.ID, event.Info["userID"])) {
					return
				}

				if event.Info["userScore"] != nil {
					assert.Equal(t, 3, event.Info["userScore"])
				} else if event.Info["opinion"] != nil {
					assert.Equal(t, "Schmeckt ziemlich gut", event.Info["opinion"])
				} else if event.Info["prices"] != nil {
					require.IsType(t, map[string]interface{}{}, event.Info["prices"])

					prices := event.Info["prices"].(map[string]interface{})
					assert.Equal(t, map[string]interface{}{
						"normalKebab": map[string]interface{}{
							"price":    4.50,
							"currency": "EUR",
						},
						"vegiKebab": map[string]interface{}{
							"price":    5.50,
							"currency": "EUR",
						},
					}, prices)
				} else {
					assert.FailNow(t, "rating payload missing")
				}
			})
		})

	test.DoAPITest(t, "a user tries to rate a kebab shop but sends invalid data", apispecs.API_V1_SpecsFile,
		func(t *testing.T, env *test.APITestEnvironment) {
			// TODO: add tests
			t.Skip("TODO: add tests")
		})

	test.DoAPITest(t, "Request kebab shop rating", apispecs.API_V1_SpecsFile,
		func(t *testing.T, env *test.APITestEnvironment) {
			user := env.CreateUser(t, "User 1")
			s1 := env.CreateKebabShop(t, "Shop1", 13, 37)
			_ = env.CreateKebabShop(t, "Shop2", -4, 20)

			// set session cookie
			cookie := env.Expect.POST("/api/_test/setSession").WithJSON(
				gin.H{
					twitch.UserDisplaySessionKey:   user.DisplayName,
					twitch.UserIDSessionKey:        strconv.Itoa(int(user.ID)),
					twitch.UserActivatedSessionKey: true,
				}).Expect().Cookie(session.SessionCookieName)

			env.Expect.POST("/api/v1/kebabshops/{shop_id}/rate", s1.ID).
				WithCookie(cookie.Name().Raw(), cookie.Value().Raw()).
				WithJSON(gin.H{
					"rating": gin.H{
						"anonymous": false,
						"prices": gin.H{
							"normalKebab": gin.H{
								"price":    4.50,
								"currency": "EUR",
							},
							"vegiKebab": gin.H{
								"price":    5.50,
								"currency": "EUR",
							},
						},
						"opinion":   "Schmeckt ziemlich gut",
						"userScore": 3,
					},
				}).Expect().Status(http.StatusOK)

			resp := env.Expect.GET("/api/v1/kebabshops/{shop_id}", s1.ID).
				Expect().Status(http.StatusOK).JSON()

			shop := resp.Path("$.shop").Object()
			shop.Path("$.id").String().Equal(strconv.Itoa(int(s1.ID)))
			rating := shop.Path("$.rating").Object()
			rating.Path("$.score").Number().Equal(3)
			rating.Path("$.prices").Array().Contains(
				gin.H{
					"price":       450e-2,
					"currency":    "EUR",
					"type":        "normalKebab",
					"order_index": 0,
				},
				gin.H{
					"price":       550e-2,
					"currency":    "EUR",
					"type":        "vegiKebab",
					"order_index": 10,
				})
			rating.Path("$.reviews").Array()
		})
}

func TestV1KebabShops_AddShop(t *testing.T) {
	test.DoAPITest(t, "a user add a kebab shop and it is logged as an event", apispecs.API_V1_SpecsFile,
		func(t *testing.T, env *test.APITestEnvironment) {
			user := env.CreateUser(t, "Test User")

			// set session cookie
			cookie := env.Expect.POST("/api/_test/setSession").WithJSON(
				gin.H{
					twitch.UserDisplaySessionKey:   user.DisplayName,
					twitch.UserIDSessionKey:        strconv.Itoa(int(user.ID)),
					twitch.UserActivatedSessionKey: true,
				}).Expect().Cookie(session.SessionCookieName)

			env.Expect.POST("/api/v1/kebabshops").
				WithCookie(session.SessionCookieName, cookie.Raw().Value).
				WithJSON(gin.H{
					"name":      "Mega D??ner",
					"lat":       13.0,
					"lng":       37.0,
					"anonymous": false,
				}).Expect().Status(http.StatusOK)

			env.Log.WaitUntil(event.EventTypeUserSubmittedAShop, time.Second, func(t *testing.T, event ent.Event) {
				_ = assert.NotNil(t, event.Info) && assert.Equal(t, user.ID, event.Info["user_id"])
				// TODO: inspect event further
			})

			env.Client.KebabShop.Query().Where(
				kebabshop.HasSubmittedByWith(twitchuser.ID(user.ID)),
				kebabshop.Name("Mega D??ner"),
				kebabshop.Lng(37.0),
				kebabshop.Lat(13.0),
				kebabshop.PostedAnonymously(false),
			)

			env.Expect.POST("/api/v1/kebabshops").
				WithCookie(session.SessionCookieName, cookie.Raw().Value).
				WithJSON(gin.H{
					"name":      "Mega D??ner 2",
					"lat":       24.0,
					"lng":       42.0,
					"anonymous": true,
				}).Expect().Status(http.StatusOK)

			env.Log.WaitUntil(event.EventTypeUserSubmittedAShop, time.Second, func(t *testing.T, event ent.Event) {
				_ = assert.NotNil(t, event.Info) && assert.Equal(t, user.ID, event.Info["user_id"])
				// TODO: inspect event further
			})

			env.Client.KebabShop.Query().Where(
				kebabshop.HasSubmittedByWith(twitchuser.ID(user.ID)),
				kebabshop.Name("Mega D??ner 2"),
				kebabshop.Lng(42.0),
				kebabshop.Lat(24.0),
				kebabshop.PostedAnonymously(true),
			)
		})
}
