tag := $(shell git describe --tags --abbrev=0)

.PHONY: all

all: clean build

.PHONY: build docker-build docker-push

build: dist/httpserver

dist/httpserver: 
	go build -o dist/httpserver ./cmd/httpserver

docker-build:
	docker build --platform linux/amd64 -t docker.io/amburskui/httpserver:${tag} .

docker-push: docker-build
	docker push docker.io/amburskui/httpserver:${tag}

clean:
	@rm -f dist/httpserver

kube-apply:
	minikube kubectl -- apply -f scripts/kube/

kube-delete:
	minikube kubectl -- delete -f scripts/kube/