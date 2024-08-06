package main

import (
	"fmt"
	"strconv"
)

// Zadanie 1
func filterSlice[T any](slice []T, fun func(T) bool) []T {
	var filteredSlice []T
	for _, elem := range slice {
		if fun(elem) {
			filteredSlice = append(filteredSlice, elem)
		}
	}
	return filteredSlice
}

// Zadanie 2
func mapSlice[T any](slice []T, fun func(T) T) []T {
	mappedSlice := make([]T, len(slice))
	for i, elem := range slice {
		mappedSlice[i] = fun(elem)
	}
	return mappedSlice
}

// Zadanie 3
func reduceSlice[T any](slice []T, fun func(T, T) T) T {
	if len(slice) == 0 {
		panic("Empty slice < 1")
	}
	result := slice[0]
	for _, elem := range slice[1:] {
		result = fun(result, elem)
	}
	return result
}

// Zadanie 4
func splitMap[K comparable, V any](m map[K]V) ([]K, []V) {
	keys := make([]K, 0, len(m))
	values := make([]V, 0, len(m))
	for k, v := range m {
		keys = append(keys, k)
		values = append(values, v)
	}
	return keys, values
}

// Zadanie 5
type Pair[T comparable, V any] struct {
	Key   T
	Value V
}

func mapToPairs[T comparable, V any](m map[T]V) []Pair[T, V] {
	pairs := make([]Pair[T, V], 0, len(m))
	for k, v := range m {
		pairs = append(pairs, Pair[T, V]{k, v})
	}
	return pairs
}

// Zadanie 6
func slicesToMap[T comparable, V any](keys []T, values []V) map[T]V {
	size := len(keys)
	if len(values) < size {
		size = len(values)
	}
	result := make(map[T]V, size)
	for i := 0; i < size; i++ {
		result[keys[i]] = values[i]
	}
	return result
}

type customNumbers interface {
	int | float64
}

// Zadanie 7
func convertSlice[T int | float64](s []string) (r []T) {
	var t interface{}
	t = *new(T)
	for _, v := range s {
		switch t.(type) {

		case int:
			i, err := strconv.Atoi(v)
			if err == nil {
				r = append(r, T(i))
			}

		case float64:
			f, err := strconv.ParseFloat(v, 64)
			if err == nil {
				r = append(r, T(f))
			}

		}
	}
	return
}

func main() {
	// Testy
	testSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	testMap := map[int]string{1: "one", 2: "two", 3: "three", 4: "four", 5: "five"}
	stringSlice := []string{"1", "2", "3", "4", "5.5", "6.6"}

	// Zadanie 1
	fmt.Println("Zadanie 1:")
	sli11 := []int{1, 2, 3, 4, 5}
	sli12 := []float64{10.1, 9.9, 30.0, 42.0, 5.6}
	isEven := func(arg int) bool { return arg%2 == 0 }
	isGreaterThan10 := func(arg float64) bool { return arg > 10 }
	res11 := filterSlice(sli11, isEven)
	fmt.Println(res11)
	res12 := filterSlice(sli12, isGreaterThan10)
	fmt.Println(res12)

	// Zadanie 2
	fmt.Println("\nZadanie 2:")
	sli21 := []int{1, 2, 3, 4, 5}
	sli22 := []float64{10.1, 9.9, 30.0, 42.0, 5.6}
	add5 := func(arg int) int { return arg + 5 }
	divide5 := func(arg float64) float64 { return arg / 5 }
	res21 := mapSlice(sli21, add5)
	fmt.Println(res21)
	res22 := mapSlice(sli22, divide5)
	fmt.Println(res22)

	// Zadanie 3
	fmt.Println("\nZadanie 3:")
	sli32 := []string{"q", "w", "e", "r", "t", "y"}
	sumFunc := func(a, b int) int { return a + b }
	concat := func(arg1 string, arg2 string) string { return arg1 + arg2 }
	sum := reduceSlice(testSlice, sumFunc)
	fmt.Println(sum)
	res32 := reduceSlice(sli32, concat)
	fmt.Println(res32)

	// Zadanie 4
	fmt.Println("\nZadanie 4:")
	keys, values := splitMap(testMap)
	fmt.Println("Keys:", keys)
	fmt.Println("Values:", values)

	// Zadanie 5
	fmt.Println("\nZadanie 5:")
	pairs := mapToPairs(testMap)
	fmt.Println("Pairs:", pairs)

	// Zadanie 6
	fmt.Println("\nZadanie 6:")
	keys2 := []int{1, 2, 3, 4, 5}
	values2 := []string{"one", "two", "three", "four"}
	resultMap := slicesToMap(keys2, values2)
	fmt.Println("Map:", resultMap)

	keys3 := []string{"Michal", "Kasia", "Olga", "Janek", "Sebastian", "Ptaryk"}
	values3 := []int{21, 34, 13, 11, 37}
	m := slicesToMap(keys3, values3)
	fmt.Println("Map: ", m)

	// Zadanie 7
	fmt.Println("Zadanie 7:")
	floatSlice := convertSlice[float64](stringSlice)
	intSlice := convertSlice[int](stringSlice)
	fmt.Println("Float slice:", floatSlice)
	fmt.Println("Int slice:", intSlice)

	n := []string{"2.5", "2", "-4", "5", "-9.9", "a", "-", "0", "0."}
	fmt.Println(convertSlice[int](n))
	fmt.Println(convertSlice[float64](n))
}
