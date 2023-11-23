#!/bin/bash
SCRIPT_PATH=$(dirname $(realpath $0))
cd $SCRIPT_PATH/api/pkg
go build -o ../container_files/chronoBro
