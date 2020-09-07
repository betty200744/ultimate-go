package _1_exchange

func Exchange(nums []int) []int {
	length := len(nums)
	if length < 2 {
		return nums
	}
	j := 0
	for i := 0; length-i-j > 0; {
		if nums[i]%2 == 0 {
			nums[i], nums[length-j-1] = nums[length-j-1], nums[i]
			j++
		} else {
			i++
		}
	}
	return nums
}
