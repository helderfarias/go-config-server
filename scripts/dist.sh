#!/usr/bin/env bash

# Build!
echo
echo "-- Building release docker images for version $VERSION --"
docker build --no-cache -t helderfarias/gcs${VERSION} -f "scripts/docker/Dockerfile" .
