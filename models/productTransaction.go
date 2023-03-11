package models

type ProductTransaction struct {
	ProductID     int    `json:"product_id"`
	ProductName   string `json:"product_name"`
	ProductPrice  int    `json:"product_price"`
	OrderQuantity int    `json:"order_quantity"`
	TransactionID int    `json:"-"`
}
