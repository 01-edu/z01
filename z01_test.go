package z01_test

import (
	"testing"

	"github.com/01-edu/z01"
)

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

func TestSizeof(t *testing.T) {
	tests := []struct {
		i    interface{}
		size int
	}{
		{"azefoiazhefzaef", 16},
		{int64(0), 8},
		{'👌', 4},
		{int16(0), 2},
		{false, 1},
		{nil, 0},
	}
	for _, test := range tests {
		actual := z01.Sizeof(test.i)
		if actual != test.size {
			t.Errorf("z01.Sizeof(%#v) == %d instead of %d\n", test.i, test.size, actual)
		}
	}
}
