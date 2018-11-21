package main

import "fmt"

func main() {
	// :=短类型引用，只能用在函数内7
	//a := "Hello World"
	//b := a[0:4]
	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Println("this first go ")

	arr := [3]int{7, 2, 3}
	for i := range arr {
		fmt.Println(arr[i])

	}

L1:
	for i := 0; ; i++ {
		for j := 0; ; j++ {
			if i >= 5 {
				fmt.Println("555")
				break L1
			}
			if j > 10 {
				fmt.Println("10 0")
				break
				//break L1
			}
		}
	}
}
