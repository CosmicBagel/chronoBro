#!/bin/bash

# note: requires inotify-tools to be installed
SCRIPT_PATH=$(dirname $(realpath $0))
cd $SCRIPT_PATH/api/pkg

inotifywait -m . -e modify -e move -e create -e delete |
    while read -r directory action file; do
        if [[ $file =~ .*go$ ]]; then
            docker exec -it chronobro-dev-tooling-1 ash /build.sh
        fi
    done
