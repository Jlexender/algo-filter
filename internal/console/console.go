package console

import "alex/bvs/internal/core"

func HandleCmd(command, path string) {
	switch command {
	case "new":
		NewFile(path)
	case "load":
		LoadFile(path)
	}

	Run(path)
}


func Run(path string) {
	bf := core.NewBloomFilter(88)

	defer SaveFile(path, bf.List())
}