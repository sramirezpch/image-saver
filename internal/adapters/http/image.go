package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	s "github.com/sramirezpch/image-saver/internal/service"
)

type Routes struct {
	Method  string
	Path    string
	Handler func(c *gin.Context)
}

func CreateHandlers(r *gin.Engine, s *s.ImageService) {
	routes := []Routes{
		{Method: "GET", Path: "/", Handler: func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}},
		{Method: "POST", Path: "/image", Handler: s.SaveImage},
	}

	for i := 0; i < len(routes); i++ {
		r.Handle(routes[i].Method, routes[i].Path, routes[i].Handler)
	}
}
