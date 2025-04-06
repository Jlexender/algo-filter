package core

import (
	"testing"
)

func TestNewFilter(t *testing.T) {
	filter := NewBloomFilter(20)
	if filter == nil {
		t.Errorf("expected filter not to be nil")
	}

	if filter.Size() != 20 {
		t.Errorf("expected filter size to be 20, got %d", filter.Size())
	}
}

func TestInsert(t *testing.T) {
	filter := NewBloomFilter(64)

	inputString := "abacaba"
	wrongString := "aboba"

	filter.Insert(inputString)

	if !filter.Exists(inputString) {
		t.Errorf("expected '%s' to be inside the filter", inputString)
	}

	if filter.Exists(wrongString) {
		t.Logf("expected '%s' to be not in the filter (weak hash?)", wrongString)
	}
}
func TestEmptyFilter(t *testing.T) {
	filter := NewBloomFilter(64)

	inputString := "test"

	if filter.Exists(inputString) {
		t.Errorf("expected '%s' to not exist in an empty filter", inputString)
	}
}

func TestMultipleInsertions(t *testing.T) {
	filter := NewBloomFilter(64)

	strings := []string{"apple", "banana", "cherry", "date"}
	for _, s := range strings {
		filter.Insert(s)
	}

	for _, s := range strings {
		if !filter.Exists(s) {
			t.Errorf("expected '%s' to exist in the filter", s)
		}
	}

	wrongString := "grape"
	if filter.Exists(wrongString) {
		t.Logf("expected '%s' to not exist in the filter (false positive?)", wrongString)
	}
}

func TestFilterSize(t *testing.T) {
	filter := NewBloomFilter(128)

	if filter.Size() != 128 {
		t.Errorf("expected filter size to be 128, got %d", filter.Size())
	}
}

func TestCollisionHandling(t *testing.T) {
	filter := NewBloomFilter(8)

	filter.Insert("foo")
	filter.Insert("bar")

	if !filter.Exists("foo") {
		t.Errorf("expected 'foo' to exist in the filter")
	}

	if !filter.Exists("bar") {
		t.Errorf("expected 'bar' to exist in the filter")
	}
}



