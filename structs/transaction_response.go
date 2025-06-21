package structs

type TransactionResponse struct {
	ID int `json:"id"`
	Date interface{} `json:"date"`
	UserID *int `json:"user_id"`
	Total *int `json:"total"`
	Paid *int `json:"paid"`
	Change *int `json:"change"`
	CreatedAt interface{} `json:"created_at"`
	UpdatedAt interface{} `json:"updated_at"`
	User interface{} `json:"user"`
	TransactionDetails interface{} `json:"transaction_details"`
}