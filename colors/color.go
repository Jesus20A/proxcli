package colors

import "fmt"

const (
	ColorDefault = "\x1b[39m"
	ColorYellow  = "\x1b[93m"
	ColorRed     = "\x1b[91m"
	ColorGreen   = "\x1b[32m"
	ColorBlue    = "\x1b[94m"
	ColorGray    = "\x1b[90m"
	ColorWhite   = "\x1b[97m"
)

func Red(s string) string {
	return fmt.Sprintf("%s%s%s", ColorRed, s, ColorDefault)
}

func Green(s string) string {
	return fmt.Sprintf("%s%s%s", ColorGreen, s, ColorDefault)
}

func White(s string) string {
	return fmt.Sprintf("%s%s%s", ColorWhite, s, ColorDefault)
}

func Blue(s string) string {
	return fmt.Sprintf("%s%s%s", ColorBlue, s, ColorDefault)
}

func Yellow(s string) string {
	return fmt.Sprintf("%s%s%s", ColorYellow, s, ColorDefault)
}

func Color(state, text string) string {

	switch state {
	case "running":
		return fmt.Sprintf("\u2705%s", Green(text))
	case "stopped":
		return fmt.Sprintf("\u274C%s", Red(text))

	}
	return text
}
