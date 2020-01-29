#!/bin/bash -u

export COUCHBASE_HOST=${COUCHBASE_HOST:-couchbase}
export COUCHBASE_PORT=${COUCHBASE_PORT:-8091}
export COUCHBASE_BUCKET=${COUCHBASE_BUCKET:-food-db}
export COUCHBASE_HOST=${COUCHBASE_USER:-couchbase}
export COUCHBASE_HOST=${COUCHBASE_PASSWORD:-foodaloo}

CB=http://${COUCHBASE_HOST}:${COUCHBASE_PORT}

set -e
set -x

# Setup Services
curl -X POST -u ${COUCHBASE_USER}:${COUCHBASE_PASSWORD} ${CB}/node/controller/setupServices \
   -d 'services=kv%2Cn1ql%2Cindex'

# Initialize Node
# TODO use /data
curl -X POST ${CB}/nodes/self/controller/settings \
   -d 'path=/opt/couchbase/var/lib/couchbase/data&index_path=/opt/couchbase/var/lib/couchbase/data'

echo "Set Couchbase admin user and password"
curl -X POST ${CB}/settings/web \
   -d "username=${COUCHBASE_USER}&password=${COUCHBASE_PASSWORD}&port=SAME"

echo "Create a bucket ${COUCHBASE_BUCKET}"
curl -X POST -u ${COUCHBASE_USER}:${COUCHBASE_PASSWORD} ${CB}/pools/default/buckets \
    -d "name=${COUCHBASE_BUCKET}&authType=none&replicaNumber=0&ramQuotaMB=100"

echo "Set Couchbase index mode type"
curl -X POST -u ${COUCHBASE_USER}:${COUCHBASE_PASSWORD} ${CB}/settings/indexes \
   -d "storageMode=memory_optimized"

# Setup Index RAM Quota
#curl -u username=[admin]&password=[password] -X POST \
#http://[localhost]:8091/pools/default -d memoryQuota=[value] \
#-d indexMemoryQuota=[value]
