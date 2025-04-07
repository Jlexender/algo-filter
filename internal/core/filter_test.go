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

func TestInsertSameSprintf(t *testing.T) {
	f := NewBloomFilter(128)

	f.Insert("[1 2 3]")
	if f.Contains([]byte{1, 2, 3}) {
		t.Errorf("expected filter not to contain the byte slice '[1 2 3]', but it does")
	}

	type pair struct {
		a int
		b string
	}

	p1 := pair{a: 1, b: "foo"}
	p2 := pair{a: 1, b: "bar"}

	f.Insert(p1)

	if !f.Contains(p1) {
		t.Errorf("expected filter to contain '%v', but it does not", p1)
	}

	if f.Contains(p2) {
		t.Errorf("expected filter not to contain '%v', but it does", p2)
	}

	slice1 := []int{1, 2, 3}
	slice2 := []int{3, 2, 1}

	f.Insert(slice1)

	if !f.Contains(slice1) {
		t.Errorf("expected filter to contain '%v', but it does not", slice1)
	}

	if f.Contains(slice2) {
		t.Errorf("expected filter not to contain '%v', but it does", slice2)
	}

	map1 := map[string]int{"x": 1, "y": 2}
	map2 := map[string]int{"x": 2, "y": 1}

	f.Insert(map1)

	if !f.Contains(map1) {
		t.Errorf("expected filter to contain '%v', but it does not", map1)
	}

	if f.Contains(map2) {
		t.Errorf("expected filter not to contain '%v', but it does", map2)
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

func TestObjectAdapterStress(t *testing.T) {
    f := NewBloomFilter(16384)

    complexObjects := []any{
        "[1 2 3]",
        []byte{1, 2, 3},
        123,
        "hello world",

        struct {
            A int
            B string
        }{A: 42, B: "foo"},
        struct {
            A []int
            B map[string]int
        }{A: []int{1, 2, 3}, B: map[string]int{"x": 10, "y": 20}},

        []int{1, 2, 3, 4, 5},
        []int{5, 4, 3, 2, 1},

        map[string]int{"a": 1, "b": 2, "c": 3},
        map[string]int{"c": 3, "b": 2, "a": 1},

        map[string]any{
            "numbers": []int{10, 20, 30},
            "nested": map[string]string{"foo": "bar", "baz": "qux"},
        },
        map[string]any{
            "numbers": []int{10, 20, 30},
            "nested": map[string]string{"baz": "qux", "foo": "bar"},
        },

        struct {
            X []float64
            Y string
        }{X: []float64{3.14, 2.718, 1.618}, Y: "constants"},
        struct {
            X []float64
            Y string
        }{X: []float64{1.618, 2.718, 3.14}, Y: "constants"},
    }

    for _, obj := range complexObjects {
        f.Insert(mapToBytes(obj))
    }

    for i := 0; i < 1000; i++ {
        obj := complexObjects[i%len(complexObjects)]
        if !f.Contains(mapToBytes(obj)) {
            t.Errorf("Iteration %d: expected filter to contain '%v', but it does not", i, obj)
        }
    }

    type testPair struct {
        a int
        b string
    }

    p1 := testPair{a: 1, b: "foo"}
    f.Insert(mapToBytes(p1))

    p2 := testPair{a: 1, b: "bar"}
    if f.Contains(mapToBytes(p2)) {
        t.Errorf("expected filter not to contain '%v', but it does", p2)
    }

    largeSlice := make([]int, 1000)
    for i := 0; i < 1000; i++ {
        largeSlice[i] = i
    }
    f.Insert(mapToBytes(largeSlice))
    if !f.Contains(mapToBytes(largeSlice)) {
        t.Errorf("expected filter to contain largeSlice, but it does not")
    }

    alteredSlice := make([]int, 1000)
    copy(alteredSlice, largeSlice)
    alteredSlice[500] = -1
    if f.Contains(mapToBytes(alteredSlice)) {
        t.Errorf("expected filter not to contain alteredSlice, but it does")
    }
}
