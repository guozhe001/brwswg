package main

import (
	"context"
	pb "github.com/guozhe001/brwswg/chapter6/serverPush/datafiles"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
)

const (
	address = "localhost:50051"
)

func ReceiveStream(client pb.MoneyTransactionClient, request *pb.TransactionRequest) {
	log.Println("started listening to the server stream!")
	stream, err := client.MakeTransaction(context.Background(), request)
	if err != nil {
		log.Fatalf("%v.MakeTransaction(_) = _, %v", client, err)
	}
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.MakeTransaction(_)=_,%v", client, err)
		}
		log.Printf("Status: %v, Operation: %v", response.Status, response.Description)
	}
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("dit not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewMoneyTransactionClient(conn)
	from := "alice"
	to := "to"
	amount := float32(123123)
	ReceiveStream(client, &pb.TransactionRequest{From: from, To: to, Amount: amount})
}
