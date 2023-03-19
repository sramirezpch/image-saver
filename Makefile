IMAGE_NAME=nft-image-saver
DOCKERFILE_PATH=./infra/.dockerfile
COMPOSE_FILE=./docker-compose.yml

build-image:
	docker build -t $(IMAGE_NAME) -f $(DOCKERFILE_PATH) . --no-cache

run-with-docker:
	docker run -p 3001:8080 -d --name nft-image-saver nft-image-saver

down: 
	docker-compose down --remove-orphans

.PHONY: run
run:
	make down
	docker-compose --file $(COMPOSE_FILE) up -d --build
