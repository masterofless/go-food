#!/bin/bash -xu

export NAMESPACE=${NAMESPACE:-afc}
export BUILD_NUMBER=${BUILD_NUMBER:-latest}
export REPO_PREFIX=${REPO_PREFIX:-masterofless/go-food}
export BRANCH=${BRANCH:-`git status | grep 'On branch' | perl -pe 's/On branch //'`}
export TEST_IMAGE=${TEST_IMAGE:-"$REPO_PREFIX/food-db-test:$BRANCH.$BUILD_NUMBER"}

kubectl -n $NAMESPACE delete pod/go-food-test
kubectl -n $NAMESPACE run go-food-test -it --rm --restart=Never --image=$TEST_IMAGE $*
