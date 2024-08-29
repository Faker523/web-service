package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{"1", "one", "zuo", 100},
	{"2", "two", "myl", 200},
	{"3", "three", "summer", 3000},
}

// var albums = []album{
// 	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
// 	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
// 	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// }

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// POST
func createAlbums(c *gin.Context) {
	var data album
	if err := c.BindJSON(&data); err != nil {
		return
	}
	albums = append(albums, data)
	c.IndentedJSON(http.StatusCreated, data)
}

// get by id
func getAlbumById(c *gin.Context) {
	id := c.Param("id")
	for _, album := range albums {
		if id == album.ID {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// update by id
func updateAlbumById(c *gin.Context) {
	id := c.Param("id")
	newPrice := 1000
	for _, album := range albums {
		if id == album.ID {
			album.Price = float64(newPrice)
			c.IndentedJSON(http.StatusOK, albums)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "id not found in albums"})
}

// delete by id
func deleteAlbumById(c *gin.Context) {
	id := c.Param("id")
	for i, album := range albums {
		if id == album.ID {
			albums = append(albums[:i], albums[i+1])
			c.IndentedJSON(http.StatusOK, albums)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "id not found"})
}

func main() {
	router := gin.Default()
	router.GET("/album", getAlbums)
	router.POST("/album", createAlbums)
	router.GET("/album/:id", getAlbumById)
	router.PUT("/album/:id", deleteAlbumById)
	router.PATCH("/album/:id", updateAlbumById)
	if err := router.Run("127.0.0.1:8000"); err != nil {
		log.Println("run server error")
	}
}
