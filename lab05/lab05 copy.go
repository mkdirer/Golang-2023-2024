package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"sort"
	"strings"
)

// readData funkcja czyta dane z wejścia (io.Reader) w zależności od ustawień flagi.
// Usuwa duplikaty lub filtruje wyrazy zaczynające się od określonego znaku.
func readData(reader io.Reader, removeDuplicates bool, filterPrefix string) ([]string, error) {
	scanner := bufio.NewScanner(reader)
	result := make([]string, 0)

	seen := make(map[string]bool)
	for scanner.Scan() {
		line := scanner.Text()

		if removeDuplicates {
			if seen[line] {
				continue
			}
			seen[line] = true
		}

		if filterPrefix != "" {
			if strings.HasPrefix(line, filterPrefix) {
				result = append(result, line)
			}
		} else {
			result = append(result, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

// generateRandomNumbers generuje losowe liczby całkowite z podanego zakresu.
func generateRandomNumbers(min, max int, count int) []int {
	numbers := make([]int, count)
	for i := 0; i < count; i++ {
		numbers[i] = rand.Intn(max-min+1) + min
	}
	return numbers
}

// calculateHistogram tworzy histogram częstości losowania określonych wartości.
func calculateHistogram(numbers []int) map[int]int {
	histogram := make(map[int]int)
	for _, num := range numbers {
		histogram[num]++
	}
	return histogram
}

func main() {
	// Definicja flag
	var (
		inputFile     string
		removeDup     bool
		filterPrefix  string
		minRange      int
		maxRange      int
		numCount      int
		sortHistogram bool
	)

	// Obsługa flag
	flag.StringVar(&inputFile, "input", "", "Nazwa pliku wejściowego. Jeśli pusta, program odczytuje z konsoli.")
	flag.BoolVar(&removeDup, "remove-duplicates", false, "Usuwanie duplikatów.")
	flag.StringVar(&filterPrefix, "filter-prefix", "", "Filtrowanie wyrazów zaczynających się od określonego znaku.")
	flag.IntVar(&minRange, "min", 0, "Minimalna wartość dla losowych liczb.")
	flag.IntVar(&maxRange, "max", 100, "Maksymalna wartość dla losowych liczb.")
	flag.IntVar(&numCount, "count", 100, "Liczba losowych liczb do wygenerowania.")
	flag.BoolVar(&sortHistogram, "sort-histogram", false, "Sortowanie histogramu.")

	flag.Parse()

	// Logowanie opcji użycia
	log.Printf("inputFile: %s\n", inputFile)
	log.Printf("removeDup: %t\n", removeDup)
	log.Printf("filterPrefix: %s\n", filterPrefix)
	log.Printf("minRange: %d\n", minRange)
	log.Printf("maxRange: %d\n", maxRange)
	log.Printf("numCount: %d\n", numCount)
	log.Printf("sortHistogram: %t\n", sortHistogram)

	// Obsługa wejścia
	var reader io.Reader
	if inputFile != "" {
		file, err := os.Open(inputFile)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		reader = file
	} else {
		reader = os.Stdin
	}

	lines, err := readData(reader, removeDup, filterPrefix)
	if err != nil {
		log.Fatal(err)
	}

	// Wypisanie przetworzonych danych
	fmt.Println("Przetworzone dane:")
	for _, line := range lines {
		fmt.Println(line)
	}

	// Generowanie losowych liczb
	randomNumbers := generateRandomNumbers(minRange, maxRange, numCount)

	// Tworzenie histogramu
	histogram := calculateHistogram(randomNumbers)

	// Wypisanie histogramu
	fmt.Println("\nHistogram:")
	keys := make([]int, 0, len(histogram))
	for k := range histogram {
		keys = append(keys, k)
	}
	if sortHistogram {
		sort.Ints(keys)
	}
	for _, k := range keys {
		fmt.Printf("%d: %d\n", k, histogram[k])
	}
}
