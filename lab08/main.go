package main

import (
	"fmt"
)

func main() {
	sli, err := split("aaa/bbb/ccc Hello   World\nGopher\tis\tawesome")
	fmt.Println(sli, err)

	fmt.Println(split("aaa bbb ccc"))
	fmt.Println(split("aaa bbb ccc", "", " "))

	fmt.Println("splitAndSort")

	fmt.Println(splitAndSort("bbb ccc aaa  Hello   World\nGopher\tis\tawesome"))

	fmt.Println(getNumberOfNonEmptyLines("aaa\n\nbbb\nccc\n\nd"))

	fmt.Println(getNumOfWords("bbb ccc aaa  Hello   World\nGopher\tis\tawesome"))

	fmt.Println(getNumOfNonWhiteChars("bbb ccc aaa  Hello   World\nGopher\tis\tawesome"))

	fmt.Println(getFrequencyForWords("aaa bbb aaa ccc bbb aaa  Hello Hello  World\nGopher\tWorld\nis\tawesome"))

	fmt.Println(getSortedPalindroms("bccb cda aaa aba dcdc Hello Hello o  World\nHelleH\tGopher\tWorld\nis\tawesome"))

}
