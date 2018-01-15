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

var (
	servAddr string
	p        *os.Process
)

func init() {
	servAddr = "127.0.0.1:8081"
	var err error
	p, err = os.FindProcess(os.Getpid())
	handleError("Error finding pid: ", err)
}

// Start starts the tcp connection
func Start() (*net.TCPConn, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	handleError("ResolveTCPAddr failed: ", err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Printf("   Error creating connection to server: %v\n", err)
		fmt.Print("   Attempt to reconnect? y/n ")
		r := bufio.NewReader(os.Stdin)
		i := getInput(r)
		if i == "y" || i == "yes" {
			conn, err = reconnect(tcpAddr)
			if err != nil {
				return conn, err
			}
		} else {
			fmt.Println("  Shutting down client...\n")
			return nil, err
		}
	}
	fmt.Println("    Connected to server!")
	return conn, nil
}

// Restart restarts the client
func Restart() (*net.TCPConn, error) {
	var conn *net.TCPConn
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	handleError("ResolveTCPAddr failed: ", err)

	t1 := make(chan *net.TCPConn)

	go func() {
		conn, err = reconnect(tcpAddr)
		if err != nil {
			fmt.Println("   Error restarting client.")
		}

		t1 <- conn
	}()

	select {
	case res := <-t1:
		return res, err
	case <-time.After(30 * time.Second):
		return conn, err
	}
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

// Reconnect attempts to reconnect to the server times out after 30 sec
func reconnect(ta *net.TCPAddr) (*net.TCPConn, error) {
	for {
		var c *net.TCPConn
		var err error
		fmt.Println("   Attempting to reconnect...")
		timeout := make(chan *net.TCPConn, 1)
		go func() {
			for {
				time.Sleep(1 * time.Second)
				c, err = net.DialTCP("tcp", nil, ta)
				if err == nil {
					break
				}
			}

			timeout <- c
		}()

		select {
		case c = <-timeout:
			fmt.Println("   Reconnected!")
			return c, nil
		case <-time.After(time.Second * 60):
			fmt.Println("   Reconnect timed out...")
			return c, err
		}
	}
}
