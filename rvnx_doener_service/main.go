package main

import (
	"embed"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	_ "github.com/lib/pq"
	"log"
	"os"
	"rvnx_doener_service/internal/api"
	"rvnx_doener_service/internal/data"
	log2 "rvnx_doener_service/internal/log"
	"rvnx_doener_service/internal/osm"
	"rvnx_doener_service/internal/services"
	"strconv"
	"strings"
	"time"
)

//go:embed all:frontend
var embedFrontend embed.FS

func main() {
	port := 8080
	portStr := os.Getenv("PORT")
	if portStr != "" {
		var err error
		port, err = strconv.Atoi(portStr)
		if err != nil {
			log.Panicln(err)
		}
	}

	debug := strings.ToLower(os.Getenv("DEBUG")) == "true"

	sslMode := "require"
	if debug {
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
	cronScheduler.StartAsync()

	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(gin.Logger())
	api.RouteAPI(engine.Group("/api"), serviceEnv)

	engine.Use(static.Serve("", &data.ServeFileSystemFS{FS: embedFrontend, Root: "frontend"}))

	log.Panicln(engine.Run(":" + strconv.Itoa(port)))
}
