package jb

import (
	"net"
	"strings"
)

// DNSResolve here
func DNSResolve(ip string) string {
	hn, err := net.LookupAddr(ip)
	if err == nil {
		return strings.Trim(hn[0], ".")
	}
	return ip
}

// ParseFQDN here
func ParseFQDN(s string) string {
	return strings.Split(s, ".")[0]
}
