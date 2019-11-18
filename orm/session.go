package orm

type Session struct {
	Id          string
	SessionType int32
	topic       string
	status      int32
	InsertTime  int64
	OpenTime    int64
	CloseTime   int64
}
