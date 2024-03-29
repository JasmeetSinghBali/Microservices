// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateDonutHubInput struct {
	Title            string `json:"title"`
	Description      string `json:"description"`
	ShopLocation     string `json:"shop_location"`
	Bestselling      string `json:"bestselling"`
	PriceBestselling int    `json:"price_bestselling"`
}

type DeleteDonutHubResponse struct {
	DeleteHubID string `json:"deleteHubId"`
}

type DonutHub struct {
	ID               string `json:"_id"`
	Title            string `json:"title"`
	Description      string `json:"description"`
	ShopLocation     string `json:"shop_location"`
	Bestselling      string `json:"bestselling"`
	PriceBestselling int    `json:"price_bestselling"`
}

type UpdateDonutHubInput struct {
	Bestselling      *string `json:"bestselling"`
	PriceBestselling *int    `json:"price_bestselling"`
}
