NAME   := riccardomc/k8sbhw-webapp
TAG    := $(shell git log -1 --pretty=%h)
IMG    := ${NAME}:${TAG}
LATEST := ${NAME}:latest

SOURCES=$(shell find . -name '*.go')
.PHONY=image push test

webapp: $(SOURCES) 
	go test -v ./...
	CGO_ENABLED=0 GOOS=linux go build -o $@ -a -ldflags '-extldflags "-static"' .

image: webapp
	docker build -t ${IMG} .
	docker tag ${IMG} ${LATEST}

push: webapp image
	docker push ${NAME}

test:
	go test -v ./...
