package services_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"rvnx_doener_service/ent"
	"rvnx_doener_service/ent/event"
	"rvnx_doener_service/ent/kebabshop"
	log2 "rvnx_doener_service/internal/log"
	"rvnx_doener_service/internal/services"
	"rvnx_doener_service/internal/test"
	"strconv"
	"testing"
	"time"
)

func TestKebabShopService_CreateKebabShop(t *testing.T) {
	test.DoTest(t, "Create a kebab shop and log an event",
		func(t *testing.T, client *ent.Client, services *services.ServiceEnvironment, log *log2.TestEventLogger) {
			kebabShop, err := services.KebabShopService.CreateKebabShop("Best Test Kebab", 13, 37)
			if !assert.NoError(t, err) {
				t.FailNow()
			}

			assert.Equal(t, float64(13), kebabShop.Point.P.X)
			assert.Equal(t, float64(37), kebabShop.Point.P.Y)
			assert.Equal(t, "Best Test Kebab", kebabShop.Name)
			assert.Nil(t, kebabShop.OsmID)

			kebabShop2, err := client.KebabShop.Query().Unique(false).Where(kebabshop.ID(kebabShop.ID)).First(context.Background())
			if !assert.NoError(t, err) {
				t.FailNow()
			}

			assert.Equal(t, kebabShop.ID, kebabShop2.ID)
			assert.Equal(t, float64(13), kebabShop2.Point.P.X)
			assert.Equal(t, float64(37), kebabShop2.Point.P.Y)
			assert.Equal(t, "Best Test Kebab", kebabShop2.Name)
			assert.Nil(t, kebabShop2.OsmID)

			log.WaitUntil(event.EventTypeKebabShopCreated, time.Second, func(t *testing.T, event ent.Event) {
				assert.Equal(t, kebabShop.ID, event.Info["id"])
				assert.Equal(t, kebabShop.Name, event.Info["name"])
				assert.Equal(t, strconv.FormatFloat(kebabShop.Point.P.X, 'E', -1, 64), event.Info["lat"])
				assert.Equal(t, strconv.FormatFloat(kebabShop.Point.P.Y, 'E', -1, 64), event.Info["long"])
			})
		})
}
