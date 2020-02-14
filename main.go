package main

import (
	"github.com/gin-gonic/gin"
	"main.go/extend/CosMos/BlockSync"
	v1 "main.go/route/v1"
)

func main() {
	go BlockSync.Syncdata()
	route := gin.Default()
	//gin.SetMode(gin.ReleaseMode)
	//gin.DefaultWriter = ioutil.Discard
	OnRoute(route)
	route.Run(":80")
}

func OnRoute(router *gin.Engine) {
	router.Any("/", func(context *gin.Context) {
		context.String(0, router.BasePath())
	})
	version1 := router.Group("/v1")
	{
		version1.Any("/", func(context *gin.Context) {
			context.String(0, version1.BasePath())
		})
		index := version1.Group("/index")
		{
			v1.IndexRouter(index)
		}
		cosmos := version1.Group("/cosmos")
		{
			v1.CosmosRouter(cosmos)
		}
	}
}
