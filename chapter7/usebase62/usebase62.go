package main

import (
	"github.com/guozhe001/brwswg/chapter7/base62"
	"log"
)

func main() {
	x := 100
	toBase62 := base62.ToBase62(x)
	log.Println(toBase62)
	base10 := base62.ToBase10(toBase62)
	log.Println(base10)
}
