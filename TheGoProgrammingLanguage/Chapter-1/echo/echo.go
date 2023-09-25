// Inspired by the Unix "echo" command
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// Version 1: prints all command line arguments
// using two local variables and a for-range loop.
func v1() {

	var s, sep string

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Printf("%s\n\n", s)
}

// Version 2: prints all command line arguments using
// strings.Join and no local variables.
func v2() {
	fmt.Printf("%s\n\n", strings.Join(os.Args[1:], " "))
}

// Exercise 1.1: modify echo to also print the name
// of the executable.
func exercise1_1() {
	fmt.Println(strings.Join(os.Args[:], " "))
}

// Exercise 1.2: Modify echo to print the index and
// value of each of the arguments, one per line.
func exercise1_2() {
	for i, arg := range os.Args[:] {
		fmt.Printf("%d: %s\n", i, arg)
	}
}

// Exercise 1.3: compare the time taken for each of
// these solutions using the time package.
func exercise1_3() {
	fmt.Println("v1: ", timeOfF(v1))
	fmt.Println("v2: ", timeOfF(v2))
	fmt.Println("ex1.1: ", timeOfF(exercise1_1))
	fmt.Println("ex1.2: ", timeOfF(exercise1_2))
}

func timeOfF(f func()) int64 {
	start := time.Now()
	f()
	return time.Since(start).Microseconds()
}

func main() {
	exercise1_3()
}
