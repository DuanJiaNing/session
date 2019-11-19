package orm

import "time"

type Session struct {
	Id          *int64
	SessionType *int32
	Topic       *string
	Status      *int32
	InsertTime  *time.Time
	OpenTime    *time.Time
	CloseTime   *time.Time
}
