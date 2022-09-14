package controller

import (
	"log"
	"net/http"
	"net/http/httptest"
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
	server  *gin.Engine
}

// New controller
func New(config *apiserver.Config, store *store.Store, service *service.Config) *Config {
	c := &Config{
		config:  config,
		store:   store,
		service: service,
		server:  gin.Default(),
	}

	c.SetUpRouters()

	return c
}

func (c *Config) SetUpRouters() *gin.Engine {
	c.server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := c.server.Group("/api/v1")
	{
		// without user
		v1.GET("/locations", c.GetLocations)

		// auth User
		v1.Use(c.Auth())

		// "show_api" middleware
		v1.Use(c.ShowApiMiddleware())
		v1.GET("/tire/analytics", c.GetTireAnalytics)

		// "manager" middleware
		v1.Use(c.ManagerMiddleware())
		v1.POST("/scan", c.AddScan)
		v1.POST("/scan/bulk", c.AddScanBulk)
		v1.POST("/vin", c.VinByPlate)
		v1.POST("/vin/bulk", c.VinByPlateBulk)
	}

	return c.server
}

func (c *Config) RunServer() error {
	return c.server.Run(c.Addr())
}

func (c *Config) Addr() string {
	return c.config.BindAddr
}

func (c *Config) GetService() *service.Config {
	return c.service
}

func (c *Config) GetConfig() *apiserver.Config {
	return c.config
}

func (c *Config) testRequest(req *http.Request) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	c.server.ServeHTTP(w, req)
	return w
}

// initRequest получить юзера авторизованного и тело запроса в req
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
