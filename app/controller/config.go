package controller

import (
	"errors"
	"log"
	"net/http"
	"scan/app/apiserver"
	"scan/app/model"
	"scan/app/service"
	"scan/app/store"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Config struct {
	config  *apiserver.Config
	store   *store.Store
	service *service.Config
}

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
	}

	return c, nil
}

// auth user 
func (c *Config) Auth() gin.HandlerFunc {
	return func(g *gin.Context) {
		authHeader := g.GetHeader("Authorization")
		if len(authHeader) == 0 {
			c.error(g, errors.New("Authorization field is required in Header"))
			g.Abort()
			return
		}
		user, err := c.service.User().FindBySession(authHeader)
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			c.error(g, errors.New("Authorization is invalid. User not found: session=" + authHeader))
			g.Abort()
			return
		}
		if err != nil {
			c.error(g, err)
			g.Abort()
			return
		}
		g.Set("use2r", user)
		g.Set("user", model.UserRoles)
		g.Next()
	}
}

// Middleware ShowApi 
func (c *Config) MiddlewareShowApi() gin.HandlerFunc {
	return func(g *gin.Context) {
		user, err := c.GetUser(g)
		if err != nil {
			c.error(g, err)
			g.Abort()
			return
		}
		if !user.IsMiddlewareApi() {
			c.error(g, errors.New("Authorization is invalid. The user does not have enough rights id=" + strconv.FormatInt(user.ID, 10)))
			g.Abort()
			return
		}
		g.Next()
	}
}

// GetUser from the gin reqeust
func (c *Config) GetUser(g *gin.Context) (*model.User, error) {
	user, found := g.Get("user")
	if found == false {
		return nil, errors.New("User not found in GetUser")
	}
	u, convert := user.(*model.User)
	if convert == false {
		return nil, errors.New("User not convert in GetUser")
	}
	return u, nil
}

func (c *Config) Addr() string {
	return c.config.BindAddr
}

func (c *Config) respond(g *gin.Context, obj any) {
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

