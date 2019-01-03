package z01

import (
	"errors"
	"os"
	"reflect"
	"unicode/utf8"
)

// PrintRune prints a single rune (Unicode code point) and returns any error
// if the encoding or the writing fails.
func PrintRune(r rune) error {
	l := utf8.RuneLen(r)
	if l == -1 {
		return errors.New("The rune is not a valid value to encode in UTF-8")
	}
	p := make([]byte, l)
	utf8.EncodeRune(p, r)
	_, err := os.Stdout.Write(p)
	return err
}

// Sizeof returns the number of bytes needed to store a value of the given dynamic type of i
func Sizeof(i interface{}) int {
	if i == nil {
		return 0
	}
	return int(reflect.TypeOf(i).Size())
}
