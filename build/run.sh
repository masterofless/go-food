#!/bin/bash -eu

export NAMESPACE=${NAMESPACE:-afc}
export BUILD_NUMBER=${BUILD_NUMBER:-latest}
export REPO_PREFIX=${REPO_PREFIX:-masterofless/go-food}
export BRANCH=${BRANCH:-`git status | grep 'On branch' | perl -pe 's/On branch //'`}

set -x
build/build.sh $* && build/deploy.sh $* && build/test.sh $*
