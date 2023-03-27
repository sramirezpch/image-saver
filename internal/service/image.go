package service

import "github.com/gin-gonic/gin"

type SaveImage interface {
	SaveImage(c *gin.Context)
}

type ImageService struct {
}

func (s *ImageService) SaveImage(c *gin.Context) {

}
