package main

import "fmt"

func main() {
	f := generator()
	for i := 0; i < 5; i++ {
		fmt.Println(f(), "\t")
	}
}

func generator() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
