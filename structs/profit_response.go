package structs

type ProfitResponse struct {
	Revenue  int         `json:"revenue"`
	Expense  int         `json:"expense"`
	Profit   int         `json:"profit"`
	Currency string      `json:"currency"`
	Period   interface{} `json:"period"`
}
