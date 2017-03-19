#!/usr/bin/env bash

# set -xe

if [[ ${TRAVIS_TAG} ]]; then
  TRAVIS_BRANCH=${TRAVIS_TAG}
fi

DOCKER_TAG=$(sed 's/\//_/g' <<< ${TRAVIS_BRANCH})

docker build -f Dockerfile.build -t mlaccetti/kube-cert-manager-r53:build-${DOCKER_TAG} .
docker run -v $(pwd):/go/src/github.com/mlaccetti//usr/src/dns01-exec-plugin-r53 mlaccetti/kube-cert-manager-r53:build-${DOCKER_TAG}

docker build -t mlaccetti/kube-cert-manager-r53:${DOCKER_TAG} .

docker login -u="${DOCKER_USERNAME}" -p="${DOCKER_PASSWORD}"
docker push mlaccetti/kube-cert-manager-r53:${DOCKER_TAG}
