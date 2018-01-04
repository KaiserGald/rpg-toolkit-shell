package server

import (
	"fmt"
	"net"
)

// Start starts the tcp connection
func Start() net.Conn {
	c, _ := net.Dial("tcp", "127.0.0.1:8081")

	return c
}

// Send sends a message to the connection
func Send(s string, c net.Conn) {
	fmt.Fprintf(c, s+"\n")
}
