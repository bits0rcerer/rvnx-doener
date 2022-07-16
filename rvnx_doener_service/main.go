package main

import (
	_ "github.com/lib/pq"
	"log"
	"rvnx_doener_service/internal/data"
	log2 "rvnx_doener_service/internal/log"
	"rvnx_doener_service/internal/services"
)

func main() {
	closeDB, dataClient, err := data.OpenPostgres("require")
	if err != nil {
		log.Panicln(err)
	}
	defer func() {
		_ = closeDB()
	}()

	serviceEnv := services.NewDefaultServiceEnvironment(dataClient)
	serviceEnv.EventService.SetLogger(log2.ConsoleEventLogger{})
}
