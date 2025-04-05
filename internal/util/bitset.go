package util

import "fmt"

// Bitset implementation, bit sequence packed in byte slice.
// Similar to little endian, we'll store bits all in reverse
// for each chunk, just to keep simpler formula.
type Bitset struct {
	bits    []byte
	bitsize int32
}

func NewBitset(bitsize int32) *Bitset {
	return &Bitset{
		bits:    make([]byte, (bitsize+7)/8),
		bitsize: bitsize,
	}
}

func (bs *Bitset) Size() int32 {
	return bs.bitsize
}

func (bs *Bitset) List() []byte {
	return bs.bits
}

func (bs *Bitset) Set(index int32) error {
	if index >= bs.Size() {
		return fmt.Errorf("index out of range: %d", index)
	}

	bs.bits[index/8] |= (1 << (index % 8))
	return nil
}

func (bs *Bitset) Unset(index int32) error {
	if index >= bs.Size() {
		return fmt.Errorf("index out of range: %d", index)
	}

	bs.bits[index/8] &^= (1 << (index % 8))
	return nil
}

func (bs *Bitset) Toggle(index int32) error {
	if index >= bs.Size() {
		return fmt.Errorf("index out of range: %d", index)
	}

	bs.bits[index/8] ^= (1 << (index % 8))
	return nil
}

func (bs *Bitset) IsSet(index int32) (bool, error) {
	if index >= bs.Size() {
		return false, fmt.Errorf("index out of range: %d", index)
	}

	return (bs.bits[index/8] & (1 << (index % 8))) != 0, nil
}
