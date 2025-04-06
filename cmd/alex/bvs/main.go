package main

import (
	"alex/bvs/internal/console"
	"fmt"
	"os"
)

func main() {
	argv := os.Args

	if len(argv) != 3 {
		usage()
		return
	}

	console.HandleCmd(argv[1], argv[2])
}

func usage() {
	fmt.Println("Usage: bvs <new|load> <file>")
}