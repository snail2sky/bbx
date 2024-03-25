package echo

import (
	"bufio"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net"
)

type Config struct {
	host    string
	port    uint
	bufSize uint
}

func NewConfig(cmd *cobra.Command, args []string) *Config {
	host, _ := cmd.Flags().GetString("host")
	port, _ := cmd.Flags().GetUint("port")
	bufSize, _ := cmd.Flags().GetUint("buf-size")
	return &Config{
		host:    host,
		port:    port,
		bufSize: bufSize,
	}
}

func (t *Config) serveTCP(addr string, protocol string) {
	log.Printf("ECHO server listen on <%s> %s\n", protocol, addr)
	sock, err := net.Listen(protocol, addr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := sock.Accept()
		log.Printf("ECHO server accept new connection <%s> %s<->%s", protocol, conn.LocalAddr(), conn.RemoteAddr())
		if err != nil {
			log.Print(err)
		}
		go t.handleTCPConn(conn)
	}
}

func (t *Config) serveUDP(addr string, protocol string) {
	log.Printf("ECHO server listen on <%s> %s", protocol, addr)
	udpAddr, err := net.ResolveUDPAddr(protocol, addr)
	if err != nil {
		log.Fatal(err)
	}
	sock, err := net.ListenUDP(protocol, udpAddr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		var data = make([]byte, t.bufSize)
		n, addr, err := sock.ReadFromUDP(data)
		if err == io.EOF || err != nil {
			log.Print(err)
		}
		go t.handleUDPConn(sock, data[:n], addr)
	}
}

func (t *Config) serve(addr string, protocol string) {
	switch protocol {
	case "tcp":
		t.serveTCP(addr, protocol)
	case "udp":
		t.serveUDP(addr, protocol)
	}
}

func (t *Config) handleTCPConn(conn net.Conn) {
	protocol := "tcp"
	r := bufio.NewReader(conn)
	w := bufio.NewWriter(conn)
	defer conn.Close()
	for {
		var data = make([]byte, t.bufSize)
		n, err := r.Read(data)
		if err == io.EOF || err != nil {
			break
		}
		log.Printf("ECHO server receive new data <%s> %s<->%s, length: %d", protocol, conn.LocalAddr(), conn.RemoteAddr(), n)
		_, _ = w.Write(data[:n])
		_ = w.Flush()
	}
	log.Printf("ECHO server close old connection <%s> %s<->%s", protocol, conn.LocalAddr(), conn.RemoteAddr())
}

func (t *Config) handleUDPConn(conn *net.UDPConn, data []byte, addr *net.UDPAddr) {
	protocol := "udp"
	log.Printf("ECHO server receive new data <%s> %s<->%s, length: %d", protocol, conn.LocalAddr(), addr, len(data))
	conn.WriteTo(data, addr)
}
