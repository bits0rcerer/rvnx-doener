package services

import (
	"context"
	"github.com/jackc/pgtype"
	"rvnx_doener_service/ent"
	"rvnx_doener_service/ent/kebabshop"
	"rvnx_doener_service/ent/shopprice"
	"rvnx_doener_service/internal/model"
)

func NewKebabShopService(client *ent.Client, eventService *EventService) *KebabShopService {
	return &KebabShopService{client: client, context: context.Background(), eventService: eventService}
}

type KebabShopService struct {
	client       *ent.Client
	eventService *EventService
	context      context.Context
}

func (s *KebabShopService) CreateKebabShop(name string, lat, long float64) (*ent.KebabShop, error) {
	kebabShop, err := s.client.KebabShop.Create().
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
	kebabShop, err := s.client.KebabShop.Create().
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
	first, err := s.client.KebabShop.Query().Unique(false).Where(kebabshop.OsmID(*ks.OsmID)).First(s.context)
	if ent.IsNotFound(err) {
		return s.importOSMKebabShop(ks)
	}
	if err != nil {
		return nil, err
	}

	if first.Name == ks.Name && first.Lat == ks.Lat && first.Lng == ks.Lng {
		return nil, err
	}

	_, err = s.client.KebabShop.Update().
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
	shops, err = s.client.KebabShop.Query().Unique(false).
		Where(
			kebabshop.LatGTE(latMin),
			kebabshop.LatLTE(latMax),
			kebabshop.LngGTE(lngMin),
			kebabshop.LngLTE(lngMax),
		).Select(fields...).
		All(s.context)

	return shops, err
}

func (s *KebabShopService) KebabShop(id uint64) (shop *ent.KebabShop, exists bool, err error) {
	shop, err = s.client.KebabShop.Query().Unique(false).
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

func (s *KebabShopService) AddUserScore(shopID uint64, userID int64, anonymous bool, score float64) (notFound bool, err error) {
	shop, author, err := s.getShopAndUser(shopID, userID)
	if err != nil {
		if ent.IsNotFound(err) {
			return true, nil
		}
		return false, err
	}

	scoreRating, err := s.client.ScoreRating.
		Create().
		SetScore(score).
		SetAnonymous(anonymous).
		SetAuthor(author).
		Save(s.context)
	if err != nil {
		return false, err
	}

	_, err = shop.Update().AddUserScores(scoreRating).Save(s.context)
	if err != nil {
		return false, err
	}

	s.eventService.LogUserRating(shopID, userID, anonymous, map[string]interface{}{
		"userScore": score,
	})
	return false, nil
}

func (s *KebabShopService) AddOpinion(shopID uint64, userID int64, anonymous bool, opinion string) (notFound bool, err error) {
	shop, author, err := s.getShopAndUser(shopID, userID)
	if err != nil {
		if ent.IsNotFound(err) {
			return true, nil
		}
		return false, err
	}

	userOpinion, err := s.client.UserOpinion.
		Create().
		SetOpinion(opinion).
		SetAnonymous(anonymous).
		SetAuthor(author).
		Save(s.context)
	if err != nil {
		return false, err
	}

	_, err = shop.Update().AddUserOpinions(userOpinion).Save(s.context)
	if err != nil {
		return false, err
	}

	s.eventService.LogUserRating(shopID, userID, anonymous, map[string]interface{}{
		"opinion": opinion,
	})
	return false, nil
}

func (s *KebabShopService) AddPrices(shopID uint64, userID int64, anonymous bool, prices map[string]model.PriceEntry) (notFound bool, err error) {
	shop, author, err := s.getShopAndUser(shopID, userID)
	if err != nil {
		if ent.IsNotFound(err) {
			return true, nil
		}
		return false, err
	}

	var shopPrices []*ent.ShopPrice
	eventPayload := map[string]interface{}{}
	for k, p := range prices {
		numericPrice := pgtype.Numeric{}
		err = numericPrice.Set(p.Price)
		if err != nil {
			return false, err
		}

		shopPrice, err := s.client.ShopPrice.
			Create().
			SetCurrency(shopprice.Currency(p.Currency)).
			SetPrice(&numericPrice).
			SetAuthor(author).
			SetAnonymous(anonymous).
			SetPriceType(shopprice.PriceType(k)).
			Save(s.context)
		if err != nil {
			return false, err
		}

		shopPrices = append(shopPrices, shopPrice)
		eventPayload[k] = map[string]interface{}{
			"price":    p.Price,
			"currency": p.Currency,
		}
	}

	_, err = shop.Update().AddUserPrices(shopPrices...).Save(s.context)
	if err != nil {
		return false, err
	}

	s.eventService.LogUserRating(shopID, userID, anonymous, map[string]interface{}{
		"prices": eventPayload,
	})
	return false, nil
}

func (s *KebabShopService) getShopAndUser(shopID uint64, userID int64) (shop *ent.KebabShop, author *ent.TwitchUser, err error) {
	shop, err = s.client.KebabShop.Get(s.context, shopID)
	if err != nil {
		return nil, nil, err
	}

	author, err = s.client.TwitchUser.Get(s.context, int64(userID))
	if err != nil {
		return nil, nil, err
	}

	return shop, author, nil
}
