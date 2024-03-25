package echo

import "fmt"

func (c *Config) UDPRun() {
	addr := fmt.Sprintf("%s:%d", c.host, c.port)
	c.serve(addr, "udp")
}
