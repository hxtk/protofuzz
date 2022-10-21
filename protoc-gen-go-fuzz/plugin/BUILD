load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "plugin",
    srcs = ["plugin.go"],
    importpath = "github.com/hxtk/protofuzz/protoc-gen-go-fuzz/plugin",
    visibility = ["//visibility:public"],
    deps = [
        "@org_golang_google_protobuf//compiler/protogen",
        "@org_golang_google_protobuf//reflect/protoreflect",
    ],
)
