package service

import (
	"log"

	"github.com/gin-gonic/gin"
)

type SaveImage interface {
	SaveImage(c *gin.Context)
}

type ImageService struct {
}

func NewService() *ImageService {
	return &ImageService{}
}

func (s *ImageService) SaveImage(c *gin.Context) {
	// Will be receiving the image as a form data
	form, _ := c.MultipartForm()
	file := form.File["Image"]
	title := form.Value["Title"]

	log.Println(title)
	log.Println(file)
}
