package main

import (
	"example/hello/app/apiserver"
	"example/hello/app/service"
	"example/hello/app/store"
	"fmt"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"rsc.io/quote"
)

// http://localhost:5555/scan?place=pokrovka&plate={plate}&datetime={datetime}&direction={direction}&image={image}


// album represents data about a record album.
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist,omitempty"`
    Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	fmt.Println(quote.Hello())

    config := apiserver.NewConfig()
	_, err := toml.DecodeFile(".env", config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(config.ApiKey)

    db, err := apiserver.ConnectDb(config.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
    
    s := service.Server{
		Router:       gin.Default(),
		Store:        store.New(db),
	}

    s.Router.GET("/scan", s.AddScan)
    s.Router.GET("/albums", getAlbums)
    s.Router.GET("/albums/:id", getAlbumByID)
    s.Router.POST("/albums", postAlbums)

    s.Router.Run("localhost:5555")
}


// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
    var newAlbum album
    newAlbum.Artist = "dd"
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
    var newAlbum album

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    // Add the new album to the slice.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
    id := c.Param("id")
    // plate := c.Query("location")
    // datetime := c.Query("datetime")
    // image := c.Query("image")

    // Loop through the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}