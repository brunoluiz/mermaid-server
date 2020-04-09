DOCKER_REPO:=brunoluiz
PROJECT_NAME:=mermaid-server

## Prepare version tags
GIT_BRANCH?=$(shell git rev-parse --abbrev-ref HEAD)
VERSION:=$(shell git rev-parse HEAD)
DOCKER_TAG:=$(subst :,-,$(subst /,-,$(GIT_BRANCH)))
ifeq ($(GIT_BRANCH), master)
	DOCKER_TAG := latest
endif

#
# Golang tooling
#
.PHONY: test
test:
	go test -race ./...

lint:
	bin/golangci-lint run --deadline=2m ./...

build:
	go build -o ./bin/go-pwa-server ./cmd

LINTER_VERSION ?= 1.21.0
install-tools:
	@if [ "`bin/golangci-lint --version | awk '{print $$4}'`" != $(LINTER_VERSION) ]; then \
		wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v$(LINTER_VERSION); fi

#
# Publishing tooling
#
publish-test:
	goreleaser --snapshot --skip-publish --rm-dist

#
# Docker tooling
#
docker-login:
	docker login --username $(DOCKER_HUB_USER) --password=$(DOCKER_HUB_PASSWORD)

docker-build:
	docker build -t $(DOCKER_REPO)/$(PROJECT_NAME):local .

docker-push:
	docker push $(DOCKER_REPO)/$(PROJECT_NAME)
