#!/bin/sh

set -ex

cd $(dirname $0)

kubectl apply -f ../deployment.yaml

