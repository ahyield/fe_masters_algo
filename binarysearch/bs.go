package binarysearch

import "fmt"

func BinarySearch(arr []int, num int) (int, error) {
	low := 0
	high := len(arr) - 1

	for i := low; low <= high; i++ {
		marker := high - low/2
		guess := arr[marker]

		if guess == num {
			return num, nil
		} else if guess < num {
			low = marker + 1
		} else {
			high = marker - 1
		}
	}

	return 0, fmt.Errorf("Number doesn't exist")
}
