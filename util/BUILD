load("@rules_go//go:def.bzl", "go_library")

go_library(
    name = "util",
    srcs = [
        "log.go",
        "log_server.go",
    ],
    importpath = "github.com/corvad/argus/util",
    visibility = ["//visibility:public"],
    deps = [
        "//proto/logging",
        "@org_golang_google_grpc//:grpc",
        "@org_golang_google_grpc//credentials/insecure",
    ],
)
