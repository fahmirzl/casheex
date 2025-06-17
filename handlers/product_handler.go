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
)

func ProductIndex(c *gin.Context) {
	products, err := repositories.GetAllProduct(configs.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Response{
			Message: "Internal server error",
			Error:   "Internal Server Error",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, structs.Response{
		Message: "Products retrieved succesfully",
		Error:   nil,
		Data: gin.H{
			"products": products,
		},
	})
}

func ProductStore(c *gin.Context) {
	var product structs.Product

	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Response{
			Message: "Invalid JSON data",
			Error:   "Bad Request",
			Data:    nil,
		})
		return
	}

	validations := map[string]string{}
	if product.Name == "" {
		validations["name"] = "The name field is required"
	}
	if product.Stock == nil {
		validations["stock"] = "The stock field is required"
	}
	if product.PurchasePrice == nil {
		validations["purchase_price"] = "The purchase price field is required"
	}
	if product.SellingPrice == nil {
		validations["selling_price"] = "The selling price field is required"
	}
	if len(validations) > 0 {
		c.JSON(http.StatusUnprocessableEntity, structs.Response{
			Message: "Validation error",
			Error:   "Unprocessable Entity",
			Data:    gin.H{
				"validations": validations,
			},
		})
		return
	}

	err = repositories.InsertProduct(configs.DB, &product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Response{
			Message: "Internal server error",
			Error:   "Internal Server Error",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, structs.Response {
		Message: "Product inserted successfully",
		Error: nil,
		Data: gin.H{
			"product": product,
		},
	})
}

func ProductFind(c *gin.Context) {
	var product structs.Product
	id, _ := strconv.Atoi(c.Param("id"))

	product.ID = id

	err := repositories.GetProductById(configs.DB, &product)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, structs.Response{
			Message: fmt.Sprintf("Product with id %d not found", id),
			Error:   "Not Found",
			Data:    nil,
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Response{
			Message: "Internal server error",
			Error:   "Internal Server Error",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, structs.Response {
		Message: fmt.Sprintf("Product with id %d successfully found", id),
		Error: nil,
		Data: gin.H{
			"product": product,
		},
	})
}

func ProductUpdate(c *gin.Context) {
	var product structs.Product
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Response{
			Message: "Invalid JSON data",
			Error:   "Bad Request",
			Data:    nil,
		})
		return
	}

	validations := map[string]string{}
	if product.Name == "" {
		validations["name"] = "The name field is required"
	}
	if product.Stock == nil {
		validations["stock"] = "The stock field is required"
	}
	if product.PurchasePrice == nil {
		validations["purchase_price"] = "The purchase price field is required"
	}
	if product.SellingPrice == nil {
		validations["selling_price"] = "The selling price field is required"
	}
	if len(validations) > 0 {
		c.JSON(http.StatusUnprocessableEntity, structs.Response{
			Message: "Validation error",
			Error:   "Unprocessable Entity",
			Data:    gin.H{
				"validations": validations,
			},
		})
		return
	}

	product.ID = id

	err = repositories.UpdateProduct(configs.DB, &product)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, structs.Response{
			Message: fmt.Sprintf("Product with id %d not found", id),
			Error:   "Not Found",
			Data:    nil,
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Response{
			Message: "Internal server error",
			Error:   "Internal Server Error",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, structs.Response {
		Message: "Product updated successfully",
		Error: nil,
		Data: gin.H{
			"product": product,
		},
	})
}

func ProductDestroy(c *gin.Context) {
	var product structs.Product
	id, _ := strconv.Atoi(c.Param("id"))

	product.ID = id

	err := repositories.DeleteProduct(configs.DB, &product)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, structs.Response{
			Message: fmt.Sprintf("Product with id %d not found", id),
			Error:   "Not Found",
			Data:    nil,
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Response{
			Message: "Internal server error",
			Error:   "Internal Server Error",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, structs.Response {
		Message: "Product deleted successfully",
		Error: nil,
		Data: nil,
	})
}