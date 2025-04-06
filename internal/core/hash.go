package core

import (
	"crypto/sha1"
	"encoding/binary"
)

func HashList() []func(string) uint32 {
	return []func(string) uint32{
		polyhash,
		sha1sum,
	}
}

func polyhash(s string) uint32 {
	var hash uint32 = 0
	for _, c := range s {
		hash = hash*31 + uint32(c)
	}
	return hash
}

func sha1sum(s string) uint32 {
	h := sha1.New()
	h.Write([]byte(s))
	sum := h.Sum(nil)
	return uint32(binary.BigEndian.Uint32(sum[:4]))
}
