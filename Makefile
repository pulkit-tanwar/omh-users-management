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


.PHONY: dep run test cover local-cover 
