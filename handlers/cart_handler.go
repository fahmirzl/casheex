package handlers

import (
	"casheex/configs"
	"casheex/repositories"
	"casheex/structs"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func CartIndex(c *gin.Context) {
	var cart structs.Cart
	claims := c.MustGet("claims").(jwt.MapClaims)
	rawUserID, _ := claims["user_id"].(float64)

	userID := int(rawUserID)
	cart.UserID = &userID

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

	c.JSON(http.StatusOK, structs.Response{
		Message: fmt.Sprintf("Carts with user id %d successfully found", userID),
		Error:   nil,
		Data: gin.H{
			"carts": cartResponse,
			"total": total,
		},
	})
}

func AddToCart(c *gin.Context) {
	var cart structs.Cart
	claims := c.MustGet("claims").(jwt.MapClaims)
	rawUserID, _ := claims["user_id"].(float64)

	err := c.BindJSON(&cart)
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

	validations := map[string]string{}
	if cart.ProductID == nil {
		validations["product_id"] = "The product id field is required"
	}
	if cart.Quantity == nil {
		validations["quantity"] = "The quantity field is required"
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

	var product structs.Product
	product.ID = *cart.ProductID
	err = repositories.GetProductById(configs.DB, &product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Response{
			Message: "Internal server error",
			Error:   "Internal Server Error",
			Data:    nil,
		})
		return
	}

	cart.SellingPrice = product.SellingPrice
	subtotal := (*cart.SellingPrice) * (*cart.Quantity)
	cart.Subtotal = &subtotal

	err = repositories.AddProductToCart(configs.DB, &cart)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Response{
			Message: "Internal server error",
			Error:   "Internal Server Error",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, structs.Response{
		Message: "Product successfully added to cart",
		Error:   nil,
		Data:    nil,
	})
}

func RemoveFromCart(c *gin.Context) {
	var cart structs.Cart
	id, _ := strconv.Atoi(c.Param("id"))
	claims := c.MustGet("claims").(jwt.MapClaims)
	rawUserID, _ := claims["user_id"].(float64)

	userID := int(rawUserID)
	cart.ID = id

	err := repositories.DeleteProductFromCart(configs.DB, &cart)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, structs.Response{
			Message: fmt.Sprintf("Cart with id %d not found", id),
			Error:   "Not Found",
			Data:    nil,
		})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Response{
			Message: "Internal server error",
			Error:   "Internal Server Error",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, structs.Response{
		Message: fmt.Sprintf("Product successfully removed from cart for user with id %d", userID),
		Error:   nil,
		Data:    nil,
	})
}
