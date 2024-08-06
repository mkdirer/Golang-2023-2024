package main

import (
	"errors"
	"sort"
	"strings"
	"unicode"
)

func split(str string, sep ...string) ([]string, error) {
	defaultSep := " \t\n"
	s := defaultSep
	if len(sep) == 1 {
		s = sep[0]
	} else if len(sep) > 1 {
		return nil, errors.New("Too many arguments")
	}
	res := strings.FieldsFunc(str, func(r rune) bool {
		return strings.ContainsRune(s, r)
	})
	return res, nil
}

func splitAndSort(str string, sep ...string) ([]string, error) {
	defaultSep := " \t\n"
	s := defaultSep
	if len(sep) == 1 {
		s = sep[0]
	} else if len(sep) > 1 {
		return nil, errors.New("Too many arguments")
	}
	res := strings.FieldsFunc(str, func(r rune) bool {
		return strings.ContainsRune(s, r)
	})
	sort.Strings(res)
	return res, nil
}

func getNumberOfNonEmptyLines(text string) int {
	lines := strings.Split(text, "\n")
	numOfNonEmptyLines := 0
	for _, line := range lines {
		if line != "" {
			numOfNonEmptyLines++
		}
	}
	return numOfNonEmptyLines
}

func getNumOfWords(text string) int {
	words := strings.FieldsFunc(text, func(r rune) bool {
		return unicode.IsSpace(r)
	})
	return len(words)
}

func getNumOfNonWhiteChars(text string) int {
	return len(text) - strings.Count(text, " ") - strings.Count(text, "\t") - strings.Count(text, "\n")
}

func getFrequencyForWords(text string) map[string]int {
	mapWordsFrequency := make(map[string]int)
	words := strings.FieldsFunc(text, func(r rune) bool {
		return unicode.IsSpace(r)
	})
	for _, word := range words {
		mapWordsFrequency[word]++
	}
	return mapWordsFrequency
}

func getSortedPalindroms(text string) []string {
	words := strings.FieldsFunc(text, func(r rune) bool {
		return unicode.IsSpace(r)
	})
	palindroms := make([]string, 0)
	for _, word := range words {
		palindrom := true
		for i := 0; i < len(word)/2; i++ {
			if word[i] != word[len(word)-1-i] {
				palindrom = false
				break
			}
		}
		if palindrom {
			palindroms = append(palindroms, word)
		}
	}
	sort.Strings(palindroms)
	return palindroms
}
