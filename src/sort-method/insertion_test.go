package sortmethod

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSort(t *testing.T) {
	t.Run("InsertionSort: empty array", func(t *testing.T) {
		assert.Equal(t, []int{}, InsertionSort([]int{}))
	})

	t.Run("InsertionSort: array with 1 element", func(t *testing.T) {
		assert.Equal(t, []int{3}, InsertionSort([]int{3}))
	})

	t.Run("InsertionSort: array with odd number of elements", func(t *testing.T) {
		assert.Equal(t, []int{-3,3,4,5,5,6,8}, InsertionSort([]int{8,-3,5,6,3,4,5}))
	})

	t.Run("InsertionSort: array with even number of elements", func(t *testing.T) {
		assert.Equal(t, []int{-3,1,3,4,5,5,6,8}, InsertionSort([]int{8,-3,5,6,3,4,5,1}))
	})
}
