# ProtoFuzz

ProtoFuzz is a `protoc` plugin to generate implementations of [`github.com/google/gofuzz.Interface`](https://pkg.go.dev/github.com/google/gofuzz#Interface) for generated Protobuf types.

## Usage

### Generating Code

The ProtoFuzz `protoc` plugin may be installed by running `go install github.com/hxtk/protofuzz/protoc-gen-go-fuzz@latest`.
This process requires that `protoc` be installed on the host machine.
Users may invoke `protoc-gen-go-fuzz` through protoc by specifying the argument `--go-fuzz_out`.
Additional options may be specified via `--go-fuzz_opt`, e.g., `--go-fuzz_opt=paths=source_relative`.

Alternatively, a docker image is provided at `docker.io/hxtk/protofuzz`.
This image is (mostly) compliant with https://google.aip.dev/client-libraries/4290.
It does not currently support the PROTO files found in `github.com/googleapis/googleapis` as default includes.
To use the docker interface, mount the input and output directories into the container at `/in` and `/out`, as shown below.

```sh
docker run --rm --user $UID \
    --mount type=bind,source=$PWD/path/to/api/v1,destination=/in/path/to/api/v1,readonly \
    --mount type=bind,source=$PWD/path/to/dest,destination=/out/ \
    docker.io/hxtk/protofuzz:latest
```

For example, to generate the proto files for this repository's `test` directory, a user with a shell in the repository root might run the following command:

```
docker run --rm --user $UID \
    --mount type=bind,source=$PWD/tests,destination=/in/github.com/hxtk/protofuzz/tests,readonly \
    --mount type=bind,source=$PWD/tests,destination=/out/github.com/hxtk/protofuzz/tests \
    docker.io/hxtk/protofuzz:latest
```

### Writing Fuzzers

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
