#!/bin/bash -xu

export DEST_HOST=${DEST_HOST:-food-db-server}
export DEST_PORT=${DEST_PORT:-27017}

echo "NC sampling; waiting for MySQL at ${DEST_HOST} ${DEST_PORT}"
let count=0
until $(ncat --wait 3 ${DEST_HOST} ${DEST_PORT} </dev/null > /dev/null); do
    let count=count+1
    if ((count >= 10)); then
        echo "Timeout waiting for ${DEST_HOST} ${DEST_PORT}"
        exit -1
    fi
    printf '.'
    sleep 5
done
echo "${DEST_HOST} ${DEST_PORT} is alive"
