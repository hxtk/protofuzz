package tests

import (
	"testing"

	fuzz "github.com/google/gofuzz"
	"google.golang.org/protobuf/proto"
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

func TestFuzzEnum(t *testing.T) {
	t.Run("package enum", func(t *testing.T) {
		fuzzer := fuzz.NewWithSeed(0)
		enumMsg := new(PackageEnum)
		fuzzer.Fuzz(enumMsg)

		for i := 0; i < 1000; i++ {
			if !proto.Equal(enumMsg, new(PackageEnum)) {
				return
			}
		}
		t.Error("PackageEnum unchanged after fuzzing.")
	})

	t.Run("embedded enum", func(t *testing.T) {
		fuzzer := fuzz.NewWithSeed(0)
		enumMsg := new(EmbeddedEnum)
		fuzzer.Fuzz(enumMsg)

		for i := 0; i < 1000; i++ {
			if proto.Equal(enumMsg, new(EmbeddedEnum)) {
				return
			}
		}
		t.Error("PackageEnum unchanged after fuzzing.")
	})
}
