#!/usr/bin/env bash

echo
echo "-- Pushing tags $VERSION, light, full and latest up to dockerhub --"

docker push helderfarias/gcs${VERSION}
