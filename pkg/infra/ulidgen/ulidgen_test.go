package ulidgen_test

import (
	"testing"

	"github.com/minguu42/simoom/pkg/infra/ulidgen"
	"github.com/stretchr/testify/assert"
)

func TestGenerator_Generate(t *testing.T) {
	g := ulidgen.Generator{}
	t.Run("IDの長さは26文字である", func(t *testing.T) {
		got := len(g.Generate())
		assert.Equal(t, 26, got)
	})
}
