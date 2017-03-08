#!/bin/sh

for env in `cat ENV`
do
  export $env
done


if `which fresh`; then
  fresh
else
  go run main.go
fi
