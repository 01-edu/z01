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
)

func init() {
	nsSince1970 := time.Now().UnixNano()
	rand.Seed(nsSince1970)
}

// RandomInt return a random int between min and max
func RandomInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

// RandomInts return a slice of n random int between min and max
func RandomInts(n, min, max int) []int {
	r := make([]int, n)
	for i := range r {
		r[i] = RandomInt(min, max)
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

// Table is the type used to pass values to Fixed & Random functions
type Table [][2][]interface{}

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

// Fixed is a test function that, for each test in tests, calls fn with the provided arguments
// If the result is different from those expected, an error is propagated through t
func Fixed(t *testing.T, fn interface{}, tests Table) {
	for _, test := range tests {
		var in = test[0]
		var out = test[1]
		vals := make([]reflect.Value, len(in))
		for i, v := range in {
			if v != nil {
				vals[i] = valueOf(v)
			} else {
				vals[i] = reflect.Zero(reflect.TypeOf((*interface{})(nil)).Elem())
			}
		}
		vals = valueOf(fn).Call(vals)
		actual := make([]interface{}, len(vals))
		for i, v := range vals {
			actual[i] = v.Interface()
		}
		if !reflect.DeepEqual(actual, out) {
			t.Errorf(
				"%s(%s) == %s instead of %s\n",
				NameOfFunc(fn),
				Format(in...),
				Format(actual...),
				Format(out...),
			)
		}
	}
}
