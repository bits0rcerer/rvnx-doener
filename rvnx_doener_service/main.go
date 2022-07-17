package main

import (
	"github.com/go-co-op/gocron"
	_ "github.com/lib/pq"
	"log"
	"os"
	"rvnx_doener_service/internal/data"
	log2 "rvnx_doener_service/internal/log"
	"rvnx_doener_service/internal/osm"
	"rvnx_doener_service/internal/services"
	"strings"
	"time"
)

func main() {
	sslMode := "require"
	if strings.ToLower(os.Getenv("DEBUG")) == "true" {
		sslMode = "disable"
	}

	closeDB, dataClient, err := data.OpenPostgres(sslMode)
	if err != nil {
		log.Panicln(err)
	}
	defer func() {
		_ = closeDB()
	}()

	serviceEnv := services.NewDefaultServiceEnvironment(dataClient)
	serviceEnv.EventService.SetLogger(log2.ConsoleEventLogger{})

	cronScheduler := gocron.NewScheduler(time.UTC)
	_, err = cronScheduler.
		SingletonMode().
		Every(1).Day().
		StartImmediately().Do(osm.SyncOSMKebabShops, serviceEnv.KebabShopService)
	if err != nil {
		log.Panicln(err)
	}
	cronScheduler.StartBlocking()
}
