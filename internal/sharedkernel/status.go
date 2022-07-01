package sharedkernel

type Status int

const (
	NEW Status = iota
	// PROCESSING
	// INVALID
	// PROCESSED
)

func (s Status) String() string {
	return [...]string{"NEW", "PROCESSING", "INVALID", "PROCESSED"}[s]
}
