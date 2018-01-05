package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

var servAddr string

// Start starts the tcp connection
func Start() (*net.TCPConn, error) {
	servAddr = "127.0.0.1:8081"
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	handleError("ResolveTCPAddr failed: ", err)

	c, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Printf("   Error creating connection to server: %v\n", err)
		fmt.Println("   Attempt to reconnect? y/n ")
		r := bufio.NewReader(os.Stdin)
		i := getInput(r)
		if i == "y" || i == "yes" {
			for err != nil {
				fmt.Println("   Attempting to reconnect in 10 seconds...")
				time.Sleep(10 * time.Second)
				fmt.Println("   Reconnecting now...")
				c, err = net.DialTCP("tcp", nil, tcpAddr)
				if err != nil {
					fmt.Printf("   Error creating connection to server: %v\n", err)
				}
			}
		}
	}

	return c, nil
}

// Send sends a message to the connection
func Send(s string, c *net.TCPConn) {
	fmt.Fprintf(c, s+"\n")
}

// Read reads a message from the connection
func Read(c *net.TCPConn) (string, error) {
	message, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		return "", err
	}
	if message != "" {
		message = strings.Trim(message, "\n")
		return message, nil
	}

	return message, nil
}

func getInput(r *bufio.Reader) string {
	t, err := r.ReadString('\n')
	handleError("Error reading from console: ", err)
	t = strings.Trim(t, "\n")
	return t
}

func handleError(m string, e error) {
	if e != nil {
		log.Println(m+"%v\n", e)
	}
}
