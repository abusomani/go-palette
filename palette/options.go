package palette

// Option is a self-referential function to make palette modular and extensible.
// Read more about self-referential function here:
// https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html
type Option func(p *Palette)

// WithForeground is a self-referential function to set the foreground color.
func WithForeground(fg Color) Option {
	return func(p *Palette) {
		p.foreground = &fg
	}
}

// WithBackground is a self-referential function to set the background color.
func WithBackground(bg Color) Option {
	return func(p *Palette) {
		p.background = &bg
	}
}

// WithSpecialEffects is a self-referential function to set the special effect styles.
func WithSpecialEffects(effects []Special) Option {
	return func(p *Palette) {
		p.specialEffects = effects
	}
}

// WithDefaults is a self-referential function to set the default styles.
func WithDefaults() Option {
	return func(p *Palette) {
		p.specialEffects = make([]Special, 0)
		p.enabled = true
		p.background = nil
		p.foreground = nil
	}
}

// WithEnabled is a self-referential function to enable/disable palette.
func WithEnabled(e bool) Option {
	return func(p *Palette) {
		p.enabled = e
	}
}
