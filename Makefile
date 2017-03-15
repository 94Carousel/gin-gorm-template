.PHONY: run go build fmt
## OS checking
OS := $(shell uname)
ifeq ($(OS),Darwin)
	BUILD_OPTS=
else
	BUILD_OPTS=CGO_ENABLED=0 GOOS=linux GOARCH=amd64
endif

VERSION=$(shell git describe --always --tags)
BUILD_TIME=$(shell date -u +%Y-%m-%d:%H-%M-%S)
GO_LDFLAGS=-ldflags "-X `go list ./version`.Version=$(VERSION) -X `go list ./version`.BuildTime=$(BUILD_TIME)"

include ENV
export

## Hotreload checking
FRESH := $(shell which fresh 2> /dev/null)
ifdef FRESH
	RUN_COMMAND=fresh
else
	RUN_COMMAND=go run main.go
endif

default: run

run:
	@echo "🐳 $@ Running Web Server with ${RUN_COMMAND} 🐳"
	${RUN_COMMAND}

go:
	@echo "🐳 $@ Running Web Server Using Go 🐳"
	go run main.go


build: fmt
	${BUILD_OPTS} go build ${GO_LDFLAGS} -v -o ./web main.go

fmt:
	go fmt .
