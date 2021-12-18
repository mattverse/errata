#! /bin/bash -x
set -e

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"

# Build the docker image first.
docker build --tag errata/protogen -f "$DIR"/Dockerfile "$DIR"/../..

# Make sure that previous container not exist.
docker rm --force errata_protogen 1>/dev/null 2>&1

docker run -v "$DIR"/../..:/app --name errata_protogen errata/protogen sh dev/proto/protogen-swagger-doc.sh

# Clear the container
docker rm --force errata_protogen 1>/dev/null 2>&1
