#!/usr/bin/env bash

# Build!
echo
echo "-- Building release docker images for version $VERSION --"

RELEASE=""
if [ ! -z $VERSION ]; then
    RELEASE=":$VERSION"
fi

docker build --no-cache -t helderfarias/gcs${RELEASE} -f "scripts/docker/Dockerfile" .
