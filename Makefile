## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

# ========================================================================= #
# DEVELOPMENT
# ========================================================================= #

## run/web: run the cmd/webserver application
.PHONY: run/webserver
run/webserver:
	go run ./cmd/webserver

## run/cli: run the cmd/cli application
.PHONY: run/cli
run/cli:
	go run ./cmd/cli

# ========================================================================= #
# QUALITY CONTROL
# ========================================================================= #

## tidy: tidy module dependencies and format all .go files
.PHONY: tidy
tidy:
	@echo 'Tidying module dependencies...'
	go mod tidy
	@echo 'Formatting .go files...'
	go fmt ./...

## test: test the application
.PHONY: test
test:
	go test .

# ==================================================================================== #
# BUILD
# ==================================================================================== #

## build/webserver: build the cmd/webserver application
.PHONY: build/webserver
build/webserver:
	@echo 'Building cmd/webserver...'
	go build -o=./bin/webserver ./cmd/webserver

## build/cli: build the cmd/cli application
.PHONY: build/webserver
build/cli:
	@echo 'Building cmd/cli...'
	go build -o=./bin/cli ./cmd/cli
