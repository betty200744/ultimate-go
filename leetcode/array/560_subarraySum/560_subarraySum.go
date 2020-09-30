package _60_subarraySum

func SubarraySum(nums []int, k int) int {
	count := 0
	for s := 0; s < len(nums); s++ {
		for e := s + 1; e < len(nums); e++ {
			sum := 0
			for i := s; i < e; i++ {
				sum += nums[i]
			}
			count++
			if sum == k {
				return count
			}
		}
	}
	return 0
}
