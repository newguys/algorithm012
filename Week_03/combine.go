func combine(n int,k int) [][]int {
	if n <1 ||k<1 ||n<k {
		return nil
	}
	cur := make([]int, k,k)
	res := [][]int{}
	helper(cur,0,1,n,k,&res)
	return res
}

func helper(cur []int,ci,start,n,k int ,res *[][]int)  {
	if k == 0 {
		tmp :=make([]int, len(cur))
		copy(tmp,cur)
		*res  = append(*res,tmp)
		return
	}
	for i := start; i <= n-k+1; i++ {
		cur[ci] = i
		helper(cur,ci+1,i+1,n,k,res)
	}
}