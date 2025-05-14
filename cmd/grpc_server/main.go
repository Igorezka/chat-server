package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/brianvoe/gofakeit"
	desc "github.com/igorezka/chat-server/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

const grpcPort = 50052

type server struct {
	desc.UnimplementedChatV1Server
}

func (*server) Create(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("CreateRequest: %v", req)

	return &desc.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}

func (*server) Delete(_ context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("DeleteRequest: %v", req)

	return &emptypb.Empty{}, nil
}

func (*server) SendMessage(_ context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("SendMessage: %s", req.GetText())

	return &emptypb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
