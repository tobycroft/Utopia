package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	v1 "main.go/route/v1"
	"net/http"
)

func main() {
	//go BlockSync.Syncdata()
	route := gin.Default()
	route.Use(cors.Default())

	//gin.SetMode(gin.ReleaseMode)
	//gin.DefaultWriter = ioutil.Discard
	OnRoute(route)
	route.Run(":80")
}

func OnRoute(router *gin.Engine) {
	//router.LoadHTMLGlob("html/*")

	router.StaticFS("/html", http.Dir("./html"))
	//http.Handle("/html/", http.StripPrefix("/html/", fsh))
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
