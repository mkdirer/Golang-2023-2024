package main

import (
	"fmt"
)

func FibChannel() chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		prev := 0
		cur := 1
		ch <- prev
		for {
			ch <-cur
			temp := cur
			cur = prev + cur
			prev = temp
		}
	}()
	return ch
}

func OddSum(max int, odd chan int) {
	ch := FibChannel()
	sum := 0
	for val := range ch {
		if val > max {
			break
		}
		if val % 2 == 1 {
			sum += val
		}
	}
	odd <- sum
}

func EvenSum(max int, even chan int) {
	ch := FibChannel()
	sum := 0
	for val := range ch {
		if val > max {
			break
		}
		if val % 2 == 0 {
			sum += val
		}
	}
	even <- sum
}

func main() {
	odd := make(chan int)
	even := make(chan int)
	go OddSum(100, odd)
	go EvenSum(100, even)
	a := <-even
	b := <-odd
	fmt.Printf("Even %d Odd %d\n", a, b)
}