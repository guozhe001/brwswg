package main

import (
	"encoding/json"
	"fmt"
	pb "github.com/guozhe001/brwswg/chapter6/protofiles"
	"google.golang.org/protobuf/proto"
)

func main() {
	p := &pb.Person{
		Id:    1234,
		Name:  "guozhe",
		Email: "guozhegz@gmail.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "12345678901", Type: pb.Person_MOBILE},
		},
	}
	p1 := &pb.Person{}
	body, _ := proto.Marshal(p)
	_ = proto.Unmarshal(body, p1)
	fmt.Println("Original struct loaded from proto file:", p, "\n")
	fmt.Println("Marshaled proto data:", body, "\n")
	fmt.Println("Unmarshaled struct:", p1)

	json_body, _ := json.Marshal(p)
	fmt.Printf("Marshal json body: %s\n", string(json_body))
}
