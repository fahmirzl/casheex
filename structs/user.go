package structs

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Gender string `json:"gender"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role string `json:"role"`
	CreatedAt interface{} `json:"created_at"`
	UpdatedAt interface{} `json:"updated_at"`
}