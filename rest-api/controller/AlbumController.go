package controller

import (
	"example/rest-api/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

var albums = []dto.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func PostAlbums(c *gin.Context) {
	var newAlbum dto.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
