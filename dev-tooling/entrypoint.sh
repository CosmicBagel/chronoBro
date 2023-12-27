#!/bin/ash
should_quit=0

# trap sigterm, and gracefully shutdown process via PID 
function on_sigterm() {
    should_quit=1
}
trap on_sigterm SIGTERM

# sleep till sigterm (should_quit updated via trap)
while [[ $should_quit -ne 1 ]]; do 
    sleep 0.5 
done
