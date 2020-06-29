package function

import (
	"fmt"
	"strings"
)

// func 함수이름 (argument) returnType {}
// func multiply(a int, b int) int {   //int축약가능
func Multiply(a, b int) int {
	return a * b
}

// return 값을 여러개 할수 있음 go의 특별한 기능
func LenAndUpper(name string) (int, string) {
	// strings.ToUpper 대문자변환(UpperCase) package
	return len(name), strings.ToUpper(name)
}

// variadic Function(가변인자함수) 다양한 숫자의 파라미터를 전달하고자 할 때 사용
func RepeatMe(words ...string) {
	fmt.Println(words)
}

// naked function와 defer
// 미리 리턴값을 지정하고 함수내 반환 부분은 인수 없이 return만 사용
func LenAndUppercase(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done") // defer는 함수가 실행되고난 후 실행
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}
