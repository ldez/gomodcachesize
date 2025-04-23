# gomodcachesize

Calculates the module cache size of dependencies (`GOMODCACHE`) for a module.

The module cache is calculated is an isolated directory, then other module dependencies are ignored. 

## Install

```bash
go install github.com/ldez/gomodcachesize@latest
```

## Usage

```console
$ gomodcachesize -h
Usage of gomodcachesize:
  -json
        JSON output format
  -mode string
        Mode: download, tidy (default "download")
```

## Examples

<details><summary>on Kubernetes</summary>

```console
$ gomodcachesize           
|                                             MODULE                                             |    SIZE     |
|------------------------------------------------------------------------------------------------|-------------|
| cache                                                                                          | 61.050 MiB  |
| google.golang.org/genproto@v0.0.0-20240123012728-ef4313101c80                                  | 43.664 MiB  |
| golang.org/x/text@v0.19.0                                                                      | 39.210 MiB  |
| github.com/gogo/protobuf@v1.3.2                                                                | 17.264 MiB  |
| k8s.io/kube-openapi@v0.0.0-20241105132330-32ad38e42d3f                                         | 15.303 MiB  |
| google.golang.org/protobuf@v1.35.1                                                             | 10.879 MiB  |
| golang.org/x/sys@v0.26.0                                                                       | 8.904 MiB   |
| golang.org/x/tools@v0.26.0                                                                     | 7.919 MiB   |
| google.golang.org/grpc@v1.65.0                                                                 | 7.403 MiB   |
| github.com/google/pprof@v0.0.0-20241029153458-d1b30febd7db                                     | 7.211 MiB   |
| golang.org/x/net@v0.30.0                                                                       | 6.178 MiB   |
| sigs.k8s.io/kustomize/kyaml@v0.18.1                                                            | 5.814 MiB   |
| go.opentelemetry.io/otel@v1.28.0                                                               | 5.497 MiB   |
| github.com/grpc-ecosystem/grpc-gateway/v2@v2.20.0                                              | 5.280 MiB   |
| golang.org/x/crypto@v0.28.0                                                                    | 3.978 MiB   |
| github.com/chai2010/gettext-go@v1.0.2                                                          | 2.889 MiB   |
| github.com/grpc-ecosystem/grpc-gateway@v1.16.0                                                 | 2.727 MiB   |
| github.com/google/cel-go@v0.22.0                                                               | 2.619 MiB   |
| go.etcd.io/etcd/server/v3@v3.5.16                                                              | 2.318 MiB   |
| github.com/onsi/ginkgo/v2@v2.21.0                                                              | 2.307 MiB   |
| go.etcd.io/etcd/pkg/v3@v3.5.16                                                                 | 2.243 MiB   |
| sigs.k8s.io/structured-merge-diff/v4@v4.4.2                                                    | 2.159 MiB   |
| github.com/opencontainers/runc@v1.2.1                                                          | 1.978 MiB   |
| github.com/containerd/containerd/api@v1.7.19                                                   | 1.818 MiB   |
| github.com/google/cadvisor@v0.51.0                                                             | 1.755 MiB   |
| github.com/prometheus/procfs@v0.15.1                                                           | 1.755 MiB   |
| sigs.k8s.io/kustomize/api@v0.18.0                                                              | 1.601 MiB   |
| cel.dev/expr@v0.18.0                                                                           | 1.575 MiB   |
| github.com/google/gnostic-models@v0.6.8                                                        | 1.523 MiB   |
| golang.org/x/exp@v0.0.0-20240719175910-8a7402abbf56                                            | 1.468 MiB   |
| github.com/onsi/gomega@v1.35.1                                                                 | 1.171 MiB   |
| github.com/vishvananda/netlink@v1.3.1-0.20240905180732-b1ce50cfa9be                            | 1.119 MiB   |
| go.etcd.io/etcd/api/v3@v3.5.16                                                                 | 1.047 MiB   |
| github.com/prometheus/client_golang@v1.19.1                                                    | 1.032 MiB   |
| google.golang.org/genproto/googleapis/api@v0.0.0-20240826202546-f6391c0de4c7                   | 994.458 KiB |
| github.com/fxamacker/cbor/v2@v2.7.0                                                            | 944.293 KiB |
| sigs.k8s.io/yaml@v1.4.0                                                                        | 890.639 KiB |
| github.com/golang/protobuf@v1.5.4                                                              | 797.319 KiB |
| go.etcd.io/etcd/raft/v3@v3.5.16                                                                | 780.661 KiB |
| gopkg.in/square/go-jose.v2@v2.6.0                                                              | 764.028 KiB |
| go.uber.org/zap@v1.27.0                                                                        | 723.969 KiB |
| go.opentelemetry.io/proto/otlp@v1.3.1                                                          | 702.530 KiB |
| github.com/container-storage-interface/spec@v1.9.0                                             | 675.715 KiB |
| github.com/libopenstorage/openstorage@v1.0.0                                                   | 647.814 KiB |
| github.com/spf13/cobra@v1.8.1                                                                  | 646.996 KiB |
| github.com/prometheus/common@v0.55.0                                                           | 604.237 KiB |
| github.com/stretchr/testify@v1.9.0                                                             | 603.282 KiB |
| k8s.io/utils@v0.0.0-20241104100929-3ea5e8cea738                                                | 547.083 KiB |
| go.etcd.io/bbolt@v1.3.11                                                                       | 541.105 KiB |
| github.com/json-iterator/go@v1.1.12                                                            | 472.025 KiB |
| golang.org/x/mod@v0.21.0                                                                       | 463.765 KiB |
| github.com/google/go-cmp@v0.6.0                                                                | 458.510 KiB |
| gopkg.in/yaml.v3@v3.0.1                                                                        | 452.741 KiB |
| github.com/antlr4-go/antlr/v4@v4.13.0                                                          | 450.822 KiB |
| sigs.k8s.io/json@v0.0.0-20241010143419-9aa6b5e7a4b3                                            | 449.049 KiB |
| github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0                                            | 440.131 KiB |
| golang.org/x/oauth2@v0.23.0                                                                    | 437.626 KiB |
| github.com/russross/blackfriday/v2@v2.1.0                                                      | 430.952 KiB |
| go.opentelemetry.io/otel/sdk@v1.28.0                                                           | 428.273 KiB |
| github.com/!microsoft/go-winio@v0.6.2                                                          | 417.321 KiB |
| k8s.io/klog/v2@v2.130.1                                                                        | 387.923 KiB |
| go.etcd.io/etcd/client/v3@v3.5.16                                                              | 384.013 KiB |
| sigs.k8s.io/kustomize/kustomize/v5@v5.5.0                                                      | 381.120 KiB |
| github.com/coreos/go-systemd/v22@v22.5.0                                                       | 361.166 KiB |
| github.com/opencontainers/runtime-spec@v1.2.0                                                  | 349.900 KiB |
| github.com/spf13/pflag@v1.0.5                                                                  | 303.436 KiB |
| github.com/godbus/dbus/v5@v5.1.0                                                               | 297.457 KiB |
| github.com/!microsoft/hnslib@v0.0.8                                                            | 268.178 KiB |
| github.com/asaskevich/govalidator@v0.0.0-20190424111038-f61b66f89f4a                           | 235.793 KiB |
| github.com/containerd/ttrpc@v1.2.5                                                             | 234.961 KiB |
| github.com/go-logr/logr@v1.4.2                                                                 | 234.140 KiB |
| github.com/emicklei/go-restful/v3@v3.11.0                                                      | 233.609 KiB |
| github.com/fsnotify/fsnotify@v1.7.0                                                            | 229.162 KiB |
| github.com/stretchr/objx@v0.5.2                                                                | 229.050 KiB |
| go.etcd.io/etcd/client/pkg/v3@v3.5.16                                                          | 224.911 KiB |
| github.com/mailru/easyjson@v0.7.7                                                              | 221.297 KiB |
| github.com/davecgh/go-spew@v1.1.2-0.20180830191138-d8f796af33cc                                | 207.580 KiB |
| k8s.io/gengo/v2@v2.0.0-20240911193312-2b36238f13e9                                             | 207.423 KiB |
| github.com/gorilla/websocket@v1.5.0                                                            | 204.744 KiB |
| sigs.k8s.io/knftables@v0.0.17                                                                  | 198.847 KiB |
| github.com/cyphar/filepath-securejoin@v0.3.4                                                   | 196.638 KiB |
| github.com/!n!y!times/gziphandler@v1.1.1                                                       | 192.861 KiB |
| github.com/go-openapi/swag@v0.23.0                                                             | 187.330 KiB |
| github.com/golang-jwt/jwt/v4@v4.5.0                                                            | 184.263 KiB |
| github.com/sirupsen/logrus@v1.9.3                                                              | 180.857 KiB |
| github.com/coredns/caddy@v1.1.1                                                                | 178.024 KiB |
| go.etcd.io/etcd/client/v2@v2.305.16                                                            | 166.831 KiB |
| google.golang.org/genproto/googleapis/rpc@v0.0.0-20240826202546-f6391c0de4c7                   | 155.799 KiB |
| github.com/go-task/slim-sprig/v3@v3.0.0                                                        | 146.100 KiB |
| github.com/moby/spdystream@v0.5.0                                                              | 143.319 KiB |
| go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp@v0.53.0                          | 141.097 KiB |
| go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc@v1.27.0                        | 133.636 KiB |
| github.com/opencontainers/selinux@v1.11.1                                                      | 126.890 KiB |
| go.opentelemetry.io/otel/metric@v1.28.0                                                        | 126.185 KiB |
| github.com/distribution/reference@v0.6.0                                                       | 125.625 KiB |
| k8s.io/system-validators@v1.9.1                                                                | 121.583 KiB |
| sigs.k8s.io/apiserver-network-proxy/konnectivity-client@v0.31.0                                | 113.054 KiB |
| github.com/moby/sys/mountinfo@v0.7.2                                                           | 112.433 KiB |
| go.opentelemetry.io/otel/trace@v1.28.0                                                         | 110.192 KiB |
| github.com/coredns/corefile-migration@v1.0.24                                                  | 107.208 KiB |
| github.com/karrick/godirwalk@v1.17.0                                                           | 97.028 KiB  |
| github.com/grpc-ecosystem/go-grpc-prometheus@v1.2.0                                            | 96.085 KiB  |
| go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc@v0.53.0            | 94.473 KiB  |
| github.com/robfig/cron/v3@v3.0.1                                                               | 90.957 KiB  |
| github.com/!azure/go-ansiterm@v0.0.0-20230124172434-306776ec8161                               | 89.945 KiB  |
| github.com/pquerna/cachecontrol@v0.1.0                                                         | 86.926 KiB  |
| github.com/coreos/go-oidc@v2.2.1+incompatible                                                  | 85.479 KiB  |
| github.com/opencontainers/go-digest@v1.0.0                                                     | 83.470 KiB  |
| github.com/felixge/httpsnoop@v1.0.4                                                            | 81.334 KiB  |
| go.uber.org/goleak@v1.3.0                                                                      | 80.493 KiB  |
| github.com/google/uuid@v1.6.0                                                                  | 77.172 KiB  |
| github.com/prometheus/client_model@v0.6.1                                                      | 76.667 KiB  |
| github.com/go-logr/zapr@v1.3.0                                                                 | 73.698 KiB  |
| gopkg.in/evanphx/json-patch.v4@v4.12.0                                                         | 73.525 KiB  |
| github.com/dustin/go-humanize@v1.0.1                                                           | 65.602 KiB  |
| github.com/gregjones/httpcache@v0.0.0-20190611155906-901d90724c79                              | 64.543 KiB  |
| go.uber.org/multierr@v1.11.0                                                                   | 61.441 KiB  |
| github.com/ishidawataru/sctp@v0.0.0-20230406120618-7ff4192f6ff2                                | 61.365 KiB  |
| github.com/moby/ipvs@v1.1.0                                                                    | 59.628 KiB  |
| golang.org/x/sync@v0.8.0                                                                       | 59.618 KiB  |
| github.com/go-openapi/jsonpointer@v0.21.0                                                      | 58.752 KiB  |
| github.com/google/gofuzz@v1.2.0                                                                | 58.650 KiB  |
| github.com/modern-go/reflect2@v1.0.2                                                           | 58.352 KiB  |
| github.com/google/btree@v1.0.1                                                                 | 58.021 KiB  |
| github.com/soheilhy/cmux@v0.1.5                                                                | 57.909 KiB  |
| github.com/x448/float16@v0.8.4                                                                 | 57.781 KiB  |
| github.com/blang/semver/v4@v4.0.0                                                              | 56.989 KiB  |
| github.com/peterbourgon/diskv@v2.0.1+incompatible                                              | 54.568 KiB  |
| github.com/pkg/errors@v0.9.1                                                                   | 53.510 KiB  |
| github.com/moby/term@v0.5.0                                                                    | 51.776 KiB  |
| github.com/mistifyio/go-zfs@v2.1.2-0.20190413222219-f784269be439+incompatible                  | 50.753 KiB  |
| go.opentelemetry.io/otel/exporters/otlp/otlptrace@v1.28.0                                      | 49.625 KiB  |
| gopkg.in/natefinch/lumberjack.v2@v2.2.1                                                        | 48.854 KiB  |
| gopkg.in/inf.v0@v0.9.1                                                                         | 46.804 KiB  |
| golang.org/x/term@v0.25.0                                                                      | 46.501 KiB  |
| github.com/go-openapi/jsonreference@v0.20.2                                                    | 41.279 KiB  |
| github.com/liggitt/tabwriter@v0.0.0-20181228230101-89fcab3d43de                                | 41.109 KiB  |
| github.com/docker/go-units@v0.5.0                                                              | 40.845 KiB  |
| github.com/jonboulle/clockwork@v0.4.0                                                          | 39.545 KiB  |
| github.com/tmc/grpc-websocket-proxy@v0.0.0-20220101234140-673ab2c3ae75                         | 39.441 KiB  |
| bitbucket.org/bertimus9/systemstat@v0.5.0                                                      | 38.871 KiB  |
| github.com/mohae/deepcopy@v0.0.0-20170929034955-c48cc78d4826                                   | 38.366 KiB  |
| golang.org/x/time@v0.7.0                                                                       | 38.354 KiB  |
| github.com/exponent-io/jsonpath@v0.0.0-20210407135951-1de76d718b3f                             | 36.229 KiB  |
| github.com/cenkalti/backoff/v4@v4.3.0                                                          | 35.556 KiB  |
| github.com/pmezard/go-difflib@v1.0.1-0.20181226105442-5d4384ee4fb2                             | 35.223 KiB  |
| github.com/coreos/go-semver@v0.3.1                                                             | 33.729 KiB  |
| github.com/cespare/xxhash/v2@v2.3.0                                                            | 32.325 KiB  |
| go.opentelemetry.io/contrib/instrumentation/github.com/emicklei/go-restful/otelrestful@v0.42.0 | 31.883 KiB  |
| github.com/xlab/treeprint@v1.2.0                                                               | 31.764 KiB  |
| github.com/beorn7/perks@v1.0.1                                                                 | 30.504 KiB  |
| github.com/go-errors/errors@v1.4.2                                                             | 30.191 KiB  |
| github.com/cpuguy83/go-md2man/v2@v2.0.4                                                        | 30.114 KiB  |
| github.com/armon/go-socks5@v0.0.0-20160902184237-e75332964ef5                                  | 30.016 KiB  |
| github.com/vishvananda/netns@v0.0.4                                                            | 29.024 KiB  |
| github.com/containerd/log@v0.1.0                                                               | 26.703 KiB  |
| github.com/go-logr/stdr@v1.2.2                                                                 | 25.719 KiB  |
| github.com/containerd/errdefs@v0.1.0                                                           | 25.085 KiB  |
| github.com/euank/go-kmsg-parser@v2.0.0+incompatible                                            | 24.321 KiB  |
| github.com/google/shlex@v0.0.0-20191202100458-e7afc7fbc510                                     | 23.898 KiB  |
| github.com/!jeff!ashton/win_pdh@v0.0.0-20161109143554-76bb4ee9f0ab                             | 21.967 KiB  |
| github.com/modern-go/concurrent@v0.0.0-20180306012644-bacd9c7ef1dd                             | 20.406 KiB  |
| github.com/mxk/go-flowrate@v0.0.0-20140419014527-cca7078d478f                                  | 19.632 KiB  |
| github.com/mrunalp/fileutils@v0.5.1                                                            | 18.121 KiB  |
| github.com/stoewer/go-strcase@v1.3.0                                                           | 15.921 KiB  |
| github.com/moby/sys/userns@v0.1.0                                                              | 14.355 KiB  |
| github.com/monochromegane/go-gitignore@v0.0.0-20200626010858-205db1a8cc00                      | 14.154 KiB  |
| github.com/inconshreveable/mousetrap@v1.1.0                                                    | 13.724 KiB  |
| Others                                                                                         | 11.797 KiB  |
| github.com/munnerz/goautoneg@v0.0.0-20191010083416-a7dc8b61c822                                | 9.917 KiB   |
| github.com/xiang90/probing@v0.0.0-20221125231312-a49e3df8f510                                  | 9.057 KiB   |
| github.com/!make!now!just/heredoc@v1.0.0                                                       | 8.491 KiB   |
| github.com/lithammer/dedent@v1.1.0                                                             | 8.158 KiB   |
| github.com/armon/circbuf@v0.0.0-20190214190532-5111143e8da2                                    | 7.408 KiB   |
| github.com/fatih/camelcase@v1.0.0                                                              | 7.325 KiB   |
| github.com/mitchellh/go-wordwrap@v1.0.1                                                        | 6.004 KiB   |
| github.com/josharian/intern@v1.0.0                                                             | 3.295 KiB   |
Size: 307.532 MiB
21518 files, 5538 directories
```

