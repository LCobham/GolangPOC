// shaCompareBits compares the number of bits that differ in two sha256 digests.
package main

import (
	"crypto/sha256"
	"fmt"
)

func shaCompareBits(digest1, digest2 *[32]uint8) int {
	var count int

	for idx, bit1 := range digest1 {
		bit2 := (*digest2)[idx]

		for i := 0; i < 8; i++ {
			count += int((bit1 >> i & 1) ^ (bit2 >> i & 1))
		}
	}
	return count
}

func main() {
	digest1 := sha256.Sum256([]byte("x"))
	digest2 := sha256.Sum256([]byte("x"))

	fmt.Println(shaCompareBits(&digest1, &digest2))

	digest2 = sha256.Sum256([]byte("X"))
	fmt.Println(shaCompareBits(&digest1, &digest2))
}
