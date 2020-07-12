//两数之和  先做两次遍历的方法
func twoSum(nums []int, taregt int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	for i := 0; i < len(nums)-1; i++ {
		for j := i+1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i,j}
			}
		}
	}
	return []int{}
}
