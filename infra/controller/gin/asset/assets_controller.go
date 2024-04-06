package assets_controller

import (
	"assets_manager/infra/controller/gin/helpers"
	assets_use_cases "assets_manager/use_cases/assets"
	"assets_manager/utils/exception"
	"assets_manager/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetAssetRoutes(r *gin.Engine) {
	r.POST("/assets", createAsset)
	r.GET("/assets", listAssets)
	r.GET("/assets/:id", findAssetById)
	r.PUT("/assets/:id", updateAssetData)
	r.DELETE("/assets/:id", deleteAsset)
	r.PUT("/assets/:id/status", changeStatus)
	r.PUT("/assets/:id/user", setCurrentUser)
}

func createAsset(c *gin.Context) {
	var body assets_use_cases.CreateAssetDto

	if err := c.BindJSON(&body); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, exception.New(err.Error(), http.StatusInternalServerError))
		return
	}

	if err := assets_use_cases.CreateAsset(&body); err != nil {
		c.IndentedJSON(err.StatusCode, err)
		return
	}

	c.IndentedJSON(http.StatusOK, response.New("Asset created successfully"))
}

func listAssets(c *gin.Context) {
	q := helpers.GetQuery(c)

	a := assets_use_cases.FindAssets(q)

	c.IndentedJSON(http.StatusOK, a)
}

func findAssetById(c *gin.Context) {
	id, err := helpers.GetIdFromParams(c)

	if err != nil {
		c.IndentedJSON(err.StatusCode, err)
		return
	}

	a, err2 := assets_use_cases.FindAssetById(id)

	if err2 != nil {
		c.IndentedJSON(err2.StatusCode, err2)
		return
	}

	c.IndentedJSON(http.StatusOK, a)
}

func updateAssetData(c *gin.Context) {
	id, err := helpers.GetIdFromParams(c)

	if err != nil {
		c.IndentedJSON(err.StatusCode, err)
		return
	}

	var body assets_use_cases.UpdateAssetDto

	if err := c.BindJSON(&body); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, exception.New(err.Error(), http.StatusInternalServerError))
		return
	}

	err2 := assets_use_cases.UpdateAsset(id, &body)

	if err2 != nil {
		c.IndentedJSON(err2.StatusCode, err2)
		return
	}

	c.IndentedJSON(http.StatusOK, response.New("Asset updated successfully"))
}

func deleteAsset(c *gin.Context) {
	id, err := helpers.GetIdFromParams(c)

	if err != nil {
		c.IndentedJSON(err.StatusCode, err)
		return
	}

	if err := assets_use_cases.DeleteAsset(id); err != nil {
		c.IndentedJSON(err.StatusCode, err)
		return
	}

	c.IndentedJSON(http.StatusOK, response.New("Asset deleted successfully"))
}

func changeStatus(c *gin.Context) {
	id, err := helpers.GetIdFromParams(c)

	if err != nil {
		c.IndentedJSON(err.StatusCode, err)
		return
	}

	var body assets_use_cases.ChangeStatusDto

	if err := c.BindJSON(&body); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, exception.New(err.Error(), http.StatusInternalServerError))
		return
	}

	if err2 := assets_use_cases.ChangeStatus(id, &body); err2 != nil {
		c.IndentedJSON(err2.StatusCode, err2)
		return
	}

	c.IndentedJSON(http.StatusOK, response.New("Assets status updated successfully"))
}

func setCurrentUser(c *gin.Context) {
	id, err := helpers.GetIdFromParams(c)

	if err != nil {
		c.IndentedJSON(err.StatusCode, err)
		return
	}

	var body assets_use_cases.SetCurrentUserDto

	if err := c.BindJSON(&body); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, exception.New(err.Error(), http.StatusInternalServerError))
		return
	}

	if err2 := assets_use_cases.SetCurrentUser(id, &body); err2 != nil {
		c.IndentedJSON(err2.StatusCode, err2)
		return
	}

	c.IndentedJSON(http.StatusOK, response.New("Assets current user updated successfully"))
}
