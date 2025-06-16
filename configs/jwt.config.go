package configs

import (
	"casheex/constants"
	"casheex/structs"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(user *structs.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString([]byte(constants.JWT_SECRET))
}
