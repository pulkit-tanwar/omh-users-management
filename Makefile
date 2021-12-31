SERVICE ?= omh-users-management
VERSION      ?= $(shell cat version | head -n 1)
REVISION     ?= $(shell git rev-parse --short HEAD)
BUILD_NUMBER ?= none
DOCKER_TAG ?= latest
COVERAGE_REPORT_SERVER_PORT ?= 3001
COVERPROFILE=.cover.out
COVERDIR=.cover

dep:
	@go get ./...

run: 
	@go run main.go serve

test: 
	$(eval export CONFIG_FILE_NAME=config.test.properties.json)
	@go test -coverprofile=$(COVERPROFILE) ./...

cover: test
	@mkdir -p $(COVERDIR)
	@go tool cover -html=$(COVERPROFILE) -o $(COVERDIR)/index.html
	@cd $(COVERDIR) && python -m SimpleHTTPServer $(COVERAGE_REPORT_SERVER_PORT)

local-cover: test
	@mkdir -p $(COVERDIR)
	@go tool cover -html=$(COVERPROFILE) -o $(COVERDIR)/index.html

clean:
	@rm -rf bin

build: clean
	@echo Compiling version: $(VERSION), revision: $(REVISION), build: $(BUILD_NUMBER)
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a --installsuffix cgo \
		-o bin/$(SERVICE) main.go	

image: build
	docker build -t $(SERVICE):$(DOCKER_TAG) -f build/docker/dockerfile .		


.PHONY: dep run test cover local-cover clean build image
