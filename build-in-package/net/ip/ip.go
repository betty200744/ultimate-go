package main

import (
	"encoding/binary"
	"net"

	"github.com/smallnest/rpcx/log"
)

func ip2int(ip net.IP) uint32 {
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}

func int2ip(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip
}

func main() {
	ipString := "192.168.1.9"
	ip := net.ParseIP(ipString)
	ipInt := ip2int(ip)
	log.Infof("ip2int:  ip string: %v -> int: %v", ipString, ipInt)
	ipFromInt := int2ip(ipInt)
	log.Infof("int2ip: ip int: %v -> ip: %v", ipInt, ipFromInt)
}
