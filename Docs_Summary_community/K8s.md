api 输出接口文档用，基本是 json 源码
    api-rules 
    openapi-spec
build 构建脚本
    run.sh
    copy-output.sh
    make-clean.sh
    shell.sh
    ...... 一些构建脚本
cluster 自动创建和配置 kubernetes 集群的脚本，包括 networking、DNS、nodes 等
cmd  所有的二进制可执行文件入口代码，也就是各种命令的接口代码。
    clicheck
    cloud-controller-manager
    controller-manager
    gendocs
    genkubedocs
    genman
    genswaggertypedocs
    genutils
    genyaml
    importverifier
    kube-apiserver
    kube-controller-manager
    kube-proxy
    kube-scheduler
    kubeadm
    kubectl
    kubelet
    kubemark
    linkcheck
    preferredimports
    verifydependencies
docs 文档
Godeps
hack 编译、构建及校验的工具类
logo kubernetes 的 logo
pkg 是 kubernetes 的主体代码，里面实现了 kubernetes 的主体逻辑
    api kubernetes api 主要包括最新版本的 Rest API 接口的类，并提供数据格式验证转换工具类，对应版本号文件夹下的文件描述了特定的版本如何序列化存储和网络
    apis  apis 是提供 api 访问接口服务
    auth 
    capabilities
    client Kubernetes 中公用的客户端部分，实现对对象的具体操作，即增删改查操作
    cloudprovider  kubernetes 提供对 aws、azure、gce、cloudstack、mesos 等云供应商提供了接口支持，目前包括负载均衡、实例、zone 信息、路由信息等
    controller kubernetes controller 主要包括各个 controller 的实现逻辑，为各类资源如 replication、endpoint、node 等的增删改等逻辑提供派发和执行
    credentialprovider kubernetes credentialprovider 为 docker 镜像仓库贡献者提供权限认证
    features
    fieldpath
    generated kubernetes generated 包是所有生成的文件的目标文件，一般这里面的文件日常是不进行改动的
    kubeapiserver API Server 负责和 etcd 交互（其他组件不会直接操作 etcd，只有 API Server 这么做），是整个 kubernetes 集群的数据中心，所有的交互都是以 API Server 为核心的。
    kubectl kuernetes kubectl 模块是 kubernetes 的命令行工具，提供 apiserver 的各个接口的命令行操作，包括各类资源的增删改查、扩容等一系列命令工具
    kubelet kuernetes kubelet 模块是 kubernetes 的核心模块，该模块负责 node 层的 pod 管理，完成 pod 及容器的创建，执行 pod 的删除同步等操作
    kubemark kubemark 测试集群性能
    master kubernetes master 负责集群中 master 节点的运行管理、api 安装、各个组件的运行端口分配、NodeRegistry、PodRegistry 等工作
    printers 
    probe kubernetes 的健康探测
    proxy 代理
    quota 资源配额
    registry 
    routes 路由
    scheduler  调度中心
    security 安全
    securitycontext
    serviceaccount 服务账户
    ssh 
    util pkg 工具类
    volume 各种存储卷
    watch 监听事件
    windows 
plugin 插件
    admission 认证
    auth 鉴权
staging 这里的代码都存放在独立的 repo 中，以引用包的方式添加到项目
test 测试代码
third_party 第三方代码，protobuf、golang-reflect 等
translations 不同国家的语言包，使用poedit查看及编辑
vendor 第三方包

# k8s-controller 源码阅读笔记



# k8s-scheduler 源码阅读笔记

kube-scheduler 是 kubernetes 的核心组件之一，主要负责整个集群资源的调度功能，根据特定的调度算法和策略，将 Pod 调度到最优的工作节点上面去，从而更加合理、更加充分的利用集群的资源。

kubernetes/pkg/scheduler
-- scheduler.go         //调度相关的具体实现
|-- algorithm
|   |-- predicates      //节点筛选策略  预选
|   |-- priorities      //节点打分策略  优选
|-- algorithmprovider
|   |-- defaults         //定义默认的调度器


一般来说，我们有4种扩展 Kubernetes 调度器的方法。

