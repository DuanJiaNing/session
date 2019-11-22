package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-xorm/xorm"
	"google.golang.org/grpc"

	pb "com/duan/session"
	"session/app"
	"session/db"
	"session/orm"
	"session/utils"
)

type sessionServer struct {
}

func RegisterSessionServiceServer(s *grpc.Server) {
	pb.RegisterSessionServiceServer(s, &sessionServer{})
}

func (s *sessionServer) Close(ctx context.Context, req *pb.CloseRequest) (*pb.CloseResponse, error) {
	client, err := db.NewClient()
	if err != nil {
		return nil, app.WithInternalError(err)
	}

	se, err := getSessionById(client, req.SessionId)
	if err != nil {
		return nil, app.WithInternalError(err)
	}
	if se == nil {
		return nil, app.Error("session not exist")
	}

	if app.CodeOfSessionStatus(pb.SessionStatus_NEWBORN) == se.Status {
		return nil, app.Error("session not open yet")
	}

	if app.CodeOfSessionStatus(pb.SessionStatus_SESSION_CLOSED) == se.Status {
		return nil, app.Error("session already closed")
	}

	se.Status = app.CodeOfSessionStatus(pb.SessionStatus_SESSION_CLOSED)
	affected, err := client.ID(se.Id).Update(se)
	if err = app.CheckDbExecuteResult(affected, err, 1); err != nil {
		return nil, err
	}

	return &pb.CloseResponse{
		Status:    pb.SessionStatus_SESSION_CLOSED,
		SessionId: se.Id,
	}, nil
}

func (s *sessionServer) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	client, err := db.NewClient()
	if err != nil {
		return nil, app.WithInternalError(err)
	}

	if utils.BlankString(req.Topic) {
		return nil, app.Error("topic can not be empty")
	}

	topic := req.Topic
	ss := &orm.Session{
		Status:      app.CodeOfSessionStatus(pb.SessionStatus_NEWBORN),
		Topic:       topic,
		SessionType: app.CodeOfSessionType(pb.SessionType_LONG_TERM),
	}
	i, err := client.InsertOne(ss)
	if err = app.CheckDbExecuteResult(i, err, 1); err != nil {
		return nil, err
	}

	return &pb.CreateResponse{
		SessionId: ss.Id,
	}, nil
}

func (s *sessionServer) Open(ctx context.Context, req *pb.OpenRequest) (
	*pb.OpenResponse, error) {
	client, err := db.NewClient()
	if err != nil {
		return nil, app.WithInternalError(err)
	}

	se, err := getSessionById(client, req.SessionId)
	if err != nil {
		return nil, app.WithInternalError(err)
	}
	if se == nil {
		return nil, app.Error("session not exist")
	}

	if app.CodeOfSessionStatus(pb.SessionStatus_SESSION_OPEN) == se.Status {
		return nil, app.Error("session already open")
	}

	if app.CodeOfSessionStatus(pb.SessionStatus_SESSION_CLOSED) == se.Status {
		return nil, app.Error("session closed, can not be open again")
	}

	se.Status = app.CodeOfSessionStatus(pb.SessionStatus_SESSION_OPEN)
	affected, err := client.ID(se.Id).Update(se)
	if err = app.CheckDbExecuteResult(affected, err, 1); err != nil {
		return nil, err
	}

	return &pb.OpenResponse{
		Status:    pb.SessionStatus_SESSION_OPEN,
		SessionId: se.Id,
	}, nil
}

func getSessionById(client *xorm.Engine, sessionId int64) (*orm.Session, error) {
	session := &orm.Session{}
	has, err := client.ID(sessionId).Get(session)
	if !has || err != nil {
		return nil, err
	}

	return session, nil
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
