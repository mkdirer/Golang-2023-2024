package main

import (
	"fmt"
	"strconv"
	"time"
)

func FindNFib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return FindNFib(n-1) + FindNFib(n-2)
	}
}

func PrintSeq(step time.Duration) {
	s := "-\\|/"
	for {
		for _, elem := range s {
			time.Sleep(step)
			fmt.Printf("\r" + string(elem))
		}
	}
}

func main() {
	go PrintSeq(time.Second)
	fmt.Printf("\r" + strconv.Itoa(FindNFib(45)) + "\n")
}
