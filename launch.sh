#!/bin/bash

for env in `cat ENV`
do
  export $env
done

COMMAND=$1
if [[ $COMMAND == "fresh" ]] ; then
  fresh
elif [[ $COMMAND == "" ]] && [[ `which fresh` ]] ; then
  fresh
else
  go run main.go
fi
