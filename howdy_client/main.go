package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/corvad/argus/proto/howdy"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	addr := flag.String("addr", "localhost:8080", "server address")
	name := flag.String("name", "David", "your name")
	major := flag.String("major", "Computer Science", "your major")
	town := flag.String("town", "League City, TX")
	class := flag.String("class", "2028", "your class year")
	flag.Parse()
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGigemClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Trigger(ctx, &pb.GigemRequest{Name: *name, Major: *major, Town: *town, Class: *class})
	if err != nil {
		log.Fatalf("Yell failed: %v", err)
	}
	log.Printf("Yell: %s", r.GetResult())

}
