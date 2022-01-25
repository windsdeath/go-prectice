package main

import "fmt"

func main() {
	var a float32 = 1234.523
	var b float32 = 3456.123
	var c float32 = a * b
	var d float32 = c * 3

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

/*
1234.523
3456.123
4.266663e+06 // 계산시 4,266,663.334329 이지만 4.266663아래로 누락
1.2799989e+07 // 위의 계산에서 *3을 하여 12,799,990.002987이 되어야 하지만 1.2799989 부분에서
값이 누락됨
-> 앞의 0.3정도의 오차에서 *3을 하였더니 1이상의 오차가 발생이됨
-> 오차는 연산을 거듭할수록 커질 수 있음
-> 고로 실수의 연산에서는 조심하여야 함
*/
