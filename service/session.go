package service

import (
	ses "com/duan/session"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	defaultMaxCallSendMsgSize = 1024 * 1024        // 1M
	defaultMaxCallRecvMsgSize = 1024 * 1024 * 1024 // 1G
	target                    = "host:port"
)

var (
	Client = newClient()
)

func newClient() ses.SessionServiceClient {
	conn, err := grpc.Dial(
		target,
		grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")),
		//grpc.WithUserAgent(userAgent),
		//grpc.WithChainUnaryInterceptor(
		//	statisticInterceptor,
		//	clientInterceptor,
		//),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallSendMsgSize(defaultMaxCallSendMsgSize),
			grpc.MaxCallRecvMsgSize(defaultMaxCallRecvMsgSize),
		),
	)
	if err != nil {
		panic(err)
	}

	return ses.NewSessionServiceClient(conn)
}
