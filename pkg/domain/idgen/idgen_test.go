package idgen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	t.Run("IDの長さは26文字である", func(t *testing.T) {
		got := len(Generate())
		assert.Equal(t, 26, got)
	})
}
