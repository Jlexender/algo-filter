package main

import (
	"fmt"
	"os"
)

func main() {
	argv := os.Args

	if len(argv) != 3 {
		fmt.Println("Usage: bvs <new|load> <file>")
		return
	}

	
}