#!/bin/sh 

set -ex

kubectl create secret generic service-account-cred \
  --from-file service-account-cred.json
