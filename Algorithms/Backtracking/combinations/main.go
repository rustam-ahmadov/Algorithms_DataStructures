package main

func combine(n int, k int) [][]int {
    res := [][]int{}  
    helper(&res, []int{}, 1, n , k)
	return res
}

func helper(res *[][]int, cur []int, l, r,  k int){
    if len(cur) == k {
		c := append([]int{}, cur...)
        *res = append(*res, c)
        return
    }

    new := append([]int{}, cur...)
    for i:=l; i <=r ; i++ {
        new = append(new, i)
        helper(res, new, i + 1, r, k)
        new = new[:len(new) - 1]
    } 
}


func main(){
	combine(4, 2)
}