package godiskalert

import (
	"net"
)

// IPアドレスを取得して返却
func IpAddress() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				// IPv4が取得できれば返却
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
