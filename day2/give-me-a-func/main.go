package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	myFunc := giveMeAFunc()
	myFunc("Hello world")
}

func giveMeAFunc() func(string) {
	return func(message string) {
		fmt.Println(message)
	}
}
