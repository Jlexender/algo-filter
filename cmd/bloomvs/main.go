package main

import (
	"fmt"
	"os"
)

func main() {
	argv := os.Args

	if len(argv) != 3 {
		fmt.Println("Usage: bloomvs <new|load> <file>")
		return
	}

	
}