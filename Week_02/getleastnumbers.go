//sort 
// func getLeastNumbers(arr []int, k int) []int {
//     if len(arr) == 0{
//         return []int{}
//     }
//     sort.Ints(arr)
//     return arr[:k]
// }

//最大堆
func getLeastNumbers(arr []int, k int) []int {
    if k==0 || k>len(arr){
        return []int{}
    }
    h := &intHeap{}
    heap.Init(h)
    for _,v:=range(arr){
        if h.Len()<k{
            heap.Push(h, v)
        }else{
            if (*h)[0]>v{
                heap.Pop(h)
                heap.Push(h, v)
            }
        }
    }
    
    res:=[]int{}
    for h.Len()>0{
        res = append(res, heap.Pop(h).(int))
    }
    return res
}

type intHeap []int

func (h intHeap) Len() int {
    return len(h)
}

func (h intHeap) Less(i, j int) bool {
    return h[i] > h[j]
}

func (h intHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *intHeap) Push(x interface{}) {
    // Push 使用 *h，是因为
    // Push 增加了 h 的长度
    *h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
    // Pop 使用 *h ，是因为
    // Pop 减短了 h 的长度
    res := (*h)[len(*h)-1]
    *h = (*h)[:len(*h)-1]
    return res
}
