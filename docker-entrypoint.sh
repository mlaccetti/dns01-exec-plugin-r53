#!/usr/bin/env bash

# For easy debugging
# set -ex

if [ "$#" -gt 0 ]; then
  echo "Running user provided command: [ $@ ]"
  exec $@
else
  echo "Building dns01-exec-plugin-r53 in $PWD..."
  rm -fr dist
  mkdir dist
  cd route53
  go get github.com/aws/aws-sdk-go
  go build -o ../dist/route53 -v
fi
