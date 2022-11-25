workspace(name = "com_github_hxtk_protofuzz")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive", "http_file")

http_archive(
    name = "bazel_skylib",
    sha256 = "74d544d96f4a5bb630d465ca8bbcfe231e3594e5aae57e1edbf17a6eb3ca2506",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-skylib/releases/download/1.3.0/bazel-skylib-1.3.0.tar.gz",
        "https://github.com/bazelbuild/bazel-skylib/releases/download/1.3.0/bazel-skylib-1.3.0.tar.gz",
    ],
)

load("@bazel_skylib//:workspace.bzl", "bazel_skylib_workspace")

bazel_skylib_workspace()

################################################################################
# Golang Rules

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "099a9fb96a376ccbbb7d291ed4ecbdfd42f6bc822ab77ae6f1b5cb9e914e94fa",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.35.0/rules_go-v0.35.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.35.0/rules_go-v0.35.0.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "efbbba6ac1a4fd342d5122cbdfdb82aeb2cf2862e35022c752eaddffada7c3f3",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.27.0/bazel-gazelle-v0.27.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.27.0/bazel-gazelle-v0.27.0.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
load("//:deps.bzl", "go_dependencies")

# gazelle:repository_macro deps.bzl%go_dependencies
go_dependencies()

go_rules_dependencies()

go_register_toolchains(
    nogo = "@com_github_hxtk_protofuzz//:my_nogo",
    version = "1.19.3",
)

gazelle_dependencies()

################################################################################
# Protobuf Rules

http_archive(
    name = "com_google_protobuf",
    sha256 = "8a815691f25c15e721a2c654ed63662017644ce8f78d895f485566555bb7b25e",
    strip_prefix = "protobuf-21.8",
    urls = ["https://github.com/protocolbuffers/protobuf/archive/refs/tags/v21.8.zip"],
)

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

################################################################################
# Container rules

http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "b1e80761a8a8243d03ebca8845e9cc1ba6c82ce7c5179ce2b295cd36f7e394bf",
    urls = ["https://github.com/bazelbuild/rules_docker/releases/download/v0.25.0/rules_docker-v0.25.0.tar.gz"],
)

load(
    "@io_bazel_rules_docker//repositories:repositories.bzl",
    container_repositories = "repositories",
)

container_repositories()

load("@io_bazel_rules_docker//repositories:deps.bzl", container_deps = "deps")

container_deps()

load(
    "@io_bazel_rules_docker//container:container.bzl",
    "container_pull",
)

container_pull(
    name = "cc_base_image",
    digest = "sha256:6865ad48467c89c3c3524d4c426f52ad12d9ab7dec31fad31fae69da40eb6445",
    registry = "gcr.io",
    repository = "distroless/cc-debian11",
    tag = "debug",
)

################################################################################
# Package Rules

http_archive(
    name = "rules_pkg",
    sha256 = "eea0f59c28a9241156a47d7a8e32db9122f3d50b505fae0f33de6ce4d9b61834",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_pkg/releases/download/0.8.0/rules_pkg-0.8.0.tar.gz",
        "https://github.com/bazelbuild/rules_pkg/releases/download/0.8.0/rules_pkg-0.8.0.tar.gz",
    ],
)

load("@rules_pkg//:deps.bzl", "rules_pkg_dependencies")

rules_pkg_dependencies()

http_file(
    name = "amd64_debian11_libc6",
    downloaded_file_path = "libc6_2.36-5_amd64.deb",
    sha256 = "768acd8ae472f3797ff21eb7aeb6ae4b59ae2fa8bb5a3a360bd3a69ebfe44df1",
    urls = ["https://snapshot.debian.org/archive/debian/20221120T085602Z/pool/main/g/glibc/libc6_2.36-5_amd64.deb"],
)

http_file(
    name = "amd64_debian11_libc-bin",
    downloaded_file_path = "libc-bin_2.31-13+deb11u5_amd64.deb",
    sha256 = "b858bfc2bdb558a3fd3482e77c7e751acdd814210b30ff8387da26c5923bca86",
    urls = ["https://snapshot.debian.org/archive/debian/20221120T085602Z/pool/main/g/glibc/libc-bin_2.36-5_amd64.deb"],
)

http_file(
    name = "amd64_debian11_libstdcpp6",
    downloaded_file_path = "libstdc++6_12.2.0-9_amd64.deb",
    sha256 = "7d70dcf69236f27d7ebb0bc0ebbd4afbb44db3c0815c29384f349881b05cd673",
    urls = ["https://snapshot.debian.org/archive/debian/20221120T085602Z/pool/main/g/gcc-12/libstdc%2B%2B6_12.2.0-9_amd64.deb"],
)
