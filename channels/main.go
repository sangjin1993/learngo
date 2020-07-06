package main

import (
	"fmt"
	"time"
)

func main() {
	// channel 만드는 방법
	c := make(chan string)
	// people := [2]string{"nico", "flynn"}
	// 사람들이 늘어나면 어떻게 할까?
	people := [5]string{"nico", "flynn", "japanguy", "larry"}
	for _, person := range people {
		// channel go 에 보냄
		go isSexy(person, c)
	}
	// resultOne := <-c
	// resultTwo := <-c
	// 오류 발생 blocking operation
	// resultThree := <-c

	// for loop를 이용해서 사용 가능 동시에 결과 return
	for i := 0; i < len(people); i++ {
		// massage를 받는 건 blocking operation
		fmt.Print("waiting for", i)
		fmt.Println(<-c)
	}
	// fmt.Println("Received this message:", resultOne)
	// fmt.Println("Received this message:", resultTwo)
	// 오류 발생
	// fmt.Println("Received this message:", resultThree)
	// goroutine 1 [chan receive] 하나를 더 넣으면 오류 발생
	// fmt.Println(<-c)
}

// channel
func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 5)
	c <- person + " is sexy"
}
