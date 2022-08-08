package services

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/jackc/pgtype"
	"rvnx_doener_service/ent"
	"rvnx_doener_service/ent/kebabshop"
	"rvnx_doener_service/ent/scorerating"
	"rvnx_doener_service/ent/shopprice"
	"rvnx_doener_service/ent/useropinion"
	"rvnx_doener_service/internal/model"
	"strconv"
	"time"
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
		"Opinion": opinion,
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

func (s *KebabShopService) shopUserScore(id uint64) (score *float64, err error) {
	var resp []struct {
		KebabShopUserScores interface{} `json:"kebab_shop_user_scores"`
		Avg                 float64     `json:"avg"`
	}

	err = s.client.ScoreRating.Query().Unique(false).
		Where(
			scorerating.HasShopWith(kebabshop.ID(id)),
		).
		GroupBy(scorerating.ShopColumn).Aggregate(func(s *sql.Selector) string {
		return sql.Avg(s.C(scorerating.FieldScore))
	}).Scan(s.context, &resp)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}

	if resp == nil {
		return nil, nil
	}

	return &resp[0].Avg, nil
}

func (s *KebabShopService) shopPrices(id uint64) (prices map[shopprice.PriceType]model.PriceEntry, err error) {
	var priceIDsWithCurrency []struct {
		ID        uint64              `json:"max"`
		PriceType shopprice.PriceType `json:"price_type"`
	}
	err = s.client.ShopPrice.Query().Where(
		shopprice.HasShopWith(kebabshop.ID(id)),
	).
		GroupBy(shopprice.FieldPriceType).
		Aggregate(ent.Max(shopprice.FieldID)).
		Scan(context.Background(), &priceIDsWithCurrency)
	if err != nil {
		return nil, err
	}

	var ids []uint64
	for _, d := range priceIDsWithCurrency {
		ids = append(ids, d.ID)
	}

	priceList, err := s.client.ShopPrice.Query().Where(shopprice.IDIn(ids...)).All(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}

	prices = make(map[shopprice.PriceType]model.PriceEntry)
	for _, p := range priceList {
		prices[p.PriceType] = model.PriceEntry{
			Price:    p.Price.Int.String() + "e" + strconv.Itoa(int(p.Price.Exp)),
			Currency: string(p.Currency),
		}
	}

	return prices, nil
}

type DatedReview struct {
	Time    time.Time
	Opinion string
}

func (s *KebabShopService) shopReviews(id uint64) (reviews []DatedReview, err error) {
	opinions, err := s.client.UserOpinion.Query().
		Unique(false).
		Where(
			useropinion.HasShopWith(kebabshop.ID(id)),
		).Order(ent.Desc(useropinion.ShopColumn)).
		Limit(5).All(s.context)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}

	for _, o := range opinions {
		reviews = append(reviews, DatedReview{
			Time:    o.Created,
			Opinion: o.Opinion,
		})
	}

	return reviews, err
}

func (s *KebabShopService) GetShopRating(id uint64) (
	score *float64,
	prices map[shopprice.PriceType]model.PriceEntry,
	reviews []DatedReview,
	err error,
) {
	score, err = s.shopUserScore(id)
	if err != nil {
		return nil, nil, nil, err
	}

	prices, err = s.shopPrices(id)
	if err != nil {
		return nil, nil, nil, err
	}

	reviews, err = s.shopReviews(id)
	if err != nil {
		return nil, nil, nil, err
	}

	return score, prices, reviews, err
}
