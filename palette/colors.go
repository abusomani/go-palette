package palette

// Color is a wrapper.
type Color int32

// Special is a wrapper.
type Special int32

const (

	/////////////////////
	// Standard Colors //
	/////////////////////

	//Black.
	Black Color = 0
	//Red.
	Red Color = 1
	//Green.
	Green Color = 2
	//Yellow.
	Yellow Color = 3
	//Blue.
	Blue Color = 4
	//Magenta.
	Magenta Color = 5
	// Cyan.
	Cyan Color = 6
	//White.
	White Color = 7

	////////////////////////////
	// Standard Bright Colors //
	////////////////////////////

	//BrightBlack.
	BrightBlack Color = 8
	//BrightRed.
	BrightRed Color = 9
	//BrightGreen.
	BrightGreen Color = 10
	//BrightYellow.
	BrightYellow Color = 11
	//BrightBlue.
	BrightBlue Color = 12
	//BrightMagenta.
	BrightMagenta Color = 13
	//Bright Cyan.
	BrightCyan Color = 14
	//BrightWhite.
	BrightWhite Color = 15

	/////////////////////
	// Special Effects //
	/////////////////////

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
	//SlowBlink.
	SlowBlink Special = 5
	//Hidden.
	Hidden Special = 8
	//Strikethrough.
	Strikethrough Special = 9
	//DefaultFont.
	DefaultFont Special = 10
)
