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

func setupLogger(filename string) *os.File {
	wd, _ := os.Getwd()
	f, err := os.OpenFile(wd+"/"+filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	return f
}
