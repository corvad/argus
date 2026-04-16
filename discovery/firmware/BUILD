load("@rules_cc//cc:defs.bzl", "cc_binary")
load("@rules_platform//platform_data:defs.bzl", "platform_data")

cc_binary(
    name = "firmware",
    srcs = glob(["*.c"]),
    deps = [
        "@pico-sdk//src/rp2_common/pico_stdlib",
    ],
)

platform_data(
    name = "rp2350_firmware",
    target = ":firmware",
    platform = "@pico-sdk//bazel/platform:rp2350",
)
