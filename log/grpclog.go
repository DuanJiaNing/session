package log

import (
	"fmt"
	"log"

	"google.golang.org/grpc/grpclog"
)

const (
	gRpcLogPrefix = "grpc-"
)

type gRpcLogger struct {
}

func NewGRpcLogger() grpclog.LoggerV2 {
	return &gRpcLogger{}
}

func (g *gRpcLogger) Info(args ...interface{}) {
	log.Print(gRpcLogPrefix + "info: " + fmt.Sprint(args))
}

func (g *gRpcLogger) Infoln(args ...interface{}) {
	log.Println(gRpcLogPrefix + "info: " + fmt.Sprint(args))
}

func (g *gRpcLogger) Infof(format string, args ...interface{}) {
	log.Print(gRpcLogPrefix + "info: " + fmt.Sprintf(format, args))
}

func (g *gRpcLogger) Warning(args ...interface{}) {
	log.Print(gRpcLogPrefix + "warn: " + fmt.Sprint(args))
}

func (g *gRpcLogger) Warningln(args ...interface{}) {
	log.Println(gRpcLogPrefix + "warn: " + fmt.Sprint(args))
}

func (g *gRpcLogger) Warningf(format string, args ...interface{}) {
	log.Print(gRpcLogPrefix + "warn: " + fmt.Sprintf(format, args))
}

func (g *gRpcLogger) Error(args ...interface{}) {
	log.Print(gRpcLogPrefix + "error: " + fmt.Sprint(args))
}

func (g *gRpcLogger) Errorln(args ...interface{}) {
	log.Println(gRpcLogPrefix + "error: " + fmt.Sprint(args))
}

func (g *gRpcLogger) Errorf(format string, args ...interface{}) {
	log.Print(gRpcLogPrefix + "error: " + fmt.Sprintf(format, args))
}

func (g *gRpcLogger) Fatal(args ...interface{}) {
	log.Fatal(gRpcLogPrefix + "fatal: " + fmt.Sprint(args))
}

func (g *gRpcLogger) Fatalln(args ...interface{}) {
	log.Fatalln(gRpcLogPrefix + "fatal: " + fmt.Sprint(args))
}

func (g *gRpcLogger) Fatalf(format string, args ...interface{}) {
	log.Fatal(gRpcLogPrefix + "fatal: " + fmt.Sprintf(format, args))
}

func (g *gRpcLogger) V(l int) bool {
	return true
}
