package jmap

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewJMap(t *testing.T) {
	h, err := NewJmap(256)
	assert.NoError(t, err)
	assert.NotEqual(t, h, nil)
}

func TestFailToNewJMap(t *testing.T) {
	_, err := NewJmap(255)
	assert.Error(t, err)
}

func TestSetJMap(t *testing.T) {
	m, err := NewJmap(256)
	assert.NoError(t, err)
	m.Set("1", "a")
	fmt.Printf("map: %+v:", m)
}

func TestSetJMapCollision(t *testing.T) {
	m, err := NewJmap(256)
	assert.NoError(t, err)
	m.Set("12", "a")
	m.Set("21", "b")
	fmt.Printf("map: %+v:", m)
}

func TestSetJMapFull(t *testing.T) {
	m, err := NewJmap(256)
	assert.NoError(t, err)
	for i := 0; i < 256; i++ {
		m.Set(strconv.Itoa(i), "a")
	}
}

func TestSetJMapFullPlusOne(t *testing.T) {
	m, err := NewJmap(256)
	assert.NoError(t, err)
	for i := 0; i < 257; i++ {
		m.Set(strconv.Itoa(i), "a")
	}
	fmt.Printf("map: %+v:", m)
}

func TestGetJMap(t *testing.T) {
	m, err := NewJmap(256)
	assert.NoError(t, err)
	for i := 0; i < 256; i++ {
		m.Set(strconv.Itoa(i), strconv.Itoa(i))
	}
	for i := 0; i < 256; i++ {
		v, ok := m.Get(strconv.Itoa(i))
		assert.True(t, ok)
		assert.Equal(t, v, strconv.Itoa(i))
	}
}
