package service

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc"

	pb "com/duan/session"
)

type sessionServer struct {
}

func RegisterSessionServiceServer(s *grpc.Server) {
	pb.RegisterSessionServiceServer(s, &sessionServer{})
}

func (s *sessionServer) UserCountInSession(context.Context, *pb.UserCountInSessionRequest) (
	*pb.UserCountInSessionResponse, error) {
	fmt.Println("UserCountInSession been called...")
	return nil, errors.New("not support yet")
}

func (s *sessionServer) Join(ctx context.Context, req *pb.JoinRequest) (*pb.JoinResponse, error) {
	fmt.Println("Join been called...")
	fmt.Println(req)
	//return nil, errors.New("not support yet")
	return &pb.JoinResponse{
		UserSession: &pb.UserSessionData{
			User: &pb.UserSession{
				UserId: "rpc-" + req.UserId,
			},
		},
	}, nil
}

func (s *sessionServer) RefreshStatus(context.Context, *pb.RefreshStatusRequest) (*pb.RefreshStatusResponse, error) {
	fmt.Println("RefreshStatus been called...")
	return nil, errors.New("not support yet")
}

func (s *sessionServer) Pause(context.Context, *pb.PauseRequest) (*pb.PauseResponse, error) {
	fmt.Println("Pause been called...")
	return nil, errors.New("not support yet")
}

func (s *sessionServer) Leave(context.Context, *pb.LeaveRequest) (*pb.LeaveResponse, error) {
	fmt.Println("Leave been called...")
	return nil, errors.New("not support yet")
}
