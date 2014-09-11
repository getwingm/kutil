package kutil

import (
	"net"
	"strings"
)

func GetAllIpAddress() (ips []string) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return ips
	}
	// handle err
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			return ips
		}
		// handle err
		for _, addr := range addrs {
			switch addr.(type) {
			case *net.IPAddr:
				// process IP address
				ips = append(ips, addr.String())
			}
		}
	}
	return ips
}

func GetPreferIpAddress() (ip string) {
	conn, err := net.Dial("udp", "google.com:80")
	if err != nil {
		return ""
	}
	defer conn.Close()
	return strings.Split(conn.LocalAddr().String(), ":")[0]
}
