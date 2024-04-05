package main

import (
	user_controller "assets_manager/infra/controller/gin/user"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	user_controller.SetUserRoutes(r)

	r.Run("localhost:8000")
}
