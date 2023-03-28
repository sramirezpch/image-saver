package main

import (
	"os"

	"github.com/sramirezpch/image-saver/internal/adapters"
	"github.com/sramirezpch/image-saver/internal/service"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func createAwsSession() (*session.Session, error) {
	creds := credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_SECRET_KEY"), "")

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: creds,
	})

	if err != nil {
		return nil, err
	}

	return sess, nil
}
func main() {
	sess, err := createAwsSession()
	if err != nil {
		panic(err)
	}

	s3Client := s3.New(sess, &aws.Config{Region: aws.String("us-east-1")})

	imageService := service.NewService(s3Client)

	router := adapters.NewRouter(imageService)

	router.AttachHandlers()

	router.Run()
}
