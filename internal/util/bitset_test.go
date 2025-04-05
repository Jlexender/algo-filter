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

func TestBitsetSet(t *testing.T) {
	bitset := NewBitset(10)
	bitset.Set(3)
	if !bitset.IsSet(3) {
		t.Errorf("Expected bit at position 3 to be set")
	}
	bitset.Set(7)
	if !bitset.IsSet(7) {
		t.Errorf("Expected bit at position 7 to be set")
	}
	if bitset.IsSet(5) {
		t.Errorf("Did not expect bit at position 5 to be set")
	}
}
func TestBitsetUnset(t *testing.T) {
	bitset := NewBitset(10)
	bitset.Set(3)
	bitset.Unset(3)
	if bitset.IsSet(3) {
		t.Errorf("Expected bit at position 3 to be unset")
	}
}

func TestBitsetToggle(t *testing.T) {
	bitset := NewBitset(10)
	bitset.Toggle(4)
	if !bitset.IsSet(4) {
		t.Errorf("Expected bit at position 4 to be set after toggle")
	}
	bitset.Toggle(4)
	if bitset.IsSet(4) {
		t.Errorf("Expected bit at position 4 to be unset after second toggle")
	}
}

func TestBitsetOutOfRange(t *testing.T) {
	bitset := NewBitset(10)
	err := bitset.Set(11)
	if err == nil {
		t.Errorf("Expected error when setting bit out of range")
	}
	err = bitset.Unset(11)
	if err == nil {
		t.Errorf("Expected error when unsetting bit out of range")
	}
	if bitset.IsSet(11) {
		t.Errorf("Did not expect bit at position 11 to be set")
	}
}

func TestBitsetEdgeCases(t *testing.T) {
	bitset := NewBitset(10)
	err := bitset.Set(9)
	if err != nil {
		t.Errorf("Did not expect error when setting bit at the last valid index")
	}
	if !bitset.IsSet(9) {
		t.Errorf("Expected bit at position 9 to be set")
	}
	err = bitset.Set(0)
	if err != nil {
		t.Errorf("Did not expect error when setting bit at the first valid index")
	}
	if !bitset.IsSet(0) {
		t.Errorf("Expected bit at position 0 to be set")
	}
}
