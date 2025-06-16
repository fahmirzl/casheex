package middlewares

import (
	"casheex/constants"
	"casheex/structs"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthMiddleware(isAdmin bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
		if tokenString == "" {
			unauthorized(c, "Invalid token")
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(constants.JWT_SECRET), nil
		})
		if err != nil || !token.Valid {
			unauthorized(c, "Invalid token")
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if isAdmin {
			if !ok || claims["role"].(string) != "admin" {
				forbidden(c, "You do not have permission to access this resource")
				return
			}
		}

		c.Set("claims", claims)
		c.Next()
	}
}

func unauthorized(c *gin.Context, msg string) {
	c.JSON(http.StatusUnauthorized, structs.Response{
		Message: msg,
		Error:   "Unauthorized",
		Data:    nil,
	})
	c.Abort()
}

func forbidden(c *gin.Context, msg string) {
	c.JSON(http.StatusForbidden, structs.Response{
		Message: msg,
		Error:   "Forbidden",
		Data:    nil,
	})
	c.Abort()
}
