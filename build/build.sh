#!/bin/bash -eu

export NAMESPACE=${NAMESPACE:-afc}
export BUILD_NUMBER=${BUILD_NUMBER:-latest}
export REPO_PREFIX=${REPO_PREFIX:-masterofless/go-food}
export BRANCH=${BRANCH:-`git status | grep 'On branch' | perl -pe 's/On branch //'`}
export TEST_IMAGE=${TEST_IMAGE:-"$REPO_PREFIX/food-db-test:$BRANCH.$BUILD_NUMBER"}

containers=${1:-`ls containers`}

for i in $containers; do
  TAG="$REPO_PREFIX/$i:master.$BUILD_NUMBER"
  dir=`echo $i | perl -pe 's/.*\///'`
  set -x
  docker build -t $TAG containers/$dir
  set +x
done
