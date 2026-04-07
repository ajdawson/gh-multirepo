package cmd

import (
	"github.com/cli/cli/v2/pkg/iostreams"
	"github.com/mgutz/ansi"
)

var (
	noThemeRepoHeader    = ansi.ColorFunc("default+bu")
	darkThemeRepoHeader  = ansi.ColorFunc("white+bu")
	lightThemeRepoHeader = ansi.ColorFunc("black+bu")
)

func RepoHeader(cs *iostreams.ColorScheme, t string) string {
	if !cs.Enabled {
		return t
	}

	switch cs.Theme {
	case iostreams.DarkTheme:
		return darkThemeRepoHeader(t)
	case iostreams.LightTheme:
		return lightThemeRepoHeader(t)
	default:
		return noThemeRepoHeader(t)
	}
}
