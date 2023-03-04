package utils

import "fmt"

// ANSI Escape codes for formatting terminal strings
// Refer: https://en.wikipedia.org/wiki/ANSI_escape_code
const (
	normalAnsiSequence     = "\033[%dm"
	foregroundAnsiSequence = "\033[38;5;%dm"
	backgroundAnsiSequence = "\033[48;5;%dm"
	resetColor             = "\033[0m"
)

// GetForeground returns formatted foreground color.
func GetForeground(code int32) string {
	return fmt.Sprintf(foregroundAnsiSequence, code)
}

// GetBackground returns formatted background color.
func GetBackground(code int32) string {
	return fmt.Sprintf(backgroundAnsiSequence, code)
}

// GetSpecialEffect returns formatted special effect.
func GetSpecialEffect(code int32) string {
	return fmt.Sprintf(normalAnsiSequence, code)
}

// GetResetColor returns the reset color.
func GetResetColor() string {
	return resetColor
}
