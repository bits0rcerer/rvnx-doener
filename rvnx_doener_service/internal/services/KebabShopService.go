package services

import (
	"context"
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
		SetLat(lat).
		SetLng(long).
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
		SetLat(ks.Lat).
		SetLng(ks.Lng).
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

	if first.Name == ks.Name && first.Lat == ks.Lat && first.Lng == ks.Lng {
		return nil, err
	}

	_, err = s.client.Update().
		Where(kebabshop.OsmID(*ks.OsmID)).
		SetName(ks.Name).
		SetLat(ks.Lat).
		SetLng(ks.Lng).
		Save(s.context)
	if err != nil {
		return nil, err
	}

	ks.ID = first.ID
	s.eventService.LogKebabShopUpdatedFromOSM(ks)

	return ks, err
}

func (s *KebabShopService) Within(latMin, latMax, lngMin, lngMax float64, fields ...string) (shops []*ent.KebabShop, err error) {
	shops, err = s.client.Query().Unique(false).
		Where(
			kebabshop.LatGTE(latMin),
			kebabshop.LatLTE(latMax),
			kebabshop.LngGTE(lngMin),
			kebabshop.LngLTE(lngMax),
		).Select(fields...).
		All(s.context)

	return shops, err
}

func (s *KebabShopService) KebabShop(id int) (shop *ent.KebabShop, exists bool, err error) {
	shop, err = s.client.Query().Unique(false).
		Where(
			kebabshop.ID(id),
		).First(s.context)

	if ent.IsNotFound(err) {
		return nil, false, nil
	}

	if err != nil {
		return nil, false, err
	}

	return shop, true, nil
}
