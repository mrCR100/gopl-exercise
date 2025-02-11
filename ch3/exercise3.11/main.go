package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

// assume input is (+/-)d+\.d+
func comma(s string) string {
	var buf bytes.Buffer
	l := len(s)
	for i := 0; i < l; i++ {
		point := strings.LastIndex(s, ".")
		sign := -1
		if s[0] == '+' || s[0] == '-' {
			sign = 0
		}
		if i > sign+1 && i < point && (point-i)%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
	}
	return buf.String()
}
