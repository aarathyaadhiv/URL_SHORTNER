package api

import (
	"github.com/gin-gonic/gin"
	handlerInterface "github.com/url-shortner/pkg/api/handler/interface"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(handler handlerInterface.Handler) *ServerHTTP {
	server := gin.Default()
	server.POST("/shorten", handler.ShortenHandler)
	server.GET("/short/:shorten",handler.RedirectHandler)
	server.GET("/count",handler.CountHandler)

	return &ServerHTTP{engine: server}

}

func (s *ServerHTTP) Start() {
	s.engine.Run(":3000")
}
