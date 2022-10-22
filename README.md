# ProtoFuzz

ProtoFuzz is a `protoc` plugin to generate implementations of [`github.com/google/gofuzz.Interface`](https://pkg.go.dev/github.com/google/gofuzz#Interface) for generated Protobuf types.

## Usage

```go
package foo_test

import (
	"context"
	"testing"

	fuzz "github.com/google/gofuzz"

	pb "example.com/your/proto/v1"
)

func FuzzSomeMethod(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		fuzzer := fuzz.FromGoFuzz(data).WithNilChance(0.01)

		var request pb.SomeMethodRequest
		fuzzer.Fuzz(&request)

		// request has now been initialized to a random value
		// suitable for use in a fuzz test.
	}
}
```
