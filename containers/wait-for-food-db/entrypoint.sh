#!/bin/bash -exu

export DEST=${DEST:-couchbase:8091}

echo "Curl sampling;waiting for $DEST"
let count=0
until $(curl -m 5 --output /dev/null --silent --head --fail $DEST); do
    let count=count+1
    if ((count >= 20)); then
        echo "Timeout waiting for $$DEST"
        exit -1
    fi
    printf '.'
    sleep 5
done
echo "$DEST is alive"
