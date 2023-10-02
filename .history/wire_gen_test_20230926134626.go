package part

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDI(t *testing.T) {
	t.Run("MustConfig", func(t *testing.T) {
		cfg := MustConfig()

		assert.Equal(t, "gos-part", cfg.String("app.name"))
	})
}
