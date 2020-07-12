//暴力
func rotate(nums []int, k int)  {
	for i := 0; i < k; i++ {
		pre := nums[len(nums)-1]
		for i := 0; i < len(nums); i++ {
			tmp := nums[j]
			nums[j] = pre
			pre = tmp
		}
	}
}

//三次循环

func rotate1(nums []int, k int)  {
	n := len(nums)
	reverse(nums,0,n-1)
	reverse(nums,0,k%n-1)
	reverse(nums,k%n,n-1)
}

func reverse(nums []int,start,end int)  {
	for start <end{
		nums[start],nums[end] = nums[end],nums[start]
		start++
		end--
	}
}