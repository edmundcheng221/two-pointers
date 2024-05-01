package main

import (
	"errors"
	"fmt"
)

func main() {
	arr := []int{2, 3, 6, 9, 14, 83, 112} // target exists in idx 3 and 4
	target := 23
	fmt.Println(twoPointer(arr, target))

	arr = []int{} // empty slice
	target = 23
	fmt.Println(twoPointer(arr, target))

	arr = []int{2, 3, 6, 9, 14, 83, 112} // no possible way to add target
	target = 43
	fmt.Println(twoPointer(arr, target))

	arr = []int{2, 3, 6, 22, 14, 83, 112} // cannot work on unsorted array
	target = 43
	fmt.Println(twoPointer(arr, target))

}

func isSorted(arr []int) (bool, error) {
	if len(arr) < 2 {
		return true, nil
	}

	prevIdx, currIdx := 0, 1

	for currIdx < len(arr) {
		if arr[prevIdx] > arr[currIdx] {
			return false, nil
		}
		prevIdx += 1
		currIdx += 1
	}

	return true, nil
}

func twoPointer(arr []int, target int) (start int, end int, err error) {

	if len(arr) == 0 {
		err = errors.New("slice length is 0")
		return
	}

	isSorted, err := isSorted(arr)

	if err != nil {
		return
	}

	if !isSorted {
		err = errors.New("can only work with sorted slices")
		return
	}

	start, end = 0, len(arr)-1
	sum := arr[start] + arr[end]

	for sum != target {
		if start == end {
			break
		}
		if sum < target {
			start += 1
			sum = arr[start] + arr[end]
		} else if sum > target {
			end -= 1
			sum = arr[start] + arr[end]
		}

		if sum == target {
			return
		}
	}

	// reset start and end
	start, end = 0, 0
	return
}
