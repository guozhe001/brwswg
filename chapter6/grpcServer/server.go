package main

import (
	"context"
	pb "github.com/guozhe001/brwswg/chapter6/datafiles"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const port = ":50051"

type server struct {
	pb.UnimplementedMoneyTransactionServer
}

func (s *server) MakeTransaction(ctx context.Context, in *pb.TransactionRequest) (*pb.TransactionResponse, error) {
	log.Println("got request for money transfer...")
	log.Printf("Amount:%f, From A/c:%s, To A/c:%s", in.Amount, in.From, in.To)
	return &pb.TransactionResponse{Confirmation: true}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMoneyTransactionServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
