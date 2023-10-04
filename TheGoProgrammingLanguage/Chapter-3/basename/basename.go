// Basename removes all file prefixes that look like
// file system path, and removes suffixes that look like
// file extensions.
package basename

import "strings"

func basename(s string) string {
	slash := strings.LastIndex(s, "/") // Returns -1 if no "/" char found
	s = s[slash+1:]                    // if no "/" char found, this becomes s = s[0:] => s = s

	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}

	return s
}
