package console

import (
	"alex/bvs/internal/core"
	"bufio"
	"fmt"
	"os"
)

func HandleCmd(command, path string) {
	var bytes []byte

	switch command {
	case "new":
		NewFile(path)
		var fsize int32
		fmt.Print("size: ")
		fmt.Scan(&fsize)

		bytes = make([]byte, (fsize+7)/8)
	case "load":
		bytes = LoadFile(path)
	}

	defer SaveFile(path, bytes)
	Run(bytes)
}

func Run(bytes []byte) {
	bf := core.BloomFilterFromBytes(bytes)
	bf.Exists("0")
	
	ls := bufio.NewScanner(os.Stdin)
	var command string
	for command != "." {
		fmt.Scan(&command)
		
		switch command {
		case "!":
			ls.Scan()
			bf.Insert(ls.Text())
			fmt.Println("ok")
		case "?":
			ls.Scan()
			fmt.Println(bf.Exists(ls.Text()))
		default:
			fmt.Println("Commands: ! to insert, ? to check existence, . to exit")
		}
	}
}



