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

func Login(c *gin.Context) {
	var user structs.User

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, structs.Response{
			Message: "Invalid JSON data",
			Error:   "Bad Request",
			Data:    nil,
		})
		return
	}

	err = repositories.CheckLogin(configs.DB, &user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, structs.Response{
			Message: "Invalid username or password",
			Error:   "Unathorized",
			Data:    nil,
		})
		return
	}

	token, _ := configs.GenerateJWT(&user)
	c.JSON(http.StatusOK, structs.Response{
		Message: "Login successful",
		Error:   nil,
		Data: gin.H{
			"token": token,
		},
	})
}

func UserIndex(c *gin.Context) {
	users, err := repositories.GetAllUser(configs.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Response{
			Message: "Internal server error",
			Error:   "Internal Server Error",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, structs.Response{
		Message: "Users retrieved succesfully",
		Error:   nil,
		Data: gin.H{
			"users": users,
		},
	})
}

func UserStore(c *gin.Context) {
	var user structs.User

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Response{
			Message: "Invalid JSON data",
			Error:   "Bad Request",
			Data:    nil,
		})
		return
	}

	validations := map[string]string{}
	if user.Name == "" {
		validations["name"] = "The name field is required"
	}
	if user.Gender == "" {
		validations["gender"] = "The gender field is required"
	}
	if user.Username == "" {
		validations["username"] = "The username field is required"
	}
	if user.Password == "" {
		validations["password"] = "The password field is required"
	}
	if user.Role == "" {
		validations["role"] = "The role field is required"
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

	var userResponse structs.UserResponse
	err = repositories.InsertUser(configs.DB, &user, &userResponse)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Response{
			Message: "Internal server error",
			Error:   "Internal Server Error",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, structs.Response {
		Message: "User inserted successfully",
		Error: nil,
		Data: gin.H{
			"user": userResponse,
		},
	})
}

func UserFind(c *gin.Context) {
	var user structs.User
	id, _ := strconv.Atoi(c.Param("id"))

	user.ID = id

	var userResponse structs.UserResponse
	err := repositories.GetUserById(configs.DB, &user, &userResponse)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, structs.Response{
			Message: fmt.Sprintf("User with id %d not found", id),
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
		Message: fmt.Sprintf("User with id %d successfully found", id),
		Error: nil,
		Data: gin.H{
			"user": userResponse,
		},
	})
}

func UserUpdate(c *gin.Context) {
	var user structs.User
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, structs.Response{
			Message: "Invalid JSON data",
			Error:   "Bad Request",
			Data:    nil,
		})
		return
	}

	validations := map[string]string{}
	if user.Name == "" {
		validations["name"] = "The name field is required"
	}
	if user.Gender == "" {
		validations["gender"] = "The gender field is required"
	}
	if user.Username == "" {
		validations["username"] = "The username field is required"
	}
	if user.Password == "" {
		validations["password"] = "The password field is required"
	}
	if user.Role == "" {
		validations["role"] = "The role field is required"
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

	user.ID = id

	var userResponse structs.UserResponse
	err = repositories.UpdateUser(configs.DB, &user, &userResponse)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, structs.Response{
			Message: fmt.Sprintf("User with id %d not found", id),
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

	c.JSON(http.StatusOK, structs.Response {
		Message: "User updated successfully",
		Error: nil,
		Data: gin.H{
			"user": userResponse,
		},
	})
}

func UserDestroy(c *gin.Context) {
	var user structs.User
	id, _ := strconv.Atoi(c.Param("id"))

	user.ID = id

	err := repositories.DeleteUser(configs.DB, &user)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, structs.Response{
			Message: fmt.Sprintf("User with id %d not found", id),
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
		Message: "User deleted successfully",
		Error: nil,
		Data: nil,
	})
}
