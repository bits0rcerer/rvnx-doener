package services

import (
	"context"
	"log"
	"strconv"

	"rvnx_doener_service/ent"
	"rvnx_doener_service/ent/event"
	log2 "rvnx_doener_service/internal/log"
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
			"id":      ks.ID,
			"name":    ks.Name,
			"lat":     strconv.FormatFloat(ks.Lat, 'E', -1, 64),
			"long":    strconv.FormatFloat(ks.Lng, 'E', -1, 64),
			"visible": ks.Visible,
		}).Save(s.context)
	if err != nil {
		log.Panicln("unable to store event: " + newEvent.String())
	}

	if s.logger != nil {
		s.logger.Handle(newEvent)
	}
}

func (s *EventService) LogKebabShopSubmitted(ks *ent.KebabShop, submitterID int64) {
	newEvent, err := s.client.Create().
		SetEventType(event.EventTypeUserSubmittedAShop).
		SetInfo(map[string]interface{}{
			"user_id": submitterID,
			"id":      ks.ID,
			"name":    ks.Name,
			"lat":     strconv.FormatFloat(ks.Lat, 'E', -1, 64),
			"lng":    strconv.FormatFloat(ks.Lng, 'E', -1, 64),
			"visible": ks.Visible,
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
			"id":      ks.ID,
			"osm_id":  ks.OsmID,
			"name":    ks.Name,
			"lat":     strconv.FormatFloat(ks.Lat, 'E', -1, 64),
			"long":    strconv.FormatFloat(ks.Lng, 'E', -1, 64),
			"visible": ks.Visible,
		}).Save(s.context)
	if err != nil {
		log.Panicln("unable to store event: " + event.EventTypeKebabShopImported)
	}

	if s.logger != nil {
		s.logger.Handle(newEvent)
	}
}

func (s *EventService) LogKebabShopUpdatedFromOSM(ks *ent.KebabShop) {
	newEvent, err := s.client.Create().
		SetEventType(event.EventTypeKebabShopUpdatedFromOsm).
		SetInfo(map[string]interface{}{
			"id":      ks.ID,
			"osm_id":  ks.OsmID,
			"name":    ks.Name,
			"lat":     strconv.FormatFloat(ks.Lat, 'E', -1, 64),
			"long":    strconv.FormatFloat(ks.Lng, 'E', -1, 64),
			"visible": ks.Visible,
		}).Save(s.context)
	if err != nil {
		log.Panicln("unable to store event: " + newEvent.String())
	}

	if s.logger != nil {
		s.logger.Handle(newEvent)
	}
}

func (s *EventService) LogFirstTimeUserLogin(tu *ent.TwitchUser) {
	newEvent, err := s.client.Create().
		SetEventType(event.EventTypeUserLoggedInFirstTime).
		SetInfo(map[string]interface{}{
			"id":         tu.ID,
			"login":      tu.Login,
			"display":    tu.DisplayName,
			"email":      tu.Email,
			"created_at": tu.CreatedAt,
		}).Save(s.context)
	if err != nil {
		log.Panicln("unable to store event: " + newEvent.String())
	}

	if s.logger != nil {
		s.logger.Handle(newEvent)
	}
}

func (s *EventService) LogUserLogin(tu *ent.TwitchUser) {
	newEvent, err := s.client.Create().
		SetEventType(event.EventTypeUserLoggedIn).
		SetInfo(map[string]interface{}{
			"id":      tu.ID,
			"login":   tu.Login,
			"display": tu.DisplayName,
			"email":   tu.Email,
		}).Save(s.context)
	if err != nil {
		log.Panicln("unable to store event: " + newEvent.String())
	}

	if s.logger != nil {
		s.logger.Handle(newEvent)
	}
}

func (s *EventService) LogUserRating(shopID uint64, userID int64, anonymous bool, payload map[string]interface{}) {
	payload["userID"] = userID
	payload["shopID"] = shopID
	payload["anonymous"] = anonymous

	newEvent, err := s.client.Create().
		SetEventType(event.EventTypeUserSubmittedARating).
		SetInfo(payload).
		Save(s.context)
	if err != nil {
		log.Panicln("unable to store event: " + newEvent.String())
	}

	if s.logger != nil {
		s.logger.Handle(newEvent)
	}
}
