package service

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"

	pb "com/duan/session"
	"session/conf"
	"session/test"
)

func init() {
	conf.Init("../app-test.yaml")
}

func Test_create_mock_data(t *testing.T) {
	server := &sessionServer{}
	ctx := context.Background()
	count := 30
	for i := 0; i < count; i++ {
		sid := int64(rand.Intn(20) + 1)
		_, _ = server.Open(ctx, &pb.OpenRequest{
			SessionId: sid,
		})
		_, _ = server.Join(ctx, &pb.JoinRequest{
			SessionId: sid,
			UserId:    fmt.Sprintf("%v", time.Now().UnixNano()),
		})
	}
}

func Test_Join(t *testing.T) {
	server := &sessionServer{}
	type args struct {
		req *pb.JoinRequest
	}
	tests := []struct {
		name    string
		args    args
		success bool
	}{
		{
			name: "success",
			args: args{
				req: &pb.JoinRequest{
					SessionId: 1,
					UserId:    "user01",
				},
			},
			success: true,
		},

		{
			name: "session not exist",
			args: args{
				req: &pb.JoinRequest{
					SessionId: 100001,
					UserId:    "user01111",
				},
			},
			success: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, err := server.Join(context.Background(), tt.args.req)
			if tt.success && err != nil {
				t.Fatal(err)
			}

			if !tt.success && err == nil {
				t.Fatal("should fail")
			}
			t.Log(response)
		})
	}
}

func Test_Create(t *testing.T) {
	server := &sessionServer{}
	type args struct {
		req *pb.CreateRequest
	}
	tests := []struct {
		name    string
		args    args
		success bool
	}{
		{
			name: "success",
			args: args{
				req: &pb.CreateRequest{
					Type:  pb.SessionType_LONG_TERM,
					Topic: test.MockTopic(),
				},
			},
			success: true,
		},

		{
			name: "topic can not be empty",
			args: args{
				req: &pb.CreateRequest{
					Type: pb.SessionType_LONG_TERM,
					// Topic: "test topic 4",
				},
			},
			success: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//ctrl := gomock.NewController(t)
			//defer ctrl.Finish()
			//dsc := mockdb.NewMockClient(ctrl)
			//tt.mock(dsc)
			//db.NewClient = func() (db.Client, error) {
			//	return dsc, nil
			//}

			response, err := server.Create(context.Background(), tt.args.req)
			if tt.success && err != nil {
				t.Fatal(err)
			}

			if !tt.success && err == nil {
				t.Fatal("should fail")
			}
			t.Log(response)
		})
	}
}

func Test_Open(t *testing.T) {
	server := &sessionServer{}

	ctx := context.Background()
	getRequest := func() *pb.OpenRequest {
		res, err := server.Create(ctx, &pb.CreateRequest{
			Type:  0,
			Topic: test.MockTopic(),
		})
		if err != nil {
			t.Fatal(err)
		}

		return &pb.OpenRequest{
			SessionId: res.SessionId,
		}
	}

	type args struct {
		req *pb.OpenRequest
	}
	tests := []struct {
		name    string
		args    args
		before  func(id int64)
		success bool
	}{
		{
			name: "success",
			args: args{
				req: getRequest(),
			},
			success: true,
		},

		{
			name: "session already open",
			args: args{
				req: getRequest(),
			},
			before: func(id int64) {
				_, err := server.Open(ctx, &pb.OpenRequest{
					SessionId: id,
				})
				if err != nil {
					t.Fatal(err)
				}
			},
			success: false,
		},

		{
			name: "session closed, can not be open again",
			args: args{
				req: getRequest(),
			},
			before: func(id int64) {
				_, err := server.Open(ctx, &pb.OpenRequest{
					SessionId: id,
				})
				if err != nil {
					t.Fatal(err)
				}

				_, err = server.Close(ctx, &pb.CloseRequest{
					SessionId: id,
				})
				if err != nil {
					t.Fatal(err)
				}
			},
			success: false,
		},

		{
			name: "session not exist",
			args: args{
				req: &pb.OpenRequest{
					SessionId: math.MaxInt64,
				},
			},
			success: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.before != nil {
				tt.before(tt.args.req.SessionId)
			}
			response, err := server.Open(ctx, tt.args.req)

			if tt.success && err != nil {
				t.Fatal(err)
			}

			if !tt.success && err == nil {
				t.Fatal("should fail")
			}
			t.Log(response)
		})
	}
}

func Test_Close(t *testing.T) {
	server := &sessionServer{}

	ctx := context.Background()
	getRequest := func() *pb.CloseRequest {
		res, err := server.Create(ctx, &pb.CreateRequest{
			Type:  0,
			Topic: test.MockTopic(),
		})
		if err != nil {
			t.Fatal(err)
		}

		return &pb.CloseRequest{
			SessionId: res.SessionId,
		}
	}

	type args struct {
		req *pb.CloseRequest
	}
	tests := []struct {
		name    string
		args    args
		before  func(id int64)
		success bool
	}{
		{
			name: "success",
			args: args{
				req: getRequest(),
			},
			before: func(id int64) {
				_, err := server.Open(ctx, &pb.OpenRequest{
					SessionId: id,
				})
				if err != nil {
					t.Fatal(err)
				}
			},
			success: true,
		},

		{
			name: "session already closed",
			args: args{
				req: getRequest(),
			},
			before: func(id int64) {
				_, err := server.Open(ctx, &pb.OpenRequest{
					SessionId: id,
				})
				if err != nil {
					t.Fatal(err)
				}

				_, err = server.Close(ctx, &pb.CloseRequest{
					SessionId: id,
				})
				if err != nil {
					t.Fatal(err)
				}
			},
			success: false,
		},

		{
			name: "session not open yet",
			args: args{
				req: getRequest(),
			},
			success: false,
		},

		{
			name: "session not exist",
			args: args{
				req: &pb.CloseRequest{
					SessionId: math.MaxInt64,
				},
			},
			success: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.before != nil {
				tt.before(tt.args.req.SessionId)
			}
			response, err := server.Close(ctx, tt.args.req)

			if tt.success && err != nil {
				t.Fatal(err)
			}

			if !tt.success && err == nil {
				t.Fatal("should fail")
			}
			t.Log(response)
		})
	}
}
