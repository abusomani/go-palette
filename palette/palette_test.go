package palette

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintln(t *testing.T) {
	t.Run("should test Println with expected styles", func(t *testing.T) {
		p := New(WithDefaults())
		n, err := p.Println("Sample text", "Second text")
		// one extra byte for newline
		assert.Equal(t, 24, n)
		assert.Nil(t, err)
		p.SetOptions(WithForeground(Color(32)))
		n, err = p.Println("Sample text", "Second text")
		// one extra byte for newline
		assert.Equal(t, 38, n)
		assert.Nil(t, err)
		p.SetOptions(WithBackground(Color(210)))
		n, err = p.Println("Sample text", "Second text")
		// one extra byte for newline
		assert.Equal(t, 49, n)
		assert.Nil(t, err)
		p.SetOptions(WithSpecialEffects([]Special{Bold}))
		n, err = p.Println("Sample text", "Second text")
		// one extra byte for newline
		assert.Equal(t, 53, n)
		assert.Nil(t, err)
		p.Flush()
		n, err = p.Println("Sample text", "Second text")
		// one extra byte for newline
		assert.Equal(t, 24, n)
		assert.Nil(t, err)
	})
}
func TestErrorf(t *testing.T) {
	t.Run("should test Errorf with expected styles", func(t *testing.T) {
		p := New(WithDefaults())
		err := p.Errorf("%s %s", "Sample text", "Second text")
		assert.NotNil(t, err)
		assert.Equal(t, "Sample text Second text", err.Error())
		p.SetOptions(WithBackground(Color(33)))
		err = p.Errorf("%s %s", "Sample text", "Second text")
		assert.NotNil(t, err)
		assert.Equal(t, "\x1b[48;5;33mSample text Second text\x1b[0m", err.Error())
	})
}

func TestFprint(t *testing.T) {
	t.Run("should test Fprint with expected styles", func(t *testing.T) {
		p := New(WithDefaults())
		n, err := p.Fprint(os.Stdout, "Sample text", "Second text")
		assert.Equal(t, 23, n)
		assert.Nil(t, err)
		p.SetOptions(WithForeground(Color(32)))
		n, err = p.Fprint(os.Stdout, "Sample text", "Second text")
		assert.Equal(t, 37, n)
		assert.Nil(t, err)
	})
}

func TestFprintf(t *testing.T) {
	t.Run("should test Fprintf with expected styles", func(t *testing.T) {
		p := New(WithDefaults())
		n, err := p.Fprintf(os.Stdout, "%s %s", "Sample text", "Second text")
		assert.Equal(t, 23, n)
		assert.Nil(t, err)
		p.SetOptions(WithForeground(Color(32)))
		n, err = p.Fprintf(os.Stdout, "%s %s", "Sample text", "Second text")
		assert.Equal(t, 37, n)
		assert.Nil(t, err)
	})
}

func TestFprintln(t *testing.T) {
	t.Run("should test Fprintln with expected styles", func(t *testing.T) {
		p := New(WithDefaults())
		n, err := p.Fprintln(os.Stdout, "Sample text", "Second text")
		// one extra byte for newline
		assert.Equal(t, 24, n)
		assert.Nil(t, err)
		p.SetOptions(WithForeground(Color(32)))
		n, err = p.Fprintln(os.Stdout, "Sample text", "Second text")
		// one extra byte for newline
		assert.Equal(t, 38, n)
		assert.Nil(t, err)
		p.SetOptions(WithBackground(Color(210)))
		n, err = p.Fprintln(os.Stdout, "Sample text", "Second text")
		// one extra byte for newline
		assert.Equal(t, 49, n)
		assert.Nil(t, err)
		p.SetOptions(WithSpecialEffects([]Special{Bold}))
		n, err = p.Fprintln(os.Stdout, "Sample text", "Second text")
		// one extra byte for newline
		assert.Equal(t, 53, n)
		assert.Nil(t, err)
		p.Flush()
		n, err = p.Fprintln(os.Stdout, "Sample text", "Second text")
		// one extra byte for newline
		assert.Equal(t, 24, n)
		assert.Nil(t, err)
	})
}

func TestPrint(t *testing.T) {
	t.Run("should test Print with expected styles", func(t *testing.T) {
		p := New(WithDefaults())
		n, err := p.Print("Sample text", "Second text")
		assert.Equal(t, 23, n)
		assert.Nil(t, err)
		p.SetOptions(WithForeground(Color(32)))
		n, err = p.Print("Sample text", "Second text")
		assert.Equal(t, 37, n)
		assert.Nil(t, err)
	})
}

func TestPrintf(t *testing.T) {
	t.Run("should test Printf with expected styles", func(t *testing.T) {
		p := New(WithDefaults())
		n, err := p.Printf("%s %s", "Sample text", "Second text")
		assert.Equal(t, 23, n)
		assert.Nil(t, err)
		p.SetOptions(WithForeground(Color(32)))
		n, err = p.Printf("%s %s", "Sample text", "Second text")
		assert.Equal(t, 37, n)
		assert.Nil(t, err)
	})
}

func TestSprint(t *testing.T) {
	t.Run("should test Sprint with expected styles", func(t *testing.T) {
		p := New(WithBackground(Color(33)))
		res := p.Sprint("Sample text", "Second text")
		assert.Equal(t, "\x1b[48;5;33mSample text Second text\x1b[0m", res)
	})
}

func TestSprintf(t *testing.T) {
	t.Run("should test Sprintf with expected styles", func(t *testing.T) {
		p := New(WithBackground(Color(33)))
		res := p.Sprintf("%s %s", "Sample text", "Second text")
		assert.Equal(t, "\x1b[48;5;33mSample text Second text\x1b[0m", res)
	})
}

func TestSprintfln(t *testing.T) {
	t.Run("should test Sprintln with expected styles", func(t *testing.T) {
		p := New(WithDefaults())
		res := p.Sprintln("Sample text", "Second text")
		assert.Equal(t, "Sample text Second text\n", res)
		p.SetOptions(WithForeground(Color(32)))
		res = p.Sprintln("Sample text", "Second text")
		assert.Equal(t, "\x1b[38;5;32mSample text Second text\x1b[0m\n", res)
		p.SetOptions(WithBackground(Color(210)))
		res = p.Sprintln("Sample text", "Second text")
		assert.Equal(t, "\x1b[38;5;32m\x1b[48;5;210mSample text Second text\x1b[0m\n", res)
		p.SetOptions(WithSpecialEffects([]Special{Bold}))
		res = p.Sprintln("Sample text", "Second text")
		assert.Equal(t, "\x1b[38;5;32m\x1b[48;5;210m\x1b[1mSample text Second text\x1b[0m\n", res)
		p.Flush()
		res = p.Sprintln("Sample text", "Second text")
		assert.Equal(t, "Sample text Second text\n", res)
		p.SetOptions(WithEnabled(false))
		res = p.Sprintln("Sample text", "Second text")
		assert.Equal(t, "Sample text Second text\n", res)
	})
}
