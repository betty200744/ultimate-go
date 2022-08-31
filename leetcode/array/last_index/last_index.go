package last_index

func LastIndex(arr []int, value int) int {
	arrMap := make(map[int]int)
	for i, i2 := range arr {
		arrMap[i2] = i
	}
	return arrMap[value]
}
