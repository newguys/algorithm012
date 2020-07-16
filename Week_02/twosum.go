
//[3,2,4]  6  两遍哈希 
func twoSum(nums []int,target int)[]int{
	if len(nums) == 0 {
		return []int{}
	}
	m := map[int]int{}
	for i:=0;i<len(nums);i++{
		m[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		sub := target - nums[i]
		v,ok := m[sub]
		if ok && v != i {
			return []int{i,v}
		}
	}
	return []int{}
}

//一遍哈希
func twoSum(nums []int,target int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		sub := target -nums[i]
		v,ok := m[sub]
		if ok && v != i {
			return []int{v,i}
		}
		m[nums[i]] = i
	}
	return []int{}
}