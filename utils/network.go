package utils

import (
	"net"
	"strings"
)

func EscapeHttpHeader(input string) string {
	if len(input) == 0 {
		return input
	}
	// trim space
	input = strings.TrimSpace(input)
	// trim http
	input = strings.TrimPrefix(input, "http://")
	// trim https
	input = strings.TrimPrefix(input, "https://")
	return input
}

func GetIp() string {
	adds, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, a := range adds {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
