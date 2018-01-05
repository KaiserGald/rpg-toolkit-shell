package server

import (
	"fmt"
	"net"
)

var servAddr string

// Start starts the tcp connection
func Start() net.Conn {
	servAddr = "127.0.0.1:8081"
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		fmt.Println("ResolveTCPAddr failed: ", err)
	}

	c, _ := net.DialTCP("tcp", nil, tcpAddr)

	return c
}

// Send sends a message to the connection
func Send(s string, c net.Conn) {
	fmt.Fprintf(c, s+"\n")
}
