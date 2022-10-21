load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = glob(
        ["*.go"],
        exclude = ["*_test.go"],
    ),
    importpath = "github.com/hxtk/yggdrasil/common/test/protofuzz/protoc-gen-go-fuzz/plugin",
    visibility = [
        "//common/test/protofuzz/protoc-gen-go-fuzz:__pkg__",
    ],
    deps = [
        "@org_golang_google_protobuf//compiler/protogen:go_default_library",
        "@org_golang_google_protobuf//proto:go_default_library",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//types/descriptorpb:go_default_library",
    ],
)

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
