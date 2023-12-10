package main

func reverse(x int) int{
	res := 0
	for x!= 0{
		res = (res * 10 ) + x % 10
		x /= 10

		if res > 2_147_483_647 || res < -2_147_483_648{
			return 0
		}
	}
	return res
}

func main(){

}