package test

import (
	"math/rand"
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
