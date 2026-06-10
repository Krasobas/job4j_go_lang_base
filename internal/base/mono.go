package base

func Mono(nums []int) bool {
	increasing := true
	decreasing := true
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			increasing = false
		}
		if nums[i] < nums[i+1] {
			decreasing = false
		}
	}
	return increasing || decreasing
}
