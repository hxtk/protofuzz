# gazelle:exclude

load("@rules_proto//proto:defs.bzl", "proto_library")
load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_skylib//rules:build_test.bzl", "build_test")

proto_library(
    name = "test_proto",
    srcs = ["test.proto"],
    visibility = ["//visibility:public"],
    deps = [
        "@com_google_protobuf//:empty_proto",
    ],
)

go_proto_library(
    name = "tests_go_proto",
    compilers = [
        "@io_bazel_rules_go//proto:go_proto",
        "//protoc-gen-go-fuzz:go_proto_fuzz",
    ],
    importpath = "github.com/hxtk/protofuzz/tests",
    proto = ":test_proto",
    visibility = ["//visibility:public"],
    deps = ["@io_bazel_rules_go//proto/wkt:empty_go_proto"],
)

go_library(
    name = "tests",
    embed = [":tests_go_proto"],
    importpath = "github.com/hxtk/protofuzz/tests",
    visibility = ["//visibility:public"],
)

go_test(
    name = "tests_test",
    srcs = ["tests_test.go"],
    embed = [":tests"],
)

build_test(
    name = "build_test",
    targets = [":tests"],
)
