package errors_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/egsam98/errors"
)

var ErrSome = errors.New("some")

func TestWrapf(t *testing.T) {
	err := errors.Wrapf(ErrSome, "prefix %d", 1)
	t.Log(err)
	assert.ErrorIs(t, err, ErrSome)

	err = errors.Wrapf(err, "prefix %d", 2)
	t.Log(err)
	assert.ErrorIs(t, err, ErrSome)
}

func TestNew(t *testing.T) {
	err := errors.New("some error")
	t.Log(err)
}

func TestErrorf(t *testing.T) {
	err := errors.Errorf("prefix: error №%d", 5)
	t.Log(err)
}
