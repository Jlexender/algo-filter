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
	for command != "exit" {
		fmt.Scan(&command)
		
		switch command {
		case "insert":
			ls.Scan()
			bf.Insert(ls.Text())
			fmt.Println("ok")
		case "check":
			ls.Scan()
			fmt.Println(bf.Exists(ls.Text()))
		default:
			fmt.Println("Commands: insert, check, exit")
		}
	}
}



