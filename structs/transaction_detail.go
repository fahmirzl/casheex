package structs

type TransactionDetail struct {
	ID int `json:"id"`
	ProductID *int `json:"product_id"`
	PurchasePrice *int `json:"purchase_price"`
	SellingPrice *int `json:"selling_price"`
	Quantity *int `json:"quantity"`
	Subtotal *int `json:"subtotal"`
	CreatedAt interface{} `json:"created_at"`
	UpdatedAt interface{} `json:"updated_at"`
}