package server

import (
	"github.com/gin-gonic/gin"
)

type server struct {
	*gin.Engine
}

func New() *server{
    router := &server{gin.New()}
	router.Use(gin.Recovery())

	return router
}

func (self *server) EnableCors(config CorsConfig) {
	self.Use(corsMiddleware(config))
}
