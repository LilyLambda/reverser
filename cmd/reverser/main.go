package main

import (
	"fmt"
	"os"

	"github.com/lilylambda/reverser"
)

func main() {
	if len(os.Args) != 2 {
		println("usage: reverser [word]")
		os.Exit(1)
	}
	text := os.Args[1]
	fmt.Println(reverser.Reverse(text))
}
