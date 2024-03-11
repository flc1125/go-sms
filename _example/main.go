package main

import (
	"context"
	"fmt"
	"os"

	"github.com/flc1125/go-sms"
	"github.com/flc1125/go-sms/driver/writer"
)

func main() {
	s := sms.New(
		writer.New(os.Stdout),
	)
	resp, err := s.Send(context.Background(), &sms.Request{
		Phone:   "1234567890",
		Content: "test",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Status, resp.Message)

	// Output: writer: send sms {1234567890 test <nil>}
}
