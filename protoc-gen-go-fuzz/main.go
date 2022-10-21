package main

import (
	"flag"

	"google.golang.org/protobuf/compiler/protogen"

	"github.com/hxtk/protofuzz/protoc-gen-go-fuzz/plugin"
)

func main() {
	var (
		flags flag.FlagSet
	)
	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if f.Generate {
				plugin.GenerateFile(gen, f)
			}
		}
		return nil
	})
}
