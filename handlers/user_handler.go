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

// Login godoc
// @Summary User login
// @Description Authenticate user and return JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body structs.Credentials true "User credentials"
// @Success 200 {object} structs.Response
// @Failure 400 {object} structs.Response
// @Failure 401 {object} structs.Response
// @Router /api/users/login [post]
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

// UserIndex godoc
// @Summary      Get all users
// @Description  Retrieve a list of all users
// @Tags         Users
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  structs.Response
// @Failure      500  {object}  structs.Response
// @Failure      401  {object}  structs.Response
// @Router       /api/users [get]
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

// UserStore godoc
// @Summary      Create a new user
// @Description  Store a new user into the database
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        user  body      structs.User  true  "User data"
// @Success      200   {object}  structs.Response
// @Failure      400   {object}  structs.Response
// @Failure      422   {object}  structs.Response
// @Failure      500   {object}  structs.Response
// @Failure      401   {object}  structs.Response
// @Router       /api/users [post]
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

// UserFind godoc
// @Summary      Find user by ID
// @Description  Retrieve a user by their ID
// @Tags         Users
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  structs.Response
// @Failure      404  {object}  structs.Response
// @Failure      500  {object}  structs.Response
// @Failure      401  {object}  structs.Response
// @Router       /api/users/{id} [get]
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

// UserUpdate godoc
// @Summary      Update a user
// @Description  Update user data by ID
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id    path      int           true  "User ID"
// @Param        user  body      structs.User  true  "Updated user data"
// @Success      200   {object}  structs.Response
// @Failure      400   {object}  structs.Response
// @Failure      404   {object}  structs.Response
// @Failure      422   {object}  structs.Response
// @Failure      500   {object}  structs.Response
// @Failure      401   {object}  structs.Response
// @Router       /api/users/{id} [put]
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

// UserDestroy godoc
// @Summary      Delete user
// @Description  Delete a user by their ID
// @Tags         Users
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  structs.Response
// @Failure      404  {object}  structs.Response
// @Failure      500  {object}  structs.Response
// @Failure      401  {object}  structs.Response
// @Router       /api/users/{id} [delete]
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
