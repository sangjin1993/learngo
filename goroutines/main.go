package main

import (
	"fmt"
	"time"
)

// go를 붙혀서 병렬적으로 처리
func main() {
	go sexyCount("nico")
	// go를 사용하면은 프로그램이 작동할 때만 움직임 결국 메인 함수가 끝남
	// 메인 함수는 goroutines을 기다리지 않음.
	go sexyCount("flynn")
	time.Sleep(time.Second * 5)
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}
