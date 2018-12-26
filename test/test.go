package test

import (
	"io"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strings"
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
