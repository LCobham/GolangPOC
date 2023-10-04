// IsAnagram returns true if two strings are anagrams of eachother.
package anagram

func IsAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	letters := make(map[rune]int)
	for _, char := range s1 {
		letters[char]++
	}
	for _, char := range s2 {
		letters[char]--
		if letters[char] < 0 {
			return false
		}
	}

	return true
}
