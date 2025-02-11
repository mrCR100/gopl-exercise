package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

// Exercise 3.10: Write a non-recursive version of comma, using bytes.Buffer instead of string concatenation.
func comma(s string) string {
	var buf bytes.Buffer
	l := len(s)
	for i := 0; i < l; i++ {
		if i > 0 && (l-i)%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
	}
	return buf.String()
}
