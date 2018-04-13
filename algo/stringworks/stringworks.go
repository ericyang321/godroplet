package stringworks

import (
	"unicode"
)

func buildDict(s string) map[string]int {
	dict := make(map[string]int)
	for _, r := range s {
		val, keyPresent := dict[string(r)]
		if keyPresent {
			dict[string(r)] = val + 1
		} else {
			dict[string(r)] = 1
		}
	}
	return dict
}

// IsPermutation checks if two strings are the same permutation
func IsPermutation(f, g string) bool {
	if f == g {
		return true
	}
	var shorterStr string
	var longerStr string

	switch {
	case len(f) < len(g):
		shorterStr = f
		longerStr = g
	default:
		shorterStr = g
		longerStr = f
	}
	dict := buildDict(shorterStr)
	for _, r := range longerStr {
		char := string(r)
		count := dict[char]
		if count == 0 {
			return false
		} else {
			dict[char] = count - 1
		}
	}
	return true
}

// IsUnique checks has all unique alphabets
func IsUnique(s string) bool {
	dict := make(map[string]bool)
	for _, char := range s {
		word := string(char)
		isAlphabet := !unicode.IsPunct(char) && !unicode.IsSpace(char)

		_, keyPresent := dict[word]
		if keyPresent {
			return false
		} else if isAlphabet {
			dict[word] = true
		}
	}
	return true
}
