// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Node interface {
	IsNode()
	GetID() string
}

type Product struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type YahooAuctionProduct struct {
	ID             string `json:"id"`
	ProductID      int    `json:"productId"`
	YahooAuctionID string `json:"yahooAuctionId"`
	Name           string `json:"name"`
	Price          int    `json:"price"`
	Published      bool   `json:"published"`
}
