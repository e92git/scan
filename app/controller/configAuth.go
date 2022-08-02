package controller

import (
	"errors"
	"fmt"
	"scan/app/model"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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
			c.error(g, errors.New("Authorization is invalid. User not found: session="+authHeader))
			g.Abort()
			return
		}
		if err != nil {
			c.error(g, err)
			g.Abort()
			return
		}
		g.Set("user", user)
		g.Next()
	}
}

// Middleware ShowApi
func (c *Config) ShowApiMiddleware() gin.HandlerFunc {
	return func(g *gin.Context) {
		user, err := c.GetUser(g)
		if err != nil {
			c.error(g, err)
			g.Abort()
			return
		}
		if !user.HasShowApi() {
			c.error(g, errors.New("Authorization is invalid. The user does not have enough rights id="+strconv.FormatInt(user.ID, 10)))
			g.Abort()
			return
		}
		g.Next()
	}
}

// Middleware Manager
func (c *Config) ManagerMiddleware() gin.HandlerFunc {
	return func(g *gin.Context) {
		user, err := c.GetUser(g)
		if err != nil {
			c.error(g, err)
			g.Abort()
			return
		}
		if !user.HasManager() {
			c.error(g, errors.New("Authorization is invalid. The user does not have enough rights id="+strconv.FormatInt(user.ID, 10)))
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
	fmt.Println("---GetUser---")
	return u, nil
}