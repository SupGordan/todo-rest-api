MODULE = $(shell go list -m)


build:  ## build the API server binary
	CGO_ENABLED=0 go build -o server
	
build-docker: ## build the API server as a docker image
	docker build -f docker/Dockerfile -t server .