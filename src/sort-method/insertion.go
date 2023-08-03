package sortmethod

func InsertionSort(arr []int) []int {
	size := len(arr)

	for i := 1; i < size; i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			tmp := arr[j-1]
			arr[j-1] = arr[j]
			arr[j] = tmp
		}
	}

	return arr
}
