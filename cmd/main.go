package main

import (
	"github.com/sramirezpch/image-saver/internal/adapters"
	"github.com/sramirezpch/image-saver/internal/service"
)

func main() {
	imageService := service.NewService()

	router := adapters.NewRouter(imageService)

	router.AttachHandlers()

	router.Run()
}
