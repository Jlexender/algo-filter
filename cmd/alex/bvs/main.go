package main

import (
	"fmt"
	"os"
	"alex/bvs/internal/console"
)

func main() {
	argv := os.Args

	if len(argv) != 3 {
		fmt.Println("Usage: bvs <new|load> <file>")
		return
	}

	console.HandleCmd(argv[1], argv[2])
}