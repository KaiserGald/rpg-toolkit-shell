package rpgAppSh

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/KaiserGald/rpgAppSh/pkg/client"
)

// All the commands that the console can read
const (
	exit string = "exit"
	y    string = "y"
	yes  string = "yes"
	n    string = "n"
	no   string = "no"
	stop string = "stop"
)

var p *os.Process

// Start starts the console
func Start() error {

	var err error
	p, err = os.FindProcess(os.Getpid())
	handleError("Error finding pid: ", err)
	go run()
	waitForSignal()
	return nil
}

// run starts the loop for the console
func run() {
	conn, err := client.Start()
	handleError("Error starting client: ", err)
	r := bufio.NewReader(os.Stdin)
	var text string

	fmt.Println("\n")
	fmt.Println("Welcome to Unlicht RPG Toolkit Shell")
	fmt.Println(strings.Repeat("-", 100))
	for {
		client.Send("\n", conn)
		_, err := client.Read(conn)
		if err != nil {
			reconnect()
		}
		fmt.Print(":> ")
		text = getInput(r)

		// check for exit first
		if text == exit {
			fmt.Print("   Are you sure you want to exit? y/n ")
			text = getInput(r)
			if text == y || text == yes {
				fmt.Println("Goodbye!")
				err := p.Signal(os.Interrupt)
				handleError("   Error emitting the interrupt signal: ", err)
			}
			// else send and receive response
		} else {
			client.Send(text, conn)
			res, err := client.Read(conn)
			if err != nil {
				reconnect()
			}

			handleResponse(res)
		}
	}
}

func reconnect() *net.TCPConn {
	fmt.Println("   Lost connection to client, attempting reconnect.")
	time.Sleep(5 * time.Second)
	conn, err := client.Start()
	handleError("Error starting client: ", err)
	return conn
}

// getInput reads input from the console
func getInput(r *bufio.Reader) string {
	t, err := r.ReadString('\n')
	handleError("Error reading from console: ", err)
	t = strings.Trim(t, "\n")
	return t
}

// handleResponse takes the comand string and routes it to the proper use case
func handleResponse(res string) {
	switch res {
	case stop:
		fmt.Printf("   Stop command received, shutting server down...\n")
		fmt.Println("Goodbye!")
		err := p.Signal(os.Interrupt)
		handleError("   Error emitting the interrupt signal: ", err)

	default:
		fmt.Printf("   Unknown command received.\n")
	}
}

// waitForSignal waits for an interrupt or terminate signal before handling shut down of the program.
func waitForSignal() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	s := <-ch

	fmt.Printf("Got signal: %v, exiting.\n", s)
	time.Sleep(2 * time.Second)
}

func handleError(m string, e error) {
	if e != nil {
		log.Println(m+"%v\n", e)
	}
}
