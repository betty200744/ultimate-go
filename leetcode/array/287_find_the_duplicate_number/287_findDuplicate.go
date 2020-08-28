package _87_find_the_duplicate_number

func FindDuplicate(nums []int) int {
	numMap := make(map[int]int)
	for i, i2 := range nums {
		if _, has := numMap[i2]; has {
			return i2
		} else {
			numMap[i2] = i
		}
	}
	return -1
}
