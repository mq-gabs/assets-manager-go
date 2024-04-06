package group_controller

import (
	"assets_manager/infra/controller/gin/helpers"
	groups_use_cases "assets_manager/use_cases/groups"
	"assets_manager/utils/exception"
	"assets_manager/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetGroupRoutes(r *gin.Engine) {
	r.POST("/groups", createGroup)
	r.GET("/groups", listGroups)
	r.GET("/groups/:id", findGroupById)
	r.PUT("/groups/:id", updateGroupData)
	r.DELETE("/groups/:id", deleteGroup)

}

func createGroup(c *gin.Context) {
	var body groups_use_cases.CreateGroupDto

	if err := c.BindJSON(&body); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, exception.New(err.Error(), http.StatusInternalServerError))
		return
	}

	if err := groups_use_cases.CreateGroup(&body); err != nil {
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
		c.IndentedJSON(err2.StatusCode, err2)
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

	var body groups_use_cases.UpdateGroupDto

	if err := c.BindJSON(&body); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, exception.New(err.Error(), http.StatusInternalServerError))
		return
	}

	if err := groups_use_cases.UpdateGroup(id, &body); err != nil {
		c.IndentedJSON(err.StatusCode, err)
		return
	}

	c.IndentedJSON(http.StatusOK, response.New("Group updated successfully"))
}

func deleteGroup(c *gin.Context) {
	id, err := helpers.GetIdFromParams(c)

	if err != nil {
		c.IndentedJSON(err.StatusCode, err)
		return
	}

	if err := groups_use_cases.DeleteGroup(id); err != nil {
		c.IndentedJSON(err.StatusCode, err)
		return
	}

	c.IndentedJSON(http.StatusOK, response.New("Group deleted successfully"))
}
