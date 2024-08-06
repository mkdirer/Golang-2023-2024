package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	n := 100
	var short []string
	var medium []string
	var long []string
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		l := r.Intn(11) + 3
		var slice []byte
		for l > 0 {
			val := uint8(r.Intn(26) + 65)
			for len(slice) > 0 && slice[len(slice)-1] == val {
				val = uint8(r.Intn(26) + 65)
			}
			slice = append(slice, val)
			l--
		}
		switch {
		case len(slice) < 6:
			short = append(short, string(slice))
		case len(slice) < 10:
			medium = append(medium, string(slice))
		case len(slice) <= 13:
			long = append(long, string(slice))
		}
	}

	fmt.Println(short)
	fmt.Println(medium)
	fmt.Println(long)

	// zad 1
	fmt.Printf("==========\nZAD %d\n==========\n", 1)
	m := map[int]([]string){}
	addStringsToMap := func(slice []string, m map[int][]string) {
		for _, s := range slice {
			(m)[len(s)] = append((m)[len(s)], s)
		}
	}
	addStringsToMap(short, m)
	addStringsToMap(medium, m)
	addStringsToMap(long, m)
	for k, v := range m {
		fmt.Printf("%d: ", k)
		fmt.Println(v)
	}
	fmt.Println("==========")
	for k, v := range m {
		fmt.Printf("%d: ", k)
		fmt.Println(v)
	}

	// zad 2
	fmt.Printf("==========\nZAD %d\n==========\n", 2)
	var irregularSlice [][]int
	rows := 6 // Liczba wierszy

	// Generowanie wartości losowych dla każdego wiersza
	for i := 0; i < rows; i++ {
		cols := rand.Intn(6) + 4 // Losowa liczba kolumn dla każdego wiersza
		var row []int
		for j := 0; j < cols; j++ {
			val := rand.Intn(6) // Losowa wartość od 0 do 5
			row = append(row, val)
		}
		// row = append(row, 0) //for check common idiom part
		irregularSlice = append(irregularSlice, row)
	}

	fmt.Println("Dwuwymiarowy nieregularny slice:")
	for _, row := range irregularSlice {
		fmt.Println(row)
	}

	// Tworzenie zbioru zawierającego wartości występujące we wszystkich wierszach
	set := make(map[int]bool)
	for _, row := range irregularSlice {
		for _, val := range row {
			if _, ok := set[val]; !ok {
				set[val] = true
			}
		}
	}

	fmt.Println(set)

	//instrukcja goto
allRows:
	for val := range set {
		for _, row := range irregularSlice {
			found := false
			for _, v := range row {
				if val == v {
					found = true
					break
				}
			}
			if !found {
				delete(set, val)
				goto allRows
			}
		}
	}

	fmt.Println("Zbiór zawierający wartości występujące we wszystkich wierszach:")
	for val, ok := range set {
		fmt.Printf("%d: %t\n", val, ok)
	}

	//zad 3
	fmt.Printf("==========\nZAD %d\n==========\n", 3)
	addSorted := func(slice []int, value int) ([]int, error) {
		if len(slice) == cap(slice) {
			return slice, fmt.Errorf("Range error, length = %d capability = %d", len(slice), cap(slice))
		}
		if len(slice) == 0 {
			slice = append(slice, value)
		} else {
			for idx := 0; idx < len(slice); idx++ {
				if slice[idx] < value {
					continue
				}
				slice = append(slice[:idx+1], slice[idx:]...)
				slice[idx] = value
				return slice, nil
			}
			slice = append(slice, value)
		}
		return slice, nil
	}
	var tab [8]int
	sortedSlice := tab[:0]

	sortedSlice, e := addSorted(sortedSlice, 5)
	if e != nil { // no error
		fmt.Println(e)
	}
	sortedSlice, _ = addSorted(sortedSlice, 8)
	sortedSlice, _ = addSorted(sortedSlice, 5)
	sortedSlice, _ = addSorted(sortedSlice, 4)
	sortedSlice, _ = addSorted(sortedSlice, 7)
	sortedSlice, _ = addSorted(sortedSlice, 5)
	sortedSlice, _ = addSorted(sortedSlice, 7)
	sortedSlice, _ = addSorted(sortedSlice, 3)
	sortedSlice, e = addSorted(sortedSlice, 7)
	if e != nil { // error
		fmt.Println(e)
	}
	fmt.Println(sortedSlice)

	//zad 4
	fmt.Printf("==========\nZAD %d\n==========\n", 4)
	withPrefix := func(prefix string) func(string) string {
		return func(s string) string {
			return prefix + s
		}
	}

	withSuffix := func(suffix string) func(string) string {
		return func(s string) string {
			return s + suffix
		}
	}

	aggregateDecorators := func(decorators ...func(string) string) func(string) string {
		return func(s string) string {
			result := s
			for _, decorator := range decorators {
				result = decorator(result)
			}
			return result
		}
	}

	preDecorator := withPrefix("Pre: ")
	fmt.Println(preDecorator("Test"))

	sufDecorator := withSuffix(" Suf")
	fmt.Println(sufDecorator("Test"))

	aggregatedDecorator := aggregateDecorators(withPrefix("Pre: "), withSuffix(" Suf"))
	fmt.Println(aggregatedDecorator("Test"))

}
