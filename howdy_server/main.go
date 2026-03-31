// Copyright (c) 2026 David Corvaglia. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/corvad/argus/proto/howdy"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGigemServer
}

func (s *server) Trigger(ctx context.Context, in *pb.GigemRequest) (*pb.GigemResponse, error) {
	return &pb.GigemResponse{Result: "Howdy! My name is " + in.GetName() + ". I'm a " + in.GetMajor() + " major from " + in.GetTown() + " , but more importantly, I am a loud and proud member of the Fightin' Texas Aggie Class of " + in.GetClass()}, nil
}

func main() {
	host := fmt.Sprintf("localhost:%d", 8080)
	lis, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGigemServer(s, &server{})
	log.Printf("gRPC server started; listening on %s", host)
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
