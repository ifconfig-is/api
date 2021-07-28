package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/url"
)

func ParseIP(address string) net.IP {
	ip := net.ParseIP(address)
	if ip != nil {
		return ip
	} else {
		var host string
		u, err := url.ParseRequestURI(address)
		if err == nil {
			host = u.Host
		} else {
			host = address
		}
		ips, err := net.LookupIP(host)
		if err != nil {
		}
		return ips[0]
	}
}

func PPrint(i interface{}) {
	s, _ := json.MarshalIndent(i, "", "  ")
	fmt.Println(string(s))
}
