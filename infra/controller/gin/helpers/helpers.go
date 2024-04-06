package helpers

import (
	"assets_manager/utils/exception"
	"assets_manager/utils/query"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetIdFromParams(c *gin.Context) (uint16, *exception.Exception) {
	id, found := c.Params.Get("id")

	if !found {
		return 0, exception.New("Id not found from params", http.StatusBadRequest)
	}

	new_id, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		return 0, exception.New("Id must be uint16", http.StatusBadRequest)
	}

	return uint16(new_id), nil
}

func GetQuery(c *gin.Context) *query.IQuery {
	var q query.IQuery

	page, foundPage := c.GetQuery("page")

	if foundPage {
		q.Page = page
	}

	pageSize, foundPageSize := c.GetQuery("pageSize")

	if foundPageSize {
		q.PageSize = pageSize
	}

	name, foundName := c.GetQuery("name")

	if foundName {
		q.Name = name
	}

	return &q
}
