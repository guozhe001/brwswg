package main

import (
	"log"
	"net/rpc"
)

type Args struct {
}

func main() {
	var reply int64
	args := Args{}
	clent, err := rpc.DialHTTP("tcp", "localhost"+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	err = clent.Call("TimeServer.GiveServerTime", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	log.Printf("%d\n", reply)
}
