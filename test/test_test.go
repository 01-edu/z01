package test_test

import (
	"fmt"
	"testing"

	"github.com/01-edu/z01/test"
)

func ExampleExecOut() {
	output, err := test.ExecOut("echo", "test")
	if err == nil {
		fmt.Print(output)
	} else {
		fmt.Print(err)
	}
	// Output: test
}

func TestNameOfFunc(t *testing.T) {
	if test.NameOfFunc(test.NameOfFunc) != "NameOfFunc" {
		t.Error("NameOfFunc failed to get its own name.")
	}
}

func ExampleFormat() {
	a := '♥'
	fmt.Println(fmt.Sprint(a))
	fmt.Println(test.Format(a))

	b := []interface{}{' ', " ", 5, true, nil, []int{1}, [...]int{1}}
	fmt.Println(fmt.Sprint(b...))
	fmt.Println(test.Format(b...))

	// Output:
	// 9829
	// '♥'
	// 32 5 true <nil> [1] [1]
	// ' ', " ", 5, true, nil, []int{1}, [1]int{1}
}
