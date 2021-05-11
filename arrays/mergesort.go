package arrays

import "fmt"

func merge(arr []int, start int, mid int, end int) {
	/*
		MERGE SORT
		Space complexity: stack O(logn) + space for merge procedure O(n)
		Therefore total = O(n)
		Time complexity: time taken to merge O(n) + time taken to sort by masters theorem O(nlogn)
		Therefore: O(nlogn)
	*/
	len1 := mid - start + 2
	len2 := end - mid + 1

	left := make([]int, len1)
	right := make([]int, len2)

	i := 0
	for i = 0; i < len1-1; i++ {
		left[i] = arr[start+i]
	}

	for i = 0; i < len2-1; i++ {
		right[i] = arr[mid+i+1]
	}

	left[len1-1] = int(0)
	right[len2-1] = int(0)

	var k, j int
	i, j = 0, 0
	for k = start; k <= end; k++ {
		if left[i] < right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
	}
}

func MergeSort(arr []int, start int, end int) {
	var mid int
	if start < end {
		mid = (start + end) / 2
		MergeSort(arr, start, mid)
		MergeSort(arr, mid+1, end)

		merge(arr, start, mid, end)
	}
}

func Display(arr []int, size int) {
	for i := 0; i < size; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println("")
}
