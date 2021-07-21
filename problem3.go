package main

import "fmt"

func main(){
	var x[]int = []int{1,2,3,4,5,}
	sum := 0
	for i := 0 ; i < 5 ; i++ {
		sum += x[i] * x[i]
	}
	fmt.Println("SUm of suqares of x[] = ",sum)

}
