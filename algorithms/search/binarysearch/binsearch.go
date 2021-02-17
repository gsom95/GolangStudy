package binarysearch

// Binary search requires a sorted array
func BinarySearchRecursive(arr []int, find, left, right int) int {
	if len(arr) == 0 {
		return -1
	}

	if left <= right {
		m := (left + right) / 2

		if find == arr[m] {
			return m
		}

		if find > arr[m] {
			left = m + 1
		} else {
			right = m - 1
		}

		return BinarySearchRecursive(arr, find, left, right)
	}

	return -1
}

func BinarySearchLoop(arr []int, find int) int {
	if len(arr) == 0 {
		return -1
	}

	l := 0
	r := len(arr)

	for l <= r {
		m := (l + r) / 2

		if find == arr[m] {
			return m
		}

		if find > arr[m] {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return -1
}
