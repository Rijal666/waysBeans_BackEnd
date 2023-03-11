package transactionsdto

type TransactionRequest struct {
	UserID        int `json:"user_id"`
	TotalQuantity int `json:"total_quantity"`
	TotalPrice    int `json:"total_price"`
}
