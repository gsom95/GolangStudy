package linearsearch

func LinearSearch(arr []int, find int) int {
	for i := range arr {
		if arr[i] == find {
			return i
		}
	}

	return -1
}
