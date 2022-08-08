package model

type PriceEntry struct {
	Price      string `json:"price"`
	Currency   string `json:"currency"`
	OrderIndex int    `json:"order_index"`
}
