package main

import (
	"fmt"
	"sync"
)

var p []int

func addMove(player int, p []int, m *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	var row, col int
	for {
		fmt.Printf("Player %d, enter row and column: ", player)
		fmt.Printf("\n")
		fmt.Scanf("%d %d", &row, &col)
		idx := row*3 + col
		if row < 0 || row > 2 || col < 0 || col > 2 || p[idx] != 0 {
			fmt.Println("Invalid move, try again.")
			continue
		}
		m.Lock()
		p[idx] = player
		m.Unlock()
		break
	}
}

func checkResult(p []int) bool {
	winPatterns := [8][3]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, // wiersze
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8}, // kolumny
		{0, 4, 8}, {2, 4, 6}, // przekątne
	}

	for _, pattern := range winPatterns {
		if p[pattern[0]] != 0 && p[pattern[0]] == p[pattern[1]] && p[pattern[1]] == p[pattern[2]] {
			return true
		}
	}
	return false
}

func printBoard(p []int) {
	for i := 0; i < 9; i += 3 {
		fmt.Printf("%d %d %d\n", p[i], p[i+1], p[i+2])
	}
}

func main() {
	p = make([]int, 9)
	var wg sync.WaitGroup
	var m sync.Mutex

	for i := 0; i < 9; i++ {
		printBoard(p)
		wg.Add(1)
		//move from 0 to 2
		if i%2 == 0 {
			go addMove(1, p, &m, &wg) // Player 1 (X)
		} else {
			go addMove(2, p, &m, &wg) // Player 2 (O)
		}
		wg.Wait()
		if checkResult(p) {
			printBoard(p)
			fmt.Printf("Player %d wins!\n", 1+i%2)
			return
		}
	}
	printBoard(p)
	fmt.Println("It's a draw!")
}
