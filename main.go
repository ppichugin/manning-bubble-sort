package main

import (
	"fmt"

	"github.com/ppichugin/manning-bubble-sort/helper"
)

func main() {
	// Get the number of items and maximum item value.
	var numItems, maxIt int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&maxIt)

	// Make and display an unsorted slice.
	slice := helper.MakeRandomSlice(numItems, maxIt)
	helper.PrintSlice(slice, 40)
	fmt.Println()

	// Sort and display the result.
	bubbleSort(slice)
	helper.PrintSlice(slice, 40)

	// Verify that it's sorted.
	helper.CheckSorted(slice)
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
