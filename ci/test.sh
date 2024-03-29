#!/usr/bin/env bash

parent_path=$( cd "$(dirname ${BASH_SOURCE[0]})" ; pwd -P )
source $parent_path/common.sh

if [ -n "$1" ]; then
	sourceFile "$1"
fi

log "Finding ${DOCKER_IMAGE}..."
list_image=$(docker images -q ${DOCKER_IMAGE})
checkError $?

if [ "$list_image" == "" ]; then
	log "Loading ${DOCKER_IMAGE} ..."
	docker load -i ${DOCKER_ARCHIVE}
	checkError $?
fi

log "Running ${DOCKER_IMAGE} ..."
docker run -t --rm --entrypoint=/tmp/test.sh ${DOCKER_IMAGE}
checkError $?

log "OK"
exit 0
