.PHONE: run build tidy test clean help

MAIN := cmd/main.go
BUILD_FOLDER := bin/server
ENV ?= dev
-include .env.$(ENV)
export

run:	## Run project in dev environment
	go run $(MAIN)

run-staging:	## Run project in staging environment
	make run ENV=staging

run-prod:	## Run project in production environment
	make run ENV=prod

build: tidy test	## Build project with tidy and test
	go build -o $(BUILD_FOLDER) $(MAIN)

tidy:	## Add missing imports and remove unused libaries
	go mod tidy

test:	## Run tests
	go test ./... -v -cover

$(BUILD_FOLDER): $(MAIN)
	go build -o $@ $<

clean:	## Clear build artifacts and test coverage report
	rm -rf bin/
	rm -f coverage.out

help:	## Show all available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
	sort | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'