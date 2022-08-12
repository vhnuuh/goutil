package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"

	pb "agpb/grpc/sample1/protocol"
)

const (
	Address = ":50052"
)

// server is used to implement hello.HelloService
type server struct {
	pb.UnimplementedHelloServiceServer
}

// Hello implements hello.HelloService
func (s *server) Hello(ctx context.Context, in *pb.String) (*pb.String, error) {
	log.Printf("Received: %v", in.GetValue())
	return &pb.String{Value: "Hello " + in.GetValue()}, nil
}

func (s *server) Channel(stream pb.HelloService_ChannelServer) error {
	for {
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		reply := &pb.String{Value: "hello: " + args.GetValue()}

		err = stream.Send(reply)
		if err != nil {
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterHelloServiceServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
