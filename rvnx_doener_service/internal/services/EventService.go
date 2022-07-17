package services

import (
	"context"
	"log"
	"rvnx_doener_service/ent"
	"rvnx_doener_service/ent/event"
	log2 "rvnx_doener_service/internal/log"
	"strconv"
)

func NewEventService(client *ent.Client) *EventService {
	return &EventService{client: client.Event, context: context.Background()}
}

type EventService struct {
	client  *ent.EventClient
	context context.Context
	logger  log2.EventLogger
}

func (s *EventService) SetLogger(logger log2.EventLogger) {
	s.logger = logger
}

func (s *EventService) LogKebabShopCreated(ks *ent.KebabShop) {
	newEvent, err := s.client.Create().
		SetEventType(event.EventTypeKebabShopCreated).
		SetInfo(map[string]interface{}{
			"id":   ks.ID,
			"name": ks.Name,
			"lat":  strconv.FormatFloat(ks.Point.P.X, 'E', -1, 64),
			"long": strconv.FormatFloat(ks.Point.P.Y, 'E', -1, 64),
		}).Save(s.context)

	if err != nil {
		log.Panicln("unable to store event: " + newEvent.String())
	}

	if s.logger != nil {
		s.logger.Handle(newEvent)
	}
}

func (s *EventService) LogKebabShopImported(ks *ent.KebabShop) {
	newEvent, err := s.client.Create().
		SetEventType(event.EventTypeKebabShopImported).
		SetInfo(map[string]interface{}{
			"id":     ks.ID,
			"osm_id": ks.OsmID,
			"name":   ks.Name,
			"lat":    strconv.FormatFloat(ks.Point.P.X, 'E', -1, 64),
			"long":   strconv.FormatFloat(ks.Point.P.Y, 'E', -1, 64),
		}).Save(s.context)

	if err != nil {
		log.Panicln("unable to store event: " + newEvent.String())
	}

	if s.logger != nil {
		s.logger.Handle(newEvent)
	}
}

func (s *EventService) LogKebabShopUpdatedFromOSM(ks *ent.KebabShop) {
	newEvent, err := s.client.Create().
		SetEventType(event.EventTypeKebabShopUpdatedFromOsm).
		SetInfo(map[string]interface{}{
			"id":     ks.ID,
			"osm_id": ks.OsmID,
			"name":   ks.Name,
			"lat":    strconv.FormatFloat(ks.Point.P.X, 'E', -1, 64),
			"long":   strconv.FormatFloat(ks.Point.P.Y, 'E', -1, 64),
		}).Save(s.context)

	if err != nil {
		log.Panicln("unable to store event: " + newEvent.String())
	}

	if s.logger != nil {
		s.logger.Handle(newEvent)
	}
}
