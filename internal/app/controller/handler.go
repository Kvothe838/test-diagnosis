package controller

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func New() *handler {
	h := &handler{
		router: gin.New(),
	}

	h.setupRouter()
	return h
}

type handler struct {
	router *gin.Engine
}

func (h *handler) setupRouter() {
	h.router.Use(cors.Default())
	h.router.Use(gin.Recovery())

	h.router.GET("/health", h.health)
}

func (h *handler) health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]string{
		"status": "healthy",
	})
}

func (h *handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	h.router.ServeHTTP(writer, request)
}
