#!/bin/bash -exu

export WORKSPACE=afc
export BUILD_NUMBER=${BUILD_NUMBER:-0}
export BUILD_NUMBER=$(($BUILD_NUMBER + 1))
export REPO_PREFIX=masterofless/go-food
export BRANCH=`git status | grep 'On branch' | perl -pe 's/On branch //'`

build/build.sh && build/deploy.sh && build/test.sh
