package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"os"
	"sort"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	graph := NewGraph()
	N := 10 // rozmiar grafu

	// dodawanie wierzchołków
	for i := 0; i < N; i++ {
		graph.AddNode(i)
	}

	// dodawanie krawędzi z prawdopodobieństwem p=4/N
	p := 4.0 / float64(N)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if i != j && rand.Float64() < p {
				graph.AddEdge(i, j)
			}
		}
	}

	// wyswietlanie
	fmt.Println("Graf:")
	graph.Print()

	//rozklad
	inDegrees, outDegrees := graph.DegreeDistribution()
	fmt.Println("Rozkład stopni wierzchołków wchodzących:", inDegrees)
	fmt.Println("Rozkład stopni wierzchołków wychodzących:", outDegrees)

	//najkrotsza sciezka
	shortestPaths := graph.ShortestPaths()
	fmt.Println("Najkrótsze ścieżki:")
	for _, row := range shortestPaths {
		fmt.Println(row)
	}

	inDegreesMap := make(map[int]int)
	for i, degree := range inDegrees {
		inDegreesMap[i] = degree
	}

	outDegreesMap := make(map[int]int)
	for i, degree := range outDegrees {
		outDegreesMap[i] = degree
	}

	//fmt.Println(len(inDegreesMap))
	// data := map[int]int{
	// 	1: 10,
	// 	2: 20,
	// 	3: 30,
	// }

	// rysowanie wykresow
	plotMap(inDegreesMap, "indegree_distribution.png", "Rozkład stopni wierzchołków", "Stopień", "Liczba wierzchołków")
	plotMap(outDegreesMap, "outdegree_distribution.png", "Rozkład stopni wierzchołków", "Stopień", "Liczba wierzchołków")
}

func plotMap(data map[int]int, name string, title string, x_name string, y_name string) {
	// Tworzenie nowego wykresu
	plt := plot.New()
	plt.Title.Text = title
	plt.X.Label.Text = x_name
	plt.Y.Label.Text = y_name

	var keys []int
	for k := range data {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	pts := make(plotter.XYs, len(data))
	for i := range keys {
		pts[i].X = float64(i)
		pts[i].Y = float64(data[i])
	}

	line, err := plotter.NewLine(pts)
	if err != nil {
		fmt.Println("Błąd tworzenia linii:", err)
		os.Exit(1)
	}
	line.LineStyle.Color = color.RGBA{B: 255, A: 255}
	plt.Add(line)

	if err := plt.Save(8*vg.Inch, 4*vg.Inch, name); err != nil {
		fmt.Println("Błąd zapisu wykresu:", err)
		os.Exit(1)
	}
	fmt.Printf("Wykres został zapisany do pliku '%s'.\n", name)
}
