PWD = $(shell pwd)

lint:
	docker run --rm -v $(PWD):/app -w /app golangci/golangci-lint golangci-lint run -v

up:
	docker-compose -f "deployments/docker-compose.yml" up -d 

down:
	docker-compose -f "deployments/docker-compose.yml" down