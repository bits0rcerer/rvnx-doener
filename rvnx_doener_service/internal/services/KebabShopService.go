package services

import (
	"context"
	"github.com/jackc/pgtype"
	"rvnx_doener_service/ent"
)

func NewKebabShopService(client *ent.Client, eventService *EventService) *KebabShopService {
	return &KebabShopService{client: client.KebabShop, context: context.Background(), eventService: eventService}
}

type KebabShopService struct {
	client       *ent.KebabShopClient
	eventService *EventService
	context      context.Context
}

func (s *KebabShopService) CreateKebabShop(name string, lat, long float64) (*ent.KebabShop, error) {
	kebabShop, err := s.client.Create().
		SetName(name).
		SetPoint(&pgtype.Point{
			P: pgtype.Vec2{
				X: lat,
				Y: long,
			},
			Status: pgtype.Present,
		}).
		Save(s.context)

	if err != nil {
		return nil, err
	}

	go s.eventService.LogKebabSopCreated(kebabShop)

	return kebabShop, err
}
