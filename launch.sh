#!/bin/sh

for env in `cat ENV`
do
  export $env
done

fresh
