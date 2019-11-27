package _interface

import "fmt"

func main() {
	sq1 := new(Square)
	sq1.side = 5

	// var areaIntf Shaper
	// areaIntf = sq1
	// shorter,without separate declaration:
	// areaIntf := Shaper(sq1)
	// or even:
	areaIntf := sq1
	fmt.Printf("The square has area: %f\n", areaIntf.Area())
}
