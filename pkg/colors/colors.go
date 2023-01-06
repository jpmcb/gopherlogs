// Colors.go defines some colors that can be used by the library
// instead of fatih/colors

package colors

type Attribute int

const Escape = "\x1b"

const Reset = 0

// Foreground text colors
const (
	FgBlack Attribute = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)
