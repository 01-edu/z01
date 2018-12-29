package z01_test

import . "."

func ExamplePrintRune() {
	PrintRune('0')
	PrintRune('1')
	PrintRune('\n')
	PrintRune('❤')
	invalidRune := rune(-1)
	PrintRune(invalidRune)
	// Output:
	// 01
	// ❤
}
