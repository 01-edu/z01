package test

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"
	"unsafe"
)

const (
	IntSize = unsafe.Sizeof(0)*8 - 1
	IntMin  = -1 << IntSize
	IntMax  = 1<<IntSize - 1
)

func init() {
	nsSince1970 := time.Now().UnixNano()
	rand.Seed(nsSince1970)
}

func RandomInt() int {
	min := IntMin / 2
	max := IntMax / 2
	return rand.Intn(max-min) + min
}

// RandomRange return a random int in range [min,max]
func RandomRange(min, max int) int {
	return rand.Intn(max-min+1) + min
}

// RandomInts return a slice of n random int
func RandomInts(n int) []int {
	r := make([]int, n)
	for i := range r {
		r[i] = RandomInt()
	}
	return r
}

// RandomRanges return a slice of n random int ranges [min,max]
func RandomRanges(n, min, max int) []int {
	r := make([]int, n)
	for i := range r {
		r[i] = RandomRange(min, max)
	}
	return r
}

// ExecOut runs the command name with its args and returns its combined stdout
// and stderr as string.
// The returned error is nil if the command runs, has no problems
// copying stdin, stdout, and stderr, and exits with a zero exit
// status.
func ExecOut(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	bytes, err := cmd.CombinedOutput()
	return string(bytes), err
}

// FnOut returns as string what the function fn prints on stdout
func FnOut(fn func()) string {
	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		log.Fatalln("Cannot create pipe.")
	}
	os.Stdout = w
	fn()
	outC := make(chan string)
	var buf strings.Builder
	go func() {
		io.Copy(&buf, r)
		outC <- buf.String()
	}()
	os.Stdout = old
	w.Close()
	return <-outC
}

// Format is more user-friendly than Sprint
func Format(a ...interface{}) (s string) {
	ss := make([]string, len(a))
	for i, v := range a {
		switch v.(type) {
		case nil:
			// instead of "<nil>"
			ss[i] = "nil"
		case byte, rune: // uint8, int32
			// a single-quoted character literal safely escaped with Go syntax
			ss[i] = fmt.Sprintf("%q", v)
		default:
			// a Go-syntax representation of the value
			ss[i] = fmt.Sprintf("%#v", v)
		}
	}
	return strings.Join(ss, ", ")
}

var valueOf = reflect.ValueOf

func NameOfFunc(fn interface{}) string {
	if rf := runtime.FuncForPC(valueOf(fn).Pointer()); rf != nil {
		name := rf.Name()
		tokens := strings.Split(name, ".")
		return tokens[len(tokens)-1]
	}
	return "unknownFunc"
}

func Call(fn interface{}, args []interface{}) []interface{} {
	// Convert args from []interface{} to []reflect.Value
	vals := make([]reflect.Value, len(args))
	for i, v := range args {
		if v != nil {
			vals[i] = valueOf(v)
		} else {
			vals[i] = reflect.Zero(reflect.TypeOf((*interface{})(nil)).Elem())
		}
	}

	vals = valueOf(fn).Call(vals)

	// Convert the return values from []reflect.Value to []interface{}
	result := make([]interface{}, len(vals))
	for i, v := range vals {
		result[i] = v.Interface()
	}
	return result
}

// Expect is a test function that calls fn with the provided args
// If the result is different from those expected, an error is propagated through t
func Expect(t *testing.T, fn interface{}, args, result []interface{}) {
	actual := Call(fn, args)
	if !reflect.DeepEqual(actual, result) {
		t.Errorf(
			"%s(%s) == %s instead of %s\n",
			NameOfFunc(fn),
			Format(args...),
			Format(actual...),
			Format(result...),
		)
	}
}

func Challenge(t *testing.T, fn, valid interface{}, args ...interface{}) {
	result := Call(valid, args)
	Expect(t, fn, args, result)
}
