#!/bin/bash

SCRIPT_PATH=$(dirname $(realpath $0))
cd $SCRIPT_PATH/api/pkg

inotifywait -m . -e modify -e move -e create -e delete |
    while read -r directory action file; do
        if [[ $file =~ .*go$ ]]; then
            ../../build.sh
        fi
    done
