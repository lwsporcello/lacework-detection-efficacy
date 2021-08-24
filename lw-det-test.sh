#!/bin/bash
# Script to simulate activity that will trigger Lacework events

EXT_IP="54.184.116.123"
SIM_1_BIN="lw-scan-brute"
SIM_2_BIN="lw-stage-1"
BIN_URL="http://${EXT_IP}/bin/"

download_exec_binary() {

	# Download the binary
	echo "Downloading $1..."
    if [[ $(curl -s -w "%{http_code}\\n" -X GET ${BIN_URL}$1 -o $1) -eq 200 ]]
    then
        echo " - Successfully downloaded $1!"
    else
        echo " - Failed to download $1!"
		echo ""
		continue
    fi

	# Change permissions so it can be executed
	echo "Changing permissins on $1 to allow execution..."
	chmod u+x $1
	if [[ $? -eq 0 ]]
	then
		echo " - Successfully changed permissions"
	else
		echo " - Failed to change permissions on $1!"
		echo ""
		continue
	fi

	# Execute the binary
	echo "Executing $1..."
	./$1$2 >/dev/null 2>&1 &
	if [[ $? -eq 0 ]]
    then
        echo " - Successfully executed $1!"
    else
        echo " - Failed to execute $1!"
		echo ""
		continue
    fi
}

# Show menu
echo "-----------------------------------------------------"
echo "| Welcome to the Lacework detection testing script. |"
echo "| This script will run attack simulation in your    |"
echo "| workload running the Lacework agent. The agent    |"
echo "| will capture this activity and generate events in |"
echo "| the Lacework UI. Begin by choosing which          |"
echo "| simulations you want to run...                    |"
echo "-----------------------------------------------------"
echo ""
echo "  0. Quit"
echo "  1. Simulation 1: Network Scan & Brute Force"
echo "  2. Simulation 2: Multi-Stage Malware"
echo ""
echo "You may enter one simulation (i.e. 2) or multiple comma separated (i.e. 1,2)"
echo ""
read -p 'Enter Selection(s): ' selection

#Error check input
while ! [[ "$selection" =~  ^[0-2](,[0-2])*$ ]]; do
	read -p 'Bad Entry. Enter selection(s): ' selection </dev/tty
done

#run each simulation selected
for i in $(echo $selection | sed "s/,/ /g")
do
	echo "Running Simulation $i"
	echo "---------------------"
	if [[ $i -eq 0 ]]; then
		echo "Script terminating..."
		exit
	elif [[ $i -eq 1 ]]; then
		download_exec_binary $SIM_1_BIN
	elif [[ $i -eq 2 ]]; then
		download_exec_binary $SIM_2_BIN
	fi
	echo ""
done

echo "Simulations Completed!"
