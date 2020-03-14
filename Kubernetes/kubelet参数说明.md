--add-dir-header

设置为 true 表示添加文件目录到 header 中

--address 0.0.0.0

kubelet 的服务 IP 地址（所有 IPv4 接口设置为 0.0.0.0 ，所有 IPv6 接口设置为 “::”）（默认值为 0.0.0.0）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--allowed-unsafe-sysctls strings

设置允许的非安全 sysctls 或 sysctl 模式(以 * 结尾) 白名单。使用此参数，风险自担。（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--alsologtostderr

设置为 true 表示将日志输出到文件的同时输出到 stderr

--anonymous-auth

设置为 true 表示 kubelet 服务器可以接受匿名请求。未被任何认证组件拒绝的请求将被视为匿名请求。匿名请求的用户名为 system:anonymous，用户组为 system:unauthenticated。（默认值为 true）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--application-metrics-count-limit int

设置每个容器应用性能度量值存储的个数上限。（默认值为 100）。（已弃用：这是一个错误地在 kubelet 中注册的 cadvisor 参数。由于遗留问题，在删除之前，它将遵循标准的 CLI 弃用时间表。）

--authentication-token-webhook

使用 TokenReview API 对持有者令牌进行身份认证。（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--authentication-token-webhook-cache-ttl duration

webhook 令牌认证器返回的响应的缓存时间。（默认值为 2m0s）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--authorization-mode string

kubelet 服务器的鉴权模式。可选值包括：AlwaysAllow、Webhook。Webhook 模式使用 SubjectAccessReview API 鉴权。（默认值：“AlwaysAllow”）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--authorization-webhook-cache-authorized-ttl duration

webhook 认证器所返回的 “己授权” 应答的缓存时间。（默认值为 5m0s）（默认值为 “AlwaysAllow”）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--authorization-webhook-cache-unauthorized-ttl duration

webhook 认证器所返回的 “未授权” 应答的缓存时间。（默认值为 30s）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--azure-container-registry-config string

Azure 云上镜像库的配置文件路径。

--boot-id-file string

以逗号分隔的文件列表，用于检查引导 id（boot-id）。使用第 1 个存在 boot-id 的文件。（默认值为“/proc/sys/kernel/random/boot_id”）（已弃用：这是一个错误地在 kubelet 中注册的 cadvisor 参数。由于遗留问题，在删除之前，它将遵循标准的 CLI 弃用时间表。）

--bootstrap-checkpoint-path string

<警告：alpha 功能> 存储检查点的目录的路径

--bootstrap-kubeconfig string

kubeconfig 文件的路径，该文件将用于获取 kubelet 的客户端证书。如果 --kubeconfig 指定的文件不存在，则使用引导 kubeconfig 从 API 服务器请求客户端证书。成功后，将引用生成的客户端证书和密钥的 kubeconfig 文件写入 --kubeconfig 所指定的路径。客户端证书和密钥文件将存储在 --cert-dir 指向的目录中。

--cert-dir string

TLS 证书所在的目录。如果设置了 --tls-cert-file 和 --tls-private-key-file，则该设置将被忽略。（默认值为 “/var/lib/kubelet/pki”）

--cgroup-driver string

kubelet 操作本机 cgroup 时使用的驱动程序。支持的选项包括 cgroupfs 或者 systemd（默认值为 cgroupfs）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。

--cgroup-root string

为 pod 设置可选的根 cgroup。容器运行时会尽力而为。默认值：‘’，意味着将使用容器运行时的默认设置。（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--cgroups-per-qos

启用创建 QoS cgroup 层次结构。此值为 true 时创建顶级的 QoS 和 Pod cgroup。（默认值为 true）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--chaos-chance float

如果大于 0.0，则引入随机客户端错误和延迟。用于测试。

--client-ca-file string

如果已设置客户端 CA 证书文件，则使用与客户端证书的 CommonName 对应的身份对任何携带 client-ca 文件中的授权机构之一签名的客户端证书的请求进行身份验证。（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--cloud-config string

云服务商的配置文件路径。

--cloud-provider string

云服务商。设置为空字符串表示在没有云服务商的情况下运行。如果已设置云服务商，则云服务商将确定节点的名称（查阅云提供商文档以确定是否以及如何使用主机名）。

--cluster-dns strings

集群内 DNS 服务的 IP 地址，以逗号分隔。仅当 Pod 设置了 “dnsPolicy=ClusterFirst” 属性时可用。注意：列表中出现的所有 DNS 服务器必须包含相同的记录组，否则集群中的名称解析可能无法正常工作。无法保证名称解析过程中会牵涉到哪些 DNS 服务器。（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--cluster-domain string

集群的域名。如果设置了此值，除了主机的搜索域外，kubelet 还将配置所有容器以搜索所指定的域名（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/）。

--cni-bin-dir string

<警告：alpha 功能> 以逗号分隔的目录的完整路径列表，kubelet 将在其中搜索 CNI 插件可执行文件。仅当容器运行环境设置为 docker 时，此特定于 docker 的参数才有效。(默认为 “/opt/cni/bin”)

--cni-cache-dir string

<警告：alpha 功能> CNI 用于缓存文件的目录的完整路径。仅当容器运行环境设置为 docker 时，此特定于 docker 的参数才有效。(默认为 “/var/lib/cni/cache”)

--cni-conf-dir string

