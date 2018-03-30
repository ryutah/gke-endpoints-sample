#!/bin/sh

set -ex

gcloud endpoints configs list --service=echo-api.endpoints.${PROJECT_ID}.cloud.goog
