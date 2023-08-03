package sortmethod

func BubbleSort(arr []int) []int {
	size := len(arr)

	for i := 0; i < size-1; i++ {
		for j := 1; j < size-i; j++ {
			if arr[j-1] > arr[j] {
				tmp := arr[j-1]
				arr[j-1] = arr[j]
				arr[j] = tmp
			}
		}
	}

	return arr
}
