package user_controller

import (
	"assets_manager/infra/controller/gin/helpers"
	users_use_cases "assets_manager/use_cases/users"
	"assets_manager/utils/exception"
	"assets_manager/utils/response"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUserRoutes(r *gin.Engine) {
	r.POST("/users", createUser)
	r.GET("/users", listUsers)
	r.GET("/users/:id", findUserById)
	r.PUT("/users/:id", updateUserData)
	r.DELETE("/users/:id", deleteUser)
}

func createUser(c *gin.Context) {
	var body users_use_cases.CreateUserDto

	if err := c.BindJSON(&body); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, exception.New(err.Error(), http.StatusInternalServerError))
		return
	}

	if err := users_use_cases.CreateUser(&body); err != nil {
		c.IndentedJSON(err.StatusCode, err)
		return
	}

	c.IndentedJSON(http.StatusCreated, response.New("User created successfully"))
}

func listUsers(c *gin.Context) {
	query := helpers.GetQuery(c)

	fmt.Print(query)

	u := users_use_cases.FindUsers(query)

	c.IndentedJSON(http.StatusOK, u)
}

func findUserById(c *gin.Context) {
	id, err := helpers.GetIdFromParams(c)

	if err != nil {
		c.IndentedJSON(err.StatusCode, err)
		return
	}

	u, err2 := users_use_cases.FindUserById(id)

	if err2 != nil {
		c.IndentedJSON(err2.StatusCode, err2)
		return
	}

	c.IndentedJSON(http.StatusOK, u)
}

func updateUserData(c *gin.Context) {
	id, err := helpers.GetIdFromParams(c)

	if err != nil {
		c.IndentedJSON(err.StatusCode, err)
		return
	}

	var updateData users_use_cases.UpdateUserDto

	if err := c.BindJSON(&updateData); err != nil {
		c.IndentedJSON(http.StatusBadRequest, exception.New(err.Error(), http.StatusBadRequest))
		return
	}

	if err := users_use_cases.UpdateUser(id, &updateData); err != nil {
		c.IndentedJSON(err.StatusCode, err)
		return
	}

	c.IndentedJSON(http.StatusOK, response.New("Usuer updated successfully"))
}

func deleteUser(c *gin.Context) {
	id, err := helpers.GetIdFromParams(c)

	if err != nil {
		c.IndentedJSON(err.StatusCode, err)
		return
	}

	if err := users_use_cases.DeleteUser(id); err != nil {
		c.IndentedJSON(err.StatusCode, err)
		return
	}

	c.IndentedJSON(http.StatusOK, response.New("User deleted successfully"))
}
