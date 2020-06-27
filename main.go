package main

import "fmt"

// name3 := "sang" 축약형 사용불가

func main() {
	const name1 string = "nico"
	// name1 = "lynn" cannot assign to name1
	// const name bool = "nico" cannot convert to bool
	fmt.Println(name1)

	var name2 string = "nico"
	// 축약형은 func안에서만 사용가능
	// name2 := "nico" 축약형 go가 type를 자동으로 찾아줌.
	name2 = "lynn"
	fmt.Println(name2)

}
