DOCKER_REGISTRY = hello-image

all: 

package:
	go build -o myapp main.go
	docker build -t "$(DOCKER_REGISTRY)" .
.PHONY:package

