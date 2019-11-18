package orm

type UserSession struct {
	id           string
	ttl          int32
	SessionId    string
	UserId       string
	status       int32
	InsertTime   int64
	JoinTime     int64
	ReactiveTime int64
	LeaveTime    int64
	CloseTime    int64
}
