package forloop

import "fmt"

// range array에 loop를 적용할 수 있도록 해줌
// range는 index를 리턴
func Add(numbers ...int) int {
	for index, number := range numbers {
		fmt.Println(index, number)
	}
	return 1
}

// index를 이용한 forloop
func AddLenght(numbers ...int) int {
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	return 1
}

func SuperAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
