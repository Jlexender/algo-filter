package console

import (
	"alex/bvs/internal/util"
	"os"
)

func NewFile(path string) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
}

func LoadFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// TODO.
}

func SaveFile(path string, bs *util.Bitset) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, serr := file.Write(bs.List())
	if serr != nil {
		panic(serr)
	}
}