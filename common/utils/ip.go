package utils

import (
	"net"
)

// GetLocalIP 获取本机IP
func GetLocalIP() (ip string) {
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for _, addr := range addresses {
		networkIp, ok := addr.(*net.IPNet)
		if ok && !networkIp.IP.IsLoopback() && networkIp.IP.To4() != nil {
			ip = networkIp.IP.String()
			return
		}
	}
	return "127.0.0.1"
}
