package main

import (
	"fmt"

	"github.com/itssadon/data-structures-and-algorithms/trees"
)

func main() {
	/*
		Arrays
	*/
	// unsortedArray := []int{9, 6, 5, 0, 3, 4, 1, 2, 14, 10}
	// arrays.InsertionSort(unsortedArray)

	// arrayToMerge := []int{1, 2, 4, 8, 10, 3, 5, 7, 9, 11}
	// arrays.Merge(arrayToMerge, 0, 4, 9)

	// unsortedArray := []int{9, 6, 5, 0, 8, 2}
	// size := reflect.TypeOf(unsortedArray).Size() / reflect.TypeOf(unsortedArray[0]).Size()
	// arrays.MergeSort(unsortedArray, 0, int(size-1))
	// arrays.Display(unsortedArray, int(size))

	/*
		Trees
	*/
	tree := &trees.Node{Key: 100}
	tree.Insert(2000)
	fmt.Println(tree)
}
