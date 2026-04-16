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

genrule(
    name = "rp2350_firmware_uf2",
    srcs = [":rp2350_firmware"],
    outs = ["firmware.uf2"],
    cmd = "$(location @picotool//:picotool) uf2 convert -t elf $$(echo $(SRCS) | cut -d' ' -f1) $@",
    tools = ["@picotool//:picotool"],
)
