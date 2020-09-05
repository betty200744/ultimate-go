package _98_house_Robber

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func Rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[1]
	}
	if len(nums) == 2 {
		return Max(nums[0], nums[1])
	}
	a, b := nums[0], Max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		a, b = b, Max(a+nums[i], b)
	}
	return b
}
