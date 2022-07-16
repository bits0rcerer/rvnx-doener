package services

import "rvnx_doener_service/ent"

type ServiceEnvironment struct {
	KebabShopService *KebabShopService
	EventService     *EventService
}

func NewDefaultServiceEnvironment(client *ent.Client) *ServiceEnvironment {
	eventService := NewEventService(client)

	return &ServiceEnvironment{
		KebabShopService: NewKebabShopService(client, eventService),
		EventService:     eventService,
	}
}
