package core

import (
	"alex/bvs/internal/util"
	"fmt"
)

type BloomFilter struct {
	bs           *util.Bitset
	hashes       []Hash
	elementCount uint32
}

// Object adapter
func mapToBytes(obj any) []byte {
	return []byte(fmt.Sprintf("%v", obj))
}

func NewBloomFilter(size uint32) *BloomFilter {
	if size == 0 {
		panic("size must be greater than 0")
	}

	return &BloomFilter{
		bs:           util.NewBitset(size),
		hashes:       NewHashList(size),
		elementCount: 0,
	}
}

func (bf *BloomFilter) Insert(data any) {
	if data == nil || bf.Contains(data) {
		return
	}

	bitset := bf.bs
	bf.elementCount++
	for _, hash := range bf.hashes {
		hashsum := hash.Compute(mapToBytes(data))
		bitset.Set(hashsum % bitset.Size())
		bf.hashes = UpdateList(bf.hashes, bf.Size(), bf.elementCount)
	}
}

func (bf *BloomFilter) Contains(data any) bool {
	bitset := bf.bs

	for _, hash := range bf.hashes {
		hashsum := hash.Compute(mapToBytes(data))

		set, _ := bitset.IsSet(hashsum % bitset.Size())
		if !set {
			return false
		}
	}

	return true
}

func (bf *BloomFilter) Size() uint32 {
	return bf.bs.Size()
}
