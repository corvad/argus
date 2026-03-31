load("@rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "howdy_client_lib",
    srcs = ["main.go"],
    importpath = "github.com/corvad/argus/howdy_client",
    visibility = ["//visibility:public"],
    deps = [
        "//proto/howdy",
        "@org_golang_google_grpc//:grpc",
        "@org_golang_google_grpc//credentials/insecure",
    ],
)

go_binary(
    name = "howdy_client",
    embed = [":howdy_client_lib"],
    visibility = ["//visibility:public"],
)
