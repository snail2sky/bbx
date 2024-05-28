package net

import (
	"fmt"
	"github.com/spf13/cobra"
	"net"
)

type IPv6Config struct {
	where string
}

func NewIPv6Config(cmd *cobra.Command) *IPv6Config {
	where, _ := cmd.Flags().GetString("where")
	return &IPv6Config{
		where: where,
	}
}

func (c *IPv6Config) GetIPv6() {
	switch c.where {
	case "local":
		c.getLocalIPv6()
	case "all":
		c.getAllIPv6()
	default:
		c.getIPv6()
	}
}

func (c *IPv6Config) getLocalIPv6() {
	address, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, addr := range address {
		if ipNet, ok := addr.(*net.IPNet); ok && ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() == nil {
				fmt.Println(ipNet.IP.String())
			}
		}
	}
}

func (c *IPv6Config) getAllIPv6() {
	address, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, addr := range address {
		if ipNet, ok := addr.(*net.IPNet); ok {
			if ipNet.IP.To4() == nil {
				fmt.Println(ipNet.IP.String())
			}
		}
	}
}

func (c *IPv6Config) getIPv6() {
	address, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, addr := range address {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() == nil {
				fmt.Println(ipNet.IP.String())
			}
		}
	}
}
