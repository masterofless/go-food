#!/bin/bash -xu

kubectl -n ${NAMESPACE} delete secret mongo-credentials
kubectl -n ${NAMESPACE} create secret generic mongo-credentials \
    --from-literal=initdb_root_username=foodstuff_${NAMESPACE}_user \
    --from-literal=initdb_root_passwd=$(pwgen -N 1 16)

kubectl -n ${NAMESPACE} delete job food-db-test
kubectl -n ${NAMESPACE} delete job food-db-setup
# kubectl -n ${NAMESPACE} delete deployment food-db-server
# kubectl -n ${NAMESPACE} delete service food-db-server

for i in deployments/*.yml; do
  kubectl -n $NAMESPACE apply -f $i
done
