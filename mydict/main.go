package main

import (
	"fmt"

	"github.com/sangjin1993/learngo/mydict/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}

	// dictionary["hello"] = "hello"
	// fmt.Println(dictionary)

	fmt.Println(dictionary["first"])
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

	// word := "hello"
	// hello := "Greeting"
	// err2 := dictionary.Add(word, definition)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// hello, _ := dictionary.Search(word)
	// fmt.Println(hello)
}
