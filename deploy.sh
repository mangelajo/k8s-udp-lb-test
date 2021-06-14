#!/bin/sh
set -e
set -x


docker build -t quay.io/mangelajo/k8s-udp-lb-test:dev .
docker push quay.io/mangelajo/k8s-udp-lb-test:dev

kubectl apply -f reproducer.yaml
