package handler

import (
	"context"
	"net"

	"parkour/pb"

	"github.com/vimcoders/go-driver/grpcx"
	"github.com/vimcoders/go-driver/log"
)

type Option struct {
	Info struct {
		Address string `yaml:"address"`
	} `yaml:"info"`
}

type Handler struct {
	pb.ParkourServer
}

// MakeHandler creates a Handler instance
func MakeHandler(ctx context.Context) *Handler {
	h := &Handler{}
	return h
}

// Handle receives and executes redis commands
func (x *Handler) Handle(ctx context.Context, c net.Conn) {
	client := grpcx.NewClient(c, grpcx.Option{ServiceDesc: pb.Parkour_ServiceDesc})
	client.Register(ctx, x)
}

// Close stops handler
func (x *Handler) Close() error {
	return nil
}

func (x *Handler) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	log.Info("Ping")
	return &pb.PingResponse{Message: req.Message}, nil
}

func (x *Handler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	log.Info("Login")
	return &pb.LoginResponse{}, nil
}