<警告：alpha 功能> 用来搜索 CNI 配置文件的目录的完整路径。仅当容器运行环境设置为 docker 时，此特定于 docker 的参数才有效。(默认为 “/etc/cni/net.d”)

--config string

kubelet 将从该文件加载其初始配置。该路径可以是绝对路径，也可以是相对路径。相对路径从 kubelet 的当前工作目录开始。省略此参数则使用内置的默认配置值。命令行参数会覆盖此文件中的配置。

--container-hints string

容器提示（hints）文件的位置。（默认值为 “/etc/cadvisor/container_hints.json”）（已弃用：这是一个错误地在 kubelet 中注册的 cadvisor 参数。由于遗留问题，在删除之前，它将遵循标准的 CLI 弃用时间表。）

--container-log-max-files int32

<警告：beta 功能> 设置容器可以存在的容器日志文件个数上限。此值必须不小于 2。此参数只能与 --container-runtime=remote 参数一起使用。（默认值为 5）。（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--container-log-max-size string

<警告：beta 功能> 设置容器日志文件在轮转生成新文件时之前的最大值。此参数只能与 --container-runtime=remote 参数一起使用。（默认值为 “10Mi”）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--container-runtime string

要使用的容器运行时。目前支持 ‘docker’、‘remote’、‘rkt(已弃用)’。（默认值为 ‘docker’）。

--container-runtime-endpoint string

[实验性特性] 容器运行时的远程服务端点。目前支持的类型包括 Linux 系统上的 UNIX 套接字、Windows 系统上支持的 npipe 和 TCP 端点。例如：‘unix:///var/run/dockershim.sock’、‘npipe:////./pipe/dockershim’。（默认值为 “unix:///var/run/dockershim.sock”）

--containerd string

设置 containerd 的端点（默认值为 “/run/containerd/containerd.sock”）（已弃用：这是一个错误地在 kubelet 中注册的 cadvisor 参数。由于遗留问题，在删除之前，它将遵循标准的 CLI 弃用时间表。）

--contention-profiling

当启用了性能分析时，启用锁竞争分析（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--cpu-cfs-quota

设置为 true 表示启用 CPU CFS 配额，用于设置容器的 CPU 限制（默认值为 true）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--cpu-cfs-quota-period duration

设置 CPU CFS 配额周期，cpu.cfs_period_us。默认使用 Linux 内核所设置的默认值 （默认值为 100ms）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--cpu-manager-policy string

设置 CPU 管理器策略。可选值包括：‘none’ 和 ‘static’。默认值：“none”。（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--cpu-manager-reconcile-period NodeStatusUpdateFrequency

<警告：alpha 功能> 设置 CPU 管理器的调和时间。例如：‘10s’ 或者 ‘1m’。如果未设置，默认使用 NodeStatusUpdateFrequency 取值（默认值为 10s）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--docker string

docker 服务的端点地址（默认值为 “unix:///var/run/docker.sock”）（已弃用：这是一个错误地在 kubelet 中注册的 cadvisor 参数。由于遗留问题，在删除之前，它将遵循标准的 CLI 弃用时间表。）

--docker-endpoint string

