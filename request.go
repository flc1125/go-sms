package sms

type RequestExtra interface {
	IsRequestExtra()
}

type UnimplementedRequestExtra struct{}

func (u *UnimplementedRequestExtra) IsRequestExtra() {}

type Request struct {
	IDDCode string       `json:"idd_code"`
	Phone   string       `json:"phone"`
	Content string       `json:"content"`
	Extra   RequestExtra `json:"extra"`
}
