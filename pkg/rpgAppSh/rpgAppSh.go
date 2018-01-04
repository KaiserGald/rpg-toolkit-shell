package rpgAppSh

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	exit string = "exit"
)

// Start starts the console
func Start() {
	run()
}

func run() {
	r := bufio.NewReader(os.Stdin)
	var text string

	fmt.Println("Welcome to Unlicht RPG Toolkit Shell")
	fmt.Println(strings.Repeat("-", 100))

	for {
		fmt.Print(":> ")
		text = getInput(r)

		if text == "exit" {
			fmt.Print("Are you sure you want to exit? y/n :> ")
			text = getInput(r)
			if text == "y" {
				fmt.Println("Goodbye!")
				os.Exit(0)
			}
		}
	}
}

func getInput(r *bufio.Reader) string {
	t, _ := r.ReadString('\n')
	t = strings.Trim(t, "\n")
	return t
}
