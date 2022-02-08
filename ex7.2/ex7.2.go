package main

import "fmt"

func Divide(a int, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}
func main() {
	c, success := Divide(9, 3)
	fmt.Println(c, success)
	d, success := Divide(9, 0) // success가 재선언되었지만 := 할수 있는 이유는 d가 새로 선언되었기 때문.
	fmt.Println(d, success)
}
