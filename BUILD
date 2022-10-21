load("@bazel_gazelle//:def.bzl", "gazelle")
load("@io_bazel_rules_go//go:def.bzl", "TOOLS_NOGO", "nogo")

nogo(
    name = "my_nogo",
    config = "nogo_config.json",
    visibility = ["//visibility:public"],
    deps = TOOLS_NOGO,
)

# gazelle:prefix github.com/hxtk/protofuzz
gazelle(name = "gazelle")

gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.mod",
        "-to_macro=deps.bzl%go_dependencies",
        "-prune",
    ],
    command = "update-repos",
)
