package constants

type SqlState int32

const (
	OPEN SqlState = 1
	CLOSED
	CONNECTING
	ABORTED
)

func (s SqlState) String() string {
	return [...]string{"OPEN", "CLOSED", "CONNECTING", "ABORTED"}[s-1]
}

func (s SqlState) Index() int32 {
	return int32(s)
}
