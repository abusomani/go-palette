package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetForeground(t *testing.T) {
	t.Run("should return the appropriate foreground color", func(t *testing.T) {
		fg := GetForeground(32)
		assert.Equal(t, "\033[38;5;32m", fg)
	})
}

func TestGetBackground(t *testing.T) {
	t.Run("should return the appropriate background color", func(t *testing.T) {
		bg := GetBackground(32)
		assert.Equal(t, "\033[48;5;32m", bg)
	})

}

func TestGetSpecialEffect(t *testing.T) {
	t.Run("should return the appropriate special effect", func(t *testing.T) {
		se := GetSpecialEffect(2)
		assert.Equal(t, "\033[2m", se)
	})
}

func TestGetResetColor(t *testing.T) {
	t.Run("should return the appropriate reset color", func(t *testing.T) {
		rc := GetResetColor()
		assert.Equal(t, "\033[0m", rc)
	})
}
