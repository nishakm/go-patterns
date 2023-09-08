package main

import (
	"fmt"
	_ "plugins/english"
	"plugins/greet"
)

func main() {
	greeting, err := greet.In("english")
	fmt.Println(greeting)
	greeting, err = greet.In("hindi")
	if err != nil {
		fmt.Println(err)
	}
}
