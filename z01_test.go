package z01_test

import (
	"fmt"

	"github.com/01-edu/z01"
)

func ExamplePrintRune() {
	z01.PrintRune('0')
	z01.PrintRune('1')
	z01.PrintRune('\n')
	z01.PrintRune('♥')
	invalidRune := rune(-1)
	err := z01.PrintRune(invalidRune)
	if err == nil {
		panic("z01.PrintRune should fail with an invalid rune")
	}
	// Output:
	// 01
	// ♥
}

func ExampleExecOut() {
	output, err := z01.ExecOut("echo", "test")
	if err == nil {
		fmt.Print(output)
	} else {
		fmt.Print(err)
	}
	// Output: test
}

func ExampleFormat() {
	a := '♥'
	fmt.Println(fmt.Sprint(a))
	fmt.Println(z01.Format(a))

	b := []interface{}{' ', " ", 5, true, nil, []int{1}, [...]int{1}}
	fmt.Println(fmt.Sprint(b...))
	fmt.Println(z01.Format(b...))

	// Output:
	// 9829
	// '♥'
	// 32 5 true <nil> [1] [1]
	// ' ', " ", 5, true, nil, []int{1}, [1]int{1}
}
