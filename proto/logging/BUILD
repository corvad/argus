load("@com_google_protobuf//bazel:proto_library.bzl", "proto_library")
load("@rules_go//go:def.bzl", "go_library")
load("@rules_go//proto:def.bzl", "go_proto_library")

proto_library(
    name = "logging_proto",
    srcs = ["logging.proto"],
    visibility = ["//visibility:public"],
)

go_proto_library(
    name = "logging_go_proto",
    compilers = ["@rules_go//proto:go_grpc"],
    importpath = "github.com/corvad/argus/proto/logging",
    proto = ":logging_proto",
    visibility = ["//visibility:public"],
)

go_library(
    name = "logging",
    embed = [":logging_go_proto"],
    importpath = "github.com/corvad/argus/proto/logging",
    visibility = ["//visibility:public"],
)
