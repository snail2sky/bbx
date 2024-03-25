package echo

import (
	"fmt"
)

func (t *Config) TCPRun() {
	addr := fmt.Sprintf("%s:%d", t.host, t.port)
	t.serve(addr, "tcp")
}
