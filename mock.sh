PROJECT_GOPATH="C:/Users/duan/projects2/src/"
mockgen -source="./db/db.go" -destination="./test/mock/db/client.go" -package=mock
#mockgen -source="./vendor/com/duan/session/service.pb.go" -destination="./test/mock/grpc/session_client.go" -package=mock