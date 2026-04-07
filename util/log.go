package util

import (
	"context"
	"fmt"
	"time"

	pb "github.com/corvad/argus/proto/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

type Logger struct {
	serverHost  string
	serviceName string
	conn        *grpc.ClientConn
}

func NewLogger(serverHost string, serviceName string) (*Logger, error) {
	l := &Logger{
		serverHost:  serverHost,
		serviceName: serviceName,
	}

	err := l.Open()
	if err != nil {
		return nil, err
	}

	return l, nil
}

func (l *Logger) Open() error {
	if l.conn != nil {
		return fmt.Errorf("Logger connection already open")
	}

	conn, err := grpc.NewClient(l.serverHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("Failed to connect: %w", err)
	}
	l.conn = conn

	c := pb.NewLoggerClient(l.conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.Register(ctx, &pb.RegisterRequest{Service: l.serviceName})
	if err != nil {
		return fmt.Errorf("Failed to register: %w", err)
	}
	if !res.GetSuccess() {
		return fmt.Errorf("Failed registering with logging service: %s", res.GetMsg())
	}

	return nil
}

func (l *Logger) Log(level LogLevel, msg string) error {
	if l.conn == nil {
		return fmt.Errorf("Logger connection is not open, cannot log message")
	}

	c := pb.NewLoggerClient(l.conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.Log(ctx, &pb.LogRequest{
		Service:   l.serviceName,
		Level:     pb.LogLevel(level),
		Msg:       msg,
		Timestamp: uint64(time.Now().UnixMilli()),
	})
	if err != nil {
		return fmt.Errorf("Failed to log message: %w", err)
	}
	if !res.GetSuccess() {
		return fmt.Errorf("Failed to log message: %s", res.GetMsg())
	}

	return nil
}

func (l *Logger) Debugf(format string, args ...any) error {
	return l.Log(DEBUG, fmt.Sprintf(format, args...))
}

func (l *Logger) Infof(format string, args ...any) error {
	return l.Log(INFO, fmt.Sprintf(format, args...))
}

func (l *Logger) Warnf(format string, args ...any) error {
	return l.Log(WARN, fmt.Sprintf(format, args...))
}

func (l *Logger) Errorf(format string, args ...any) error {
	return l.Log(ERROR, fmt.Sprintf(format, args...))
}

func (l *Logger) Fatalf(format string, args ...any) error {
	return l.Log(FATAL, fmt.Sprintf(format, args...))
}

func (l *Logger) Close() error {
	if l.conn == nil {
		return fmt.Errorf("Logger connection is not open, cannot close")
	}

	err := l.conn.Close()
	if err != nil {
		return fmt.Errorf("Failed to close logger connection: %w", err)
	}
	l.conn = nil

	return nil
}
