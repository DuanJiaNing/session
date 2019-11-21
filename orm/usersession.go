package orm

import "time"

type UserSession struct {
	id           int64 `xorm: "pk autoincr"`
	ttl          int32
	SessionId    int64
	UserId       string
	status       int32
	InsertTime   time.Time
	JoinTime     time.Time
	ReactiveTime time.Time
	LeaveTime    time.Time
	CloseTime    time.Time
}
