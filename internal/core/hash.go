package core

func HashList() []func(string) int32 {
	return []func(string) int32{
		mock,
	}
}

// Note: this is a mock hash function. 
// Use it only for dev!
func mock(s string) int32 {
	return 1488_666
}
