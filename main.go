package main

import (
	"fmt"
	"strings"
)

// func 함수이름 (argument) returnType {}
// func multiply(a int, b int) int {   //int축약가능
func multiply(a, b int) int {
	return a * b
}

// return 값을 여러개 할수 있음 go의 특별한 기능
func lenAndUpper(name string) (int, string) {
	// strings.ToUpper 대문자변환(UpperCase) package
	return len(name), strings.ToUpper(name)
}

// variadic Function(가변인자함수) 다양한 숫자의 파라미터를 전달하고자 할 때 사용
func repeatMe(words ...string) {
	fmt.Println(words)
}

// naked function와 defer
// 미리 리턴값을 지정하고 함수내 반환 부분은 인수 없이 return만 사용
func lenAndUppercase(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done") // defer는 함수가 실행되고난 후 실행
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	fmt.Println(multiply(2, 2))
	// lenAndUpper함수의 리턴값이 totalLenght, upperName 변수에 들어감

	totalLenght, upperName := lenAndUpper("nico")
	// 순서대로 len(), strings.ToUpper()를 totalLenght, upperName에 값 리턴
	// totalLenght = len("nico") => 4
	// upperName = strings.ToUpper("nico") => NICO

	fmt.Println(totalLenght, upperName) // 4 NICO
	// 사용 불가 cannot initialize 1 variables with 2 values
	// totalLenght := lenAndUpper("nico")
	// upperName := lenAndUpper("nico")

	// underscore(_) 사용하면 ignore value
	// len()은 사용 strings.ToUpper()은 무시함
	lenght, _ := lenAndUpper("sang")
	_, upper := lenAndUpper("sang")
	fmt.Println(lenght) // 4
	fmt.Println(upper)  // SANG

	repeatMe("nico", "lynn", "dal", "marl", "flynn") //[nico lynn dal marl flynn]

	// naked function
	totalLenght, up := lenAndUppercase("nico")
	fmt.Println(totalLenght, up)
}
