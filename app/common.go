package app

import (
	pb "com/duan/session"
)

func CodeOfSessionStatus(ss pb.SessionStatus) int32 {
	return int32(ss)
}

func CodeOfSessionType(st pb.SessionType) int32 {
	return int32(st)
}
