GIT_TAG := $(shell git describe --tags --abbrev=0)
MIGRATE_VERSION := v4.15.0

.PHONY: all

all: clean dist/httpserver

dist/httpserver: 
	go build -o dist/httpserver ./cmd/httpserver

docker-build: docker-build-httpserver docker-build-migrations

docker-build-httpserver:
	docker build --platform linux/amd64 -t docker.io/amburskui/httpserver:${GIT_TAG} -f build/docker/httpserver.Dockerfile . 

docker-build-migrations:
	docker build --platform linux/amd64 -t docker.io/amburskui/httpserver-migrations:${GIT_TAG} -f build/docker/migrations.Dockerfile .

docker-push: docker-push-httpserver docker-push-migrations

docker-push-httpserver:
	docker push docker.io/amburskui/httpserver:${GIT_TAG}

docker-push-migrations:
	docker push docker.io/amburskui/httpserver-migrations:${GIT_TAG}

clean:
	@rm -f dist/httpserver

kube-apply:
	kubectl apply -f deployments/kube/

kube-delete:
	kubectl delete -f deployments/kube/

helm-install:
	helm install httpserver-chart deployments/helm/httpserver-chart

helm-uninstall:
	helm uninstall httpserver-chart  

tools/bin/migrate:
	GOBIN=$(shell pwd)/tools/bin go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@$(MIGRATE_VERSION)