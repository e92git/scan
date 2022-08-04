package main

import (
	"fmt"
	"log"
	"scan/app/controller"
	_ "scan/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"rsc.io/quote"
)

// @title           Дискаунтер автозачастей е92
// @version         1.0
// @description     Здесь представлены все методы для работы админстраторов и менеджеров магазинов.
// @description     Вопросы на info@e92.ru.

// @BasePath  /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
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
		// without user
		v1.GET("/locations", c.GetLocations)

		v1.Use(c.Auth())

		// user + show_api middleware
		v1.Use(c.ShowApiMiddleware())

		// user + show_api middleware
		v1.Use(c.ManagerMiddleware())
		v1.POST("/scan", c.AddScan)
		v1.POST("/vin", c.VinByPlate)

		// v1.GET("/scan", c.AddScanGet)
		// v1.GET("/users/:id", apis.GetUser)
	}

	err = r.Run(c.Addr())
	if err != nil {
		log.Fatal(err)
	}
}

// // getAlbums responds with the list of all albums as JSON.
// func getAlbums(c *gin.Context) {
// 	var newAlbum album
// 	newAlbum.Artist = "dd"
// 	albums = append(albums, newAlbum)
// 	c.IndentedJSON(http.StatusOK, albums)
// }

// // postAlbums adds an album from JSON received in the request body.
// func postAlbums(c *gin.Context) {
// 	var newAlbum album

// 	// Call BindJSON to bind the received JSON to
// 	// newAlbum.
// 	if err := c.BindJSON(&newAlbum); err != nil {
// 		return
// 	}

// 	// Add the new album to the slice.
// 	albums = append(albums, newAlbum)
// 	c.IndentedJSON(http.StatusCreated, newAlbum)
// }

// // getAlbumByID locates the album whose ID value matches the id
// // parameter sent by the client, then returns that album as a response.
// func getAlbumByID(c *gin.Context) {
// 	id := c.Param("id")
// 	// plate := c.Query("location")
// 	// datetime := c.Query("datetime")
// 	// image := c.Query("image")

// 	// Loop through the list of albums, looking for
// 	// an album whose ID value matches the parameter.
// 	for _, a := range albums {
// 		if a.ID == id {
// 			c.IndentedJSON(http.StatusOK, a)
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
// }