</details>

<details><summary>on golangci-lint</summary>

```console
$ gomodcachesize           
|                                   MODULE                                   |    SIZE     |
|----------------------------------------------------------------------------|-------------|
| golang.org/x/text@v0.24.0                                                  | 39.208 MiB  |
| cache                                                                      | 36.636 MiB  |
| google.golang.org/protobuf@v1.36.6                                         | 12.995 MiB  |
| golang.org/x/sys@v0.32.0                                                   | 8.980 MiB   |
| golang.org/x/tools@v0.32.0                                                 | 8.205 MiB   |
| github.com/alecthomas/chroma/v2@v2.16.0                                    | 6.535 MiB   |
| github.com/rivo/uniseg@v0.4.7                                              | 4.284 MiB   |
| github.com/ccojocar/zxcvbn-go@v1.0.2                                       | 2.265 MiB   |
| github.com/lucasb-eyer/go-colorful@v1.2.0                                  | 2.244 MiB   |
| github.com/yeya24/promlinter@v0.3.0                                        | 2.152 MiB   |
| honnef.co/go/tools@v0.6.1                                                  | 2.010 MiB   |
| golang.org/x/exp@v0.0.0-20240909161429-701f63a606c0                        | 1.468 MiB   |
| github.com/pelletier/go-toml/v2@v2.2.4                                     | 1.335 MiB   |
| github.com/tommy-muehle/go-mnd/v2@v2.5.1                                   | 1.261 MiB   |
| github.com/mgechev/revive@v1.9.0                                           | 1.239 MiB   |
| github.com/prometheus/procfs@v0.7.3                                        | 997.859 KiB |
| github.com/golangci/misspell@v0.6.0                                        | 983.675 KiB |
| github.com/shirou/gopsutil/v4@v4.25.3                                      | 839.261 KiB |
| github.com/prometheus/client_golang@v1.12.1                                | 796.905 KiB |
| github.com/golang/protobuf@v1.5.3                                          | 790.116 KiB |
| github.com/securego/gosec/v2@v2.22.3                                       | 763.407 KiB |
| github.com/quasilyte/go-ruleguard@v0.4.4                                   | 726.593 KiB |
| github.com/go-critic/go-critic@v0.13.0                                     | 711.302 KiB |
| go.uber.org/zap@v1.24.0                                                    | 676.543 KiB |
| github.com/spf13/cobra@v1.9.1                                              | 674.261 KiB |
| github.com/stretchr/testify@v1.10.0                                        | 628.087 KiB |
| github.com/dave/dst@v0.27.3                                                | 582.427 KiB |
| github.com/dlclark/regexp2@v1.11.5                                         | 557.831 KiB |
| github.com/rogpeppe/go-internal@v1.14.1                                    | 507.465 KiB |
| github.com/sashamelentyev/usestdlibvars@v1.28.0                            | 501.656 KiB |
| github.com/muesli/termenv@v0.16.0                                          | 492.802 KiB |
| github.com/prometheus/common@v0.32.1                                       | 487.959 KiB |
| golang.org/x/mod@v0.24.0                                                   | 471.154 KiB |
| github.com/pelletier/go-toml@v1.9.5                                        | 461.961 KiB |
| github.com/google/go-cmp@v0.7.0                                            | 460.904 KiB |
| gopkg.in/yaml.v3@v3.0.1                                                    | 452.741 KiB |
| github.com/sourcegraph/go-diff@v0.7.0                                      | 450.434 KiB |
| github.com/!antonboom/testifylint@v1.6.1                                   | 446.162 KiB |
| github.com/catenacyber/perfsprint@v0.9.1                                   | 434.830 KiB |
| github.com/charmbracelet/lipgloss@v1.1.0                                   | 390.291 KiB |
| github.com/!burnt!sushi/toml@v1.5.0                                        | 376.265 KiB |
| github.com/ebitengine/purego@v0.8.2                                        | 361.604 KiB |
| github.com/spf13/viper@v1.12.0                                             | 346.017 KiB |
| github.com/charmbracelet/x/ansi@v0.8.0                                     | 345.666 KiB |
| gopkg.in/yaml.v2@v2.4.0                                                    | 335.161 KiB |
| mvdan.cc/gofumpt@v0.8.0                                                    | 324.043 KiB |
| github.com/spf13/pflag@v1.0.6                                              | 317.407 KiB |
| github.com/tklauser/go-sysconf@v0.3.12                                     | 294.685 KiB |
| github.com/spf13/afero@v1.14.0                                             | 281.533 KiB |
| github.com/quasilyte/gogrep@v0.5.0                                         | 255.348 KiB |
| gopkg.in/ini.v1@v1.67.0                                                    | 244.938 KiB |
| github.com/hashicorp/hcl@v1.0.0                                            | 240.469 KiB |
| github.com/santhosh-tekuri/jsonschema/v6@v6.0.1                            | 231.602 KiB |
| github.com/valyala/quicktemplate@v1.8.0                                    | 230.010 KiB |
| github.com/stretchr/objx@v0.5.2                                            | 229.050 KiB |
| github.com/go-viper/mapstructure/v2@v2.2.1                                 | 220.085 KiB |
| github.com/davecgh/go-spew@v1.1.1                                          | 207.567 KiB |
| github.com/ldez/usetesting@v0.4.3                                          | 202.688 KiB |
| github.com/tomarrell/wrapcheck/v2@v2.11.0                                  | 196.396 KiB |
| github.com/go-ole/go-ole@v1.2.6                                            | 189.871 KiB |
| github.com/matttproud/golang_protobuf_extensions@v1.0.1                    | 184.011 KiB |
| github.com/sirupsen/logrus@v1.9.3                                          | 180.857 KiB |
| github.com/nunnatsa/ginkgolinter@v0.19.1                                   | 180.822 KiB |
| github.com/ldez/grignotin@v0.9.0                                           | 167.636 KiB |
| github.com/bombsimon/wsl/v4@v4.7.0                                         | 161.735 KiB |
| github.com/nishanths/exhaustive@v0.12.0                                    | 157.020 KiB |
| github.com/power-devops/perfstat@v0.0.0-20210106213030-5aafc221ea8c        | 154.851 KiB |
| github.com/mitchellh/mapstructure@v1.5.0                                   | 148.820 KiB |
| github.com/jgautheron/goconst@v1.8.1                                       | 143.433 KiB |
| github.com/butuzov/mirror@v1.3.0                                           | 142.098 KiB |
| github.com/golangci/gofmt@v0.0.0-20250106114630-d62b90e6713d               | 135.386 KiB |
| github.com/fsnotify/fsnotify@v1.5.4                                        | 129.938 KiB |
| github.com/magiconair/properties@v1.8.6                                    | 125.456 KiB |
| github.com/xo/terminfo@v0.0.0-20220910002029-abceb7e1c41e                  | 124.971 KiB |
| github.com/charmbracelet/x/cellbuf@v0.0.13-0.20250311204145-2c3ea96c31dd   | 124.643 KiB |
| github.com/gobwas/glob@v0.2.3                                              | 121.548 KiB |
| golang.org/x/exp/typeparams@v0.0.0-20250210185358-939b2ce775ac             | 117.583 KiB |
| github.com/!masterminds/semver/v3@v3.3.1                                   | 110.953 KiB |
| github.com/hashicorp/go-immutable-radix/v2@v2.1.0                          | 110.095 KiB |
| github.com/daixiang0/gci@v0.13.6                                           | 108.219 KiB |
| go.uber.org/atomic@v1.7.0                                                  | 101.853 KiB |
| github.com/polyfloyd/go-errorlint@v1.8.0                                   | 95.669 KiB  |
| go.uber.org/automaxprocs@v1.6.0                                            | 93.625 KiB  |
| github.com/!open!pee!dee!p/depguard/v2@v2.2.1                              | 92.729 KiB  |
| github.com/olekukonko/tablewriter@v0.0.5                                   | 90.927 KiB  |
| github.com/chavacava/garif@v0.1.0                                          | 86.147 KiB  |
| github.com/hashicorp/golang-lru/v2@v2.0.7                                  | 85.333 KiB  |
| github.com/kulti/thelper@v0.6.3                                            | 83.234 KiB  |
| github.com/ettle/strcase@v0.2.0                                            | 82.972 KiB  |
| github.com/breml/errchkjson@v0.4.1                                         | 80.800 KiB  |
| github.com/spf13/cast@v1.5.0                                               | 80.358 KiB  |
| github.com/hexops/gotextdiff@v1.0.3                                        | 79.958 KiB  |
| github.com/golangci/dupl@v0.0.0-20250308024227-f665c8d69b32                | 78.193 KiB  |
| github.com/kisielk/errcheck@v1.9.0                                         | 74.599 KiB  |
| github.com/tetafro/godot@v1.5.0                                            | 73.854 KiB  |
| go-simpler.org/sloglint@v0.11.0                                            | 72.893 KiB  |
| github.com/denis-tingaikin/go-header@v0.5.0                                | 68.310 KiB  |
| github.com/leonklingele/grouper@v1.1.2                                     | 67.748 KiB  |
| github.com/hashicorp/go-version@v1.7.0                                     | 63.962 KiB  |
| github.com/xen0n/gosmopolitan@v1.3.0                                       | 63.950 KiB  |
| github.com/lasiar/canonicalheader@v1.1.2                                   | 63.933 KiB  |
| github.com/quasilyte/regex/syntax@v0.0.0-20210819130434-b3f0c404a727       | 63.021 KiB  |
| github.com/golangci/golines@v0.0.0-20250217134842-442fd0091d95             | 62.189 KiB  |
| github.com/ryancurrah/gomodguard@v1.4.1                                    | 61.983 KiB  |
| github.com/firefart/nonamedreturns@v1.0.6                                  | 60.212 KiB  |
| github.com/uudashr/iface@v1.3.1                                            | 59.928 KiB  |
| mvdan.cc/unparam@v0.0.0-20250301125049-0df0534333a4                        | 59.714 KiB  |
| github.com/mattn/go-runewidth@v0.0.16                                      | 59.319 KiB  |
| golang.org/x/sync@v0.13.0                                                  | 58.636 KiB  |
| github.com/ashanbrown/forbidigo@v1.6.0                                     | 56.784 KiB  |
| github.com/blizzy78/varnamelen@v0.8.0                                      | 54.577 KiB  |
| github.com/!antonboom/errname@v1.1.0                                       | 53.622 KiB  |
| github.com/gofrs/flock@v0.12.1                                             | 51.610 KiB  |
| github.com/timonwong/loggercheck@v0.11.0                                   | 51.099 KiB  |
| github.com/alingse/nilnesserr@v0.2.0                                       | 51.081 KiB  |
| github.com/go-toolsmith/astcast@v1.1.0                                     | 50.982 KiB  |
| github.com/butuzov/ireturn@v0.4.0                                          | 50.372 KiB  |
| github.com/yusufpapurcu/wmi@v1.2.4                                         | 48.692 KiB  |
| github.com/4meepo/tagalign@v1.4.2                                          | 47.728 KiB  |
| github.com/fatih/color@v1.18.0                                             | 47.419 KiB  |
| github.com/!antonboom/nilnil@v1.1.0                                        | 46.888 KiB  |
| github.com/gostaticanalysis/analysisutil@v0.7.1                            | 46.732 KiB  |
| github.com/jjti/go-spancheck@v0.6.4                                        | 46.346 KiB  |
| github.com/uudashr/gocognit@v1.2.0                                         | 46.189 KiB  |
| github.com/!abirdcfly/dupword@v0.1.3                                       | 45.778 KiB  |
| github.com/charmbracelet/colorprofile@v0.2.3-0.20250311203215-f60798e515dc | 45.114 KiB  |
| github.com/kk!h!a!i!k!e/contextcheck@v1.1.6                                | 44.522 KiB  |
| github.com/manuelarte/funcorder@v0.2.1                                     | 44.267 KiB  |
| github.com/ldez/tagliatelle@v0.7.1                                         | 43.264 KiB  |
| github.com/golangci/revgrep@v0.8.0                                         | 42.904 KiB  |
| github.com/!gaijin!entertainment/go-exhaustruct/v3@v3.3.1                  | 42.547 KiB  |
| github.com/ldez/gomoddirectives@v0.6.1                                     | 42.161 KiB  |
| github.com/gordonklaus/ineffassign@v0.1.0                                  | 41.911 KiB  |
| github.com/prometheus/client_model@v0.2.0                                  | 41.142 KiB  |
| github.com/kunwardeep/paralleltest@v1.0.14                                 | 40.666 KiB  |
| github.com/alingse/asasalint@v0.0.11                                       | 40.215 KiB  |
| github.com/alecthomas/go-check-sumtype@v0.3.1                              | 39.922 KiB  |
| go.uber.org/multierr@v1.6.0                                                | 39.860 KiB  |
| github.com/alexkohler/nakedret/v2@v2.0.6                                   | 39.447 KiB  |
| go-simpler.org/musttag@v0.13.0                                             | 38.799 KiB  |
| github.com/golangci/plugin-module-register@v0.1.1                          | 36.288 KiB  |
| github.com/pmezard/go-difflib@v1.0.0                                       | 35.225 KiB  |
| github.com/mattn/go-colorable@v0.1.14                                      | 35.007 KiB  |
| github.com/nishanths/predeclared@v0.2.2                                    | 34.807 KiB  |
| github.com/tklauser/numcpus@v0.6.1                                         | 33.589 KiB  |
| github.com/go-toolsmith/astequal@v1.2.0                                    | 33.240 KiB  |
| github.com/ldez/exptostd@v0.4.3                                            | 33.113 KiB  |
| github.com/timakin/bodyclose@v0.0.0-20241222091800-1db5c5ca4d67            | 32.525 KiB  |
| github.com/cespare/xxhash/v2@v2.3.0                                        | 32.325 KiB  |
| github.com/alexkohler/prealloc@v1.0.0                                      | 32.264 KiB  |
| github.com/subosito/gotenv@v1.4.1                                          | 31.624 KiB  |
| github.com/maratori/testpackage@v1.1.1                                     | 31.547 KiB  |
| github.com/beorn7/perks@v1.0.1                                             | 30.504 KiB  |
| github.com/nakabonne/nestif@v0.3.1                                         | 30.493 KiB  |
| github.com/lufia/plan9stats@v0.0.0-20211012122336-39d0f177ccd0             | 29.761 KiB  |
| github.com/golangci/unconvert@v0.0.0-20250410112200-a129a6e6413e           | 29.751 KiB  |
| github.com/julz/importas@v0.2.0                                            | 29.288 KiB  |
| go.augendre.info/fatcontext@v0.8.0                                         | 28.710 KiB  |
| gitlab.com/bosi/decorder@v0.4.2                                            | 27.979 KiB  |
| github.com/ryanrolds/sqlclosecheck@v0.5.1                                  | 27.789 KiB  |
| github.com/ckaznocha/intrange@v0.3.1                                       | 27.337 KiB  |
| github.com/fzipp/gocyclo@v0.6.0                                            | 26.342 KiB  |
| github.com/maratori/testableexamples@v1.0.0                                | 26.221 KiB  |
| github.com/moricho/tparallel@v0.3.2                                        | 25.468 KiB  |
| github.com/ghostiam/protogetter@v0.3.15                                    | 25.439 KiB  |
| github.com/charithe/durationcheck@v0.0.10                                  | 24.962 KiB  |
| github.com/gostaticanalysis/comment@v1.5.0                                 | 24.419 KiB  |
| github.com/go-toolsmith/astcopy@v1.1.0                                     | 23.905 KiB  |
| 4d63.com/gochecknoglobals@v0.2.2                                           | 23.570 KiB  |
| github.com/sonatard/noctx@v0.1.0                                           | 23.377 KiB  |
| github.com/quasilyte/stdinfo@v0.0.0-20220114132959-f7386bf02567            | 22.624 KiB  |
| github.com/!djarvur/go-err113@v0.0.0-20210108212216-aea10b59be24           | 22.388 KiB  |
| github.com/go-xmlfmt/xmlfmt@v1.1.3                                         | 21.261 KiB  |
| github.com/quasilyte/go-ruleguard/dsl@v0.3.22                              | 20.878 KiB  |
| github.com/ashanbrown/makezero@v1.2.0                                      | 20.184 KiB  |
| github.com/spf13/jwalterweatherman@v1.1.0                                  | 19.904 KiB  |
| github.com/fatih/structtag@v1.2.0                                          | 19.863 KiB  |
| github.com/matoous/godox@v1.1.0                                            | 19.440 KiB  |
| github.com/go-toolsmith/typep@v1.1.0                                       | 19.074 KiB  |
| github.com/breml/bidichk@v0.3.3                                            | 18.989 KiB  |
| github.com/ykadowak/zerologlint@v0.1.5                                     | 18.852 KiB  |
| github.com/jingyugao/rowserrcheck@v1.1.1                                   | 18.752 KiB  |
| github.com/ultraware/whitespace@v0.2.0                                     | 18.726 KiB  |
| github.com/aymanbagabas/go-osc52/v2@v2.0.1                                 | 18.673 KiB  |
| github.com/curioswitch/go-reassign@v0.3.0                                  | 16.103 KiB  |
| github.com/sanposhiho/wastedassign/v2@v2.1.0                               | 16.047 KiB  |
| github.com/gostaticanalysis/nilerr@v0.1.1                                  | 15.753 KiB  |
| github.com/karamaru-alpha/copyloopvar@v1.2.1                               | 15.381 KiB  |
| github.com/yagipy/maintidx@v1.0.0                                          | 15.365 KiB  |
| github.com/gostaticanalysis/forcetypeassert@v0.2.0                         | 15.185 KiB  |
| github.com/raeperd/recvcheck@v0.2.0                                        | 14.531 KiB  |
| github.com/inconshreveable/mousetrap@v1.1.0                                | 13.724 KiB  |
| Others                                                                     | 13.555 KiB  |
| github.com/sivchari/containedctx@v1.0.3                                    | 13.552 KiB  |
| github.com/valyala/bytebufferpool@v1.0.0                                   | 13.525 KiB  |
| github.com/bkielbasa/cyclop@v1.2.3                                         | 13.372 KiB  |
| github.com/go-toolsmith/astp@v1.1.0                                        | 13.171 KiB  |
| github.com/tdakkota/asciicheck@v0.4.1                                      | 12.327 KiB  |
| github.com/mattn/go-isatty@v0.0.20                                         | 12.207 KiB  |
| github.com/ssgreg/nlreturn/v2@v2.2.1                                       | 11.194 KiB  |
| github.com/charmbracelet/x/term@v0.2.1                                     | 11.189 KiB  |
| 4d63.com/gocheckcompilerdirectives@v1.3.0                                  | 10.321 KiB  |
| github.com/ultraware/funlen@v0.2.0                                         | 10.278 KiB  |
| github.com/golangci/go-printf-func-name@v0.1.0                             | 10.226 KiB  |
| github.com/stbenjam/no-sprintf-host-port@v0.2.0                            | 10.044 KiB  |
| github.com/go-toolsmith/astfmt@v1.1.0                                      | 8.025 KiB   |
| github.com/macabu/inamedparam@v0.2.0                                       | 7.806 KiB   |
| github.com/mitchellh/go-homedir@v1.1.0                                     | 7.473 KiB   |
| github.com/sashamelentyev/interfacebloat@v1.1.0                            | 7.033 KiB   |
| github.com/go-toolsmith/strparse@v1.1.0                                    | 6.881 KiB   |
Size: 158.764 MiB
18417 files, 5310 directories
```

</details>
