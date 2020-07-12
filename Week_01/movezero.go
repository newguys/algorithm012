//移动零
func MoveZeros(nums[]int)  {
	index :=0
	for i := 0; i < len(nums); i++ {
		if nums[i] !=0 {
			nums[index] = nums[i]
			if i != index {
				nums[i] = 0
			}
			j++
		}
	}
}



func isValid(s string) bool {
	if s == "" {
		return true
	}

	var stack []uint8

	m := map[uint8]uint8{
		'}': '{',
		')': '(',
		']': '[',
	}

	for i := 0; i < len(s); i++ {
		if s[i] == '{' || s[i] == '[' || s[i] == '(' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}
			
			if m[s[i]] != stack[len(stack)-1] {
				return false
			}
			
			stack = stack[:len(stack)-1]
		}
	}
	
	return len(stack) == 0
}

