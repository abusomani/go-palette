package palette

// Style is a wrapper.
type Style int32

// Special is a wrapper.
type Special int32

const (
	/////////////////////
	// Special Effects //
	///////////////////.//.

	//Reset.
	Reset Special = 0
	//Bold.
	Bold Special = 1
	//Dim.
	Dim Special = 2
	//Italic.
	Italic Special = 3
	//Underline.
	Underline Special = 4
	//Strikethrough.
	Strikethrough Special = 9
	//DefaultFont.
	DefaultFont Special = 10
	//AltFont1.
	AltFont1 Special = 11
	//AltFont2.
	AltFont2 Special = 12
	//AltFont3.
	AltFont3 Special = 13
	//AltFont4.
	AltFont4 Special = 14
	//AltFont5.
	AltFont5 Special = 15
	//AltFont6.
	AltFont6 Special = 16
	//AltFont7.
	AltFont7 Special = 17
	//AltFont8.
	AltFont8 Special = 18
	//AltFont9.
	AltFont9 Special = 19
)
