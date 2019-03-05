#!/usr/bin/env bash
#
# Copyright 2018 ScyllaDB
#

set -eu -o pipefail

if [[ -z ${SCYLLA_API_DOC} ]]; then
    echo "set SCYLLA_API_DOC env variable pointing to Scylla api/api-doc folder"
    exit 1
fi

jq -M -s 'reduce .[] as $x (.[0]; .apis += $x.apis? | .models += $x.models) | .resourcePath="/"' ${SCYLLA_API_DOC}/*.json \
| sed -e 's/#\/utils\///g' \
| sed -e 's/"bool"/"boolean"/g' \
| sed -e 's/"int"/"integer"/g' \
> scylla-api_10.json

echo "Go to https://www.apimatic.io/transformer transform '$(pwd)/scylla-api_10.json' to scylla-api.json"