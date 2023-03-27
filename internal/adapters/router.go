package adapters

import (
	"github.com/gin-gonic/gin"
	"github.com/sramirezpch/image-saver/internal/adapters/http"
	"github.com/sramirezpch/image-saver/internal/service"
)

type Router struct {
	Router       *gin.Engine
	ImageService *service.ImageService
}

func (router *Router) AttachHandlers() {
	http.CreateHandlers(router.Router, router.ImageService)
}

func (router *Router) Run() {
	if err := router.Router.Run(":8080"); err != nil {
		panic(err)
	}
}

func NewRouter(is *service.ImageService) *Router {
	r := gin.Default()
	return &Router{Router: r, ImageService: is}
}
