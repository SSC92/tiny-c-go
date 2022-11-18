package tinycgo

import (
	"testing"
)

func TestTiny(t *testing.T) {
	t.Run("Tiny", func(t *testing.T) {
		result := Tiny(1)
		if result != 33 {
			t.Errorf("Expected 33, got %d", result)
		}
	})
}
