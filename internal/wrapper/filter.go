package wrapper

import "alex/bvs/internal/core"

type BloomFilter[T any] struct {
	filter *core.BloomFilter
}

func NewBloomFilter[T any](size uint32) *BloomFilter[T] {
	return &BloomFilter[T]{
		filter: core.NewBloomFilter(size),
	}
}

func (bf *BloomFilter[T]) Insert(data T) {
	bf.filter.Insert(data)
}

func (bf *BloomFilter[T]) Contains(data T) bool {
	return bf.filter.Contains(data)
}

func (bf *BloomFilter[T]) Size() uint32 {
	return bf.filter.Size()
}