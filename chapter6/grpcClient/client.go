package main

import (
	"context"
	pb "github.com/guozhe001/brwswg/chapter6/datafiles"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const address = "localhost:50051"

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("connect error %v", err)
	}
	defer conn.Close()
	c := pb.NewMoneyTransactionClient(conn)
	from := "alice"
	to := "bob"
	amount := float32(1000000)
	r, err := c.MakeTransaction(context.Background(), &pb.TransactionRequest{From: from, To: to, Amount: amount})
	if err != nil {
		log.Fatalf("could not transact %v", err)
	}
	log.Printf("transaction confirmed: %t", r.Confirmation)
}
