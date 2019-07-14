#!/usr/bin/env bash

# Build!
echo "-- Building release docker images for version $VERSION --"
echo ""
docker build --no-cache -t helderfarias/gcs${VERSION} -f "scripts/docker/Dockerfile" .
