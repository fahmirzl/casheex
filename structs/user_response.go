package structs

type UserResponse struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Gender string `json:"gender"`
	Username string `json:"username"`
	Role string `json:"role"`
	CreatedAt interface{} `json:"created_at"`
	UpdatedAt interface{} `json:"updated_at"`
}