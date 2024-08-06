package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"testing"
	"time"
)

func Test_split(t *testing.T) {
	res, err := split("aaa/bbb/ccc", "/")
	expected := []string{"aaa", "bbb", "ccc"}
	if len(res) != len(expected) {
		t.Error("Error in function getSortedPalindroms, res = ", res,
			"expected = ", expected)
	}
	for i, val := range res {
		if val != expected[i] {
			t.Error("Error in function split, res = ", res,
				"expected = ", expected)
		}
	}
	if err != nil {
		t.Error("Error in function split - error = ", err)
	}

	res, err = split("aaa bbb ccc")
	if len(res) != len(expected) {
		t.Error("Error in function getSortedPalindroms, res = ", res,
			"expected = ", expected)
	}
	for i, val := range res {
		if val != expected[i] {
			t.Error("Error in function split, res = ", res,
				"expected = ", expected)
		}
	}
	if err != nil {
		t.Error("Error in function split - error = ", err)
	}

	res, err = split("aaa bbb ccc", "", " ")

	if err == nil {
		t.Error("Error in function split - no error")
	}
}

func Test_split_error(t *testing.T) {
	s := "aaa/bbb/ccc "
	sep1 := "/"
	sep2 := " "
	_, err := split(s, sep1, sep2)
	if err == nil {
		t.Error("Incorrect result: expected error")
	}
}

func Test_splitAndSort(t *testing.T) {
	res, err := splitAndSort("bbb ccc aaa")
	expected := []string{"aaa", "bbb", "ccc"}
	if len(res) != len(expected) {
		t.Error("Error in function getSortedPalindroms, res = ", res,
			"expected = ", expected)
	}
	for i, val := range res {
		if val != expected[i] {
			t.Error("Error in function splitAndSort, res = ", res,
				"expected = ", expected)
		}
	}
	if err != nil {
		t.Error("Error in function splitAndSort - error = ", err)
	}
}

func Test_splitAndSort_error(t *testing.T) {
	s := "aaa/bbb/ccc "
	sep1 := "/"
	sep2 := " "
	_, err := split(s, sep1, sep2)
	if err == nil {
		t.Error("Incorrect result: expected error")
	}
}

func Test_getNumberOfNonEmptyLines(t *testing.T) {
	res := getNumberOfNonEmptyLines("aaa\n\nbbb\nccc\n\nd")
	expected := 4
	if res != expected {
		t.Error("Error in function getNumberOfNonEmptyLines, res = ", res,
			"expected = ", expected)
	}
}

func Test_getNumOfWords(t *testing.T) {
	res := getNumOfWords("bbb ccc aaa  Hello   World\nGopher\tis\tawesome")
	expected := 8
	if res != expected {
		t.Error("Error in function getNumOfWords, res = ", res,
			"expected = ", expected)
	}
}

func Test_getNumOfNonWhiteChars(t *testing.T) {
	res := getNumOfNonWhiteChars("bbb ccc aaa  Hello   World\nGopher\tis\tawesome")
	expected := 34
	if res != expected {
		t.Error("Error in function getNumOfNonWhiteChars, res = ", res,
			"expected = ", expected)
	}
}

func Test_getFrequencyForWords(t *testing.T) {
	res := getFrequencyForWords("aaa bbb aaa ccc bbb aaa")
	expected := map[string]int{"aaa": 3, "bbb": 2, "ccc": 1}
	if len(res) != len(expected) {
		t.Error("Error in function getSortedPalindroms, res = ", res,
			"expected = ", expected)
	}
	for key, val := range res {
		if val != expected[key] {
			t.Error("Error in function getFrequencyForWords, res = ", res,
				"expected = ", expected)
		}
	}
}

func Test_getSortedPalindroms(t *testing.T) {
	res := getSortedPalindroms("bccb cda aaa aba dcdc Hello Hello o  World\nHelleH\tGopher\tWorld\nis\tawesome")
	expected := []string{"HelleH", "aaa", "aba", "bccb", "o"}
	if len(res) != len(expected) {
		t.Error("Error in function getSortedPalindroms, res = ", res,
			"expected = ", expected)
	}
	for i, val := range res {
		if val != expected[i] {
			t.Error("Error in function getSortedPalindroms, res = ", res,
				"expected = ", expected)
		}
	}
}

func Test_getSortedPalindroms_random(t *testing.T) {
	const letters = "abcdef"
	l := 3
	n := 100
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	s := ""

	i := 0
	for i < n {
		j := 0
		for j < l {
			index := r.Intn(len(letters))
			s += string(letters[index])
			j++
		}
		s += " "
		i++
	}
	res := getSortedPalindroms(s)
	fmt.Println(res)

	for _, elem := range res {
		if elem[0] != elem[2] {
			t.Error("Word ", elem, " is not a palindrome")
		}
	}
}

var blackhole []string

func BenchmarkFuns(b *testing.B) {
	for _, filename := range []string{"Latin-Lipsum_5.txt", "Latin-Lipsum_13.txt", "Latin-Lipsum_20.txt"} {
		content, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}
		s := string(content)
		b.Run(fmt.Sprintf(filename), func(b *testing.B) {
			res, _ := splitAndSort(s)
			blackhole = res
		})
	}
}
