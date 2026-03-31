# Argus
My new public Go monorepo.

## Build
```bazel build //...```

## Regen build files
```bazel run //:gazelle```

## Generate .pb.go in place (fixes issues with autocomplete and go mod)
```./scripts/sync_proto.sh```

## Add license headers
```./scripts/license.sh```

## Run go mod tidy
```./scripts/mod_tidy.sh```

## License
Project is under an MIT License.
