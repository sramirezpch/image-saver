package main

import (
	"context"
	"os"

	"github.com/sramirezpch/image-saver/internal/adapters"
	"github.com/sramirezpch/image-saver/internal/service"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion("us-east-1"), config.WithCredentialsProvider(
		credentials.NewStaticCredentialsProvider(os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_SECRET_KEY"), "")))
	if err != nil {
		panic(err)
	}

	s3Client := s3.NewFromConfig(cfg)

	imageService := service.NewService(s3Client)

	router := adapters.NewRouter(imageService)

	router.AttachHandlers()

	router.Run()
}
