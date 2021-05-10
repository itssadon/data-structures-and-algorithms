package arrays

import "fmt"

func InsertionSort(unsortedArray []int) {
	/*
		INSERTION SORT
		Can be remembered as deck of cards taken one by one

		Worst case:
		Time complexity (comparisons and movements): O(n^2)

		Best case:
		Time complexity: omega(n)

		For both cases:
		Space complexity: O(1)

		The time complexity of this algo cannot be reduced even by using binary search or linked list (instead of linear
		search and arrays)
	*/

	arrayValues := unsortedArray
	fmt.Printf("Array to insert-sort is: %d, and length is: %d\n", arrayValues, len(arrayValues))
	for i := 0; i < len(arrayValues); i++ {
		key := arrayValues[i]
		j := i - 1
		for j >= 0 && key < arrayValues[j] {
			arrayValues[j+1] = arrayValues[j]
			j--
		}
		arrayValues[j+1] = key
		//j has been decremented so in the end will be at one less than the position desired, so adding one
	}

	fmt.Printf("The sorted array is: %d\n", arrayValues)
}
