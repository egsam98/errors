package errors

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestMarshalStack(t *testing.T) {
	zerolog.ErrorStackMarshaler = MarshalStack
	err := New("ERROR")
	var buf bytes.Buffer
	l := zerolog.New(&buf)
	l.Error().Stack().Err(err).Msg("Test")

	var row struct {
		Stack []logFrame `json:"stack"`
	}
	assert.NoError(t, json.NewDecoder(&buf).Decode(&row))
	if assert.NotEmpty(t, row.Stack) {
		for _, frame := range row.Stack {
			assert.NotEmpty(t, frame.File)
			assert.NotEmpty(t, frame.Function)
			assert.NotEmpty(t, frame.Line)
		}
	}
}
