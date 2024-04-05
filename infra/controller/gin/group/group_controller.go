package group_controller

import (
	"assets_manager/domain/entities/group"
	"assets_manager/infra/controller/gin/helpers"
	groups_use_cases "assets_manager/use_cases/groups"
	"assets_manager/utils/exception"
	"assets_manager/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetGroupRoutes(r *gin.Engine) {
	r.GET("/groups", createGroup)

}

func createGroup(c *gin.Context) {
	var body group.Group

	if err := c.BindJSON(&body); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, exception.New(err.Error(), http.StatusInternalServerError))
		return
	}

	if err := groups_use_cases.CreateGroup(body.Name); err != nil {
		c.IndentedJSON(err.StatusCode, err)
		return
	}

	c.IndentedJSON(http.StatusCreated, response.New("Group created successfully"))
}

func listGroups(c *gin.Context) {
	var q = helpers.GetQuery(c)

	g := groups_use_cases.FindGroups(q)

	c.IndentedJSON(http.StatusOK, g)
}

func findGroupById(c *gin.Context) {
	id, err := helpers.GetIdFromParams(c)

	if err != nil {
		c.IndentedJSON(err.StatusCode, err)
		return
	}

	g, err2 := groups_use_cases.FindGroupById(id)

	if err2 != nil {
		c.IndentedJSON(err.StatusCode, err)
		return
	}

	c.IndentedJSON(http.StatusOK, g)
}

func updateGroupData(c *gin.Context) {
	id, err := helpers.GetIdFromParams(c)

	if err != nil {
		c.IndentedJSON(err.StatusCode, err)
		return
	}

	var body group.Group

	if err := c.BindJSON(&body); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, exception.New(err.Error(), http.StatusInternalServerError))
	}
}
