package services

import (
	"sort"
	"strings"
)

// Anagram service
type Anagram struct {
}

// NewAnagram service constructor
func NewAnagram() *Anagram {
	return &Anagram{}
}

// Check two strings
func (service *Anagram) Check(a, b string) bool {
	a = strings.ReplaceAll(a, " ", "")
	b = strings.ReplaceAll(b, " ", "")
	return service.sortRunes(a) == service.sortRunes(b)
}

func (service *Anagram) sortRunes(str string) string {
	slice := strings.Split(str, "")
	sort.Strings(slice)
	return strings.Join(slice, "")
}