docker 端点使用该参数值进行通信。仅当容器运行环境设置为 docker 时，此特定于 docker 的参数才有效。(默认为 “unix:///var/run/docker.sock”)

--docker-env-metadata-whitelist string

docker 容器需要使用的以逗号分隔的环境变量键名列表（已弃用：这是一个错误地在 kubelet 中注册的 cadvisor 参数。由于遗留问题，在删除之前，它将遵循标准的 CLI 弃用时间表。）

--docker-only

设置为 true 表示除了根统计信息外，仅报告 Docker 容器的统计信息（已弃用：这是一个错误地在 kubelet 中注册的 cadvisor 参数。由于遗留问题，在删除之前，它将遵循标准的 CLI 弃用时间表。）

--docker-root string

已弃用：docker 根目录的路径（备用，默认值：/var/lib/docker）（默认值为 “/var/lib/docker”）

--docker-tls

使用 TLS 连接 docker。（已弃用：这是一个错误地在 kubelet 中注册的 cadvisor 参数。由于遗留问题，在删除之前，它将遵循标准的 CLI 弃用时间表。）

--docker-tls-ca string

可信 CA 的路径（默认值为 “ca.pem”）（已弃用：这是一个错误地在 kubelet 中注册的 cadvisor 参数。由于遗留问题，在删除之前，它将遵循标准的 CLI 弃用时间表。）

--docker-tls-cert string

客户端证书的路径（默认值为 “cert.pem”）（已弃用：这是一个错误地在 kubelet 中注册的 cadvisor 参数。由于遗留问题，在删除之前，它将遵循标准的 CLI 弃用时间表。）

--docker-tls-key string

私钥文件路径（默认值为 “key.pem”）（已弃用：这是一个错误地在 kubelet 中注册的 cadvisor 参数。由于遗留问题，在删除之前，它将遵循标准的 CLI 弃用时间表。）

--dynamic-config-dir string

设置 kubelet 使用此目录来保存所下载的配置和跟踪配置运行状况。如果目录不存在，则创建该目录。该路径可以是绝对路径，也可以是相对路径。相对路径从 kubelet 的当前工作目录开始。设置此参数将启用动态 kubelet 配置。必须启用 DynamicKubeletConfig 特性开关才能传递此参数；由于该功能为 beta，此特性开关当前默认为 true。

--enable-cadvisor-json-endpoints

启用 cAdvisor JSON 数据的 /spec 和 /stats/* 端点。（默认值为 true）

--enable-controller-attach-detach

设置为 true 表示启用 Attach/Detach 控制器进行来挂接和摘除调度到该节点的卷，同时禁用 kubelet 执行挂接和摘除操作（默认值为 true）

--enable-debugging-handlers

设置为 true 表示启用服务器端点进行日志收集和在本地运行容器和命令（默认值为 true）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--enable-load-reader

设置为 true 表示启用读取 CPU 负载（已弃用：这是一个错误地在 kubelet 中注册的 cadvisor 参数。由于遗留问题，在删除之前，它将遵循标准的 CLI 弃用时间表。）

--enable-server

启动 kubelet 服务器（默认值为 true）

--enforce-node-allocatable stringSlice

用逗号分隔的列表，包含由 kubelet 强制执行的节点可分配资源级别。可选配置为：‘none’、‘pods’、‘system-reserved’ 和 ‘kube-reserved’。在设置 ‘system-reserved’ 和 ‘kube-reserved’ 这两个值时，同时要求设置 ‘--system-reserved-cgroup’ 和 ‘--kube-reserved-cgroup’ 这两个参数。如果设置为 ‘none’，则不需要设置其他参数。更多信息请参考 https://kubernetes.io/docs/tasks/administer-cluster/reserve-compute-resources/。（默认值为 pods）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--event-burst int32

突发事件记录的个数上限，在遵从 event-qps 阈值约束的前提下临时允许事件记录达到此数目。仅在 --event-qps 大于 0 时使用（默认值为 10）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--event-qps int32

设置大于 0 的值表示限制每秒可生成的事件数量。设置为 0 表示不限制。（默认值为 5）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--event-storage-age-limit string

不同类型事件的最长保存时间。取值是键值对（key=value）的逗号分隔列表，其中键名是事件类型（例如：creation、oom）或者 “default”，键值是持续时间（duration）。所有未指定的事件类型都使用默认值（默认值为 “default=0”）（已弃用：这是一个错误地在 kubelet 中注册的 cadvisor 参数。由于遗留问题，在删除之前，它将遵循标准的 CLI 弃用时间表。）

--event-storage-event-limit string

每种事件类型的最大保存数量。取值是键值对（key=value）的逗号分隔列表，其中键名是事件类型（例如：creation、oom）或者 “default”，键值是一个整数（integer）。所有未指定的事件类型都使用默认值（默认值为 “default=0”）（已弃用：这是一个错误地在 kubelet 中注册的 cadvisor 参数。由于遗留问题，在删除之前，它将遵循标准的 CLI 弃用时间表。）

--eviction-hard mapStringString

触发 Pod 驱逐操作的一组硬性门限（例如：memory.available < 1Gi（内存可用值小于 1 G））设置。（默认值为 imagefs.available<15%,memory.available<100Mi,nodefs.available<10%,nodefs.inodesFree<5%）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）
--eviction-max-pod-grace-period int32

响应满足软驱逐阈值（soft eviction threshold）而终止 Pod 时使用的最长宽限期（以秒为单位）。如果设置为负数，则遵循 Pod 的指定值。（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--eviction-minimum-reclaim mapStringString

当本节点压力过大时，kubelet 执行软性驱逐操作。此参数设置软性驱逐操作需要回收的资源的最小数量（例如：imagefs.available=2Gi）。（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--eviction-pressure-transition-period duration

kubelet 在触发软性 Pod 驱逐操作之前的最长等待时间。（默认值为 5m0s）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--eviction-soft mapStringString

设置一组驱逐阈值（例如：memory.available<1.5Gi）。如果在相应的宽限期内达到该阈值，则会触发软性 Pod 驱逐操作。（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--eviction-soft-grace-period mapStringString

设置一组驱逐宽限期，对应于触发软性 Pod 驱逐操作之前软性驱逐阈值所需持续的时间长短。（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--exit-on-lock-contention

设置为 true 表示当发生锁文件竞争时 kubelet 可以退出。

--experimental-allocatable-ignore-eviction

设置为 true 表示在计算节点可分配资源数量时忽略硬性逐出阈值设置。请参考 https://kubernetes.io/docs/tasks/administer-cluster/reserve-compute-resources/。（默认值为 false）。

--experimental-bootstrap-kubeconfig string

（已弃用：使用 --bootstrap-kubeconfig 参数）

--experimental-check-node-capabilities-before-mount

[实验性特性] 设置为 true 表示 kubelet 在进行挂载卷操作之前对本节点上所需的组件（如可执行文件等）进行检查

--experimental-kernel-memcg-notification

设置为 true 表示 kubelet 将会集成内核的 memcg 通知机制而不是使用轮询机制来判断是否达到了内存驱逐阈值。

--experimental-mounter-path string

[实验性特性] 卷挂载器（mounter）可执行文件的路径。设置为空表示使用默认挂载器 mount。

--fail-swap-on

设置为 true 表示如果主机启用了交换分区，kubelet 将无法使用。（默认值为 true）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--feature-gates mapStringBool

    用于 alpha 实验性质的特性开关组，每个开关以 key=value 形式表示。当前可用开关包括：
    APIListChunking=true|false (BETA - 默认值=true)
    APIResponseCompression=true|false (BETA - 默认值=true)
    AllAlpha=true|false (ALPHA - 默认值=false)
    AppArmor=true|false (BETA - 默认值=true)
    AttachVolumeLimit=true|false (BETA - 默认值=true)
    BalanceAttachedNodeVolumes=true|false (ALPHA - 默认值=false)
    BlockVolume=true|false (BETA - 默认值=true)
    BoundServiceAccountTokenVolume=true|false (ALPHA - 默认值=false)
    CPUManager=true|false (BETA - 默认值=true)
    CRIContainerLogRotation=true|false (BETA - 默认值=true)
    CSIBlockVolume=true|false (BETA - 默认值=true)
    CSIDriverRegistry=true|false (BETA - 默认值=true)
    CSIInlineVolume=true|false (BETA - 默认值=true)
    CSIMigration=true|false (ALPHA - 默认值=false)
    CSIMigrationAWS=true|false (ALPHA - 默认值=false)
    CSIMigrationAzureDisk=true|false (ALPHA - 默认值=false)
    CSIMigrationAzureFile=true|false (ALPHA - 默认值=false)
    CSIMigrationGCE=true|false (ALPHA - 默认值=false)
    CSIMigrationOpenStack=true|false (ALPHA - 默认值=false)
    CSINodeInfo=true|false (BETA - 默认值=true)
    CustomCPUCFSQuotaPeriod=true|false (ALPHA - 默认值=false)
    CustomResourceDefaulting=true|false (BETA - 默认值=true)
    DevicePlugins=true|false (BETA - 默认值=true)
    DryRun=true|false (BETA - 默认值=true)
    DynamicAuditing=true|false (ALPHA - 默认值=false)
    DynamicKubeletConfig=true|false (BETA - 默认值=true)
    EndpointSlice=true|false (ALPHA - 默认值=false)
    EphemeralContainers=true|false (ALPHA - 默认值=false)
    EvenPodsSpread=true|false (ALPHA - 默认值=false)
    ExpandCSIVolumes=true|false (BETA - 默认值=true)
    ExpandInUsePersistentVolumes=true|false (BETA - 默认值=true)
    ExpandPersistentVolumes=true|false (BETA - 默认值=true)
    ExperimentalHostUserNamespaceDefaulting=true|false (BETA - 默认值=false)
    HPAScaleToZero=true|false (ALPHA - 默认值=false)
    HyperVContainer=true|false (ALPHA - 默认值=false)
    IPv6DualStack=true|false (ALPHA - 默认值=false)
    KubeletPodResources=true|false (BETA - 默认值=true)
    LegacyNodeRoleBehavior=true|false (ALPHA - 默认值=true)
    LocalStorageCapacityIsolation=true|false (BETA - 默认值=true)
    LocalStorageCapacityIsolationFSQuotaMonitoring=true|false (ALPHA - 默认值=false)
    MountContainers=true|false (ALPHA - 默认值=false)
    NodeDisruptionExclusion=true|false (ALPHA - 默认值=false)
    NodeLease=true|false (BETA - 默认值=true)
    NonPreemptingPriority=true|false (ALPHA - 默认值=false)
    PodOverhead=true|false (ALPHA - 默认值=false)
    PodShareProcessNamespace=true|false (BETA - 默认值=true)
    ProcMountType=true|false (ALPHA - 默认值=false)br/>QOSReserved=true|false (ALPHA - 默认值=false)
    RemainingItemCount=true|false (BETA - 默认值=true)
    RemoveSelfLink=true|false (ALPHA - 默认值=false)
    RequestManagement=true|false (ALPHA - 默认值=false)
    ResourceLimitsPriorityFunction=true|false (ALPHA - 默认值=false)
    ResourceQuotaScopeSelectors=true|false (BETA - 默认值=true)
    RotateKubeletClientCertificate=true|false (BETA - 默认值=true)
    RotateKubeletServerCertificate=true|false (BETA - 默认值=true)
    RunAsGroup=true|false (BETA - 默认值=true)
    RuntimeClass=true|false (BETA - 默认值=true)
    SCTPSupport=true|false (ALPHA - 默认值=false)
    ScheduleDaemonSetPods=true|false (BETA - 默认值=true)
    ServerSideApply=true|false (BETA - 默认值=true)
    ServiceLoadBalancerFinalizer=true|false (BETA - 默认值=true)
    ServiceNodeExclusion=true|false (ALPHA - 默认值=false)
    StartupProbe=true|false (BETA - 默认值=true)
    StorageVersionHash=true|false (BETA - 默认值=true)
    StreamingProxyRedirects=true|false (BETA - 默认值=true)
    SupportNodePidsLimit=true|false (BETA - 默认值=true)
    SupportPodPidsLimit=true|false (BETA - 默认值=true)
    Sysctls=true|false (BETA - 默认值=true)
    TTLAfterFinished=true|false (ALPHA - 默认值=false)
    TaintBasedEvictions=true|false (BETA - 默认值=true)
    TaintNodesByCondition=true|false (BETA - 默认值=true)
    TokenRequest=true|false (BETA - 默认值=true)
    TokenRequestProjection=true|false (BETA - 默认值=true)
    TopologyManager=true|false (ALPHA - 默认值=false)
    ValidateProxyRedirects=true|false (BETA - 默认值=true)
    VolumePVCDataSource=true|false (BETA - 默认值=true)
    VolumeSnapshotDataSource=true|false (ALPHA - 默认值=false)
    VolumeSubpathEnvExpansion=true|false (BETA - 默认值=true)
    WatchBookmark=true|false (BETA - 默认值=true)
    WinDSR=true|false (ALPHA - 默认值=false)
    WinOverlay=true|false (ALPHA - 默认值=false)
    WindowsGMSA=true|false (BETA - 默认值=true)
    WindowsRunAsUserName=true|false (ALPHA - 默认值=false)（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--file-check-frequency duration

检查配置文件中新数据的时间间隔（默认值为 20s）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--global-housekeeping-interval duration

全局资源清理（housekeeping）操作的时间间隔。（默认值为 1m0s）（已弃用：这是一个错误地在 kubelet 中注册的 cadvisor 参数。由于遗留问题，在删除之前，它将遵循标准的 CLI 弃用时间表。）

--hairpin-mode string

设置 kubelet 执行发夹模式（hairpin）网络地址转译的方式。该模式允许后端端点对其自身服务的访问能够再次经由负载均衡转发回自身。可选项包括 “promiscuous-bridge”、“hairpin-veth” 和 “none”。（默认值为 “promiscuous-bridge”）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--healthz-bind-address 0.0.0.0

用于运行 healthz 服务器的 IP 地址（对于所有 IPv4 接口，设置为 0.0.0.0；对于所有 IPv6 接口，设置为 `::`）（默认值为 127.0.0.1）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--healthz-port int32

本地 healthz 端点使用的端口（设置为 0 表示禁用）（默认值为 10248）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

-h, --help

kubelet 操作的帮助命令

--hostname-override string

如果为非空，将使用此字符串而不是实际的主机名作为节点标识。如果设置了 --cloud-provider，则云服务商将确定节点的名称（请查询云服务商文档以确定是否以及如何使用主机名）。

--housekeeping-interval duration

清理容器操作的时间间隔（默认值为 10 s）

--http-check-frequency duration

HTTP 服务以获取新数据的时间间隔（默认值为 20 s）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--image-gc-high-threshold int32

镜像垃圾回收上限。磁盘使用空间达到该百分比时，镜像垃圾回收将持续工作。值必须在 [0，100] 范围内。要禁用镜像垃圾回收，请设置为 100。（默认值为 85）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--image-gc-low-threshold int32

镜像垃圾回收下限。磁盘使用空间在达到该百分比之前，镜像垃圾回收操作不会运行。值必须在 [0，100] 范围内，并且不得大于 --image-gc-high-threshold 的值。（默认值为 80）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--image-pull-progress-deadline duration

如果在该参数值所设置的期限之前没有拉取镜像的进展，镜像拉取操作将被取消。仅当容器运行环境设置为 docker 时，此特定于 docker 的参数才有效。（默认值为 1m0s）

--image-service-endpoint string

[实验性特性] 远程镜像服务的端点。若未设定则默认情况下使用 container-runtime-endpoint 的值。目前支持的类型包括在 Linux 系统上的 UNIX 套接字端点和 Windows 系统上的 npipe 和 TCP 端点。例如：‘unix:///var/run/dockershim.sock’、‘npipe:////./pipe/dockershim’。

--iptables-drop-bit int32

标记数据包将被丢弃的 fwmark 位设置。必须在 [0，31] 范围内。（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--iptables-masquerade-bit int32

标记数据包将进行 SNAT 的 fwmark 位设置。必须在 [0，31] 范围内。请将此参数与 kube-proxy 中的相应参数匹配。（默认值为 14）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--keep-terminated-pod-volumes

设置为 true 表示 Pod 终止后仍然保留之前挂载过的卷，常用于调试与卷有关的问题。（已弃用：未来版本将会移除该参数）

--kube-api-burst int32

每秒发送到 apiserver 的请求数量上限（默认值为 10）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--kube-api-content-type string

发送到 apiserver 的请求的内容类型。（默认值为 “application/vnd.kubernetes.protobuf”）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--kube-api-qps int32

与 apiserver 通信的每秒查询数（QPS） 值（默认值为 5）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--kube-reserved mapStringString

kubernetes 系统预留的资源配置，以一组 ResourceName=ResourceQuantity 格式表示。（例如：cpu=200m,memory=500Mi,ephemeral-storage=1Gi）。当前支持用于根文件系统的 CPU、内存（memory）和本地临时存储。请参阅 http://kubernetes.io/docs/user-guide/compute-resources 获取更多信息。（默认值为 none）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--kube-reserved-cgroup string

给出某个顶层 cgroup 绝对名称，该 cgroup 用于管理带 ‘--kube-reserved’ 标签的 kubernetes 组件的计算资源。例如：‘/kube-reserved’。（默认值为 ‘’）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--kubeconfig string

kubeconfig 配置文件的路径，指定如何连接到 API 服务器。提供 --kubeconfig 将启用 API 服务器模式，而省略 --kubeconfig 将启用独立模式。

--kubelet-cgroups string

用于创建和运行 kubelet 的 cgroup 的绝对名称。（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--lock-file string

<警告：alpha 功能> kubelet 使用的锁文件的路径。

--log-backtrace-at traceLocation

当日志逻辑执行到命中 file 的第 N 行时，转储调用堆栈（默认值为：0）

--log-cadvisor-usage

设置为 true 表示将 cAdvisor 容器的使用情况写入日志（已弃用：这是一个错误地在 kubelet 中注册的 cadvisor 参数。由于遗留问题，在删除之前，它将遵循标准的 CLI 弃用时间表。）

--log-dir string

如果此值为非空，则在所指定的目录中写入日志文件

--log-file string

如果此值非空，使用所给字符串作为日志文件名

--log-file-max-size uint

定义日志文件的最大值。单位为兆字节（M）。如果值为 0，则最大文件大小表示无限制。（默认值为 1800）

--log-flush-frequency duration

两次日志刷新之间的最大秒数（默认值为 5s）

--logtostderr

日志输出到 stderr 而不是文件（默认值为 true）

--machine-id-file string

以逗号分隔的文件列表，用于检查 machine-id。kubelet 使用存在的第一个 machine-id。（已弃用：这是一个错误地在 kubelet 中注册的 cadvisor 参数。由于遗留问题，在删除之前，它将遵循标准的 CLI 弃用时间表。）

--make-iptables-util-chains

设置为 true 表示 kubelet 将确保 Iptables 规则在主机上存在。（默认值为 true）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--manifest-url string

用于访问要运行的其他 Pod 规范的 URL（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--manifest-url-header --manifest-url-header 'a:hello,b:again,c:world' --manifest-url-header 'b:beautiful'

取值为由 HTTP 头部组成的逗号分隔列表，在访问 --manifest-url 所给出的 URL 时使用。名称相同的多个头部将按所列的顺序添加。该参数可以多次使用。例如：--manifest-url-header ‘a:hello,b:again,c:world’ --manifest-url-header ‘b:beautiful’（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--master-service-namespace string

kubelet 向 Pod 注入 Kubernetes 主控服务信息时使用的命名空间（默认值为 “default”）（已弃用：此参数将在未来的版本中删除。）

--max-open-files int

kubelet 进程可以打开的最大文件数量（默认值为 1000000）。（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--max-pods int32

kubelet 能运行的 Pod 最大数量。（默认值为 110）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--maximum-dead-containers int32

设置全局可保留的已停止容器实例个数上限。每个实例会占用一些磁盘空间。要禁用，请设置为负数。（默认值为 -1）（已弃用：请改用 --eviction-hard 或者 --eviction-soft。此参数将在未来的版本中删除。）

--maximum-dead-containers-per-container int32

可以保留的每个已停止容器的最大实例数量。每个容器占用一些磁盘空间。（默认值为 1）（已弃用：请改用 --eviction-hard 或者 --eviction-soft。此参数将在未来的版本中删除。）

--minimum-container-ttl-duration duration

已结束的容器在被垃圾回收清理之前的最少存活时间。例如：‘300ms’、‘10s’ 或者 ‘2h45m’（已弃用：请改用 --eviction-hard 或者 --eviction-soft。此参数将在未来的版本中删除。）

--minimum-image-ttl-duration duration

不再使用的镜像在被垃圾回收清理之前的最少存活时间。例如：‘300ms’、‘10s’ 或者 ‘2h45m’。（默认值为 2m0s）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--network-plugin string

<警告：alpha 功能> 设置 kubelet/pod 生命周期中各种事件调用的网络插件的名称。仅当容器运行环境设置为 docker 时，此特定于 docker 的参数才有效。

--network-plugin-mtu int32

<警告：alpha 功能> 传递给网络插件的 MTU 值，将覆盖默认值。设置为 0 则使用默认的 1460 MTU。仅当容器运行环境设置为 docker 时，此特定于 docker 的参数才有效。

--node-ip string

节点的 IP 地址。如果设置，kubelet 将使用该 IP 地址作为节点的 IP 地址。

--node-labels mapStringString

<警告：alpha 功能> kubelet 在集群中注册本节点时设置的标签。标签以 key=value 的格式表示，多个标签以逗号分隔。命名空间 ‘kubernetes.io’ 中的标签必须以 kubelet.kubernetes.io 或 node.kubernetes.io 为前缀，或者在以下明确允许范围内（beta.kubernetes.io/arch、beta.kubernetes.io/instance-type、beta.kubernetes.io/os、failure-domain.beta.kubernetes.io/region、 failure-domain.beta.kubernetes.io/zone、failure-domain.kubernetes.io/region、failure-domain.kubernetes.io/zone、kubernetes.io/arch、kubernetes.io/hostname、kubernetes.io/instance-type、kubernetes.io/os）

--node-status-max-images int32

<警告：alpha 功能> 在 Node.Status.Images 中可以报告的最大镜像数量。如果指定为 -1，则不设上限。（默认值为 50）

--node-status-update-frequency duration

指定 kubelet 向主控节点汇报节点状态的时间间隔。注意：更改此常量时请务必谨慎，它必须与 nodecontroller 中的 nodeMonitorGracePeriod 一起使用。（默认值为 10s）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--non-masquerade-cidr string

kubelet 向该 IP 段之外的 IP 地址发送的流量将使用 IP 伪装技术。设置为 “0.0.0.0/0” 则不会使用伪装技术。（默认值为 “10.0.0.0/8”）（已弃用：该参数将在未来版本中删除。）

--oom-score-adj int32

kubelet 进程的 oom-score-adj 参数值。有效范围为 [-1000，1000]（默认值为 -999）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--pod-cidr string

用于给 Pod 分配 IP 地址的 CIDR 地址池，仅在单机模式下使用。在集群模式下，CIDR 设置是从主服务器获取的。对于 IPv6，分配的 IP 的最大数量为 65536（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--pod-infra-container-image string

指定基础设施镜像，Pod 内所有容器与其共享网络和 IPC 命名空间。仅当容器运行环境设置为 docker 时，此特定于 docker 的参数才有效。（默认值为 “k8s.gcr.io/pause:3.1”）

--pod-manifest-path string

设置包含要运行的静态 Pod 的文件的路径，或单个静态 Pod 文件的路径。以点（.）开头的文件将被忽略。（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--pod-max-pids int

设置每个 Pod 中的最大进程数目。如果为 -1，则 kubelet 使用节点可分配的 PID 容量作为默认值。（默认值为 -1）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--pods-per-core int32

kubelet 在每个处理器核上可运行的 Pod 数量。此 kubelet 上的 Pod 总数不能超过 max-pods 值。因此，如果此计算结果导致在 kubelet 上允许更多数量的 Pod，则使用 max-pods 值。值为 0 表示不做限制。（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--port int32

kubelet 服务监听的本机端口号。（默认值为 10250）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--protect-kernel-defaults

设置 kubelet 的默认内核调整行为。如果已设置该参数，当任何内核可调参数与 kubelet 默认值不同时，kubelet 都会出错。（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--provider-id string

设置主机数据库中用来标识节点的唯一标识，即 cloudprovider

--qos-reserved mapStringString

<警告：alpha 功能> 设置在指定的 QoS 级别预留的 Pod 资源请求，以一组 “ResourceName=Percentage（资源名称=百分比）” 的形式进行设置，例如 memory=50%。当前仅支持内存（memory）。要求启用 QOSReserved 特性开关。（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--read-only-port int32

kubelet 可以在没有身份验证/鉴权的情况下提供只读服务的端口（设置为 0 表示禁用）（默认值为 10255）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--really-crash-for-testing

设置为 true 表示发生内核失效时崩溃。仅用于测试。

--redirect-container-streaming

启用容器流数据重定向。如果设置为 false，则 kubelet 将在 apiserver 和容器运行时之间转发容器流数据；如果设置为 true，则 kubelet 将返回指向 apiserver 的 HTTP 重定向指令，而 apiserver 将直接访问容器运行时。代理方法更安全，但会带来一些开销。重定向方法性能更高，但安全性较低，因为 apiserver 和容器运行时之间的连接可能未通过身份验证。

--register-node

将本节点注册到 apiserver。如果未提供 --kubeconfig 参数，则此参数无关紧要，因为 kubelet 将没有要注册的 apiserver。（默认值为 true）

--register-schedulable

注册本节点为可调度的。register-node 为 false 时此设置无效。（默认值为 true）（已弃用：此参数将在未来的版本中删除。）

--register-with-taints []api.Taint

设置本节点的污点标记，格式为 “<key>=<value>:<effect>” ，以逗号分隔。当 register-node 为 false 时此标志无效。

--registry-burst int32

设置突发性镜像拉取的个数上限，在不超过 registration-qps 设置值的前提下暂时允许此参数所给的镜像拉取个数。仅在 --registry-qps 大于 0（默认值为 10）时使用（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--registry-qps int32

如果 --registry-qps 大于 0，用来限制镜像仓库的 QPS 上限。设置为 0，表示不受限制。（默认值为 5）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--resolv-conf string

名字解析服务的配置文件名，用作容器 DNS 解析配置的基础。（默认值为 “/etc/resolv.conf”）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--root-dir string

设置用于管理 kubelet 文件的根目录（例如挂载卷的相关文件）（默认值为 “/var/lib/kubelet”）
--rotate-certificates

<警告：alpha 功能> 设置当客户端证书即将过期时 kubelet 自动从 kube-apiserver 请求新的证书进行轮换。（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--rotate-server-certificates

当证书即将过期时自动从 kube-apiserver 请求新的证书进行轮换。要求启用 RotateKubeletServerCertificate 特性开关，以及对提交的 CertificateSigningRequest 对象进行批复（approve）操作。（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--runonce

设置为 true 表示从本地清单或远程 URL 创建完 Pod 后立即退出 kubelet 进程，与 --enable-server 参数互斥

--runtime-cgroups string

设置用于创建和运行容器运行时的 cgroup 的绝对名称。

--runtime-request-timeout duration

除了长时间运行的请求（包括 pull、logs、exec 和 attach 等操作），设置其他请求的超时时间。到达超时时间时，请求会被取消，抛出一个错误并会等待重试。（默认值为 2m0s）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--seccomp-profile-root string

<警告：alpha 功能> seccomp 配置文件目录。（默认值为 “/var/lib/kubelet/seccomp”）

--serialize-image-pulls

逐一拉取镜像。建议 *不要* 在 docker 守护进程版本低于 1.9 或启用了 Aufs 存储后端的节点上更改默认值。（默认值为 true）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--skip-headers

设置为 true，则在日志消息中去掉标头前缀

--skip-log-headers

设置为 true，打开日志文件时去掉标头

--stderrthreshold severity

设置严重程度达到或超过此阈值的日志输出到标准错误输出（默认值为 2）

--storage-driver-buffer-duration duration

设置存储驱动程序中写操作的缓冲时长，超过时长的操作会作为单一事务提交到非内存后端。（默认值为 1m0s）（已弃用：这是一个错误地在 kubelet 中注册的 cadvisor 参数。由于遗留问题，在删除之前，它将遵循标准的 CLI 弃用时间表。）

--storage-driver-db string

后端存储的数据库名称（默认值为 “cadvisor”）（已弃用：这是一个错误地在 kubelet 中注册的 cadvisor 参数。由于遗留问题，在删除之前，它将遵循标准的 CLI 弃用时间表。）

--storage-driver-host string

后端存储的数据库连接 URL 地址（默认值为 “localhost:8086”）（已弃用：这是一个错误地在 kubelet 中注册的 cadvisor 参数。由于遗留问题，在删除之前，它将遵循标准的 CLI 弃用时间表。）

--storage-driver-password string

后端存储的数据库密码（默认值为 “root”）（已弃用：这是一个错误地在 kubelet 中注册的 cadvisor 参数。由于遗留问题，在删除之前，它将遵循标准的 CLI 弃用时间表。）

--storage-driver-secure

后端存储的数据库是否用安全连接（已弃用：这是一个错误地在 kubelet 中注册的 cadvisor 参数。由于遗留问题，在删除之前，它将遵循标准的 CLI 弃用时间表。）

--storage-driver-table string

后端存储的数据库表名（默认值为 “stats”）（已弃用：这是一个错误地在 kubelet 中注册的 cadvisor 参数。由于遗留问题，在删除之前，它将遵循标准的 CLI 弃用时间表。）

--storage-driver-user string

后端存储的数据库用户名（默认值为 “root”）（已弃用：这是一个错误地在 kubelet 中注册的 cadvisor 参数。由于遗留问题，在删除之前，它将遵循标准的 CLI 弃用时间表。）

--streaming-connection-idle-timeout duration

设置流连接在自动关闭之前可以空闲的最长时间。0 表示没有超时限制。例如：‘5m’（默认值为 4h0m0s）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--sync-frequency duration

在运行中的容器与其配置之间执行同步操作的最长时间间隔（默认值为 1m0s）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--system-cgroups /

cgroup 的绝对名称，用于所有尚未放置在根目录下某 cgroup 内的非内核进程。空值表示不指定 cgroup。回滚该参数需要重启机器。（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--system-reserved mapStringString

系统预留的资源配置，以一组 ”ResourceName=ResourceQuantity“ 的格式表示，（例如：cpu=200m,memory=500Mi,ephemeral-storage=1Gi）。目前仅支持 CPU 和内存（memory）的设置。请参考 http://kubernetes.io/docs/user-guide/compute-resources 获取更多信息。（默认值为 ”none“）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--system-reserved-cgroup string

给出一个顶层 cgroup 绝对名称，该 cgroup 用于管理非 kubernetes 组件，这些组件的计算资源通过 ‘--system-reserved’ 标志进行预留。例如 ‘/system-reserved’。（默认值为 ‘’）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--tls-cert-file string

包含 x509 证书的文件路径，用于 HTTPS 认证。如果有中间证书，则中间证书要串接在在服务器证书之后。如果未提供 --tls-cert-file 和 --tls-private-key-file，kubelet 会为公开地址生成自签名证书和密钥，并将其保存到通过 --cert-dir 指定的目录中。（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--tls-cipher-suites stringSlice

服务器端加密算法列表，以逗号分隔，如果不设置，则使用 Go 语言加密包的默认算法列表。可选加密算法包括：TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256,TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,TLS_ECDHE_ECDSA_WITH_RC4_128_SHA,TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA,TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256,TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,TLS_ECDHE_RSA_WITH_RC4_128_SHA,TLS_RSA_WITH_3DES_EDE_CBC_SHA,TLS_RSA_WITH_AES_128_CBC_SHA,TLS_RSA_WITH_AES_128_CBC_SHA256,TLS_RSA_WITH_AES_128_GCM_SHA256,TLS_RSA_WITH_AES_256_CBC_SHA,TLS_RSA_WITH_AES_256_GCM_SHA384,TLS_RSA_WITH_RC4_128_SHA （已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--tls-min-version string

设置支持的最小 TLS 版本号，可选的版本号包括：VersionTLS10、VersionTLS11、VersionTLS12 和 VersionTLS13 （已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--tls-private-key-file string

包含与 --tls-cert-file 对应的 x509 私钥文件路径。（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

--topology-manager-policy string

设置拓扑管理策略（Topology Manager policy）。可选值包括：‘none’、‘best-effort’、‘restricted’ 和 ‘single-numa-node’。（默认值为 “none”）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）

-v, --v Level

设置 kubelet 日志级别详细程度的数值

--version version[=true]

打印 kubelet 版本信息并退出

--vmodule moduleSpec

以逗号分隔的 pattern=N 设置列表，用于文件过滤的日志记录

--volume-plugin-dir string

<警告：alpha 功能> 用来搜索第三方存储卷插件的目录（默认值为 “/usr/libexec/kubernetes/kubelet-plugins/volume/exec/”）

--volume-stats-agg-period duration

指定 kubelet 计算和缓存所有 Pod 和卷的磁盘用量总值的时间间隔。要禁用磁盘用量计算，请设置为 0。（默认值为 1m0s）（已弃用：在 --config 指定的配置文件中进行设置。有关更多信息，请参阅 https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/。）