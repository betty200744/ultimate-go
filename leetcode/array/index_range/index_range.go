package index_range

func IndexRange(arr []int, value int) []int {
	arrMap := make(map[int][]int)
	for i, i2 := range arr {
		arrMap[i2] = append(arrMap[i2], i)
	}
	return []int{arrMap[value][0], arrMap[value][len(arrMap[value])-1]}
}
