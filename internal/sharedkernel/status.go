package sharedkernel

type Status int

const (
	NEW Status = iota + 1
	PROCESSING
	INVALID
	PROCESSED
)

func (s Status) String() string {
	return [...]string{"NEW", "PROCESSING", "INVALID", "PROCESSED"}[s+1]
}
