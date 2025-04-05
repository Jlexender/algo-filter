package console

import (
	"alex/bvs/internal/core"
	"log"
	"time"
)

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
	// TODO: config load
	bf := core.NewBloomFilter(88)
	defer SaveFile(path, bf.List())

	log.Printf("App is running...")	
	time.Sleep(time.Second * 5)

	log.Printf("App stopped.")
}