package _78_kthSmallest

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	left := make([]int, 0)
	right := make([]int, 0)
	for i := 1; i < len(arr); i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	left = QuickSort(left)
	right = QuickSort(right)
	return append(append(left, pivot), right...)
}
func KthSmallest(matrix [][]int, k int) int {
	arr := make([]int, 0)
	for i, i2 := range matrix {
		for i3, _ := range i2 {
			arr = append(arr, matrix[i][i3])
		}
	}
	arr = QuickSort(arr)
	return arr[k-1]
}
