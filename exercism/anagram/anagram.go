package anagram

import (
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	//slice of string to store anagrams of string selected from candidates
	var output []string

	//sort subject alphabetically
	ss := SortString(subject)

	//sort each candidate string and compare with the sorted input subject string
	for _, v := range candidates {
		//words are not anagrams of themselves && sorted strings are the same
		if !strings.EqualFold(subject, v) && SortString(v) == ss {
			output = append(output, v)
		}
	}
	return output

}

// SortString takes a string and returns the same string sorted alphabetically
func SortString(s string) string {
	r := []rune(strings.ToLower(s))

	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})

	return string(r)
}
