NAME   := containersol/k8sbhw-webapp
TAG    := $(shell git log -1 --pretty=%h)
IMG    := ${NAME}:${TAG}
LATEST := ${NAME}:latest

.PHONY=image push

image:
	docker build -t ${IMG} .
	docker tag ${IMG} ${LATEST}

push: image
	docker push ${NAME}
