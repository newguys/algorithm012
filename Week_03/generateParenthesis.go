

func generateParenthesis(n int) []string {
	generate(0,0,n,"")
	return []string{}	
}

func generate(left int,right int, n int, s string) []string {
	//terminator
	res := []string{}
	if left ==n &&  right == n {
		//filter
		/*
		left 随时可以加 只要别超
		right 左>右
		*/
		res = append(res,s)
		//fmt.Println(s)
		return res
	}

	//process curent logic,left right

	//dill down
	if left < n {
		generate(left+1,right,n,s+"(")
	}
	if left >right {
		generate(left,right+1,n,s+")")
	}
	
	//reserse states
}