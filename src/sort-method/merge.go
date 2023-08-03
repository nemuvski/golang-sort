package sortmethod

/*
	分割単位を2として、それぞれの配列をマージしていく
*/
const SPLICE = 2

func MergeSort(arr []int) []int {
	if len(arr) < SPLICE {
		return arr
	}
	mid := len(arr) / SPLICE
	return merge(MergeSort(arr[:mid]), MergeSort(arr[mid:]))
}


func merge(arrA []int, arrB []int) []int {
	total := len(arrA) + len(arrB)
	conquered := make([]int, total, total)
	iA, iB := 0, 0

	for i := 0; i < total; i++ {
		if iA >= len(arrA) && iB < len(arrB) { // arrAが末尾に達したら、arrBの値を入れる
			conquered[i] = arrB[iB]
			iB++
		} else if iA < len(arrA) && iB >= len(arrB){ // arrBが末尾に達したら、arrAの値を入れる
			conquered[i] = arrA[iA]
			iA++
		} else if arrA[iA] < arrB[iB] {
			conquered[i] = arrA[iA]
			iA++
		} else {
			conquered[i] = arrB[iB]
			iB++
		}
	}

	return conquered
}
