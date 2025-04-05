package core

import "testing"

func TestNewFilter(t *testing.T) {
	filter := NewBloomFilter(20)
	if filter == nil {
		t.Errorf("expected filter not to be nil")
	}

	if filter.Size() != 20 {
		t.Errorf("expected filter size to be 20, got %d", filter.Size())
	}
}

