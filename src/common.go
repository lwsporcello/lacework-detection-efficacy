package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
)

func httpPost(url string, jsonStr []byte) error {

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	_ = resp.Body.Close()

	return err
}

func httpGet(url string, filename string) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	out, err := os.Create(filename)
	if err != nil {
		return err
	}

	_, err = io.Copy(out, resp.Body)
	_ = resp.Body.Close()
	_ = out.Close()
	return err
}

func setupLogger(filename string) *log.Logger {
	wd, _ := os.Getwd()
	logpath := wd + "/" + filename
	var file, err1 = os.Create(logpath)

	if err1 != nil {
		panic(err1)
	}

	Log := log.New(file, "", log.LstdFlags|log.Lshortfile)
	Log.Println("LogFile : " + logpath)
	return Log
}
