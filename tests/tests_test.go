package tests

import (
	"testing"

	fuzz "github.com/google/gofuzz"
)

func TestFuzzOneof(t *testing.T) {
	fuzzer := fuzz.NewWithSeed(0)
	profile := new(Profile)
	fuzzer.Fuzz(profile)
}

func TestFuzzNil(t *testing.T) {
	fuzzer := fuzz.NewWithSeed(0)
	profile := (*Profile)(nil)
	fuzzer.Fuzz(&profile)
}
