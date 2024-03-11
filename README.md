# Go SMS

[![Go Version](https://badgen.net/github/release/flc1125/go-sms/stable)](https://github.com/flc1125/go-sms/releases)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-kratos-ecosystem)](https://pkg.go.dev/github.com/flc1125/go-sms)
[![codecov](https://codecov.io/gh/flc1125/go-sms/graph/badge.svg?token=i8RgDRbbDD)](https://codecov.io/gh/flc1125/go-sms)
[![Go Report Card](https://goreportcard.com/badge/github.com/flc1125/go-sms)](https://goreportcard.com/report/github.com/flc1125/go-sms)
[![lint](https://github.com/flc1125/go-sms/actions/workflows/lint.yml/badge.svg)](https://github.com/flc1125/go-sms/actions/workflows/lint.yml)
[![test](https://github.com/flc1125/go-sms/actions/workflows/test.yml/badge.svg)](https://github.com/flc1125/go-sms/actions/workflows/test.yml)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

## Installation

```bash
go get github.com/flc1125/go-sms
```

## Usage

### sms with driver

```go
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
		writer.NewDriver(os.Stdout),
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
```

### Drivers

- [writer](https://github.com/flc1125/go-sms/tree/main/driver/writer)
- [mitake](https://github.com/flc1125/go-sms/tree/main/driver/mitake)

## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.