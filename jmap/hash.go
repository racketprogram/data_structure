package jmap

import "fmt"

type hash struct {
	hashRange int
}

func NewHash(hashRange int) (*hash, error) {
	if !isPowerOfTwo(hashRange) {
		return nil, fmt.Errorf("not power of two")
	}
	return &hash {
		hashRange: hashRange,
	}, nil
}

func (h *hash) stringHash(key string) int {
	byteArray := []byte(key)
	var n int
	for _, v := range byteArray {
		n += int(v)
	}
	return h.modPowerOfTwo(n)
}

func (h *hash) modPowerOfTwo(number int) int {
	return (number & (h.hashRange - 1))
}

func isPowerOfTwo(number int) bool {
	return (number & (number - 1)) == 0
}
