package services_test

import (
	"context"
	"rvnx_doener_service/ent"
	"rvnx_doener_service/ent/event"
	"rvnx_doener_service/ent/kebabshop"
	"rvnx_doener_service/internal/test"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	test.CommonTestMain(m)
}

func TestKebabShopService_CreateKebabShop(t *testing.T) {
	test.DoTest(t, "Create a kebab shop and log an event",
		func(t *testing.T, env *test.BaseTestEnvironment) {
			kebabShop := env.CreateKebabShop(t, "Best Test Kebab", 13, 37)

			kebabShop2, err := env.Client.KebabShop.Query().Unique(false).Where(kebabshop.ID(kebabShop.ID)).First(context.Background())
			require.NoError(t, err)
			assert.Equal(t, kebabShop.ID, kebabShop2.ID)
			assert.Equal(t, float64(13), kebabShop2.Lat)
			assert.Equal(t, float64(37), kebabShop2.Lng)
			assert.Equal(t, "Best Test Kebab", kebabShop2.Name)
			assert.Nil(t, kebabShop2.OsmID)
		})
}

func TestKebabShopService_UpdateOrInsertKebabShop(t *testing.T) {
	test.DoTest(t, "Import or update new kebab shop from osm and log an event",
		func(t *testing.T, env *test.BaseTestEnvironment) {
			osmID := 42

			kebabShop, err := env.Services.KebabShopService.UpdateOrInsertKebabShop(&ent.KebabShop{
				OsmID: &osmID,
				Name:  "Best Test Kebab",
				Lat:   13,
				Lng:   37,
			})
			if err != nil {
				return
			}
			require.NoError(t, err)

			env.Log.WaitUntil(event.EventTypeKebabShopImported, time.Second, func(t *testing.T, event ent.Event) {
				assert.Equal(t, kebabShop.ID, event.Info["id"])
				assert.Equal(t, kebabShop.Name, event.Info["name"])
				assert.Equal(t, strconv.FormatFloat(kebabShop.Lat, 'E', -1, 64), event.Info["lat"])
				assert.Equal(t, strconv.FormatFloat(kebabShop.Lng, 'E', -1, 64), event.Info["long"])
			})

			// new name on OSM
			kebabShop, err = env.Services.KebabShopService.UpdateOrInsertKebabShop(&ent.KebabShop{
				OsmID: &osmID,
				Name:  "Very Best Test Kebab",
				Lat:   13,
				Lng:   37,
			})
			if err != nil {
				return
			}
			require.NoError(t, err)

			env.Log.WaitUntil(event.EventTypeKebabShopUpdatedFromOsm, time.Second, func(t *testing.T, event ent.Event) {
				assert.Equal(t, kebabShop.ID, event.Info["id"])
				assert.Equal(t, kebabShop.Name, event.Info["name"])
				assert.Equal(t, strconv.FormatFloat(kebabShop.Lat, 'E', -1, 64), event.Info["lat"])
				assert.Equal(t, strconv.FormatFloat(kebabShop.Lng, 'E', -1, 64), event.Info["long"])
			})

			// new geo cords on OSM
			kebabShop, err = env.Services.KebabShopService.UpdateOrInsertKebabShop(&ent.KebabShop{
				OsmID: &osmID,
				Name:  "Very Best Test Kebab",
				Lat:   42,
				Lng:   24,
			})
			if err != nil {
				return
			}
			require.NoError(t, err)

			env.Log.WaitUntil(event.EventTypeKebabShopUpdatedFromOsm, time.Second, func(t *testing.T, event ent.Event) {
				assert.Equal(t, kebabShop.ID, event.Info["id"])
				assert.Equal(t, kebabShop.Name, event.Info["name"])
				assert.Equal(t, strconv.FormatFloat(kebabShop.Lat, 'E', -1, 64), event.Info["lat"])
				assert.Equal(t, strconv.FormatFloat(kebabShop.Lng, 'E', -1, 64), event.Info["long"])
			})
		})
}

func TestKebabShopService_Within(t *testing.T) {
	test.DoTest(t, "Request kebab shops within a specific area",
		func(t *testing.T, env *test.BaseTestEnvironment) {
			kebabShop := env.CreateKebabShop(t, "Best Test Kebab", 13, 37)
			kebabShop2 := env.CreateKebabShop(t, "Best Test Kebab2", -4, -20)
			_ = env.CreateKebabShop(t, "Best Test Kebab3", -8, 20)
			_ = env.CreateKebabShop(t, "Best Test Kebab4", -4, 80)

			shopsWithin := []*ent.KebabShop{kebabShop, kebabShop2}

			shops, err := env.Services.KebabShopService.Within(-4, 13, -20, 37, false)
			require.NoError(t, err)
			assert.Len(t, shops, 2)

			for _, shopToFind := range shopsWithin {
				shopFound := false
				for _, shop := range shops {
					if shop.ID == shopToFind.ID {
						shopFound = true
						break
					}
				}
				assert.True(t, shopFound)
			}
		})
}

func TestKebabShopService_KebabShop(t *testing.T) {
	test.DoTest(t, "Request a kebab shop by its id",
		func(t *testing.T, env *test.BaseTestEnvironment) {
			kebabShop := env.CreateKebabShop(t, "Best Test Kebab", 13, 37)
			kebabShop2 := env.CreateKebabShop(t, "Best Test Kebab2", -4, -20)

			shop, exists, err := env.Services.KebabShopService.KebabShop(kebabShop.ID - 1)
			assert.Nil(t, shop)
			assert.False(t, exists)
			assert.Nil(t, err)

			shop, exists, err = env.Services.KebabShopService.KebabShop(kebabShop.ID)
			assert.NotNil(t, shop)
			assert.True(t, exists)
			assert.Nil(t, err)

			test.AssertKebabShop(t, kebabShop, shop)

			shop, exists, err = env.Services.KebabShopService.KebabShop(kebabShop2.ID)
			assert.NotNil(t, shop)
			assert.True(t, exists)
			assert.Nil(t, err)

			test.AssertKebabShop(t, kebabShop2, shop)
		})
}
