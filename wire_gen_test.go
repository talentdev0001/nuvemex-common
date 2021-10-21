package part

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDI(t *testing.T) {
	t.Run("ProvideLibConfig", func(t *testing.T) {
		currentAppEnv := os.Getenv("app_env")
		currentLogWriters := os.Getenv("log_writers")
		currentLogLevel := os.Getenv("log_level")

		_ = os.Setenv("app_env", "xxx")
		_ = os.Setenv("log_writers", "")
		_ = os.Setenv("log_level", "x")

		appConfig := ProvideAppConfig(RootPath + "/resources/config")
		cfg := ProvideLibConfig(appConfig)

		assert.Equal(t, "goseanto", cfg.String("app.name"))
		assert.Equal(t, "xxx", cfg.String("app.env"))
		assert.Equal(t, "", cfg.String("log.writers"))
		assert.Equal(t, "x", cfg.String("log.level"))

		_ = os.Setenv("app_env", currentAppEnv)
		_ = os.Setenv("log_writers", currentLogWriters)
		_ = os.Setenv("log_level", currentLogLevel)

	})

	t.Run("ProvideAppConfig", func(t *testing.T) {
		cfg := ProvideAppConfig(RootPath + "/resources/config")

		assert.Equal(t, "gos-part", cfg.String("app.name"))
	})
}
