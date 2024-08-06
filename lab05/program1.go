package main

//example use: go run program1.go --input=plik_wejsciowy.txt --remove-duplicates=true
//example use: go run program1.go --input=plik_wejsciowy.txt --filter-prefix=t
//example use: go run program1.go --input=plik_wejsciowy.txt --remove-duplicates=true --filter-prefix=t

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func readData(reader io.Reader, removeDuplicates bool, filterPrefix string) ([]string, error) {
	scanner := bufio.NewScanner(reader)
	var result []string
	seenLines := make(map[string]bool)
	seenWords := make(map[string]bool)

	for scanner.Scan() {
		line := scanner.Text()

		if removeDuplicates {
			if seenLines[line] {
				continue
			}
			seenLines[line] = true
		}

		words := strings.Fields(line)
		for _, word := range words {
			if removeDuplicates {
				if seenWords[word] {
					continue
				}
				seenWords[word] = true
			}

			if filterPrefix != "" {
				if strings.HasPrefix(word, filterPrefix) {
					result = append(result, word)
				}
			} else {
				result = append(result, word)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func main() {
	var (
		inputFile    string
		removeDup    bool
		filterPrefix string
	)

	flag.StringVar(&inputFile, "input", "", "Nazwa pliku wejściowego. Puste oznacza konsolę.")
	flag.BoolVar(&removeDup, "remove-duplicates", false, "Usuwanie duplikatów.")
	flag.StringVar(&filterPrefix, "filter-prefix", "", "Filtrowanie wyrazów zaczynających się od określonego znaku.")
	flag.Parse()

	flag.Visit(func(f *flag.Flag) {
		log.Printf("Użyto flagi: %s=%v", f.Name, f.Value)
	})

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

	fmt.Println("Przetworzone dane:")
	for _, line := range lines {
		fmt.Println(line)
	}
}
