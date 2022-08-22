package test

import (
	_ "embed"
	"encoding/xml"
	"log"
	"math/rand"
	"rvnx_doener_service/ent"
	"rvnx_doener_service/ent/event"
	_ "rvnx_doener_service/ent/runtime"
	log2 "rvnx_doener_service/internal/log"
	osm2 "rvnx_doener_service/internal/osm"
	"rvnx_doener_service/internal/services"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/paulmach/osm"
	"github.com/stretchr/testify/assert"
)

//go:embed osmTestData.xml
var osmTestDataXML []byte

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

func (e *BaseTestEnvironment) LoadOSMTestData(t *testing.T) {
	t.Helper()

	var osmData osm.OSM
	err := xml.Unmarshal(osmTestDataXML, &osmData)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	shops, err := osm2.ParseOSMKebabShops(&osmData)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	for _, shop := range shops {
		_, err := e.Services.KebabShopService.UpdateOrInsertKebabShop(&shop)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
	}
}

func (e *BaseTestEnvironment) CreateKebabShop(t *testing.T, name string, lan, lng float64) *ent.KebabShop {
	t.Helper()

	kebabShop, err := e.Services.KebabShopService.CreateKebabShop(name, lan, lng, true, nil)
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

func (e BaseTestEnvironment) CreateUser(t *testing.T, name string) *ent.TwitchUser {
	user, err := e.Services.TwitchUserService.CreateOrUpdateUser(&ent.TwitchUser{
		ID:          rand.Int63(),
		Login:       strings.ReplaceAll(strings.ToLower(name), " ", "_"),
		Email:       name + "@test.org",
		DisplayName: name,
		CreatedAt:   time.Now(),
	})
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	return user
}
