load("@rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "log_server_lib",
    srcs = ["main.go"],
    importpath = "github.com/corvad/argus/util/cmd/log_server",
    visibility = ["//visibility:private"],
    deps = ["//util"],
)

go_binary(
    name = "log_server",
    embed = [":log_server_lib"],
    visibility = ["//visibility:public"],
)
