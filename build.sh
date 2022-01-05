#!/usr/bin/env bash

docker build -t tekton-sample  .
docker run -p 8080:8080 127.0.0.1:5000/tekton-sample:latest tekton-sample-01 --restart=Never --rm
