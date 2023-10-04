// comma takes a number in a string format and returns the number with
// commas to separate thousands.
package main

import (
	"fmt"
	"strings"
)

func comma(s string) (string, error) {
	pointIdx := strings.Index(s, ".")
	lastIdx := strings.LastIndex(s, ".")

	if lastIdx >= 0 && pointIdx >= 0 && lastIdx != pointIdx {
		return "", fmt.Errorf("invalid string: multiple floating points")
	}

	if pointIdx == -1 {
		pointIdx = len(s)
	}
	var accum string
	i := pointIdx

	for ; i >= 3; i -= 3 {
		accum = s[i-3:i] + accum
		if i > 3 {
			accum = "," + accum
		} else if i < 3 {
			accum = s[:i] + accum
		}
	}
	if i > 0 {
		accum = s[:i] + accum
	}

	return accum + s[pointIdx:], nil
}

func main() {
	numbers := []string{"1", "12", "123", "1234", "12345",
		"123456", "1234567", "12345678", "123456789", "1.2",
		"1.23", "1.234", "1234.56", "1234567.89", "12345678.9"}

	for _, n := range numbers {
		conv, _ := comma(n)
		fmt.Printf("%s: %s\n", n, conv)
	}
}
