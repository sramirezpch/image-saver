package service

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/gin-gonic/gin"
)

type SaveImage interface {
	SaveImage(c *gin.Context)
}

type ImageService struct {
	s3Client *s3.Client
}

func NewService(s3 *s3.Client) *ImageService {
	return &ImageService{s3Client: s3}
}

func (s *ImageService) SaveImage(c *gin.Context) {
	form, _ := c.MultipartForm()
	file := form.File["Image"]
	title := form.Value["Title"]
	contentType := form.Value["FileType"]

	reader, err := file[0].Open()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer reader.Close()

	cont, err := ioutil.ReadAll(reader)
	buf := bytes.NewBuffer(cont)

	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not read the file, please check the logs"})
	}

	resp, err := s.s3Client.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket:      aws.String("nft-real-estate-images"),
		Key:         aws.String(title[0]),
		Body:        buf,
		ACL:         types.ObjectCannedACLPublicRead,
		ContentType: aws.String(contentType[0]),
	})
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not upload the image, please check the logs"})
	}
	log.Println(resp)
}
