package controller

import (
	"example/hello/app/service"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router  *gin.Engine
	Service *service.Server
}
