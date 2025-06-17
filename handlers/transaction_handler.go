package handlers

import (
	"casheex/configs"
	"casheex/repositories"
	"casheex/structs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func TransactionStore(c *gin.Context) {
	var transaction structs.Transaction
	var cart structs.Cart
	claims := c.MustGet("claims").(jwt.MapClaims)
	rawUserID, _ := claims["user_id"].(float64)

	err := c.BindJSON(&transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Response{
			Message: "Invalid JSON data",
			Error:   "Bad Request",
			Data:    nil,
		})
		return
	}

	userID := int(rawUserID)
	cart.UserID = &userID
	transaction.UserID = &userID

	validations := map[string]string{}
	if transaction.Paid == nil {
		validations["paid"] = "The paid field is required"
	}
	if len(validations) > 0 {
		c.JSON(http.StatusUnprocessableEntity, structs.Response{
			Message: "Validation error",
			Error:   "Unprocessable Entity",
			Data: gin.H{
				"validations": validations,
			},
		})
		return
	}

	cartResponse, err := repositories.CartWithProductByUserId(configs.DB, &cart)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Response{
			Message: "Internal server error",
			Error:   "Internal Server Error",
			Data:    nil,
		})
		return
	}

	var total int
	for _, response := range cartResponse {
		total = total + *response.Subtotal
	}

	transaction.Total = &total
	change := *(transaction.Paid) - *(transaction.Total)
	transaction.Change = &change

	err = repositories.InsertTransaction(configs.DB, &transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Response{
			Message: "Internal server error",
			Error:   "Internal Server Error",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, structs.Response{
		Message: "Transaction inserted successfully",
		Error:   nil,
		Data:    gin.H{
			"total": transaction.Total,
			"paid": transaction.Paid,
			"change": transaction.Change,
		},
	})
}