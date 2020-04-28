package jmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringHash(t *testing.T)  {
	h, err := NewHash(256)
	assert.NoError(t, err)
	assert.Less(t, h.stringHash("Jimmy Huang"), 257)
}

func TestStringHashCollision(t *testing.T)  {
	h, err := NewHash(256)
	assert.NoError(t, err)
	assert.Equal(t, h.stringHash("ab"), h.stringHash("ba"))
}

func TestIsPowerOfTwo(t *testing.T) {
	assert.Equal(t, isPowerOfTwo(256), true)
	assert.Equal(t, isPowerOfTwo(1024), true)
	assert.Equal(t, isPowerOfTwo(12), false)
}

func BenchmarkStringHash(b *testing.B) {
	h, _ := NewHash(256)
    for i := 0; i < b.N; i++ {
		h.stringHash("Jimmy Huang")
    }
}
