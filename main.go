package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Get the number of items and maximum item value.
	var numItems, maxIt int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&maxIt)

	// Make and display an unsorted slice.
	slice := MakeRandomSlice(numItems, maxIt)
	PrintSlice(slice, 40)
	fmt.Println()

	// Sort and display the result.
	bubbleSort(slice)
	PrintSlice(slice, 40)

	// Verify that it's sorted.
	CheckSorted(slice)
}

// MakeRandomSlice makes a slice containing pseudorandom numbers in [0, max).
func MakeRandomSlice(numItems, max int) []int {
	// Initialize a pseudorandom number generator.
	random := rand.New(rand.NewSource(time.Now().UnixNano())) // Initialize with a changing seed

	slice := make([]int, numItems)
	for i := 0; i < numItems; i++ {
		slice[i] = random.Intn(max)
	}
	return slice
}

// PrintSlice prints at most numItems items.
func PrintSlice(slice []int, numItems int) {
	if len(slice) <= numItems {
		fmt.Println(slice)
	} else {
		fmt.Println(slice[:numItems])
	}
}

// CheckSorted verifies that the slice is sorted.
func CheckSorted(slice []int) {
	for i := 1; i < len(slice); i++ {
		if slice[i-1] > slice[i] {
			println("The slice is NOT sorted!")
			return
		}
	}
	println("The slice is sorted")
}

// Use bubble sort to sort the slice.
func bubbleSort(slice []int) {
	for i := 0; i < len(slice); i++ {
		f := 0
		for j := 0; j < len(slice)-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
				f = 1
			}
		}
		if f == 0 {
			break
		}
	}
}
