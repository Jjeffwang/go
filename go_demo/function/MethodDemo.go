package main

import "fmt"

func main() {
	var i1 int
	var f1 float64
	i1, _, f1 = getIntANDFloat(10)
	fmt.Print(i1, f1)

}
func getIntANDFloat(int int) (int, bool, float64) {
	return int, false, float64(int) * 1.12
}
