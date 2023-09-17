package shared

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Custom validation
func TestNewUUID(t *testing.T) {
	t.Run("", func(t *testing.T) {
		got := NewUUID()
		assert.Len(t, got, 36)
	})
}
