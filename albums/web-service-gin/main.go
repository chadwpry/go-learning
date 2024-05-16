package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"rest/data"
)

type GinAppConfig struct{}

func (ac *GinAppConfig) healthCheck(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"OK": true})
}

func (ac *GinAppConfig) albums(c *gin.Context) {
	albums, err := data.Albums()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Albums found %v\n", albums)

	c.IndentedJSON(http.StatusOK, albums)
}

func (ac *GinAppConfig) albumByArtists(c *gin.Context) {
	name := c.Param("name")

	albums, err := data.AlbumByArists(name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Albums found %v\n", albums)

	c.IndentedJSON(http.StatusOK, albums)
}

func (ac *GinAppConfig) albumById(c *gin.Context) {
	id := c.Param("id")

	album, err := data.AlbumByID(id)
	if err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, album)
}

func main() {
	var routes GinAppConfig

	router := gin.Default()
	router.GET("/", routes.healthCheck)
	router.GET("/albums", routes.albums)
	router.GET("/albums/artist/:name", routes.albumByArtists)
	router.GET("/albums/:id", routes.albumById)

	router.Run("0.0.0.0:8080")
}
