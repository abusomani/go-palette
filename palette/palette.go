package palette

import (
	"fmt"
	"io"

	"github.com/abusomani/go-palette/utils"
)

// Implements the Paletter interface
type Palette struct {
	// if the palette is enabled
	enabled bool
	// foreground decides the color of the text
	foreground *Style
	// background decides the background color of the text
	background *Style
	// adds special effects like bold, italic, underline etc. to the text
	specialEffects []Special
}

// New returns a new palette instance with variable options to set the styling configurations
func New(opts ...Option) *Palette {
	p := &Palette{}

	// set the default configurations
	defaults := WithDefaults()
	defaults(p)

	// apply the options
	for _, opt := range opts {
		opt(p)
	}
	return p
}

// Errorf formats according to a format specifier and decorates the input with the styling configurations.
// It returns the string as a value that satisfies error.
func (p *Palette) Errorf(input string, a ...any) error {
	return fmt.Errorf(wrap(p, fmt.Sprintf(input, a...)))
}

// Fprint formats using the default formats and decorates the input with the styling configurations. Finally writes to w.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func (p *Palette) Fprint(w io.Writer, input ...any) (n int, err error) {
	return fmt.Fprint(w, wrap(p, input...))
}

// Fprintf formats according to a format specifier and decorates the input with the styling configurations. Finally writes to w.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func (p *Palette) Fprintf(w io.Writer, input string, a ...any) (n int, err error) {
	return fmt.Fprint(w, wrap(p, fmt.Sprintf(input, a...)))
}

// Fprintln formats using the default formats for its operands and decorates the input with the styling configurations. Finally writes to w.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func (p *Palette) Fprintln(w io.Writer, input ...any) (n int, err error) {
	return fmt.Fprintln(w, wrap(p, input...))
}

// Print formats using the default formats for its operands and decorates the input with the styling configurations. Finally writes to standard output.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func (p *Palette) Print(input ...any) (n int, err error) {
	return fmt.Print(wrap(p, input...))
}

// Printf formats according to a format specifier and decorates the input with the styling configurations. Finally writes to standard output.
// It returns the number of bytes written and any write error encountered.
func (p *Palette) Printf(input string, a ...any) (n int, err error) {
	return fmt.Print(wrap(p, fmt.Sprintf(input, a...)))
}

// Println formats using the default formats for its operands and decorates the input with the styling configurations. Finally writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func (p *Palette) Println(input ...any) (n int, err error) {
	return fmt.Println(wrap(p, input...))
}

// Sprint formats using the default formats for its operands and decorates the input with the styling configurations before returning the resulting string.
// Spaces are added between operands when neither is a string.
func (p *Palette) Sprint(input ...any) string {
	return fmt.Sprint(wrap(p, input...))
}

// Printf formats according to a format specifier and decorates the input with the styling configurations before returning the resulting string.
func (p *Palette) Sprintf(input string, a ...any) string {
	return wrap(p, fmt.Sprintf(input, a...))
}

// Sprintln formats using the default formats for its operands and decorates the input with the styling configurations before returning the resulting string.
// Spaces are always added between operands and a newline is appended.
func (p *Palette) Sprintln(input ...any) string {
	return fmt.Sprintln(wrap(p, input...))
}

// SetOptions adds self-referential options to change the properties of Palette
func (p *Palette) SetOptions(opts ...Option) *Palette {
	for _, opt := range opts {
		opt(p)
	}
	return p
}

// Flush resets the Palette options with default values and disables the Palette
func (p *Palette) Flush() *Palette {
	// set the default configurations
	df := WithDefaults()
	df(p)
	// disable palette use until activated
	p.enabled = false
	return p
}

// getEscapeSequence formats the input string based on palette configurations
func getEscapeSequence(p *Palette, input string) string {
	// if palette configuration is not enabled then return the input as is
	if !p.enabled {
		return input
	}

	// to check if the input is decorated based on palette configuration
	isModified := false
	// resulting string
	res := ""

	// apply foreground style
	if p.foreground != nil {
		isModified = true
		res += utils.GetForeground(int32(*p.foreground))
	}

	// apply background style
	if p.background != nil {
		isModified = true
		res += utils.GetBackground(int32(*p.background))
	}

	// apply special effects style
	if len(p.specialEffects) > 0 {
		for _, se := range p.specialEffects {
			isModified = true
			res += utils.GetSpecialEffect(int32(se))
		}
	}
	// Add the original input string
	res += input

	// if the resulting string had any of the style decorator then we need to add a reset ansi sequence
	// the reset ansi code resets all colors and text effects
	if isModified {
		res += utils.GetResetColor()
	}
	return res
}

// wrap returns a styled string based on palette configuration
func wrap(p *Palette, input ...any) string {
	res := ""
	for idx, s := range input {
		if idx > 0 {
			res += " "
		}
		res += fmt.Sprint(s)
	}
	return getEscapeSequence(p, res)
}
