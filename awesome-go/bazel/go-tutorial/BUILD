load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

go_binary(
    name = "main",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)

# gazelle:prefix github.com/example/project
gazelle(
    name = "gazelle",
    prefix = "go-tutorial",
)

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/example/project",
    visibility = ["//visibility:private"],
)
