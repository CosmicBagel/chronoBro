#!/bin/ash

chronoBro_pid=0
inotifyd_pid=0
should_quit=0

# trap sigterm, and gracefully shutdown process via PID 
function on_sigterm() {
    should_quit=1
    echo 'API - sigterm caught'
    kill -SIGTERM $inotifyd_pid 
    kill -SIGTERM $chronoBro_pid 
}
trap on_sigterm SIGTERM

# launch process (background) and grab PID
/api/chronoBro &
chronoBro_pid=$!

# watch using inotifyd, will run until file can't be watched any more (x event)
# c change event, n create event
inotifyd - /api:cn |
    while read -r event directory file; do
        if [[ "$file" == "chronoBro" ]]; then
            echo "API - exe modified ($event $directory $file)"
            # restart chronoBro process if chronoBro exe changes
            kill -SIGTERM $chronoBro_pid 
            /api/chronoBro &
            chronoBro_pid=$!
        # else
        #     echo "busybox inotifyd event: ($event $directory $file)"
        fi
    done &
# grab its pid so we can shut it down on SIGTERM
inotifyd_pid=$!

# sleep till sigterm (should_quit updated via trap)
while [[ $should_quit -ne 1 ]]; do 
    sleep 500
done

