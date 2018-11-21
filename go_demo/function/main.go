package main

import "fmt"

func changeValue(a int) int {
	a = a + 1
	return a
}

func ChangePointer(a *int) {
	*a = *a + 1
	return
}

func sum(arr ...int) (sum int) {
	for _, v := range arr {
		sum += v
	}
	return
}

//匿名函数被直接赋值函数变量
var sumInt = func(a, b int) int {
	return a + b
}

func doInput(f func(int, int) int, a, b int) int {
	return f(a, b)
}

func swap(op string) func(int, int) int {
	switch op {
	case "add":
		return func(a int, b int) int {
			return a + b
		}
	case "mult":
		{
			return func(a int, b int) int {
				return a * b
			}
		}
	default:
		return nil
	}
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println(sum(1, 2))

	//匿名函数作为实参
	fmt.Println(doInput(func(i int, i2 int) int {
		return i + i2
	}, 5, 2))

	opFunc := swap("mult")
	re := opFunc(2, 3)
	fmt.Println(re)
}
