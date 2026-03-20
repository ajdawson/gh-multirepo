package iostreams

import (
	"io"

	"github.com/cli/go-gh/v2/pkg/term"
)

type IO struct {
	term term.Term
}

func NewIO(t term.Term) *IO {
	return &IO{
		term: t,
	}
}

func (io *IO) IsTTY() bool {
	return io.term.IsTerminalOutput()
}

func (io *IO) ColorScheme() *ColorScheme {
	return &ColorScheme{
		Enabled: io.term.IsColorEnabled(),
		Theme:   io.term.Theme(),
	}
}

func (io *IO) Out() io.Writer {
	return io.term.Out()
}
