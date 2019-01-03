package z01_test

import "github.com/01-edu/z01"

func ExamplePrintRune() {
	z01.PrintRune('0')
	z01.PrintRune('1')
	z01.PrintRune('\n')
	z01.PrintRune('❤')
	invalidRune := rune(-1)
	z01.PrintRune(invalidRune)
	// Output:
	// 01
	// ❤
}

func ExampleSizeof() {
	eightBytes := int64(0)
	tests := []bool{
		z01.Sizeof(eightBytes) == 8,
		z01.Sizeof('👌') == 4,
		z01.Sizeof(nil) == 0,
	}
	for _, test := range tests {
		if !test {
			return
		}
	}
	z01.PrintRune('👌')
	// Output: 👌
}
