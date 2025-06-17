package structs

type CartResponse struct {
	ID int `json:"id"`
	ProductID *int `json:"product_id"`
	SellingPrice *int `json:"selling_price"`
	Quantity *int `json:"quantity"`
	Subtotal *int `json:"subtotal"`
	CreatedAt interface{} `json:"created_at"`
	UpdatedAt interface{} `json:"updated_at"`
	Product interface{} `json:"product"`
}