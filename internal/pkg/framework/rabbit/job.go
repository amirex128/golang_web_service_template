package rabbit

type Job interface {
	Encode() ([]byte, error)
	Length() int
	Topic() string
}
