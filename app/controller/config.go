package controller

import (
	"log"
	"net/http"
	"scan/app/apiserver"
	"scan/app/model"
	"scan/app/service"
	"scan/app/store"
	_ "scan/docs"

	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Config struct {
	config  *apiserver.Config
	store   *store.Store
	service *service.Config
	router  *gin.Engine
}

// New controller
func New() (*Config, error) {
	config, err := apiserver.LoadConfig()
	if err != nil {
		return nil, err
	}

	db, err := apiserver.ConnectGorm(config.Dsn, config.LogLevel)
	if err != nil {
		return nil, err
	}

	store := store.New(db)

	c := &Config{
		config:  config,
		store:   store,
		service: service.New(store),
		router:  gin.Default(),
	}

	return c, nil
}

// New controller for test
func NewTest() (*Config, error) {
	config := apiserver.NewConfig()
	config.Dsn = "gen_user:0fgxqh8bc@tcp(85.193.83.246:3306)/default_db?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := apiserver.ConnectGorm(config.Dsn, config.LogLevel)
	if err != nil {
		return nil, err
	}

	store := store.New(db)

	c := &Config{
		config:  config,
		store:   store,
		service: service.New(store),
		router:  gin.Default(),
	}

	return c, nil
}

func (c *Config) SetUpRouters() *gin.Engine {
	c.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := c.router.Group("/api/v1")
	{
		// without user
		v1.GET("/locations", c.GetLocations)

		// auth User
		v1.Use(c.Auth())

		// "show_api" middleware
		v1.Use(c.ShowApiMiddleware())

		// "manager" middleware
		v1.Use(c.ManagerMiddleware())
		v1.POST("/scan", c.AddScan)
		v1.POST("/scan_batches", c.AddScanBatches)
		v1.POST("/vin", c.VinByPlate)
	}

	return c.router
}

func (c *Config) RunServer() error {
	return c.router.Run(c.Addr())
}

func (c *Config) Addr() string {
	return c.config.BindAddr
}

func (c *Config) initRequest(g *gin.Context, req any) (*model.User, error) {
	if err := g.BindJSON(req); err != nil {
		return nil, err
	}
	v := validate.Struct(req)
	if !v.Validate() {
		return nil, v.Errors
	}
	user, err := c.GetCurrentUser(g)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (c *Config) respond(g *gin.Context, obj any) {
	if obj == nil {
		obj = map[string]string{"message": "success"}
	}
	g.IndentedJSON(http.StatusOK, obj)
}

type ActionError struct {
	Error string `json:"error" example:"User not found"`
	Url   string `json:"url" example:"scan.e92.ru/api/v1/scan"`
}

func (c *Config) error(g *gin.Context, err error) {
	actionError := ActionError{
		Url:   g.Request.Host + g.Request.URL.Path,
		Error: err.Error(),
	}
	log.Print(actionError)

	g.IndentedJSON(http.StatusBadRequest, actionError)
}
