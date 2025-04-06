package console

import (
	"alex/bvs/internal/core"
	"bytes"
	"fmt"
	"log"
	"time"
)

func HandleCmd(command, path string) {
	var bytes []byte

	switch command {
	case "new":
		NewFile(path)
		var fsize int32
		fmt.Scanf("Please, enter the filter size: %d", &fsize)

		bytes = make([]byte, (fsize+7)/8)
	case "load":
		bytes = LoadFile(path)
	}

	defer SaveFile(path, bytes)
	Run(bytes)
}

func Run(bytes []byte) {
	bf := core.BloomFilterFromBytes(bytes)
	

	log.Printf("App is running...")
	time.Sleep(time.Second * 5)

	log.Printf("App stopped.")
}
