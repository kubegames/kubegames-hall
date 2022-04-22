#!/bin/bash
set -ex
cp ~/.kube/config ./kube.config
kubectl create configmap -n kubegames k8s-client-config --from-file=./kube.config