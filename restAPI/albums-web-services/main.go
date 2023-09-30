package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// album represents data about a record album
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// Write  a handler to return all items
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// post albums
func postAlbums(c *gin.Context) {
	var newAlbums album
	// call BinJSON to bin the recieved JOSN to
	//NewAlbum
	if err := c.BindJSON(&newAlbums); err != nil {
		return
	}
	albums = append(albums, newAlbums)
	c.IndentedJSON(http.StatusOK, newAlbums)
}

// get album by id
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	//Loop over the list of albums, looking for
	// find id matches the parameter
	for _, val := range albums {
		if val.ID == id {
			c.IndentedJSON(http.StatusOK, val)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})

}
func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.Run("localhost:8080")
}
