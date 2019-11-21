package service

import (
	"context"
	"errors"
	"fmt"
	"session/orm"
	"session/utils"

	"google.golang.org/grpc"

	pb "com/duan/session"
	"session/app"
	"session/db"
)

type sessionServer struct {
}

func RegisterSessionServiceServer(s *grpc.Server) {
	pb.RegisterSessionServiceServer(s, &sessionServer{})
}

func (s *sessionServer) CreateSession(ctx context.Context, req *pb.CreateSessionRequest) (*pb.CreateSessionResponse, error) {
	client, err := db.NewClient()
	if err != nil {
		return nil, app.WithInternalError(err)
	}

	if utils.BlankString(req.Topic) {
		return nil, app.Error("topic can not be empty")
	}

	code := int32(pb.SessionType_LONG_TERM)
	topic := req.Topic
	status := int32(pb.SessionStatus_NEWBORN)
	ss := &orm.Session{
		Status:      status,
		Topic:       topic,
		SessionType: code,
	}
	i, err := client.InsertOne(ss)
	if err != nil {
		return nil, app.WithInternalError(err)
	}
	if i != 1 {
		return nil, app.DbExecuteEffectRowsIncorrect()
	}

	return &pb.CreateSessionResponse{
		SessionId: ss.Id,
	}, nil
}

func (s *sessionServer) UpdateSessionStatus(ctx context.Context, req *pb.UpdateSessionStatusRequest) (
	*pb.UpdateSessionStatusResponse, error) {
	se := getSessionById(req.SessionId)
	if se == nil {
		return nil, app.Error("session not exist")
	}
	//TODO
	return nil, nil
}

func getSessionById(sessionId int64) *orm.Session {
	return nil
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
