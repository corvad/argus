load("@rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "log_client_lib",
    srcs = ["main.go"],
    importpath = "github.com/corvad/argus/util/cmd/log_client",
    visibility = ["//visibility:private"],
    deps = ["//util"],
)

go_binary(
    name = "log_client",
    embed = [":log_client_lib"],
    visibility = ["//visibility:public"],
)
