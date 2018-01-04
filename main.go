package main

import (
	"fmt"

	"github.com/KaiserGald/rpgAppSh/pkg/rpgAppSh"
)

func main() {
	if err := rpgAppSh.Start(); err != nil {
		fmt.Println("Error in main(): %v\n", err)
	}
}
