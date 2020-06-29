package main

import (
	"fmt"

	"github.com/sangjin1993/learngo/function"

	forloop "github.com/sangjin1993/learngo/for"
)

func main() {

	fmt.Println("================== function ======================")

	fmt.Println(function.Multiply(2, 2))
	// lenAndUpper함수의 리턴값이 totalLenght, upperName 변수에 들어감

	totalLenght, upperName := function.LenAndUpper("nico")
	// 순서대로 len(), strings.ToUpper()를 totalLenght, upperName에 값 리턴
	// totalLenght = len("nico") => 4
	// upperName = strings.ToUpper("nico") => NICO

	fmt.Println(totalLenght, upperName) // 4 NICO
	// 사용 불가 cannot initialize 1 variables with 2 values
	// totalLenght := lenAndUpper("nico")
	// upperName := lenAndUpper("nico")

	// underscore(_) 사용하면 ignore value
	// len()은 사용 strings.ToUpper()은 무시함
	lenght, _ := function.LenAndUpper("sang")
	_, upper := function.LenAndUpper("sang")
	fmt.Println(lenght) // 4
	fmt.Println(upper)  // SANG

	function.RepeatMe("nico", "lynn", "dal", "marl", "flynn") //[nico lynn dal marl flynn]

	// naked function
	totalLenght, up := function.LenAndUppercase("nico")
	fmt.Println(totalLenght, up)

	fmt.Println("========================================")

	fmt.Println("=================== for =====================")

	// forloop Add 함수 실행
	forloop.Add(1, 2, 3, 4, 5, 6)

	// forloop AddLength 실행 len함수를 이용해서 출력
	forloop.AddLenght(1, 2, 3, 4, 5, 6)

	// total := add(1, 2, 3, 4, 5, 6)

	result := forloop.SuperAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}
