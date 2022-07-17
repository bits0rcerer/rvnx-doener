package test

import (
	"log"
	"rvnx_doener_service/ent"
	log2 "rvnx_doener_service/internal/log"
	"rvnx_doener_service/internal/services"
	"testing"
)

type BaseTestEnvironment struct {
	client    *ent.Client
	services  *services.ServiceEnvironment
	log       *log2.TestEventLogger
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
		client:    client,
		services:  envServices,
		log:       logger,
	}
}

func (e *BaseTestEnvironment) Cleanup() {
	e.log.Close()
	_ = e.client.Close()
	e.cleanupDB()
}

func DoTest(
	t *testing.T,
	name string,
	testCase func(t *testing.T, client *ent.Client, services *services.ServiceEnvironment, log *log2.TestEventLogger)) {
	t.Helper()

	env := NewBaseTestEnvironment(t)

	t.Run(name, func(t *testing.T) {
		t.Helper()
		t.Parallel()

		defer env.Cleanup()
		testCase(t, env.client, env.services, env.log)
	})
}
