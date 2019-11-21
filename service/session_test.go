package service

import (
	"context"
	"session/conf"
	"testing"

	"github.com/golang/mock/gomock"

	pb "com/duan/session"
	"session/db"
	mockdb "session/test/mock/db"
)

func init() {
	conf.Init("../app.yaml")
}

func Test_CreateSession(t *testing.T) {
	server := &sessionServer{}
	type args struct {
		req *pb.CreateSessionRequest
	}
	tests := []struct {
		name    string
		args    args
		success bool
		mock    func(mc *mockdb.MockClient)
	}{
		{
			name:    "success",
			args:    args{},
			success: true,
			mock:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			dsc := mockdb.NewMockClient(ctrl)
			tt.mock(dsc)
			db.NewClient = func() (db.Client, error) {
				return dsc, nil
			}

			response, err := server.CreateSession(context.Background(), tt.args.req)
			if tt.success && err != nil {
				t.Fatal(err)
			}
			t.Log(response)
		})
	}
}
