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

type SessionServer struct {
}

func (s *SessionServer) CreateSession(ctx context.Context, req *pb.CreateSessionRequest) (*pb.CreateSessionResponse, error) {
	engine, err := db.Engine()
	if err != nil {
		return nil, app.RpcError(err, "Internal error")
	}

	if utils.BlankString(req.Topic) {
		return nil, app.NewError("topic can not be empty")
	}

	code := int32(pb.SessionType_LONG_TERM)
	topic := req.Topic
	status := int32(pb.SessionStatus_NEWBORN)
	ss := &orm.Session{
		Status:      &status,
		Topic:       &topic,
		SessionType: &code,
	}
	i, err := engine.InsertOne(ss)
	if err != nil {
		return nil, app.RpcError(err, "Internal error")
	}
	if i != 1 {
		return nil, app.DbError()
	}

	return &pb.CreateSessionResponse{
		SessionId: *ss.Id,
	}, nil
}

func (s *SessionServer) UpdateSessionStatus(context.Context, *pb.UpdateSessionStatusRequest) (*pb.UpdateSessionStatusResponse, error) {
	panic("implement me")
}

func RegisterSessionServiceServer(s *grpc.Server) {
	pb.RegisterSessionServiceServer(s, &SessionServer{})
}

func (s *SessionServer) UserCountInSession(context.Context, *pb.UserCountInSessionRequest) (
	*pb.UserCountInSessionResponse, error) {
	fmt.Println("UserCountInSession been called...")
	return nil, errors.New("not support yet")
}

func (s *SessionServer) Join(ctx context.Context, req *pb.JoinRequest) (*pb.JoinResponse, error) {
	engine, err := db.Engine()
	if err != nil {
		return nil, app.RpcError(err, "Internal error")
	}

	// check session exist
	//insert

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

func (s *SessionServer) RefreshStatus(context.Context, *pb.RefreshStatusRequest) (*pb.RefreshStatusResponse, error) {
	fmt.Println("RefreshStatus been called...")
	return nil, errors.New("not support yet")
}

func (s *SessionServer) Pause(context.Context, *pb.PauseRequest) (*pb.PauseResponse, error) {
	fmt.Println("Pause been called...")
	return nil, errors.New("not support yet")
}

func (s *SessionServer) Leave(context.Context, *pb.LeaveRequest) (*pb.LeaveResponse, error) {
	fmt.Println("Leave been called...")
	return nil, errors.New("not support yet")
}
