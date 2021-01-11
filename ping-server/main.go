package main

import (
	"context"
	"flag"
	"github.com/NicholasBadyal/multi-container-app/ping-server/v1/protos"
	"google.golang.org/grpc"
	"log"
	"net"
)

type PingServer struct {}

func (server *pb.PingPongServiceServer) Ping(context.Context, *pb.PingRequest) (*pb.PongResponse, error) {
	req := &pb.PingResponse{}
}

func main() {
	port := flag.Int("port", 0, "the server port")
	flag.Parse()
	log.Printf("started server on port %d", *port)


}
