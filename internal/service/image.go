package service

import (
	"log"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
)

type SaveImage interface {
	SaveImage(c *gin.Context)
}

type ImageService struct {
	s3Client *s3.S3
}

func NewService(s3 *s3.S3) *ImageService {
	return &ImageService{s3Client: s3}
}

func (s *ImageService) SaveImage(c *gin.Context) {
	form, _ := c.MultipartForm()
	file := form.File["Image"]
	title := form.Value["Title"]

	log.Println(title)
	log.Println(file)
}
