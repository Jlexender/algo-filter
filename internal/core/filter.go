package core

import (
	"alex/bvs/internal/util"
)

type BloomFilter struct {
	bs *util.Bitset
	H  []func(string) int32
}

func NewBloomFilter(size int32) *BloomFilter {
	return &BloomFilter{
		bs: util.NewBitset(size),
		H:  HashList(),
	}
}

func BloomFilterFromBytes(bytes []byte) *BloomFilter {
	return &BloomFilter{
		bs: util.BitsetFromBytes(bytes),
		H:  HashList(),
	}
}

func (bf *BloomFilter) Size() int32 {
	return bf.bs.Size()
}

func (bf *BloomFilter) List() *util.Bitset {
	return bf.bs
}

func (bf *BloomFilter) Insert(data string) {
	bitset := bf.List()

	for _, hash := range bf.H {
		hashsum := hash(data)
		bitset.Set(hashsum % bitset.Size())
	}
}

func (bf *BloomFilter) Exists(data string) bool {
	bitset := bf.List()

	for _, hash := range bf.H {
		hashsum := hash(data)

		set, _ := bitset.IsSet(hashsum % bitset.Size())
		if !set {
			return false
		}
	}

	return true
}
