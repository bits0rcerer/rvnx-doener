package test

import (
	"github.com/stretchr/testify/assert"
	"log"
	"rvnx_doener_service/ent"
	"rvnx_doener_service/ent/event"
	log2 "rvnx_doener_service/internal/log"
	"rvnx_doener_service/internal/services"
	"strconv"
	"testing"
	"time"
)

type BaseTestEnvironment struct {
	Client    *ent.Client
	Services  *services.ServiceEnvironment
	Log       *log2.TestEventLogger
	cleanupDB func()
}

func NewBaseTestEnvironment(t *testing.T) *BaseTestEnvironment {
	client, cleanUp, err := OpenSharedTestDB()
	if err != nil {
		log.Panic(err)
	}

	logger := log2.NewTestEventLogger(t)
	envServices := services.NewDefaultServiceEnvironment(client)
	envServices.EventService.SetLogger(logger)

	return &BaseTestEnvironment{
		cleanupDB: cleanUp,
		Client:    client,
		Services:  envServices,
		Log:       logger,
	}
}

func (e *BaseTestEnvironment) Cleanup() {
	e.Log.Close()
	_ = e.Client.Close()
	e.cleanupDB()
}

func DoTest(
	t *testing.T,
	name string,
	testCase func(t *testing.T, env *BaseTestEnvironment)) {
	t.Helper()

	env := NewBaseTestEnvironment(t)

	t.Run(name, func(t *testing.T) {
		t.Helper()
		t.Parallel()

		defer env.Cleanup()
		testCase(t, env)
	})
}

/*
	Test Helpers below
*/

func (e *BaseTestEnvironment) CreateKebabShop(t *testing.T, name string, lan, lng float64) *ent.KebabShop {
	t.Helper()

	kebabShop, err := e.Services.KebabShopService.CreateKebabShop(name, lan, lng)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	assert.Equal(t, lan, kebabShop.Lat)
	assert.Equal(t, lng, kebabShop.Lng)
	assert.Equal(t, name, kebabShop.Name)
	assert.Nil(t, kebabShop.OsmID)

	e.Log.WaitUntil(event.EventTypeKebabShopCreated, time.Second, func(t *testing.T, event ent.Event) {
		assert.Equal(t, kebabShop.ID, event.Info["id"])
		assert.Equal(t, kebabShop.Name, event.Info["name"])
		assert.Equal(t, strconv.FormatFloat(kebabShop.Lat, 'E', -1, 64), event.Info["lat"])
		assert.Equal(t, strconv.FormatFloat(kebabShop.Lng, 'E', -1, 64), event.Info["long"])
	})

	return kebabShop
}

func AssertKebabShop(t *testing.T, expected, actual *ent.KebabShop) {
	t.Helper()

	if !assert.NotNil(t, expected) || !assert.NotNil(t, actual) {
		t.FailNow()
	}

	assert.Equal(t, expected.ID, actual.ID)
	assert.Equal(t, expected.Lat, actual.Lat)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.OsmID, actual.OsmID)
	assert.Equal(t, expected.Created.Unix(), actual.Created.Unix())
}
