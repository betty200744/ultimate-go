package _78_kthSmallest

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestQuickSort(t *testing.T) {
	res := QuickSort([]int{2, 1, 6, 3})
	assert.Equal(t, []int{1, 2, 3, 6}, res)
}
func TestKthSmallest(t *testing.T) {
	matrix := [][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}
	res := KthSmallest(matrix, 8)
	assert.Equal(t, 13, res)
}
func TestKthSmallest2(t *testing.T) {
	matrix := [][]int{{1, 2}, {1, 3}}
	res := KthSmallest(matrix, 2)
	assert.Equal(t, 1, res)
}
