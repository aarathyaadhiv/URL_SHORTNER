package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	handlerInterface "github.com/url-shortner/pkg/api/handler/interface"
	"github.com/url-shortner/pkg/config"
	interfaceUsecase "github.com/url-shortner/pkg/usecase/interface"
)

type Handler struct {
	Usecase interfaceUsecase.Usecase
	config  config.Config
}

func NewHandler(usecase interfaceUsecase.Usecase, config config.Config) handlerInterface.Handler {
	return &Handler{Usecase: usecase, config: config}
}

func (h *Handler) ShortenHandler(c *gin.Context) {
	url := c.Request.FormValue("url")
	shorten, err := h.Usecase.Shorten(url)
	if err != nil {

		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, shorten)
}

func (h *Handler) RedirectHandler(c *gin.Context) {

	short := c.Param("shorten")
	if short == "" {
		errRes := errors.New("please provide correct url")
		c.JSON(http.StatusBadRequest, errRes)
	}
	url, err := h.Usecase.Redirect(short)
	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{err.Error(): err})
		return
	}
	c.Redirect(http.StatusMovedPermanently, url)
}

func (h *Handler) CountHandler(c *gin.Context) {
	url := c.Request.FormValue("url")

	prefix := h.config.BASE_URL + "/short/"

	shorten := url[len(prefix):]

	count, err := h.Usecase.Count(shorten)
	if err != nil {

		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, count)
}
