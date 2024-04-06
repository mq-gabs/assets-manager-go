package main

import (
	assets_controller "assets_manager/infra/controller/gin/asset"
	group_controller "assets_manager/infra/controller/gin/group"
	user_controller "assets_manager/infra/controller/gin/user"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	user_controller.SetUserRoutes(r)
	group_controller.SetGroupRoutes(r)
	assets_controller.SetAssetRoutes(r)

	r.Run(":8080")
}
