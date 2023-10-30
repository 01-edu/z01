# z01

##  PrintRune Function

This is a Go package containing a function called `PrintRune` that allows you to print a single Unicode rune (Unicode code point) to the standard output. The function handles encoding and writing of the rune, and it returns an error in case of encoding or writing failures.

## Usage

You can use the `PrintRune` function in your Go code to display Unicode runes, which is particularly useful when working with non-ASCII characters and emojis.

```go
package main

import (
	"github.com/01-edu/z01"
)

func main() {
    // Print a Unicode rune
    z01.PrintRune('🚀')
    
}
