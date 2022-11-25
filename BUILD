load("@bazel_gazelle//:def.bzl", "gazelle")
load("@io_bazel_rules_go//go:def.bzl", "TOOLS_NOGO", "nogo")
load("@io_bazel_rules_docker//container:container.bzl", "container_image", "container_push")
load("@rules_pkg//:pkg.bzl", "pkg_tar")

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

# https://google.aip.dev/client-libraries/4290
container_image(
    name = "docker_interface_image",
    base = "@cc_base_image//image",
    debs = [
        "@amd64_debian11_libc6//file",
        "@amd64_debian11_libc-bin//file",
        "@amd64_debian11_libstdcpp6//file",
    ],
    tars = [
        ":protoc_gen_go_fuzz_tar",
        ":protoc_include_tar",
    ],
)

container_push(
    name = "push_latest_docker_interface",
    image = ":docker_interface_image",
    registry = "docker.io",
    repository = "hxtk/protofuzz",
    tag = "latest",
)

pkg_tar(
    name = "protoc_gen_go_fuzz_tar",
    srcs = [
        "//protoc-gen-go-fuzz",
        "@com_google_protobuf//:protoc",
    ],
    package_dir = "/bin",
)

pkg_tar(
    name = "protoc_include_tar",
    srcs = ["@com_google_protobuf//:well_known_type_protos"],
    package_dir = "/usr/include/google/protobuf",
)
