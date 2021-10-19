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

.PHONY: help
## help: prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
