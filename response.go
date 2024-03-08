package sms

type ResponseExtra interface {
	IsResponseExtra()
}

type UnimplementedResponseExtra struct{}

func (u *UnimplementedResponseExtra) IsResponseExtra() {}

type Response struct {
	Status  Status        `json:"status"`
	Message string        `json:"message"`
	Extra   ResponseExtra `json:"extra"`
}
