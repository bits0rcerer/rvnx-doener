package services

import (
	"log"
	"os"
	"rvnx_doener_service/ent"
)

type ServiceEnvironment struct {
	TwitchUserService *TwitchUserService
	KebabShopService  *KebabShopService
	EventService      *EventService
}

const (
	clientIDKey     = "TWITCH_CLIENT_ID"
	clientSecretKey = "TWITCH_CLIENT_SECRET"
)

func NewDefaultServiceEnvironment(client *ent.Client) *ServiceEnvironment {
	eventService := NewEventService(client)

	clientID := os.Getenv(clientIDKey)
	if clientID == "" {
		log.Panic("$" + clientIDKey + " is not set")
	}
	clientSecret := os.Getenv(clientSecretKey)
	if clientSecret == "" {
		log.Panic("$" + clientSecretKey + " is not set")
	}

	return &ServiceEnvironment{
		TwitchUserService: NewTwitchUserService(client, eventService, clientID, clientSecret),
		KebabShopService:  NewKebabShopService(client, eventService),
		EventService:      eventService,
	}
}
