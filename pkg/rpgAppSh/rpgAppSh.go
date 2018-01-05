package rpgAppSh

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

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

// Start starts the console
func Start() error {
	go run()
	waitForSignal()
	return nil
}

// run starts the loop for the console
func run() {
	conn := client.Start()
	r := bufio.NewReader(os.Stdin)
	var text string

	fmt.Println("Welcome to Unlicht RPG Toolkit Shell")
	fmt.Println(strings.Repeat("-", 100))
	for {
		fmt.Print(":> ")
		text = getInput(r)

		if text == exit {
			fmt.Print("Are you sure you want to exit? y/n :> ")
			text = getInput(r)
			if text == y || text == yes {
				fmt.Println("Goodbye!")
				os.Exit(0)
			}
		} else {
			client.Send(text, conn)
			client.Read(conn)
		}
	}
}

func getInput(r *bufio.Reader) string {
	t, _ := r.ReadString('\n')
	t = strings.Trim(t, "\n")
	return t
}

func waitForSignal() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	s := <-ch

	fmt.Println("Got signal: %v, exiting.\n", s)
}
