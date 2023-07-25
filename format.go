// Based on https://github.com/rs/zerolog/blob/4612e098/console.go
package main

import (
	"fmt"

	"github.com/rs/zerolog"
)

const (
	colorBlack = iota + 30
	colorRed
	colorGreen
	colorYellow
	colorBlue
	colorMagenta
	colorCyan
	colorWhite

	styleNormal = 0
	styleBold   = 1
)

// colorize returns the string s wrapped in ANSI code c, unless disabled is true.
func colorize(s interface{}, style, color int, disabled bool) string {
	if disabled {
		return fmt.Sprintf("%s", s)
	}
	return fmt.Sprintf("\x1b[%d;%dm%v\x1b[0m", style, color, s)
}

func formatLevel(noColor bool) zerolog.Formatter {
	return func(i interface{}) string {
		var l string
		if ll, ok := i.(string); ok {
			switch ll {
			case zerolog.LevelTraceValue:
				l = colorize("TRC", styleNormal, colorMagenta, noColor)
			case zerolog.LevelDebugValue:
				l = colorize("DBG", styleNormal, colorBlue, noColor)
			case zerolog.LevelInfoValue:
				l = colorize("INF", styleNormal, colorGreen, noColor)
			case zerolog.LevelWarnValue, "warning":
				l = colorize("WRN", styleNormal, colorYellow, noColor)
			case zerolog.LevelErrorValue:
				l = colorize("ERR", styleNormal, colorRed, noColor)
			case zerolog.LevelFatalValue:
				l = colorize("FTL", styleNormal, colorRed, noColor)
			case zerolog.LevelPanicValue:
				l = colorize("PNC", styleNormal, colorRed, noColor)
			default:
				l = colorize(ll, styleBold, colorWhite, noColor)
			}
		}
		return l
	}
}
