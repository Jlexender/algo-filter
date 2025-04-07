package core

import (
	"alex/bvs/internal/util"
	"fmt"
	"reflect"
)

type BloomFilter struct {
	bs           *util.Bitset
	hashes       []Hash
	elements uint32
}

// Object adapter
func mapToBytes(obj any) []byte {
	return []byte(fmt.Sprintf("%v.%v", reflect.TypeOf(obj), obj))
}

func NewBloomFilter(size uint32) *BloomFilter {
	if size == 0 {
		panic("size must be greater than 0")
	}

	return &BloomFilter{
		bs:           util.NewBitset(size),
		hashes:       NewHashList(size),
		elements: 0,
	}
}

func (bf *BloomFilter) Insert(data any) {
	if data == nil || bf.Contains(data) {
		return
	}

	bitset := bf.bs
	bf.elements++
	for _, hash := range bf.hashes {
		hashsum := hash.Compute(mapToBytes(data))
		bitset.Set(hashsum % bitset.Size())
		bf.hashes = UpdateList(bf.hashes, bf.Size(), bf.elements)
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
