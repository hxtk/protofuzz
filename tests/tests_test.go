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