package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Album struct {
	ID     string   `json:"id"     gorm:"primaryKey"`
	Title  string   `json:"title"`
	Artist string   `json:"artist"`
	Price  *float32 `json:"price"`
}

type GinAppConfig struct {
	db *gorm.DB
}

func (ac *GinAppConfig) healthCheck(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"OK": true})
}

func (ac *GinAppConfig) list(c *gin.Context) {
	var albums []Album

	result := ac.db.Find(&albums)

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, result.Error)
		return
	}

	c.IndentedJSON(http.StatusOK, albums)
}

func (ac *GinAppConfig) albumByArtists(c *gin.Context) {
	name := c.Param("name")

	var albums []Album

	result := ac.db.Where("artist = ?", name).Find(&albums)

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, result.Error)
		return
	}

	c.IndentedJSON(http.StatusOK, albums)
}

func (ac *GinAppConfig) show(c *gin.Context) {
	id := c.Param("id")

	var album Album

	result := ac.db.Where("id = ?", id).First(&album)

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, result.Error)
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}

func (ac *GinAppConfig) create(c *gin.Context) {
	var album Album

	if err := c.BindJSON(&album); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		log.Fatalf("failed to bind values to album %v", err)
		return
	}

	album.ID = uuid.New().String()

	if result := ac.db.Create(&album); result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, result.Error)
		log.Fatalf("failed to create %v", result.Error)
		return
	}

	c.IndentedJSON(http.StatusCreated, album)
}

func main() {
	var routes GinAppConfig

	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		log.Fatalf("error opening database connection %v", err)
	}

	routes.db = db

	router := gin.Default()
	router.GET("/", routes.healthCheck)
	router.GET("/albums", routes.list)
	router.GET("/albums/artist/:name", routes.albumByArtists)
	router.GET("/albums/:id", routes.show)
	router.POST("/albums", routes.create)

	router.Run("0.0.0.0:8080")
}
