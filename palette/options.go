package palette

// Self-referential function to make palette modular and extensible
// Read more about self-referential function here:
// https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html
type Option func(p *Palette)

// Self-referential function to set the foreground style
func WithForeground(fg Style) Option {
	return func(p *Palette) {
		p.foreground = &fg
	}
}

// Self-referential function to set the background style
func WithBackground(bg Style) Option {
	return func(p *Palette) {
		p.background = &bg
	}
}

// Self-referential function to set the special effect styles
func WithSpecialEffects(effects []Special) Option {
	return func(p *Palette) {
		p.specialEffects = effects
	}
}

// Self-referential function to set the default styles
func WithDefaults() Option {
	return func(p *Palette) {
		p.specialEffects = make([]Special, 0)
		p.enabled = true
		p.background = nil
		p.foreground = nil
	}
}

// Self-referential function to enable/disable palette
func WithEnabled(e bool) Option {
	return func(p *Palette) {
		p.enabled = e
	}
}
