package arrays

import (
	"fmt"
)

func Merge(arr []int, p int, q int, r int) {
	/*
		MERGE PROCEDURE

		Time complexity: O(n)
		Space complexity: O(n)
		as we are nearly creating DS of sizes same as input
	*/

	fmt.Printf("Array to merge is: %d, and length is: %d\n", arr, len(arr))

	n1 := q - p + 1
	n2 := r - q

	var i, j, f int
	arr1 := make([]int, n1+1)
	arr2 := make([]int, n2+1)

	for i = 0; i < n1; i++ {
		arr1[i] = arr[i]
	}

	for j = 0; j < n2; j++ {
		arr2[j] = arr[q+j+1]
	}

	arr1[n1], arr2[n2] = arr2[n2], f

	i, j = 0, 0

	for z := 0; z < len(arr); z++ {
		if arr1[i] <= arr2[j] {
			arr[z] = arr1[i]
			i++
		} else {
			arr[z] = arr2[j]
			j++
		}
	}

	printArray(arr, len(arr))
}

func printArray(a []int, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", a[i])
	}
}
