package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	next := incrementor()
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
}

func incrementor() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
