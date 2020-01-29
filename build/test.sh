#!/bin/bash -xu

IMAGE="$REPO_PREFIX-food-db-test:$BRANCH.$BUILD_NUMBER"

kubectl -n $WORKSPACE delete pod/go-food-test
kubectl -n $WORKSPACE run go-food-test -it --rm --restart=Never --image=$IMAGE
