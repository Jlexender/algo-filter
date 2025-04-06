package core

import (
	"crypto/sha1"
	"encoding/binary"
)

func HashList() []func(string) int32 {
	return []func(string) int32{
		polyhash,
		sha1sum,
	}
}

func polyhash(s string) int32 {
	var hash int32 = 0
	for _, c := range s {
		hash = hash*31 + int32(c)
	}
	return hash
}

func sha1sum(s string) int32 {
	h := sha1.New()
	h.Write([]byte(s))
	sum := h.Sum(nil)
	return int32(binary.BigEndian.Uint32(sum[:4]))
}
