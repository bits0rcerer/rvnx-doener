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

			assert.Equal(t, float64(13), kebabShop.Lat)
			assert.Equal(t, float64(37), kebabShop.Lng)
			assert.Equal(t, "Best Test Kebab", kebabShop.Name)
			assert.Nil(t, kebabShop.OsmID)

			kebabShop2, err := client.KebabShop.Query().Unique(false).Where(kebabshop.ID(kebabShop.ID)).First(context.Background())
			if !assert.NoError(t, err) {
				t.FailNow()
			}

			assert.Equal(t, kebabShop.ID, kebabShop2.ID)
			assert.Equal(t, float64(13), kebabShop2.Lat)
			assert.Equal(t, float64(37), kebabShop2.Lng)
			assert.Equal(t, "Best Test Kebab", kebabShop2.Name)
			assert.Nil(t, kebabShop2.OsmID)

			log.WaitUntil(event.EventTypeKebabShopCreated, time.Second, func(t *testing.T, event ent.Event) {
				assert.Equal(t, kebabShop.ID, event.Info["id"])
				assert.Equal(t, kebabShop.Name, event.Info["name"])
				assert.Equal(t, strconv.FormatFloat(kebabShop.Lat, 'E', -1, 64), event.Info["lat"])
				assert.Equal(t, strconv.FormatFloat(kebabShop.Lng, 'E', -1, 64), event.Info["long"])
			})
		})

	test.DoTest(t, "Import or update new kebab shop from osm and log an event",
		func(t *testing.T, client *ent.Client, services *services.ServiceEnvironment, log *log2.TestEventLogger) {
			osmID := 42

			kebabShop, err := services.KebabShopService.UpdateOrInsertKebabShop(&ent.KebabShop{
				OsmID: &osmID,
				Name:  "Best Test Kebab",
				Lat:   13,
				Lng:   37,
			})
			if err != nil {
				return
			}
			if !assert.NoError(t, err) {
				t.FailNow()
			}

			log.WaitUntil(event.EventTypeKebabShopImported, time.Second, func(t *testing.T, event ent.Event) {
				assert.Equal(t, kebabShop.ID, event.Info["id"])
				assert.Equal(t, kebabShop.Name, event.Info["name"])
				assert.Equal(t, strconv.FormatFloat(kebabShop.Lat, 'E', -1, 64), event.Info["lat"])
				assert.Equal(t, strconv.FormatFloat(kebabShop.Lng, 'E', -1, 64), event.Info["long"])
			})

			// new name on OSM
			kebabShop, err = services.KebabShopService.UpdateOrInsertKebabShop(&ent.KebabShop{
				OsmID: &osmID,
				Name:  "Very Best Test Kebab",
				Lat:   13,
				Lng:   37,
			})
			if err != nil {
				return
			}
			if !assert.NoError(t, err) {
				t.FailNow()
			}

			log.WaitUntil(event.EventTypeKebabShopUpdatedFromOsm, time.Second, func(t *testing.T, event ent.Event) {
				assert.Equal(t, kebabShop.ID, event.Info["id"])
				assert.Equal(t, kebabShop.Name, event.Info["name"])
				assert.Equal(t, strconv.FormatFloat(kebabShop.Lat, 'E', -1, 64), event.Info["lat"])
				assert.Equal(t, strconv.FormatFloat(kebabShop.Lng, 'E', -1, 64), event.Info["long"])
			})

			// new geo cords on OSM
			kebabShop, err = services.KebabShopService.UpdateOrInsertKebabShop(&ent.KebabShop{
				OsmID: &osmID,
				Name:  "Very Best Test Kebab",
				Lat:   42,
				Lng:   24,
			})
			if err != nil {
				return
			}
			if !assert.NoError(t, err) {
				t.FailNow()
			}

			log.WaitUntil(event.EventTypeKebabShopUpdatedFromOsm, time.Second, func(t *testing.T, event ent.Event) {
				assert.Equal(t, kebabShop.ID, event.Info["id"])
				assert.Equal(t, kebabShop.Name, event.Info["name"])
				assert.Equal(t, strconv.FormatFloat(kebabShop.Lat, 'E', -1, 64), event.Info["lat"])
				assert.Equal(t, strconv.FormatFloat(kebabShop.Lng, 'E', -1, 64), event.Info["long"])
			})
		})
}
