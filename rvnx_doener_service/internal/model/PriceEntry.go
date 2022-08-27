package model

type PriceEntry struct {
	Price      float64 `json:"price"`
	Currency   string  `json:"currency"`
	OrderIndex int     `json:"order_index"`
}
