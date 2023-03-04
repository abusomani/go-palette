package palette

import "io"

type Paletter interface {
	Errorf(string, ...any) error
	Fprint(io.Writer, ...any) (n int, err error)
	Fprintf(io.Writer, string, ...any) (n int, err error)
	Fprintln(io.Writer, ...any) (n int, err error)
	Print(...any) (n int, err error)
	Printf(string, ...any) (n int, err error)
	Println(...any) (n int, err error)
	Sprint(...any) string
	Sprintf(string, ...any) string
	Sprintln(...any) string
}
