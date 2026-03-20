package iostreams

import (
	"github.com/mgutz/ansi"
)

const (
	NoTheme    = "none"
	DarkTheme  = "dark"
	LightTheme = "light"
)

var (
	noThemeRepoHeader    = ansi.ColorFunc("default+bu")
	darkThemeRepoHeader  = ansi.ColorFunc("white+bu")
	lightThemeRepoHeader = ansi.ColorFunc("black+bu")
)

type ColorScheme struct {
	Enabled bool
	Theme   string
}

func (c *ColorScheme) RepoHeader(t string) string {
	if !c.Enabled {
		return t
	}

	switch c.Theme {
	case DarkTheme:
		return darkThemeRepoHeader(t)
	case LightTheme:
		return lightThemeRepoHeader(t)
	default:
		return noThemeRepoHeader(t)
	}
}
