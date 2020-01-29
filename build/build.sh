#!/bin/bash -exu

for i in `ls containers`; do
  TAG="$REPO_PREFIX-$i:master.$BUILD_NUMBER"
  docker build -t $TAG containers/$i
done
