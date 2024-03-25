package echo

import "fmt"

func (c *Config) HTTPRun() {
	addr := fmt.Sprintf("%s:%d", c.host, c.port)
	c.serve(addr, "http")
}
