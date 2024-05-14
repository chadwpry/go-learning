package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"rest/data"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

func main() {
	config := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}

	router := gin.Default()
	router.GET("/", healthCheck)
	router.GET("/albums", albums(config))
	router.GET("/albums/artist/:name", albumByArtists(config))
	router.GET("/albums/:id", albumById(config))

	router.Run("localhost:8080")
}

func healthCheck(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"OK": true})
}

func albums(config mysql.Config) func(*gin.Context) {
	return func(c *gin.Context) {
		albums, err := data.Albums(config)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Albums found %v\n", albums)

		c.IndentedJSON(http.StatusOK, albums)
	}
}

func albumByArtists(config mysql.Config) func(*gin.Context) {
	return func(c *gin.Context) {
		name := c.Param("name")

		albums, err := data.AlbumByArists(config, name)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Albums found %v\n", albums)

		c.IndentedJSON(http.StatusOK, albums)
	}
}

func albumById(config mysql.Config) func(*gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")

		album, err := data.AlbumByID(config, id)
		if err != nil {
			log.Fatal(err)
		}

		c.IndentedJSON(http.StatusOK, album)
	}
}
