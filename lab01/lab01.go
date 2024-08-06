package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	fmt.Printf("%T | %T\n", letters, letters[0])
	rand.Seed(time.Now().UnixNano())

	var short, medium, long []string

	i := 0
	for i < 100 {
		length := rand.Intn(11) + 3 // Losowa długość ciągu (3-13)
		randomString := generateRandomString(letters, length)

		if !hasConsecutiveDuplicates(randomString) {
			switch length {
			case 3, 4, 5:
				short = append(short, randomString)
			case 6, 7, 8, 9:
				medium = append(medium, randomString)
			case 10, 11, 12, 13:
				if cap(long) != cap(append(long, randomString)) {
					fmt.Printf("Pojemność slice'a 'long' zmieniła się na %d, długość: %d\n", cap(long), len(long))
				}
				long = append(long, randomString)
			}

			// switch {
			// case length < 6:
			// 	short = append(short, randomString)
			// case length < 10:
			// 	medium = append(medium, randomString)
			// case length < 13:
			// 	previousCapacity := cap(long)
			// 	long = append(long, randomString)
			// 	currentCapacity := cap(long)

			// 	if previousCapacity != currentCapacity {
			// 		fmt.Printf("Pojemność slice'a 'long' zmieniła się z %d na %d, długość: %d\n", previousCapacity, currentCapacity, len(long))
			// 	}
			// }

			i++
		}
	}

	fmt.Println("Krótkie ciągi:", short)
	fmt.Println("Średnie ciągi:", medium)
	fmt.Println("Długie ciągi:", long)
}

func generateRandomString(letters string, length int) string {
	result := make([]byte, length)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func hasConsecutiveDuplicates(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			return true
		}
	}
	return false
}
