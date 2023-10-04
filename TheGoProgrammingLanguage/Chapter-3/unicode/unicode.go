// Defines functions to determine if a prefix is in a given string,
// if a string contains a substring, or if a string has a given suffix.
package unicode

func hasPrefix(str, prefix string) bool {
	return len(str) >= len(prefix) && str[:len(prefix)] == prefix
}

func hasSuffix(str, suffix string) bool {
	return len(str) >= len(suffix) && str[len(str)-len(suffix):] == suffix
}

func containsSubstr(str, substring string) bool {
	for i := 0; i < len(str); i++ {
		if hasPrefix(str[i:], substring) {
			return true
		}
	}
	return false
}
