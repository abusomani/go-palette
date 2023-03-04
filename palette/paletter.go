package palette

import "io"

// Paletter interface.
type Paletter interface {
	// Wrapper on fmt.Errorf
	Errorf(string, ...any) error

	// Wrapper on fmt.Fprint
	Fprint(io.Writer, ...any) (n int, err error)

	// Wrapper on fmt.Fprintf
	Fprintf(io.Writer, string, ...any) (n int, err error)

	// Wrapper on fmt.Fprintln
	Fprintln(io.Writer, ...any) (n int, err error)

	// Wrapper on fmt.Print
	Print(...any) (n int, err error)

	// Wrapper on fmt.Printf
	Printf(string, ...any) (n int, err error)

	// Wrapper on fmt.Println
	Println(...any) (n int, err error)

	// Wrapper on fmt.Sprint
	Sprint(...any) string

	// Wrapper on fmt.Sprintf
	Sprintf(string, ...any) string

	// Wrapper on fmt.Sprintln
	Sprintln(...any) string
}
