#!/usr/bin/env bash

echo
echo "-- Pushing tags $VERSION, light, full and latest up to dockerhub --"

RELEASE=""
if [ ! -z $VERSION ]; then
    RELEASE=":$VERSION"
fi

docker push helderfarias/gcs${RELEASE}
