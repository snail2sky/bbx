package echo

import (
	"fmt"
)

func (c *Config) TCPRun() {
	addr := fmt.Sprintf("%s:%d", c.host, c.port)
	c.serve(addr, "tcp")
}
