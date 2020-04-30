package jmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewJMap(t *testing.T) {
	h, err := NewJmap(256)
	assert.NoError(t, err)
	assert.NotEqual(t, h, nil)
}
