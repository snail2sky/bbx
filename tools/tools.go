package tools

import (
	"log"
	"net"
)

func GetNetInterface() {

}

func IP(family string) []string {
	// 获取本机所有网络接口的信息
	var ipList = make([]string, 0)
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Println("Error:", err)
		return nil
	}

	// 遍历所有网络接口
	for _, iFace := range interfaces {
		// 获取接口的 IP 地址
		addresses, err := iFace.Addrs()
		if err != nil {
			log.Println("Error:", err)
			continue
		}

		// 遍历接口的 IP 地址
		for _, addr := range addresses {
			// 检查是否是 IP 地址结构
			if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
				switch family {
				case "v4":
					if ipNet.IP.To4() != nil {
						ipList = append(ipList, ipNet.IP.To4().String())
						log.Println("IPv4:", ipNet.IP.String())
					}
				case "v6":
					if ipNet.IP.To16() != nil {
						log.Println("IPv6:", ipNet.IP.String())
						ipList = append(ipList, ipNet.IP.To16().String())
					}
				case "all":
					if ipNet.IP.To4() != nil && ipNet.IP.To16() != nil {
						ipList = append(ipList, ipNet.IP.To4().String(), ipNet.IP.To16().String())

					}
				}
			}
		}
	}
	return ipList
}
