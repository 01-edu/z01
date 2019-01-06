package test_test

import (
	"fmt"
	"testing"

	"github.com/01-edu/z01/test"
)

func TestRandomInt(t *testing.T) {
	min := 55
	max := 58
	nb := test.RandomInt(min, max)
	if nb < min || nb > max {
		t.Errorf("RandomInt(%d, %d) not in the [%d, %d] range", min, max, min, max)
	}
}

func ExampleExecOut() {
	output, err := test.ExecOut("echo", "test")
	if err == nil {
		fmt.Print(output)
	} else {
		fmt.Print(err)
	}
	// Output: test
}

func ExampleFnOut() {
	fn := func() {
		fmt.Println("first line")
		fmt.Println("second line")
	}
	output := test.FnOut(fn)
	fmt.Print(output)
	// Output:
	// first line
	// second line
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

func divMod(a, b int) (div, mod int) {
	return a / b, a % b
}

func TestFixed(t *testing.T) {
	tests := test.Table{
		//,--------------- first argument (a)
		//|  ,------------ second argument (b)
		//|  |    ,------- first return value (div)
		//|  |    |  ,---- second return value (mod)
		//|  |    |  |
		//V  V    V  V
		{{0, 1}, {0, 0}},
		{{1, 1}, {1, 0}},
		{{1, 2}, {0, 1}},
		{{3, 2}, {1, 1}},
		{{9, 2}, {4, 1}},
	}
	test.Fixed(t, divMod, tests)
}
