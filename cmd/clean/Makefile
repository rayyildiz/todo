VERSION=$(shell git describe --tags --long --dirty 2>/dev/null)
TAG := latest

.PHONY: build docker

build:
	GO111MODULE=on GOOS=linux CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags "-w -X main.version=$(VERSION)" -o tmp/app main.go

docker:
	docker build -t eu.gcr.io/serverless-307509/clean:$(TAG) .

push: docker
	docker push eu.gcr.io/serverless-307509/clean:$(TAG)
