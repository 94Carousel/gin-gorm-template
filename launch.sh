#!/bin/bash

# for env in `cat ENV`
# do
#   export $env
# done

function Run() {
  if [[ `which fresh` ]] ; then
    echo "🐳  Running Web Server with fresh 🐳"
    fresh
  else
    echo "🐳  Running Web Server with Golang 🐳"
    go run main.go
  fi
}

function Build() {
  go fmt ./
  OS=`uname`
  if [[ $OS == "Darwin" ]]; then
    BUILD_OPTS=
  else
    BUILD_OPTS=CGO_ENABLED=0 GOOS=linux GOARCH=amd64
  fi
  VERSION=`git describe --always --tags`
  BUILD_TIME=`date -u +%Y-%m-%d:%H-%M-%S`
  ${BUILD_OPTS} go build -ldflags "-X `go list ./version`.Version=$VERSION -X `go list ./version`.BuildTime=$BUILD_TIME" -v -o ./web main.go
}

COMMAND=$1
case $COMMAND in
  "run")
    Run
    ;;
  "go")
    echo "🐳  Running Web Server with Golang 🐳"
    go run main.go
    ;;
  "build")
    echo "🐳  Building...... 🐳"
    Build
    ;;
  "bin")
    echo "🐳  Runing with Binary...... 🐳"
    if [ -f web ]; then
      ./web
    else
      Build
      ./web
    fi
    ;;
  *)
    Run
    ;;
esac
exit 0
