#!/bin/bash
SCRIPT_PATH=$(dirname $(realpath $0))

# urandom is cryptographic random generation, so this should be pretty good
PASSWORD=$(LC_ALL=C tr -dc '[:alnum:]' < /dev/urandom | head -c64)

if [ ! -d "$SCRIPT_PATH/secrets" ] 
then 
    mkdir $SCRIPT_PATH/secrets 
fi
echo $PASSWORD > $SCRIPT_PATH/secrets/dbpw.txt
