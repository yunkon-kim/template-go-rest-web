package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/yunkon-kim/my-template/pkg/api/rest/model"
)

// [Note]
// No RequestBody required for "GET /users"

type GetUsersResponse struct {
	Users []model.MyUser `json:"users"`
}

// GetUsers godoc
// @Summary Get a list of users
// @Description Get information of all users.
// @Tags [Sample] Users
// @Accept  json
// @Produce  json
// @Success 200 {object} GetUsersResponse "(sample) This is a sample description for success response in Swagger UI"
// @Failure 404 {object} object "User Not Found"
// @Router /sample/users [get]
func GetUsers(c echo.Context) error {

	// In this example, hardcoded data is returned
	users := []model.MyUser{
		{ID: 1, Name: "John Doe", Email: "john@example.com"},
		{ID: 2, Name: "Anne Jacqueline Hathaway", Email: "Anne@example.com"},
		{ID: 3, Name: "Robert John Downey Jr.", Email: "Robert@example.com"},
	}
	return c.JSON(http.StatusOK, users)
}

// [Note]
// No RequestBody required for "GET /users"

// [Note]
// Struct Embedding is used to inherit the fields of MyUser
type GetUserResponse struct {
	model.MyUser
}

// GetUser godoc
// @Summary Get specific user information
// @Description Get information of a user with a specific ID.
// @Tags [Sample] Users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} GetUserResponse "(Sample) This is a sample description for success response in Swagger UI"
// @Failure 404 {object} object "User Not Found"
// @Router /sample/users/{id} [get]
func GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID format")
	}

	// Implement user retrieval logic (this is a simple example)
	if id != 1 {
		return c.JSON(http.StatusNotFound, "User not found")
	}

	// In this example, hardcoded data is returned
	user := model.MyUser{ID: 1, Name: "John Doe", Email: "john@example.com"}
	return c.JSON(http.StatusOK, user)
}

// [Note]
// Struct Embedding is used to inherit the fields of MyUser
type CreateUserRequest struct {
	model.MyUser
}

// [Note]
// Struct Embedding is used to inherit the fields of MyUser
type CreateUserResponse struct {
	model.MyUser
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the given information.
// @Tags [Sample] Users
// @Accept  json
// @Produce  json
// @Param User body CreateUserRequest true "User information"
// @Success 201 {object} GetUserResponse "(Sample) This is a sample description for success response in Swagger UI"
// @Failure 400 {object} object "Invalid Request"
// @Router /sample/users [post]
func CreateUser(c echo.Context) error {
	u := new(model.MyUser)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid Request")
	}

	// Implement user creation logic (this is a simple example)
	u.ID = 100 // Unique ID generation logic needed in actual implementation

	return c.JSON(http.StatusCreated, u)
}

// [Note]
// Struct Embedding is used to inherit the fields of MyUser
type UpdateUserRequest struct {
	model.MyUser
}

// [Note]
// Struct Embedding is used to inherit the fields of MyUser
type UpdateUserResponse struct {
	model.MyUser
}

// UpdateUser godoc
// @Summary Update a user
// @Description Update a user with the given information.
// @Tags [Sample] Users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param User body UpdateUserRequest true "User information to update"
// @Success 201 {object} UpdateUserResponse "(Sample) This is a sample description for success response in Swagger UI"
// @Failure 400 {object} object "Invalid Request"
// @Router /sample/users/{id} [put]
func UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID format")
	}

	u := new(model.MyUser)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	// Implement user update logic (this is a simple example)
	u.ID = id // Update the information of the user with the corresponding ID in the actual implementation

	return c.JSON(http.StatusOK, u)
}

// [Note]
// Struct Embedding is used to inherit the fields of MyUser
type PatchUserRequest struct {
	model.MyUser
}

// [Note]
// Struct Embedding is used to inherit the fields of MyUser
type PatchUserResponse struct {
	model.MyUser
}

// PatchUser godoc
// @Summary Patch a user
// @Description Patch a user with the given information.
// @Tags [Sample] Users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param User body PatchUserRequest true "User information to update"
// @Success 200 {object} PatchUserResponse "(Sample) This is a sample description for success response in Swagger UI"
// @Failure 400 {object} object "Invalid Request"
// @Failure 404 {object} object "User Not Found"
// @Router /sample/users/{id} [patch]
func PatchUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID format")
	}

	u := new(model.MyUser)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	// Implement user update logic (this is a simple example)
	u.ID = id // Update the information of the user with the corresponding ID in the actual implementation

	return c.JSON(http.StatusOK, u)
}

// [Note]
// No RequestBody required for "DELETE /users/{id}"

// [Note]
// No ResponseBody required for "DELETE /users/{id}"

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user with the given information.
// @Tags [Sample] Users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {string} string "User deletion successful"
// @Failure 400 {object} object "Invalid Request"
// @Failure 404 {object} object "User Not Found"
// @Router /sample/users/{id} [delete]
func DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID format")
	}

	// Implement user update logic (this is a simple example)
	// In this example, hardcoded data is returned
	user := model.MyUser{ID: 1, Name: "John Doe", Email: "john@example.com"}
	if id != user.ID {
		return c.JSON(http.StatusNotFound, "User not found")
	}

	return c.JSON(http.StatusOK, "User deletion successful")
}
