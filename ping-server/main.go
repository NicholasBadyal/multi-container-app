package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/NicholasBadyal/multi-container-app/ping-server/v1/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

type PingPongServer struct {}

func NewPingPongServer() *PingPongServer{
	return &PingPongServer{}
}

func (server *PingPongServer) Ping(ctx context.Context, req *pb.PingRequest)  (*pb.PongResponse, error) {
	ping := req.GetPing()
	log.Printf("Received ping: %s", ping)

	if ctx.Err() == context.Canceled {
		log.Print("request canceled")
		return nil, status.Error(codes.Canceled, "request canceled")
	}
	if ctx.Err() == context.DeadlineExceeded {
		log.Print("deadline exceeded")
		return nil, status.Error(codes.DeadlineExceeded, "deadline exceeded")
	}

	res := &pb.PongResponse{Pong: "Pong"}
	return res, nil
}


func main() {
	port := flag.Int("port", 8080, "the server port")
	flag.Parse()
	log.Printf("started server on port: %d", *port)

	pingServer := NewPingPongServer()

	grpcServer := grpc.NewServer()
	pb.RegisterPingPongServiceServer(grpcServer, pingServer)

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("cannot start server: %v", err)
	}

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("cannot start server: %v", err)
	}
}
