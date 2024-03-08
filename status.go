package sms

const (
	Success Status = 0
	Failed  Status = -1
)

type Status int

func (s Status) Int() int {
	return int(s)
}
