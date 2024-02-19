package handlerInterface

import "github.com/gin-gonic/gin"

type Handler interface {
	ShortenHandler(c *gin.Context)
	RedirectHandler(c *gin.Context)
	CountHandler(c *gin.Context)
}
