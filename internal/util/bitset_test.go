package util

import "testing"

func TestNewBitset(t *testing.T) {
	bitset := NewBitset(10)
	if bitset == nil {
		t.Errorf("Expected a new Bitset instance, got nil")
	}
	if bitset.Size() != 10 {
		t.Errorf("Expected Bitset size to be 10, got %d", bitset.Size())
	}
}

