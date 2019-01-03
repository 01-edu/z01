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
