package core

import "alex/bvs/internal/util"

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

func (bf *BloomFilter) Size() int32 {
	return bf.bs.Size()
}

func (bf *BloomFilter) Insert(data string) {
	for _, hash := range bf.H {
		bf.bs.Set(hash(data))
	}
}

func (bf *BloomFilter) Exists(data string) bool {
	for _, hash := range bf.H {
		v := hash(data)
		if !bf.bs.IsSet(v) {
			return false
		}
	}
	return true
}
