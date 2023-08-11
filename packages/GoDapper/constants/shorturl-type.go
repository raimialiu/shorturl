package constants

type ShortUrlType int

const (
	STANDARD ShortUrlType = 1
	ONE_OFF
	EXPIRING
	PERMANENT
)

func (s ShortUrlType) String() string {
	return [...]string{"STANDARD", "ONE_OFF", "EXPIRING", "PERMANENT"}[s-1]
}

func (s ShortUrlType) Index() int32 {
	return int32(s)
}
