package controller

import (
	"errors"
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

// ShowApiMiddleware
func (c *Config) ShowApiMiddleware() gin.HandlerFunc {
	return func(g *gin.Context) {
		user, err := c.GetCurrentUser(g)
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

// ManagerMiddleware
func (c *Config) ManagerMiddleware() gin.HandlerFunc {
	return func(g *gin.Context) {
		user, err := c.GetCurrentUser(g)
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

// GetCurrentUser from the gin reqeust
func (c *Config) GetCurrentUser(g *gin.Context) (*model.User, error) {
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
