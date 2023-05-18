package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewInt(t *testing.T) {
	val := NewInt(5)
	assert.Equal(t, 5, val.value)
	assert.False(t, val.null)
}

func Test_NewNull(t *testing.T) {
	val := NewNull()
	assert.Equal(t, 0, val.value)
	assert.True(t, val.null)
}
