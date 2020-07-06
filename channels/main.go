package main

import (
	"fmt"
	"time"
)

func main() {
	// channel 만드는 방법
	c := make(chan bool)
	people := [2]string{"nico", "flynn"}
	for _, person := range people {
		// channel go 에 보냄
		go isSexy(person, c)
	}
	fmt.Println(<-c)
	fmt.Println(<-c)
	// goroutine 1 [chan receive] 하나를 더 넣으면 오류 발생
	// fmt.Println(<-c)
}

// channel
func isSexy(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	fmt.Println(person)
	c <- true
}
