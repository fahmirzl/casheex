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

// ProductIndex godoc
// @Summary Get all products
// @Description Retrieve a list of all products
// @Tags Products
// @Security BearerAuth
// @Produce json
// @Success 200 {object} structs.Response
// @Failure 401 {object} structs.Response
// @Failure 500 {object} structs.Response
// @Router /api/products [get]
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

// ProductStore godoc
// @Summary Create a new product
// @Description Store a new product into the database
// @Tags Products
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param product body structs.Product true "Product Data"
// @Success 200 {object} structs.Response
// @Failure 401 {object} structs.Response
// @Failure 400 {object} structs.Response
// @Failure 422 {object} structs.Response
// @Failure 500 {object} structs.Response
// @Router /api/products [post]
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

// ProductFind godoc
// @Summary Find product by ID
// @Description Get a product detail by its ID
// @Tags Products
// @Security BearerAuth
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} structs.Response
// @Failure 401 {object} structs.Response
// @Failure 404 {object} structs.Response
// @Failure 500 {object} structs.Response
// @Router /api/products/{id} [get]
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

// ProductUpdate godoc
// @Summary Update a product
// @Description Update an existing product by ID
// @Tags Products
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body structs.Product true "Product Data"
// @Success 200 {object} structs.Response
// @Failure 401 {object} structs.Response
// @Failure 404 {object} structs.Response
// @Failure 422 {object} structs.Response
// @Failure 500 {object} structs.Response
// @Router /api/products/{id} [put]
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

// ProductDestroy godoc
// @Summary Delete a product
// @Description Delete a product by ID
// @Tags Products
// @Security BearerAuth
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} structs.Response
// @Failure 401 {object} structs.Response
// @Failure 404 {object} structs.Response
// @Failure 500 {object} structs.Response
// @Router /api/products/{id} [delete]
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