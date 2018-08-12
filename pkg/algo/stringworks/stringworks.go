package stringworks

import (
	"fmt"
	"strings"
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

func longShort(f, g string) (string, string) {
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
	return shorterStr, longerStr
}

// IsPermutation checks if two strings are the same permutation
func IsPermutation(f, g string) bool {
	if f == g {
		return true
	}
	shorterStr, longerStr := longShort(f, g)
	dict := buildDict(shorterStr)
	for _, r := range longerStr {
		char := string(r)
		count := dict[char]
		if count == 0 {
			return false
		}
		dict[char] = count - 1
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

// Urlify converts spaces in %20 from urls
func Urlify(s string) string {
	url := make([]string, len(s))
	for _, r := range s {
		word := string(r)
		if url[len(url)-1] == "%20" && unicode.IsSpace(r) {
			continue
		}
		if unicode.IsSpace(r) {
			url = append(url, "%20")
		} else {
			url = append(url, word)
		}
	}
	return strings.Join(url, "")
}

func findSegmentIndex(s string, startIndex int) int {
	endIndex := startIndex + 1
	if endIndex >= len(s) {
		return endIndex
	}
	startChar := string(s[startIndex])

	for endIndex < len(s) && string(s[endIndex]) == startChar {
		endIndex++
	}
	return endIndex
}

// Zip compresses given string with repeated characters
func Zip(s string) string {
	i := 0
	var toBeAppended string
	compressedStr := make([]string, len(s))

	for i < len(s) {
		currChar := string(s[i])
		nextI := i + 1
		if i+1 >= len(s) {
			nextI = i
		}
		nextChar := string(s[nextI])
		if currChar == nextChar && nextI != i {
			endIndex := findSegmentIndex(s, i)
			toBeAppended = currChar + fmt.Sprintf("%v", endIndex-i)
			i = endIndex
		} else {
			i++
			toBeAppended = currChar
		}
		compressedStr = append(compressedStr, toBeAppended)
	}
	return strings.Join(compressedStr, "")
}
