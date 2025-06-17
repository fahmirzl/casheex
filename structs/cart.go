package structs

type Cart struct {
	ID int `json:"id"`
	ProductID *int `json:"product_id"`
	SellingPrice *int `json:"selling_price"`
	Quantity *int `json:"quantity"`
	Subtotal *int `json:"subtotal"`
	UserID *int `json:"user_id"`
	CreatedAt interface{} `json:"created_at"`
	UpdatedAt interface{} `json:"updated_at"`
}