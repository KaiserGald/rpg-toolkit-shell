package client

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

var servAddr string

// Start starts the tcp connection
func Start() *net.TCPConn {
	servAddr = "127.0.0.1:8081"
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		fmt.Println("ResolveTCPAddr failed: ", err)
	}

	c, _ := net.DialTCP("tcp", nil, tcpAddr)

	return c
}

// Send sends a message to the connection
func Send(s string, c *net.TCPConn) {
	fmt.Fprintf(c, s+"\n")
}

// Read reads a message from the connection
func Read(c *net.TCPConn) {
	message, _ := bufio.NewReader(c).ReadString('\n')
	if message != "" {
		message = strings.Trim(message, "\n")
		fmt.Println("-> " + message)
	}
}
