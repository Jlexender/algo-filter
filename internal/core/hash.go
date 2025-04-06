package core

import (
	"crypto/sha1"
	"encoding/binary"
)

func HashList() []func(string) int32 {
	return []func(string) int32{
		adolf,
		murmur,
		kiskis,
		sha1hash,
	}
}

func adolf(s string) int32 {
	var hash int32 = 0
	for _, c := range s {
		hash += ((c >> 14) * 0x1488) ^ 0x666
	}
	return hash
}

func murmur(s string) int32 {
	var hash int32 = 0
	for _, c := range s {
		hash = hash*31 + int32(c)
	}
	return hash
}

func kiskis(s string) int32 {
	var hash int32 = 0
	for i, c := range s {
		hash ^= int32(c) * int32(i+1)
		hash = (hash << 5) | (hash >> 27)
	}
	return hash
}

func sha1hash(s string) int32 {
	h := sha1.New()
	h.Write([]byte(s))
	sum := h.Sum(nil)
	return int32(binary.BigEndian.Uint32(sum[:4]))
}
