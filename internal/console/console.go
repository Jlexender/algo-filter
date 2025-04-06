package console

import (
	"alex/bvs/internal/core"
	"fmt"
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
	
	var command string
	for command != "." {
		fmt.Print("> ")
		fmt.Scanln(&command)

		switch (command) {
		case "?":
			fmt.Println("HELP")
		case "i":
			var val string
			fmt.Scan(&val)
			fmt.Printf("insert '%s'\n", val)
		case ".":
			fmt.Println("exit")
		}
	}
}



