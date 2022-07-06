package main

import (
	"fmt"
	"log"
	"net/http"
	"scan/app/controller"
	_ "scan/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func main() {
	fmt.Println(quote.Hello())

	c, err := controller.New()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{
		// v1.Use(auth())
		v1.GET("/locations", c.GetLocations)
		v1.POST("/scan", c.AddScans)
		v1.GET("/scan", c.AddScan)
		// v1.GET("/users/:id", apis.GetUser)
	}

	r.GET("/albums", getAlbums)
	r.GET("/albums/:id", getAlbumByID)
	r.POST("/albums", postAlbums)

	err = r.Run(c.Addr())
	if err != nil {
		log.Fatal(err)
	}
}

// func auth() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		authHeader := c.GetHeader("Authorization")
// 		if len(authHeader) == 0 {
// 			httputil.NewError(c, http.StatusUnauthorized, errors.New("Authorization is required Header"))
// 			c.Abort()
// 		}
// 		if authHeader != config.Config.ApiKey {
// 			httputil.NewError(c, http.StatusUnauthorized, fmt.Errorf("this user isn't authorized to this operation: api_key=%s", authHeader))
// 			c.Abort()
// 		}
// 		c.Next()
// 	}
// }

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
