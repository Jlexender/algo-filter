package console

import (
	"os"
)

func NewFile(path string) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
}

func LoadFile(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return data
}

func SaveFile(path string, bytes []byte) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, serr := file.Write(bytes)
	if serr != nil {
		panic(serr)
	}
}
