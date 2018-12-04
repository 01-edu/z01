package z01

import "os"

// Putchar prints a single ASCII character (if 0 <= c <= 127)
func Putchar(c byte) {
	os.Stdout.Write([]byte{c})
}
