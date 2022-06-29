package main

import (
	"example/rest-api/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	webEngine := gin.Default()

	registerEndpoints(webEngine)

	webEngine.Run("localhost:8080")
}

func registerEndpoints(webEngine *gin.Engine) {
	webEngine.GET("/albums", controller.GetAlbums)
	webEngine.GET("/albums/:id", controller.GetAlbumByID)
	webEngine.POST("/albums", controller.PostAlbums)
}
