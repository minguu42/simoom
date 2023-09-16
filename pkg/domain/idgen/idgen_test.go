package idgen

import "testing"

func TestGenerate(t *testing.T) {
	t.Run("IDの長さは26文字である", func(t *testing.T) {
		if got := Generate(); len(got) != 26 {
			t.Errorf("ID is a 26-character string, but got %d-character string", len(got))
		}
	})
}
