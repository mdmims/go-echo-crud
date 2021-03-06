APP=goapi


.PHONY: build
## build: build the application
build: clean
	go build -o ${APP}

.PHONY: run
## run: runs application via main.go
run:
	go run -race main.go

.PHONY: clean
## clean: cleans binary
clean:
	go clean

.PHONY: build-swagger
## build-swagger: builds swagger documentation in /docs
build-swagger:
	swag init

.PHONY: build-mocks
## build-mocks: builds interface mocks via mockgen cli
build-mocks:
	mockgen -source=repository/interfaces.go -destination=mocks/items.go -package=mocks

.PHONY: unit
## unit: run unit tests
unit:
	docker compose -f docker-compose.yml run --rm goapi go test ./...

.PHONY: help
## help: prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
