package util

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/corvad/argus/proto/logging"
	"google.golang.org/grpc"
)

type LoggerServer struct {
	pb.UnimplementedLoggerServer
	serverHost         string
	conn               *grpc.Server
	registeredServices map[string]bool
}

func NewLoggerServer(serverHost string) (*LoggerServer, error) {
	l := &LoggerServer{serverHost: serverHost, registeredServices: make(map[string]bool)}

	err := l.Open()
	if err != nil {
		return nil, err
	}

	return l, nil
}

func (l *LoggerServer) Open() error {
	lis, err := net.Listen("tcp", l.serverHost)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	l.conn = grpc.NewServer()

	pb.RegisterLoggerServer(l.conn, l)

	log.Printf("Logging gRPC server started; listening on %s\n", l.serverHost)

	go func() {
		err = l.conn.Serve(lis)
		if err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	return nil
}

func (l *LoggerServer) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.Response, error) {
	service := in.GetService()
	if service == "" {
		return &pb.Response{Success: false, Msg: "Service name is required"}, nil
	}
	if l.registeredServices[service] {
		//ignore duplicate registration
		return &pb.Response{Success: true}, nil
	}
	l.registeredServices[service] = true
	log.Printf("Registered service: %s\n", service)
	return &pb.Response{Success: true}, nil
}

func (l *LoggerServer) Log(ctx context.Context, in *pb.LogRequest) (*pb.Response, error) {
	service := in.GetService()
	if service == "" {
		return &pb.Response{Success: false, Msg: "Service name is required"}, nil
	}
	if !l.registeredServices[service] {
		return &pb.Response{Success: false, Msg: "Service not registered"}, nil
	}
	level := in.GetLevel()
	if level != pb.LogLevel_DEBUG && level != pb.LogLevel_INFO && level != pb.LogLevel_WARN && level != pb.LogLevel_ERROR && level != pb.LogLevel_FATAL {
		return &pb.Response{Success: false, Msg: "Invalid log level"}, nil
	}
	msg := in.GetMsg()
	if msg == "" {
		return &pb.Response{Success: false, Msg: "Log message is required"}, nil
	}
	timestamp := in.GetTimestamp()
	if timestamp == 0 {
		return &pb.Response{Success: false, Msg: "Timestamp is required"}, nil
	}
	//print date for timestamp
	timeDate := time.UnixMilli(int64(timestamp)).Format(time.RFC3339)
	fmt.Printf("Log message: %s %s %s %s\n", service, level, msg, timeDate)
	return &pb.Response{Success: true}, nil
}

func (l *LoggerServer) Close() error {
	l.conn.Stop()
	return nil
}
