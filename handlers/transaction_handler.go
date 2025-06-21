package handlers

import (
	"casheex/configs"
	"casheex/repositories"
	"casheex/structs"
	"fmt"
	"net/http"
	"time"

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
	if *transaction.Paid < *transaction.Total {
		c.JSON(http.StatusConflict, structs.Response{
			Message: "Insufficient funds",
			Error:   "Conflict",
			Data:    nil,
		})
		return
	}
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
		Data: gin.H{
			"total":  transaction.Total,
			"paid":   transaction.Paid,
			"change": transaction.Change,
		},
	})
}

func TransactionList(c *gin.Context) {
	var user structs.User
	claims := c.MustGet("claims").(jwt.MapClaims)
	rawUserID, _ := claims["user_id"].(float64)
	date := c.Query("date")

	userID := int(rawUserID)
	user.ID = userID

	if date != "" {
		_, err := time.Parse("2006-01-02", date)
		if err != nil {
			c.JSON(http.StatusBadRequest, structs.Response{
				Message: "Invalid date format. Use yyyy-mm-dd",
				Error:   "Bad Request",
				Data:    nil,
			})
			return
		}
	}

	transactions, err := repositories.GetTransactionByUserIdAndDate(configs.DB, &user, date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Response{
			Message: "Internal server error",
			Error:   "Internal Server Error",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, structs.Response{
		Message: fmt.Sprintf("Transaction with user id %d successfully found", userID),
		Error:   nil,
		Data:    transactions,
	})
}

func TransactionAll(c *gin.Context) {
	date := c.Query("date")

	if date != "" {
		_, err := time.Parse("2006-01-02", date)
		if err != nil {
			c.JSON(http.StatusBadRequest, structs.Response{
				Message: "Invalid date format. Use yyyy-mm-dd",
				Error:   "Bad Request",
				Data:    nil,
			})
			return
		}
	}

	transactions, err := repositories.GetAllTransaction(configs.DB, date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Response{
			Message: "Internal server error",
			Error:   "Internal Server Error",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, structs.Response{
		Message: "Transactions retrieved succesfully",
		Error:   nil,
		Data: gin.H{
			"transactions": transactions,
		},
	})
}

func Profit(c *gin.Context) {
	var profitResponse structs.ProfitResponse
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if startDate == "" && endDate == "" {
		c.JSON(http.StatusBadRequest, structs.Response{
			Message: "Missing required parameters: start_date and end_date.",
			Error:   "Bad Request",
			Data:    nil,
		})
		return
	}

	if startDate != "" {
		_, err := time.Parse("2006-01-02", startDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, structs.Response{
				Message: "Invalid start date format. Use yyyy-mm-dd",
				Error:   "Bad Request",
				Data:    nil,
			})
			return
		}
	}

	if endDate != "" {
		_, err := time.Parse("2006-01-02", endDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, structs.Response{
				Message: "Invalid start end format. Use yyyy-mm-dd",
				Error:   "Bad Request",
				Data:    nil,
			})
			return
		}
	}

	err := repositories.GetProfitBetweenDate(configs.DB, &profitResponse, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Response{
			Message: "Internal server error",
			Error:   "Internal Server Error",
			Data:    nil,
		})
		return
	}

	profitResponse.Period = gin.H{
		"start_date": startDate,
		"end_date": endDate,
	}

	c.JSON(http.StatusOK, structs.Response{
		Message: "Profit counted succesfully",
		Error:   nil,
		Data: gin.H{
			"profit": profitResponse,
		},
	})
}
