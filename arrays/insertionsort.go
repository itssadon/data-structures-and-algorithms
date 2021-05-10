package arrays

import "fmt"

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

func InsertionSort() []int {
	arrayValues := []int{9, 6, 5, 0, 3, 4, 1, 2, 14, 10}
	for i := 0; i < len(arrayValues); i++ {
		key := arrayValues[i]
		j := i - 1
		for j >= 0 && key < arrayValues[j] {
			arrayValues[j+1] = arrayValues[j]
		}
		arrayValues[j+1] = key
	}

	fmt.Println("The sorted array is:")
	// printArray(arrayValues, len(arrayValues))
	return arrayValues
}

func printArray(a []int, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d", a[i])
	}
}
