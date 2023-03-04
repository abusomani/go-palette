package palette

// Wrapper for Styles
type Style int32

// Wrapper for Special effects
type Special int32

const (
	/////////////////////
	// Special Effects //
	/////////////////////
	Reset         Special = 0
	Bold          Special = 1
	Dim           Special = 2
	Italic        Special = 3
	Underline     Special = 4
	Strikethrough Special = 9
	DefaultFont   Special = 10
	AltFont1      Special = 11
	AltFont2      Special = 12
	AltFont3      Special = 13
	AltFont4      Special = 14
	AltFont5      Special = 15
	AltFont6      Special = 16
	AltFont7      Special = 17
	AltFont8      Special = 18
	AltFont9      Special = 19
)
