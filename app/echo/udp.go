package echo

import "fmt"

func (t *Config) UDPRun() {
	addr := fmt.Sprintf("%s:%d", t.host, t.port)
	t.serve(addr, "udp")
}
