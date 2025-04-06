package util

import "testing"

func TestNewBitset(t *testing.T) {
	bs := NewBitset(10)

	if bs == nil {
		t.Errorf("expected bitset, got nil")
	}

	if bs.Size() != 10 {
		t.Errorf("expected size=10, got %d", bs.Size())
	}
}

func TestSetBit(t *testing.T) {
	bs := NewBitset(8)

	bs.Set(0)

	if bs.bits[0] != 1 {
		t.Errorf("expected bits[0]=1, got %d", bs.bits[0])
	}

	bs.Set(7)

	if bs.bits[0] != 129 {
		t.Errorf("expected bits[0]=129, got %d", bs.bits[0])
	}
}

func TestUnset(t *testing.T) {
	bs := BitsetFromBytes([]byte{0xFF})

	t.Logf("Bitset: %v", bs.bits)

	bs.Unset(0)
	if bs.bits[0] != 254 {
		t.Errorf("expected bits[0]=254, got %d", bs.bits[0])
	}

	bs.Unset(6)
	if bs.bits[0] != 190 {
		t.Errorf("expected bits[0]=190, got %d", bs.bits[0])
	}
}
func TestIsSet(t *testing.T) {
	bs := NewBitset(8)

	bs.Set(0)
	isSet, err := bs.IsSet(0)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !isSet {
		t.Errorf("expected bit 0 to be set, but it is not")
	}

	isSet, err = bs.IsSet(7)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if isSet {
		t.Errorf("expected bit 7 to be unset, but it is set")
	}

	_, err = bs.IsSet(8)
	if err == nil {
		t.Errorf("expected error for out-of-range index, got nil")
	}
}
