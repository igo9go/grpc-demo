package main

import (
	"log"
	"net"
	pb "grpc-demo/go/helloworld"
	"grpc-demo/go/meta"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)                    // server is used to implement helloworld.GreeterServer.
type server struct{} // SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name + "--from go server"}, nil
}

func (s *server) GetInfo(ctx context.Context, in *meta.MetaRequest) (*meta.MetaResponse, error) {
	return &meta.MetaResponse{
		AppId: in.AppId,
		Title: "my title",
		Name:  "My Name",
		Logo:  "https://xxx.com",
	}, nil
}
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	meta.RegisterAgentServer(s, &server{})
	s.Serve(lis)
}
