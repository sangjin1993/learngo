package main

import (
	"fmt"

	"github.com/sangjin1993/learngo/something"
)

func main() {
	fmt.Println("Hello world!")
	something.SayHello()
	//something.sayBye() 소문자로 시작하는 함수는 export가 안됌.
}
