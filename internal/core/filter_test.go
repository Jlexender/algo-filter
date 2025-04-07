package core

import (
	"testing"
)

func TestNewFilter(t *testing.T) {
	f := NewBloomFilter(20)
	if f == nil {
		t.Errorf("expected filter not to be nil")
	}

	if f.Size() != 20 {
		t.Errorf("expected filter size to be 20, got %d", f.Size())
	}
}

func TestStringInsert(t *testing.T) {
	f := NewBloomFilter(16)

	str := "hello!"
	f.Insert(str)

	if f.elements != 1 {
		t.Errorf("expected element count to be 1, got %d", f.elements)
	}

	if !f.Contains(str) {
		t.Errorf("expected filter to contain the element '%s', but it does not", str)
	}

	wStr := "oo"
	if f.Contains(wStr) {
		t.Logf("expected filter not to contain the element '%s' (false positive?)", wStr)
	}

	f.Insert(wStr)

	if !f.Contains(str) || !f.Contains(wStr) {
		t.Errorf("expected filter to contain both '%s' and '%s', but it does not", str, wStr)
	}
}

func TestAnyInsert(t *testing.T) {
	f := NewBloomFilter(16)

	o := []int{1, 2, 3, 88}
	f.Insert(o)

	if !f.Contains(o) {
		t.Errorf("expected filter to contain the element '%v', but it does not", o)
	}

	p := map[string]int{"a": 1, "b": 2}
	f.Insert(p)

	if !f.Contains(p) {
		t.Errorf("expected filter to contain the element '%v', but it does not", p)
	}

	if f.elements != 2 {
		t.Errorf("expected element count to be 2, got %d", f.elements)
	}

	f.Insert(o)
	if f.elements != 2 {
		t.Errorf("expected element count to remain 2 after inserting duplicate, got %d", f.elements)
	}
}

func TestCollisionResistance(t *testing.T) {
	f := NewBloomFilter(16)

	// 16 * ln 2 = 11
	if len(f.hashes) != 11 {
		t.Errorf("expected number of hash functions to be 11, got %d", len(f.hashes))
	}

	f.Insert(0x42)
	// 16 * ln 2 / 1 = 11
	if len(f.hashes) != 11 {
		t.Errorf("expected number of hash functions to be 11, got %d", len(f.hashes))
	}

	f.Insert("x+yi")
	// 16 * ln 2 / 2 = 5
	if len(f.hashes) != 5 {
		t.Errorf("expected number of hash functions to be 5, got %d", len(f.hashes))
	}
}

func TestNilContains(t *testing.T) {
	f := NewBloomFilter(16)

	if f.Contains(nil) {
		t.Errorf("expected filter not to contain nil, but it does")
	}
}

// NOTE: further tests are AI-generated.
func TestEmptyFilter(t *testing.T) {
	f := NewBloomFilter(10)

	if f.Contains("test") {
		t.Errorf("expected empty filter not to contain any elements")
	}

	if f.elements != 0 {
		t.Errorf("expected element count to be 0, got %d", f.elements)
	}
}

func TestNilInsert(t *testing.T) {
	f := NewBloomFilter(10)

	f.Insert(nil)

	if f.elements != 0 {
		t.Errorf("expected element count to remain 0 after inserting nil, got %d", f.elements)
	}
}

func TestDuplicateInsert(t *testing.T) {
	f := NewBloomFilter(10)

	data := "duplicate"
	f.Insert(data)
	f.Insert(data)

	if f.elements != 1 {
		t.Errorf("expected element count to be 1 after inserting duplicate, got %d", f.elements)
	}
}

func TestLargeDataInsert(t *testing.T) {
	f := NewBloomFilter(100)

	largeData := make([]byte, 1000)
	for i := range largeData {
		largeData[i] = byte(i % 256)
	}

	f.Insert(largeData)

	if !f.Contains(largeData) {
		t.Errorf("expected filter to contain the large data, but it does not")
	}

	if f.elements != 1 {
		t.Errorf("expected element count to be 1, got %d", f.elements)
	}
}

func TestFilterSizeZero(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic when creating filter with size 0, but it did not panic")
		}
	}()

	NewBloomFilter(0)
}

func TestInsertAfterContains(t *testing.T) {
	f := NewBloomFilter(10)

	data := "test"
	if f.Contains(data) {
		t.Errorf("expected filter not to contain '%s' before insertion", data)
	}

	f.Insert(data)

	if !f.Contains(data) {
		t.Errorf("expected filter to contain '%s' after insertion", data)
	}
}
