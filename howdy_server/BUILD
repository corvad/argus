load("@rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "howdy_server_lib",
    srcs = ["main.go"],
    importpath = "github.com/corvad/argus/howdy_server",
    visibility = ["//visibility:private"],
    deps = [
        "//proto/howdy",
        "@org_golang_google_grpc//:grpc",
    ],
)

go_binary(
    name = "howdy_server",
    embed = [":howdy_server_lib"],
    visibility = ["//visibility:public"],
)
