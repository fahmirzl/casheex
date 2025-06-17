package structs

type Product struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Stock *int `json:"stock"`
	PurchasePrice *int `json:"purchase_price"`
	SellingPrice *int `json:"selling_price"`
	CreatedAt interface{} `json:"created_at"`
	UpdatedAt interface{} `json:"updated_at"`
}