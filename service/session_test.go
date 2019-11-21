package service

import (
	"context"
	"session/conf"
	"testing"

	pb "com/duan/session"
	mockdb "session/test/mock/db"
)

func init() {
	conf.Init("../app-test.yaml")
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
			name: "success",
			args: args{
				req: &pb.CreateSessionRequest{
					Type:  pb.SessionType_LONG_TERM,
					Topic: "test topic 4",
				},
			},
			success: true,
			mock:    nil,
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

			response, err := server.CreateSession(context.Background(), tt.args.req)
			if tt.success && err != nil {
				t.Fatal(err)
			}
			t.Log(response)
		})
	}
}
