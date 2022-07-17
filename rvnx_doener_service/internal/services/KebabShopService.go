package services

import (
	"context"
	"github.com/jackc/pgtype"
	"rvnx_doener_service/ent"
	"rvnx_doener_service/ent/kebabshop"
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

	go s.eventService.LogKebabShopCreated(kebabShop)

	return kebabShop, err
}

func (s *KebabShopService) importOSMKebabShop(ks *ent.KebabShop) (*ent.KebabShop, error) {
	kebabShop, err := s.client.Create().
		SetName(ks.Name).
		SetOsmID(*ks.OsmID).
		SetPoint(ks.Point).
		Save(s.context)

	if err != nil {
		return nil, err
	}

	go s.eventService.LogKebabShopImported(kebabShop)

	return kebabShop, err
}

func (s *KebabShopService) UpdateOrInsertKebabShop(ks *ent.KebabShop) (*ent.KebabShop, error) {
	first, err := s.client.Query().Unique(false).Where(kebabshop.OsmID(*ks.OsmID)).First(s.context)
	if ent.IsNotFound(err) {
		return s.importOSMKebabShop(ks)
	}
	if err != nil {
		return nil, err
	}

	if first.Name == ks.Name && first.Point.P.X == ks.Point.P.X && first.Point.P.Y == ks.Point.P.Y {
		return nil, err
	}

	_, err = s.client.Update().
		Where(kebabshop.OsmID(*ks.OsmID)).
		SetName(ks.Name).SetPoint(ks.Point).
		Save(s.context)
	if err != nil {
		return nil, err
	}

	ks.ID = first.ID
	s.eventService.LogKebabShopUpdatedFromOSM(ks)

	return ks, err
}
