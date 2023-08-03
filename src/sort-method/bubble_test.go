package sortmethod

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	t.Run("BubbleSort: empty array", func(t *testing.T) {
		got := BubbleSort([]int{})
		want := []int{}
		assert.Equal(t, want, got)
	})

	t.Run("BubbleSort: array with 1 element", func(t *testing.T) {
		got := BubbleSort([]int{3})
		want := []int{3}
		assert.Equal(t, want, got)
	})

	t.Run("BubbleSort: array with odd number of elements", func(t *testing.T) {
		got := BubbleSort([]int{8,-3,5,6,3,4,5})
		want := []int{-3,3,4,5,5,6,8}
		assert.Equal(t, want, got)
	})

	t.Run("MergeSort: array with even number of elements", func(t *testing.T) {
		got := BubbleSort([]int{8,-3,5,6,3,4,5,1})
		want := []int{-3,1,3,4,5,5,6,8}
		assert.Equal(t, want, got)
	})
}
