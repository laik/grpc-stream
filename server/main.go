package main

import (
	"github.com/laik/grpc-stream/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

var _ proto.StreamServiceServer = &StreamServer{}

type StreamServer struct{}

func (s StreamServer) List(request *proto.StreamRequest, server proto.StreamService_ListServer) error {
	n := 0
	for {
		err := server.Send(&proto.StreamResponse{
			Pt: &proto.StreamPoint{
				Name:  request.Pt.Name,
				Value: request.Pt.Value + int32(n),
			},
		})
		if err != nil {
			return err
		}
		time.Sleep(10000 * time.Second)
		n++

		if n >= 100 {
			break
		}
	}

	return nil
}

func (s StreamServer) Record(server proto.StreamService_RecordServer) error {
	panic("implement me")
}

func (s StreamServer) Route(server proto.StreamService_RouteServer) error {
	panic("implement me")
}

func main() {
	server := grpc.NewServer()
	proto.RegisterStreamServiceServer(server, &StreamServer{})

	listen, err := net.Listen("tcp", ":"+"8080")
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	if err := server.Serve(listen); err != nil {
		panic(err)
	}
}
