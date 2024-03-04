package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var ErrSome = New("some")

func TestWrapf(t *testing.T) {
	err := Wrapf(ErrSome, "prefix %d", 1)
	t.Log(err)
	assert.ErrorIs(t, err, ErrSome)

	err = Wrapf(err, "prefix %d", 2)
	t.Log(err)
	assert.ErrorIs(t, err, ErrSome)
}

func TestNew(t *testing.T) {
	err := New("some error")
	t.Log(err)
}

func TestErrorf(t *testing.T) {
	err := Errorf("prefix: error №%d", 5)
	t.Log(err)
}

func TestWrapRight(t *testing.T) {
	err := WrapRight(ErrSome, "suffix %d", 1)
	assert.ErrorIs(t, err, ErrSome)
	assert.EqualError(t, err, "some: suffix 1")

	err = WrapRight(err, "suffix %d", 2)
	assert.ErrorIs(t, err, ErrSome)
	assert.EqualError(t, err, "some: suffix 1: suffix 2")
}
