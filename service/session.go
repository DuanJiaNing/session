package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-xorm/xorm"
	"google.golang.org/grpc"

	pb "com/duan/session"
	"session/app"
	"session/db"
	"session/orm"
	"session/utils"
)

const (
	defaultUserInSessionTTL = 0 // no limit
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

	if int32(pb.SessionStatus_NEWBORN) == se.Status {
		return nil, app.Error("session not open yet")
	}

	if int32(pb.SessionStatus_SESSION_CLOSED) == se.Status {
		return nil, app.Error("session already closed")
	}

	se.Status = int32(pb.SessionStatus_SESSION_CLOSED)
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
		Status:      int32(pb.SessionStatus_NEWBORN),
		Topic:       topic,
		SessionType: int32(pb.SessionType_LONG_TERM),
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

	if int32(pb.SessionStatus_SESSION_OPEN) == se.Status {
		return nil, app.Error("session already open")
	}

	if int32(pb.SessionStatus_SESSION_CLOSED) == se.Status {
		return nil, app.Error("session closed, can not be open again")
	}

	se.Status = int32(pb.SessionStatus_SESSION_OPEN)
	affected, err := client.ID(se.Id).Update(se)
	if err = app.CheckDbExecuteResult(affected, err, 1); err != nil {
		return nil, err
	}

	return &pb.OpenResponse{
		Status:    pb.SessionStatus_SESSION_OPEN,
		SessionId: se.Id,
	}, nil
}

func getParliamentary(client *xorm.Engine, sessionId int64, userId string) (*orm.Parliamentary, error) {
	pa := &orm.Parliamentary{}
	pa.SessionId = sessionId
	pa.UserId = userId
	has, err := client.Get(pa)
	if !has || err != nil {
		return nil, err
	}

	return pa, nil
}

func getSessionById(client *xorm.Engine, sessionId int64) (*orm.Session, error) {
	session := &orm.Session{}
	has, err := client.ID(sessionId).Get(session)
	if !has || err != nil {
		return nil, err
	}

	return session, nil
}

func (s *sessionServer) ParliamentaryCountInSession(ctx context.Context, req *pb.ParliamentaryCountInSessionRequest) (
	*pb.ParliamentaryCountInSessionResponse, error) {
	// SELECT COUNT(*) FROM parliamentary WHERE `status`=1 AND  session_id=1
	client, err := db.NewClient()
	if err != nil {
		return nil, app.WithInternalError(err)
	}

	count, err := client.Where("status=? AND session_id=?", int32(pb.ParliamentaryStatus_IN), req.SessionId).Count()
	if err != nil {
		return nil, app.WithInternalError(err)
	}

	return &pb.ParliamentaryCountInSessionResponse{
		Count: count,
	}, nil
}

func (s *sessionServer) Join(ctx context.Context, req *pb.JoinRequest) (*pb.JoinResponse, error) {
	client, err := db.NewClient()
	if err != nil {
		return nil, app.WithInternalError(err)
	}

	// check session status
	se, err := getSessionById(client, req.SessionId)
	if err != nil {
		return nil, app.WithInternalError(err)
	}
	if se == nil {
		return nil, app.Error("session not exist")
	}

	if se.Status == int32(pb.SessionStatus_SESSION_CLOSED) {
		return nil, app.Error("session closed, can not join")
	}

	if se.Status == int32(pb.SessionStatus_NEWBORN) {
		return nil, app.Error("session not open yet, open session first")
	}

	if se.Status != int32(pb.SessionStatus_SESSION_OPEN) {
		return nil, app.InternalErrorf("session status error, status code: %v", se.Status)
	}

	// check parliamentary
	pd, err := getParliamentary(client, req.SessionId, req.UserId)
	if err != nil {
		return nil, app.WithInternalError(err)
	}
	if pd != nil {
		pd.Status = int32(pb.ParliamentaryStatus_IN)
		pd.ReactiveTime = time.Now()
		i, err := client.Update(pd)
		if err = app.CheckDbExecuteResult(i, err, 1); err != nil {
			return nil, err
		}
	} else {
		pd = &orm.Parliamentary{
			SessionId: req.SessionId,
			UserId:    req.UserId,
			Ttl:       defaultUserInSessionTTL,
			Status:    int32(pb.ParliamentaryStatus_IN),
		}
		i, err := client.InsertOne(pd)
		if err = app.CheckDbExecuteResult(i, err, 1); err != nil {
			return nil, err
		}
	}

	return &pb.JoinResponse{
		Data: &pb.SessionData{
			Session: &pb.Session{
				Id:     se.Id,
				Type:   pb.SessionType(se.SessionType),
				Status: pb.SessionStatus(se.Status),
				Topic:  se.Topic,
			},
			User: &pb.Parliamentary{
				Id:        pd.Id,
				Status:    pb.ParliamentaryStatus(pd.Status),
				UserId:    pd.UserId,
				SessionId: pd.SessionId,
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