一种方法就是直接 clone 官方的 kube-scheduler 源代码，在合适的位置直接修改代码，然后重新编译运行修改后的程序，当然这种方法是最不建议使用的，也不实用，因为需要花费大量额外的精力来和上游的调度程序更改保持一致。

第二种方法就是和默认的调度程序一起运行独立的调度程序，默认的调度器和我们自定义的调度器可以通过 Pod 的 spec.schedulerName 来覆盖各自的 Pod，默认是使用 default 默认的调度器，但是多个调度程序共存的情况下也比较麻烦，比如当多个调度器将 Pod 调度到同一个节点的时候，可能会遇到一些问题，因为很有可能两个调度器都同时将两个 Pod 调度到同一个节点上去，但是很有可能其中一个 Pod 运行后其实资源就消耗完了，并且维护一个高质量的自定义调度程序也不是很容易的，因为我们需要全面了解默认的调度程序，整体 Kubernetes 的架构知识以及各种 Kubernetes API 对象的各种关系或限制。

第三种方法是调度器扩展程序，这个方案目前是一个可行的方案，可以和上游调度程序兼容，所谓的调度器扩展程序其实就是一个可配置的 Webhook 而已，里面包含 过滤器 和 优先级 两个端点，分别对应调度周期中的两个主要阶段（过滤和打分）。

第四种方法是通过调度框架（Scheduling Framework），Kubernetes v1.15 版本中引入了可插拔架构的调度框架，使得定制调度器这个任务变得更加的容易。调库框架向现有的调度器中添加了一组插件化的 API，该 API 在保持调度程序“核心”简单且易于维护的同时，使得大部分的调度功能以插件的形式存在，而且在我们现在的 v1.16 版本中上面的 调度器扩展程序 也已经被废弃了，所以以后调度框架才是自定义调度器的核心方式。

## 调度流程

## 自定义调度器

默认调度流程：
我们使用 kubeadm 搭建的集群，启动配置文件位于 /etc/kubernetes/manifests/kube-schdueler.yaml
watch apiserver，将 spec.nodeName 为空的 Pod 放入调度器内部的调度队列中
从调度队列中 Pop 出一个 Pod，开始一个标准的调度周期
从 Pod 属性中检索“硬性要求”（比如 CPU/内存请求值，nodeSelector/nodeAffinity），然后过滤阶段发生，在该阶段计算出满足要求的节点候选列表
从 Pod 属性中检索“软需求”，并应用一些默认的“软策略”（比如 Pod 倾向于在节点上更加聚拢或分散），最后，它为每个候选节点给出一个分数，并挑选出得分最高的最终获胜者
和 apiserver 通信（发送绑定调用），然后设置 Pod 的 spec.nodeName 属性以表示将该 Pod 调度到的节点。

sudo cat /etc/kubernetes/manifests/kube-schdueler.yaml

```
apiVersion: v1
kind: Pod
metadata:
  creationTimestamp: null
  labels:
    component: kube-scheduler
    tier: control-plane
  name: kube-scheduler
  namespace: kube-system
spec:
  containers:
  - command:
    - kube-scheduler
    - --authentication-kubeconfig=/etc/kubernetes/scheduler.conf
    - --authorization-kubeconfig=/etc/kubernetes/scheduler.conf
    - --bind-address=127.0.0.1
    - --kubeconfig=/etc/kubernetes/scheduler.conf
    - --leader-elect=true
    image: registry.aliyuncs.com/google_containers/kube-scheduler:v1.16.3
    imagePullPolicy: IfNotPresent
    livenessProbe:
      failureThreshold: 8
      httpGet:
        host: 127.0.0.1
        path: /healthz
        port: 10251
        scheme: HTTP
      initialDelaySeconds: 15
      timeoutSeconds: 15
    name: kube-scheduler
    resources:
      requests:
        cpu: 100m
    volumeMounts:
    - mountPath: /etc/kubernetes/scheduler.conf
      name: kubeconfig
      readOnly: true
  hostNetwork: true
  priorityClassName: system-cluster-critical
  volumes:
  - hostPath:
      path: /etc/kubernetes/scheduler.conf
      type: FileOrCreate
    name: kubeconfig
status: {}
```



see 
https://www.qikqiak.com/post/custom-kube-scheduler/
