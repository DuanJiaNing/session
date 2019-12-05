package main

import (
	"log"
	"net"
	"os"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"

	"session/conf"
	mlog "session/log"
	"session/service"
)

func main() {
	conf.Init(getConfig())
	lis, err := net.Listen("tcp", conf.GRpc().ListenAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	service.RegisterSessionServiceServer(s)

	// Register reflection service on gRPC server.
	reflection.Register(s)
	log.Printf("rpc server listeing on %v", conf.GRpc().ListenAddress)
	grpclog.SetLoggerV2(mlog.NewGRpcLogger())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func getConfig() string {
	defaultConfig := "app.yaml"

	if len(os.Args) == 2 {
		configFile := os.Args[1]
		if len(strings.TrimSpace(configFile)) == 0 {
			return defaultConfig
		}

		return configFile
	}

	return defaultConfig
}
