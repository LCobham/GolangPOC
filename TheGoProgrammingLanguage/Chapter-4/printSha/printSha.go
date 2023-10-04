// Print the sha256 of an input. Optional flags if sha384 or sha512 are preferred.
package main

import (
	"crypto/sha512"
	"flag"
	"fmt"
	"log"
	"os"
)

var SHA384 = flag.Bool("SHA384", false, "Print the SHA384 instead of SHA256")
var SHA512 = flag.Bool("SHA512", false, "Print the SHA512 instead of SHA256")

func main() {
	flag.Parse()

	arg := os.Args[1]
	if arg == "" {
		log.Fatal("Usage: ./printSha [-384=true, -512=true] value")
	}

	if *SHA512 {
		fmt.Printf("%x\n", sha512.Sum512([]byte(arg)))
		return
	}
	if *SHA384 {
		fmt.Printf("%x\n", sha512.Sum384([]byte(arg)))
		return
	}
	fmt.Printf("%x\n", sha512.Sum512_256([]byte(arg)))
}
