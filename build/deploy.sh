#!/bin/bash -exu

for i in deployments/*.yml; do
  kubectl -n $WORKSPACE apply -f $i
done
