#!/bin/bash -exu

export DB_USER=${DB_USER:-foodstuff_${KUBERNETES_NAMESPACE:-afc}_user}
export DB_PASSWD=${DB_PASSWD:-ieHeweghepag6Eel}

go help gopath
go test -v .
go run foodstuff.go
