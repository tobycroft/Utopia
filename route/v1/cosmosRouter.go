package v1

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/cosmos/controller"
)

func CosmosRouter(route *gin.RouterGroup) {
	route.Any("/", func(context *gin.Context) {
		context.String(0, route.BasePath())
	})

	controller.IndexController(route.Group("/index"))
}
