#!/usr/bin/env bash

if [ -n "$DOCKER_USERNAME" ] && [ -n "$DOCKER_PASSWORD" ]; then
    echo "Login to the docker..."
    docker login -u $DOCKER_USERNAME -p $DOCKER_PASSWORD $DOCKER_REGISTRY
fi

if [ -n "$GITHUB_TERRAFORM_GITHUB_TOKEN" ] ; then
  export GITHUB_TOKEN=$GITHUB_TERRAFORM_GITHUB_TOKEN
fi

github-terraform $@