/*
1:1
2:2
3:f(1)+f(2)
4:f(2)+f(3)

f(n) = f(n-1)+f(n-2)
*/
func climbStaires(n int) int {
	if n <=2 {
		return n
	}
	f1,f2,f3 := 0,0,1
	for i := 1; i < n; i++ {
		f1 = f2;
		f2 = f3;
		f3 = f1+f2
	}
	return f3
}