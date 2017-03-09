## OS checking
OS := $(shell uname)
ifeq ($(OS),Darwin)
	BUILD_OPTS=Darwin
else
	BUILD_OPTS=CGO_ENABLED=0 GOOS=linux GOARCH=amd64
endif


include ENV
export


default: run

run:
	@echo "🐳 $@ ${BUILD_OPTS}"
	go run main.go
