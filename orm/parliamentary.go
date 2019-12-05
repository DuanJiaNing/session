package orm

import "time"

type Parliamentary struct {
	Id           int64 `xorm: "pk autoincr"`
	Ttl          int32
	SessionId    int64
	UserId       string
	Status       int32
	InsertTime   time.Time
	JoinTime     time.Time
	ReactiveTime time.Time
	LeaveTime    time.Time
	CloseTime    time.Time
}
