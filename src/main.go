package main

import (
	"fmt"
	sortmethod "golang-sort/src/sort-method"
	"reflect"
	"runtime"
	"strings"
	"time"
)

func main() {
	sampleArr := []int{38,53,50,56,13,-4,-53,99,-23,2,4,100,10,90,43,41,8,90,-1,39,40,6,7,9,11,12,14,15,16,17,3,18,19,20,21,22,23,24,15}

	fmt.Println("Target Array\t", sampleArr)

	runMethod(sortmethod.BubbleSort, sampleArr)
	runMethod(sortmethod.InsertionSort, sampleArr)
	runMethod(sortmethod.MergeSort, sampleArr)
}

/*
	関数の実行に加えて、実行時間を出力する
*/
func runMethod(method func([]int) []int, arr []int) {
	start := time.Now()
	res := method(arr)
	t := time.Now()
	fmt.Println(getFnName(method, true), "\t", res, "\t", t.Sub(start))
}

/*
	関数名を取得する
*/
func getFnName(fn interface{}, shortName bool) string {
	name := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	if shortName {
		splitted := strings.Split(name, ".")
		return splitted[len(splitted)-1]
	}
	return name
}
