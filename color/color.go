package color

// source: https://ahmadrosid.com/cheatsheet/go/golang-terminal-color

import "fmt"

const escape = '\x1b'

const (
	NONE = iota
	RED
	GREEN
	YELLOW
	BLUE
	PURPLE
	CYAN
)

func color(c int) string {
	if c == NONE {
		return fmt.Sprintf("%c[%dm", escape, c)
	}
	return fmt.Sprintf("%c[3%dm", escape, c)
}

func Format(c int, text string) string {
	return color(c) + text + color(NONE)
}
