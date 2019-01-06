package z01_test

import "github.com/01-edu/z01"

func ExamplePrintRune() {
	z01.PrintRune('0')
	z01.PrintRune('1')
	z01.PrintRune('\n')
	z01.PrintRune('♥')
	invalidRune := rune(-1)
	z01.PrintRune(invalidRune)
	// Output:
	// 01
	// ♥
}
