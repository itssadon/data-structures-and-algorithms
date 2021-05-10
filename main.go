package main

import (
	"github.com/itssadon/data-structures-and-algorithms/arrays"
)

func main() {
	/*
		Arrays
	*/
	// unsortedArray := []int{9, 6, 5, 0, 3, 4, 1, 2, 14, 10}
	// arrays.InsertionSort(unsortedArray)

	arrayToMerge := []int{1, 2, 4, 8, 10, 3, 5, 7, 9, 11}
	arrays.Merge(arrayToMerge, 0, 4, 9)
}
