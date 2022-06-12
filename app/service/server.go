package service

import (
	"example/hello/app/store"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router *gin.Engine
	Store  *store.Store
}