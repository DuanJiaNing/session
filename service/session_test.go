package service

import (
	"context"
	"testing"

	pb "com/duan/session"
	"session/conf"
)

func init() {
	conf.Init("../app-test.yaml")
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
					Topic: "test topic 4",
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
	type args struct {
		req *pb.OpenRequest
	}
	tests := []struct {
		name    string
		args    args
		success bool
	}{
		{
			name: "success",
			args: args{
				req: &pb.OpenRequest{
					SessionId: 1,
				},
			},
			success: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			response, err := server.Open(context.Background(), tt.args.req)
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
