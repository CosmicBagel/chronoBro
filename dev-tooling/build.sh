#!/bin/bash
SCRIPT_PATH=$(dirname $(realpath $0))
cd $SCRIPT_PATH/api/pkg
# note: doing a fully static build so exe can be deployed on minimal images, or 
# images not using clib

# CGO_ENABLED=0
# go build -o ../container_files/chronoBro -ldflags "-linkmode 'external' -extldflags '-static'"
go build -o ../container_files/chronoBro 
