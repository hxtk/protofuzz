load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@io_bazel_rules_go//proto:compiler.bzl", "go_proto_compiler")

go_binary(
    name = "protoc-gen-go-fuzz",
    embed = [":protoc-gen-go-fuzz_lib"],
    visibility = ["//visibility:public"],
)

go_proto_compiler(
    name = "go_proto_fuzz",
    plugin = ":protoc-gen-go-fuzz",
    suffix = ".pb.fuzz.go",
    valid_archive = False,
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_google_gofuzz//:go_default_library",
    ],
)

go_library(
    name = "protoc-gen-go-fuzz_lib",
    srcs = ["main.go"],
    importpath = "github.com/hxtk/protofuzz/protoc-gen-go-fuzz",
    visibility = ["//visibility:private"],
    deps = [
        "//protoc-gen-go-fuzz/plugin",
        "@org_golang_google_protobuf//compiler/protogen",
    ],
)
