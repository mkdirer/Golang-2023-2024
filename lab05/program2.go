package main

// example use: go run program2.go --min=0 --max=100 --num=1000 --show-sorted-hist=false

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"
)

// HistogramEntry reprezentuje pojedynczy wpis w histogramie.
type HistogramEntry struct {
	Value int
	Count int
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Not enough argumets. Use --help for help")
		return
	}
	var (
		minRange       int
		maxRange       int
		numSamples     int
		showSortedHist bool
	)

	flag.IntVar(&minRange, "min", 0, "Minimalna wartość zakresu losowania")
	flag.IntVar(&maxRange, "max", 100, "Maksymalna wartość zakresu losowania")
	flag.IntVar(&numSamples, "num", 1000, "Liczba losowych próbek do wygenerowania")
	flag.BoolVar(&showSortedHist, "show-sorted-hist", false, "Wyświetl histogram posortowany po kluczach")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())
	randomValues := make([]int, numSamples)
	for i := 0; i < numSamples; i++ {
		randomValues[i] = rand.Intn(maxRange-minRange+1) + minRange
	}

	histogram := make(map[int]int)
	for _, value := range randomValues {
		histogram[value]++
	}

	// Tworzenie histogramu w postaci posortowanej (jeśli wybrana opcja).
	var sortedHistogram []HistogramEntry
	if showSortedHist {
		sortedHistogram = make([]HistogramEntry, 0, len(histogram))
		for value, count := range histogram {
			sortedHistogram = append(sortedHistogram, HistogramEntry{Value: value, Count: count})
		}
		sort.Slice(sortedHistogram, func(i, j int) bool {
			return sortedHistogram[i].Value < sortedHistogram[j].Value
		})
	}

	// Obliczanie średniej.
	var sum int
	for value, count := range histogram {
		sum += value * count
	}
	average := float64(sum) / float64(numSamples)

	// Obliczanie odchylenia standardowego.
	var sumOfSquares float64
	for value, count := range histogram {
		sumOfSquares += float64(count) * (float64(value) - average) * (float64(value) - average)
	}
	standardDeviation := sumOfSquares / float64(numSamples)

	// Wyświetlanie wyników.
	fmt.Println("Histogram częstości losowania wartości:")
	if showSortedHist {
		for _, entry := range sortedHistogram {
			fmt.Printf("%3d: %d\n", entry.Value, entry.Count)
		}
	} else {
		for value, count := range histogram {
			fmt.Printf("%3d: %d\n", value, count)
		}
	}
	fmt.Printf("Średnia: %.2f\n", average)
	fmt.Printf("Odchylenie standardowe: %.2f\n", standardDeviation)
}
