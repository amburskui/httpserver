tag=v0.1

all: clean dist/httpserver

.PHONY: build build-docker

build: dist/httpserver

dist/httpserver: 
	go build -o dist/httpserver .

build-docker:
	docker build -t docker.io/amburskui/httpserver:${tag} .

push:
	docker push docker.io/amburskui/httpserver:${tag}

clean:
	rm dist/httpserver