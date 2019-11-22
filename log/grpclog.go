package log

import (
	"fmt"

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
	Info(gRpcLogPrefix, fmt.Sprint(args...))
}

func (g *gRpcLogger) Infoln(args ...interface{}) {
	Info(gRpcLogPrefix, fmt.Sprint(args...))
}

func (g *gRpcLogger) Infof(format string, args ...interface{}) {
	Info(gRpcLogPrefix, fmt.Sprintf(format, args...))
}

func (g *gRpcLogger) Warning(args ...interface{}) {
	Warning(gRpcLogPrefix, fmt.Sprint(args...))
}

func (g *gRpcLogger) Warningln(args ...interface{}) {
	Warning(gRpcLogPrefix, fmt.Sprint(args...))
}

func (g *gRpcLogger) Warningf(format string, args ...interface{}) {
	Warning(gRpcLogPrefix, fmt.Sprintf(format, args...))
}

func (g *gRpcLogger) Error(args ...interface{}) {
	Error(gRpcLogPrefix, fmt.Sprint(args...))
}

func (g *gRpcLogger) Errorln(args ...interface{}) {
	Error(gRpcLogPrefix, fmt.Sprint(args...))
}

func (g *gRpcLogger) Errorf(format string, args ...interface{}) {
	Error(gRpcLogPrefix, fmt.Sprintf(format, args...))
}

func (g *gRpcLogger) Fatal(args ...interface{}) {
	Error(gRpcLogPrefix, fmt.Sprint(args...))
}

func (g *gRpcLogger) Fatalln(args ...interface{}) {
	Error(gRpcLogPrefix, fmt.Sprint(args...))
}

func (g *gRpcLogger) Fatalf(format string, args ...interface{}) {
	Error(gRpcLogPrefix, fmt.Sprintf(format, args...))
}

func (g *gRpcLogger) V(l int) bool {
	return true
}
