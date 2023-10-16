package part

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDI(t *testing.T) {
	t.Run("MustConfig", func(t *testing.T) {
		cfg := MustConfig()

		assert.Equal(t, "gos-part", cfg.String("app.name"))
	})
}
