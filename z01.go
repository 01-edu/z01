package z01

import "os"

// PrintRune prints a single rune (Unicode code point)
func PrintRune(r rune) {
	os.Stdout.Write([]byte(string(r)))
}
