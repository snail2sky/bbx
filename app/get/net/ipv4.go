package net

import (
	"fmt"
	"github.com/spf13/cobra"
	"net"
)

type IPv4Config struct {
	where string
}

func NewIPv4Config(cmd *cobra.Command) *IPv4Config {
	where, _ := cmd.Flags().GetString("where")
	return &IPv4Config{
		where: where,
	}
}

func (c *IPv4Config) GetIPv4() {
	switch c.where {
	case "local":
		c.getLocalIPv4()
	case "all":
		c.getAllIPv4()
	default:
		c.getIPv4()
	}
}

func (c *IPv4Config) getLocalIPv4() {
	address, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, addr := range address {
		if ipNet, ok := addr.(*net.IPNet); ok && ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				fmt.Println(ipNet.IP.String())
			}
		}
	}
}

func (c *IPv4Config) getAllIPv4() {
	address, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, addr := range address {
		if ipNet, ok := addr.(*net.IPNet); ok {
			if ipNet.IP.To4() != nil {
				fmt.Println(ipNet.IP.String())
			}
		}
	}
}

func (c *IPv4Config) getIPv4() {
	address, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, addr := range address {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				fmt.Println(ipNet.IP.String())
			}
		}
	}
}
