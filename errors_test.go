package errors

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var errSome = New("some")

func TestNew(t *testing.T) {
	err := New("some error")
	assert.EqualError(t, err, "some error")
}

func TestErrorf(t *testing.T) {
	err := Errorf("prefix: error №%d", 5)
	assert.EqualError(t, err, "prefix: error №5")
}

func TestWrap(t *testing.T) {
	err := Wrap(errSome, "prefix %d", 1)
	assert.ErrorIs(t, err, errSome)
	assert.True(t, strings.HasPrefix(err.Error(), "prefix 1:"))

	err = Wrap(err, "prefix %d", 2)
	assert.ErrorIs(t, err, errSome)
	assert.True(t, strings.HasPrefix(err.Error(), "prefix 2: prefix 1:"))
}

func TestWrapR(t *testing.T) {
	err := WrapR(errSome, "suffix %d", 1)
	assert.ErrorIs(t, err, errSome)
	assert.EqualError(t, err, "some: suffix 1")

	err = WrapR(err, "suffix %d", 2)
	assert.ErrorIs(t, err, errSome)
	assert.EqualError(t, err, "some: suffix 1: suffix 2")
}
