package main

import "fmt"

// range array에 loop를 적용할 수 있도록 해줌
// range는 index를 리턴
func add(numbers ...int) int {
	for index, number := range numbers {
		fmt.Println(index, number)
	}
	return 1
}

// index를 이용한 forloop
func addLenght(numbers ...int) int {
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	return 1
}

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	add(1, 2, 3, 4, 5, 6)

	addLenght(1, 2, 3, 4, 5, 6)

	// total := add(1, 2, 3, 4, 5, 6)

	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}
