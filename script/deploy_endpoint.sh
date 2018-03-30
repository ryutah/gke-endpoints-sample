#!/bin/sh

set -ex

cd $(dirname $0)

gcloud endpoints services deploy ./../openapi.yaml
