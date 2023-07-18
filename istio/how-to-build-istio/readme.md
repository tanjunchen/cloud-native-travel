# 如何编译 Istio

## 前提条件

准备工具：
```
docker
git
make
```

本地能够拉取到编译镜像，如本次实验使用的编译镜像工具是：
`gcr.io/istio-testing/build-tools:master-7c2cba5679671f40312e6324d270a0d70ad097d0`

go 环境：
```
➜  istio git:(master) go version
go version go1.20.3 darwin/arm64
```

## 实验

1. 克隆 istio 存储库到 $GOPATH/src/istio.io/
```
git clone https://github.com/istio/istio.git
```
2. 设置环境变量
```
当前目录：~/opensource/istio
export USER="tanjunchen"
export HUB="docker.io/$USER"
export TAG="06-25-dev"
```
3. 编译二进制
istio 源根目录，编译与构建 istio 所有组件，如 pilot、istioctl 等。执行如下命名：
```
make build
```
其中构建镜像的日志如下所示：
```
➜  istio git:(master) make build
TARGET_OUT=/work/out/darwin_arm64 ISTIO_BIN=/Users/chentanjun/software/go/go1.20.3/bin GOOS_LOCAL=darwin bin/retry.sh SSL_ERROR_SYSCALL bin/init.sh
Skipping envoy debug. Set DEBUG_IMAGE to download.
/work/out/linux_arm64/release /work
Downloading envoy: https://storage.googleapis.com/istio-build/proxy/envoy-alpha-d654519fc7b0023bc32efd0a1f2c1770f4c96a9a-arm64.tar.gz to /work/out/linux_arm64/release/envoy-d654519fc7b0023bc32efd0a1f2c1770f4c96a9a

real	0m11.994s
user	0m0.985s
sys	0m1.131s
Copying /work/out/linux_arm64/release/envoy-d654519fc7b0023bc32efd0a1f2c1770f4c96a9a to /work/out/linux_arm64/release/envoy
/work
/work/out/linux_arm64/release /work
Downloading envoy: https://storage.googleapis.com/istio-build/proxy/envoy-centos-alpha-d654519fc7b0023bc32efd0a1f2c1770f4c96a9a.tar.gz to /work/out/linux_arm64/release/envoy-centos-d654519fc7b0023bc32efd0a1f2c1770f4c96a9a

real	0m12.261s
user	0m1.028s
sys	0m0.947s
Copying /work/out/linux_arm64/release/envoy-centos-d654519fc7b0023bc32efd0a1f2c1770f4c96a9a to /work/out/linux_arm64/release/envoy-centos
/work
Copying /work/out/linux_arm64/release/envoy-d654519fc7b0023bc32efd0a1f2c1770f4c96a9a to /work/out/darwin_arm64/envoy
Copying /work/out/linux_arm64/release/envoy-centos-d654519fc7b0023bc32efd0a1f2c1770f4c96a9a to /work/out/linux_arm64/envoy-centos
Copying /work/out/linux_arm64/release/envoy-d654519fc7b0023bc32efd0a1f2c1770f4c96a9a to /work/out/linux_arm64/envoy
touch /work/out/darwin_arm64/istio_is_init
TARGET_OUT=/work/out/darwin_arm64 bin/build_ztunnel.sh
/work/out/linux_arm64/release /work
Downloading ztunnel: https://storage.googleapis.com/istio-build/ztunnel/ztunnel-96de2673c928b55323f1a15ac4f492d030aa7509-arm64 to /work/out/linux_arm64/release/ztunnel-96de2673c928b55323f1a15ac4f492d030aa7509

real	0m9.417s
user	0m0.150s
sys	0m0.211s
Copying /work/out/linux_arm64/release/ztunnel-96de2673c928b55323f1a15ac4f492d030aa7509 to /work/out/linux_arm64/release/ztunnel
/work
Copying '/work/out/linux_arm64/release/ztunnel-96de2673c928b55323f1a15ac4f492d030aa7509' to /work/out/linux_arm64/ztunnel
GOOS=darwin GOARCH=arm64 LDFLAGS='-extldflags -static -s -w' common/scripts/gobuild.sh /work/out/darwin_arm64/ ./istioctl/cmd/istioctl ./pilot/cmd/pilot-discovery ./pkg/test/echo/cmd/client ./pkg/test/echo/cmd/server ./samples/extauthz/cmd/extauthz ./operator/cmd/operator ./tools/bug-report

real	1m46.614s
user	1m31.319s
sys	0m37.436s
GOOS=darwin GOARCH=arm64 LDFLAGS='-extldflags -static -s -w' common/scripts/gobuild.sh /work/out/darwin_arm64/ -tags=agent ./pilot/cmd/pilot-agent

real	0m7.315s
user	0m3.482s
sys	0m2.066s

➜  istio git:(master) tree out -L 2
out
├── darwin_arm64
│   ├── bug-report
│   ├── client
│   ├── envoy
│   ├── extauthz
│   ├── istio_is_init
│   ├── istioctl
│   ├── logs
│   ├── operator
│   ├── pilot-agent
│   ├── pilot-discovery
│   ├── release
│   └── server
└── linux_arm64
    ├── envoy
    ├── envoy-centos
    ├── logs
    ├── release
    └── ztunnel
```
4. 构建与推送镜像
```
BUILD_WITH_CONTAINER=0 make docker.push
```
其中构建镜像的日志如下所示：
```
➜  istio git:(master) BUILD_WITH_CONTAINER=0 make docker.push
./tools/docker --push

real	0m4.561s
user	0m3.121s
sys	0m0.862s
2023-06-25T03:16:56.780655Z	info	Args: Push:              true
Save:              false
NoClobber:         false
NoCache:           false
Targets:           [app app_sidecar_centos_7 app_sidecar_debian_11 app_sidecar_ubuntu_jammy app_sidecar_ubuntu_xenial ext-authz install-cni istioctl operator pilot proxyv2 ztunnel]
Variants:          [default]
Architectures:     [linux/amd64]
BaseVersion:       master-2023-06-15T19-01-36
BaseImageRegistry: gcr.io/istio-release
ProxyVersion:      d654519fc7b0023bc32efd0a1f2c1770f4c96a9a
IstioVersion:      1.19-dev
Tags:              [06-25-dev]
Hubs:              [docker.io/tanjunchen]
Builder:           docker

2023-06-25T03:16:56.781136Z	info	Skipping app_sidecar_centos_7 for linux/amd64 as --qemu is not passed
2023-06-25T03:16:56.781143Z	info	Skipping app_sidecar_debian_11 for linux/amd64 as --qemu is not passed
2023-06-25T03:16:56.781145Z	info	Skipping app_sidecar_ubuntu_jammy for linux/amd64 as --qemu is not passed
2023-06-25T03:16:56.781146Z	info	Skipping app_sidecar_ubuntu_xenial for linux/amd64 as --qemu is not passed
2023-06-25T03:16:56.781149Z	info	building for architectures: [linux/amd64]
2023-06-25T03:16:56.781492Z	info	Running make for linux/amd64: client extauthz install-cni istio-cni istioctl operator pilot-agent pilot-discovery server init
2023-06-25T03:16:56.781565Z	info	env: [TARGET_OS=linux TARGET_ARCH=amd64 TARGET_OUT=/Users/chentanjun/opensource/istio/out/linux_amd64 TARGET_OUT_LINUX=/Users/chentanjun/opensource/istio/out/linux_amd64]
TARGET_OUT=/Users/chentanjun/opensource/istio/out/linux_amd64 bin/build_ztunnel.sh
GOOS=linux GOARCH=amd64 LDFLAGS='-extldflags -static -s -w' common/scripts/gobuild.sh /Users/chentanjun/opensource/istio/out/linux_amd64/ ./istioctl/cmd/istioctl ./pilot/cmd/pilot-discovery ./pkg/test/echo/cmd/client ./pkg/test/echo/cmd/server ./samples/extauthz/cmd/extauthz ./operator/cmd/operator ./tools/bug-report

real	0m6.457s
user	0m18.731s
sys	0m3.956s
GOOS=linux GOARCH=amd64 LDFLAGS='-extldflags -static -s -w' common/scripts/gobuild.sh /Users/chentanjun/opensource/istio/out/linux_amd64/ -tags=agent ./pilot/cmd/pilot-agent

real	0m5.889s
user	0m15.073s
sys	0m2.122s
GOOS=linux GOARCH=amd64 LDFLAGS='-extldflags -static -s -w' common/scripts/gobuild.sh /Users/chentanjun/opensource/istio/out/linux_amd64/ -tags=agent ./cni/cmd/istio-cni ./cni/cmd/install-cni

real	0m3.474s
user	0m5.507s
sys	0m1.353s
make[1]: Nothing to be done for `init'.
2023-06-25T03:17:13.344571Z	info	Running command: tools/docker-copy.sh pkg/test/echo/docker/Dockerfile.app tests/testdata/certs /Users/chentanjun/opensource/istio/out/linux_amd64/client /Users/chentanjun/opensource/istio/out/linux_amd64/server /Users/chentanjun/opensource/istio/out/darwin_arm64/dockerx_build/build.docker.app
2023-06-25T03:17:13.475118Z	info	Running command: tools/docker-copy.sh samples/extauthz/docker/Dockerfile /Users/chentanjun/opensource/istio/out/linux_amd64/extauthz /Users/chentanjun/opensource/istio/out/darwin_arm64/dockerx_build/build.docker.ext-authz
2023-06-25T03:17:13.530677Z	info	Running command: tools/docker-copy.sh cni/deployments/kubernetes/Dockerfile.install-cni /Users/chentanjun/opensource/istio/out/linux_amd64/istio-cni /Users/chentanjun/opensource/istio/out/linux_amd64/install-cni /Users/chentanjun/opensource/istio/out/darwin_arm64/dockerx_build/build.docker.install-cni
2023-06-25T03:17:13.618718Z	info	Running command: tools/docker-copy.sh istioctl/docker/Dockerfile.istioctl /Users/chentanjun/opensource/istio/out/linux_amd64/istioctl /Users/chentanjun/opensource/istio/out/darwin_arm64/dockerx_build/build.docker.istioctl
2023-06-25T03:17:13.679772Z	info	Running command: tools/docker-copy.sh operator/docker/Dockerfile.operator manifests /Users/chentanjun/opensource/istio/out/linux_amd64/operator /Users/chentanjun/opensource/istio/out/darwin_arm64/dockerx_build/build.docker.operator
2023-06-25T03:17:13.822693Z	info	Running command: tools/docker-copy.sh pilot/docker/Dockerfile.pilot tools/packaging/common/envoy_bootstrap.json tools/packaging/common/gcp_envoy_bootstrap.json /Users/chentanjun/opensource/istio/out/linux_amd64/pilot-discovery /Users/chentanjun/opensource/istio/out/darwin_arm64/dockerx_build/build.docker.pilot
2023-06-25T03:17:13.949963Z	info	Running command: tools/docker-copy.sh pilot/docker/Dockerfile.proxyv2 tools/packaging/common/envoy_bootstrap.json tools/packaging/common/gcp_envoy_bootstrap.json /Users/chentanjun/opensource/istio/out/linux_amd64/release/envoy /Users/chentanjun/opensource/istio/out/linux_amd64/pilot-agent /Users/chentanjun/opensource/istio/out/darwin_arm64/dockerx_build/build.docker.proxyv2
2023-06-25T03:17:14.157398Z	info	Running command: tools/docker-copy.sh pilot/docker/Dockerfile.ztunnel /Users/chentanjun/opensource/istio/out/linux_amd64/ztunnel /Users/chentanjun/opensource/istio/out/darwin_arm64/dockerx_build/build.docker.ztunnel
2023-06-25T03:17:14.227915Z	info	make complete	runtime=17.446529834s
2023-06-25T03:17:14.228093Z	info	Running command: docker buildx bake -f /Users/chentanjun/opensource/istio/out/darwin_arm64/dockerx_build/docker-bake.json all
[+] Building 114.3s (55/56)
[+] Building 114.4s (55/56)
[+] Building 114.6s (55/56)
[+] Building 264.1s (60/60) FINISHED
 => [ext-authz-default internal] load build definition from Dockerfile                                                       0.0s
 => => transferring dockerfile: 273B                                                                                         0.0s
 => [ztunnel-default internal] load build definition from Dockerfile.ztunnel                                                 0.0s
 => => transferring dockerfile: 1.37kB                                                                                       0.0s
 => [proxyv2-default internal] load build definition from Dockerfile.proxyv2                                                 0.0s
 => => transferring dockerfile: 1.76kB                                                                                       0.0s
 => [operator-default internal] load build definition from Dockerfile.operator                                               0.0s
 => => transferring dockerfile: 948B                                                                                         0.0s
 => [pilot-default internal] load build definition from Dockerfile.pilot                                                     0.0s
 => => transferring dockerfile: 1.07kB                                                                                       0.0s
 => [istioctl-default internal] load build definition from Dockerfile.istioctl                                               0.0s
 => => transferring dockerfile: 353B                                                                                         0.0s
 => [install-cni-default internal] load build definition from Dockerfile.install-cni                                         0.0s
 => => transferring dockerfile: 1.05kB                                                                                       0.0s
 => [app-default internal] load build definition from Dockerfile.app                                                         0.0s
 => => transferring dockerfile: 386B                                                                                         0.0s
 => [ext-authz-default internal] load .dockerignore                                                                          0.0s
 => => transferring context: 2B                                                                                              0.0s
 => [ztunnel-default internal] load .dockerignore                                                                            0.0s
 => => transferring context: 2B                                                                                              0.0s
 => [proxyv2-default internal] load .dockerignore                                                                            0.0s
 => => transferring context: 2B                                                                                              0.0s
 => [operator-default internal] load .dockerignore                                                                           0.0s
 => => transferring context: 2B                                                                                              0.0s
 => [pilot-default internal] load .dockerignore                                                                              0.0s
 => => transferring context: 2B                                                                                              0.0s
 => [istioctl-default internal] load .dockerignore                                                                           0.0s
 => => transferring context: 2B                                                                                              0.0s
 => [install-cni-default internal] load .dockerignore                                                                        0.0s
 => => transferring context: 2B                                                                                              0.0s
 => [app-default internal] load .dockerignore                                                                                0.0s
 => => transferring context: 2B                                                                                              0.0s
 => [install-cni-default internal] load metadata for gcr.io/istio-release/base:master-2023-06-15T19-01-36                    1.2s
 => CACHED [ext-authz-default 1/2] FROM gcr.io/istio-release/base:master-2023-06-15T19-01-36@sha256:d8c8b994550f4ffa05f1584  0.2s
 => => resolve gcr.io/istio-release/base:master-2023-06-15T19-01-36@sha256:d8c8b994550f4ffa05f1584a15b7d894a7e7db3e620ebfda  0.0s
 => [istioctl-default internal] load build context                                                                           6.4s
 => => transferring context: 85.92MB                                                                                         6.4s
 => [app-default internal] load build context                                                                                4.2s
 => => transferring context: 57.29MB                                                                                         4.2s
 => [pilot-default internal] load build context                                                                              6.3s
 => => transferring context: 84.40MB                                                                                         6.3s
 => [ext-authz-default internal] load build context                                                                          1.2s
 => => transferring context: 13.70MB                                                                                         1.2s
 => [proxyv2-default internal] load build context                                                                            9.4s
 => => transferring context: 151.53MB                                                                                        9.3s
 => [ztunnel-default internal] load build context                                                                            2.1s
 => => transferring context: 24.17MB                                                                                         2.0s
 => [install-cni-default internal] load build context                                                                        9.1s
 => => transferring context: 124.11MB                                                                                        9.1s
 => [operator-default internal] load build context                                                                           6.5s
 => => transferring context: 84.53MB                                                                                         6.4s
 => [ext-authz-default 2/2] COPY amd64/extauthz /usr/local/bin/extauthz                                                      0.2s
 => [proxyv2-default] exporting to image                                                                                   260.5s
 => => exporting layers                                                                                                      6.0s
 => => exporting manifest sha256:2dcc791a73484503bfb12e27387777660a6c03bea5cea8def8862133631f7410                            0.0s
 => => exporting manifest sha256:fe9450cb14fb2e6f325613a3b9b3a59ad53a4761942c238c848c467cb37c46ca                            0.1s
 => => exporting config sha256:f027254e6035beb3d5e66f89831d7cde8d619a3a402aa8ac2aed20248982e438                              0.1s
 => => pushing layers                                                                                                      245.1s
 => => exporting config sha256:4679ce6464569a6d61d831e7815895e80e3e550c75debca98d3701213ed7808b                              0.0s
 => => exporting manifest sha256:4b7ee7a3b009b3f47831aedad77e2aa8b1fbd8bc867a57ef4e7597e6629d1c8f                            0.0s
 => => exporting config sha256:2c413cb6800dc00c6b6105649e5c79e23234ab644286db8fb5656f905df86193                              0.0s
 => => exporting manifest sha256:0d3996514089fbbef80a5b1431d2a369549b2d4b387d03faeb981d7b979ed5e3                            0.0s
 => => exporting config sha256:723770387a66969048cfa95fca93ace62aca8bfa7d6949ef4779ffdd5da3a003                              0.0s
 => => exporting manifest sha256:7a0158464e1cfc48d7084f552f93fbbf927f6863bf57a84ea8e01d5b0ba9b4d8                            0.0s
 => => exporting manifest sha256:4e5781453e9e145612b754ec124224c0d791a056873290acabf2ee4d0a17c806                            0.0s
 => => exporting config sha256:391455ffe87580c007e4d199939f751303c8da347d11ac73b7f5387f360d3b6c                              0.0s
 => => exporting config sha256:a968659282b8f992631e3ea31dbe25bdcb1273f7f1c782917b69820b232ee5df                              0.0s
 => => exporting manifest sha256:04e261cf8ad9f828346d2511b63ce55f3dc043249f8643d62bfc9953cdc246d2                            0.0s
 => => exporting config sha256:84e8a05603d24dbc3fb38f02b8cdc5af137c7378c904d4d7d32ae18d30cff117                              0.0s
 => => exporting manifest sha256:a483319c4bca1ff300525636e503e793e89228d6aa538f198f4f874baa05e110                            0.0s
 => => exporting config sha256:0fd0fc6bc9967033bf9447b91425d40d413e0f5594b9462f1a86c4293fceda7e                              0.0s
 => => pushing manifest for docker.io/tanjunchen/ztunnel:06-25-dev@sha256:2dcc791a73484503bfb12e27387777660a6c03bea5cea8de  24.3s
 => => pushing manifest for docker.io/tanjunchen/ext-authz:06-25-dev@sha256:fe9450cb14fb2e6f325613a3b9b3a59ad53a4761942c23  10.7s
 => => pushing manifest for docker.io/tanjunchen/istioctl:06-25-dev@sha256:0d3996514089fbbef80a5b1431d2a369549b2d4b387d03fa  6.0s
 => => pushing manifest for docker.io/tanjunchen/app:06-25-dev@sha256:4b7ee7a3b009b3f47831aedad77e2aa8b1fbd8bc867a57ef4e759  5.5s
 => => pushing manifest for docker.io/tanjunchen/operator:06-25-dev@sha256:4e5781453e9e145612b754ec124224c0d791a056873290ac  2.0s
 => => pushing manifest for docker.io/tanjunchen/pilot:06-25-dev@sha256:7a0158464e1cfc48d7084f552f93fbbf927f6863bf57a84ea8e  1.9s
 => => pushing manifest for docker.io/tanjunchen/install-cni:06-25-dev@sha256:04e261cf8ad9f828346d2511b63ce55f3dc043249f864  3.5s
 => => pushing manifest for docker.io/tanjunchen/proxyv2:06-25-dev@sha256:a483319c4bca1ff300525636e503e793e89228d6aa538f198  1.0s
 => [ztunnel-default stage-2 1/2] COPY amd64/ztunnel /usr/local/bin/ztunnel                                                  0.2s
 => [app-default 2/5] COPY amd64/client /usr/local/bin/client                                                                0.4s
 => [app-default 3/5] COPY amd64/server /usr/local/bin/server                                                                0.1s
 => [app-default 4/5] COPY certs/cert.crt /cert.crt                                                                          0.1s
 => [app-default 5/5] COPY certs/cert.key /cert.key                                                                          0.1s
 => [pilot-default stage-2 1/3] COPY amd64/pilot-discovery /usr/local/bin/pilot-discovery                                    1.0s
 => [istioctl-default 2/2] COPY amd64/istioctl /usr/local/bin/istioctl                                                       0.9s
 => [operator-default stage-2 1/2] COPY amd64/operator /usr/local/bin/operator                                               0.8s
 => [pilot-default stage-2 2/3] COPY envoy_bootstrap.json /var/lib/istio/envoy/envoy_bootstrap_tmpl.json                     0.2s
 => [operator-default stage-2 2/2] COPY manifests/ /var/lib/istio/manifests/                                                 0.4s
 => [pilot-default stage-2 3/3] COPY gcp_envoy_bootstrap.json /var/lib/istio/envoy/gcp_envoy_bootstrap_tmpl.json             0.1s
 => [auth] tanjunchen/ztunnel:pull,push token for registry-1.docker.io                                                       0.0s
 => [install-cni-default stage-2 1/3] COPY amd64/istio-cni /opt/cni/bin/istio-cni                                            0.3s
 => [install-cni-default stage-2 2/3] COPY amd64/install-cni /usr/local/bin/install-cni                                      0.3s
 => CACHED [proxyv2-default stage-2 1/5] COPY envoy_bootstrap.json /var/lib/istio/envoy/envoy_bootstrap_tmpl.json            0.0s
 => CACHED [proxyv2-default stage-2 2/5] COPY gcp_envoy_bootstrap.json /var/lib/istio/envoy/gcp_envoy_bootstrap_tmpl.json    0.0s
 => [proxyv2-default stage-2 3/5] COPY amd64/envoy /usr/local/bin/envoy                                                      0.3s
 => [install-cni-default stage-2 3/3] WORKDIR /opt/cni/bin                                                                   0.1s
 => [proxyv2-default stage-2 4/5] COPY amd64/pilot-agent /usr/local/bin/pilot-agent                                          0.3s
 => [auth] tanjunchen/ztunnel:pull,push token for registry-1.docker.io                                                       0.0s
 => [auth] tanjunchen/ext-authz:pull,push token for registry-1.docker.io                                                     0.0s
 => [auth] tanjunchen/app:pull,push token for registry-1.docker.io                                                           0.0s
 => [auth] tanjunchen/istioctl:pull,push token for registry-1.docker.io                                                      0.0s
 => [auth] tanjunchen/operator:pull,push token for registry-1.docker.io                                                      0.0s
 => [auth] tanjunchen/pilot:pull,push token for registry-1.docker.io                                                         0.0s
 => [auth] tanjunchen/install-cni:pull,push token for registry-1.docker.io                                                   0.0s
 => [auth] tanjunchen/proxyv2:pull,push token for registry-1.docker.io                                                       0.0s
 => [auth] tanjunchen/istioctl:pull,push tanjunchen/ztunnel:pull token for registry-1.docker.io                              0.0s
 => [auth] tanjunchen/operator:pull,push tanjunchen/ztunnel:pull token for registry-1.docker.io                              0.0s
 => [auth] tanjunchen/pilot:pull,push tanjunchen/ztunnel:pull token for registry-1.docker.io                                 0.0s
 => [auth] tanjunchen/install-cni:pull,push tanjunchen/ztunnel:pull token for registry-1.docker.io                           0.0s
 => [auth] tanjunchen/proxyv2:pull,push tanjunchen/ztunnel:pull token for registry-1.docker.io                               0.0s
2023-06-25T03:21:41.261921Z	info	images complete	runtime=4m27.039379959s
2023-06-25T03:21:41.262766Z	info	build complete	runtime=4m44.489377083s
```
查看 docker hub 镜像仓库，已存在上述镜像。
5. 构建并推送特定的镜像
```
BUILD_WITH_CONTAINER=0 make push.docker.pilot
```

6. 清理二进制与镜像
```
make clean
```

## 参考

1. https://github.com/istio/istio/wiki/Preparing-for-Development
1. https://github.com/istio/istio/wiki/Using-the-Code-Base
