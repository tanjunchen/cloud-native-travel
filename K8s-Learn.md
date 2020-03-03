# 容器相关概念

容器相关故事

为什么容器里只能跑“一个进程”？
为什么我原先一直在用的某个 JVM 参数，在容器里就不好使了？
为什么 Kubernetes 就不能固定 IP 地址？
容器网络连不通又该如何去 Debug？
Kubernetes 中 StatefulSet 和 Operator 到底什么区别？
PV 和 PVC 这些概念又该怎么用？

网络、存储、调度、操作系统、分布式原理

你是否曾经研发过类似 PaaS 的项目？你碰到过应用打包的问题吗，又是如何解决的呢？

而 Docker 项目之所以能取得如此高的关注，一方面正如前面我所说的那样，
它解决了应用打包和发布这一困扰运维人员多年的技术难题；
而另一方面，就是因为它第一次把一个纯后端的技术概念，通过非常友好的设计和封装，交到了最广大的开发者群体手里。

今天，我着重介绍了 Docker 项目在短时间内迅速崛起的三个重要原因：
Docker 镜像通过技术手段解决了 PaaS 的根本性问题；
Docker 容器同开发者之间有着与生俱来的密切关系；
PaaS 概念已经深入人心的完美契机。


# 容器技术入门

## 从进程说起

对于 Docker 等大多数 Linux 容器来说，Cgroups 技术是用来制造约束的主要手段，而 Namespace 技术则是用来修改进程视图的主要方法。

除了我们刚刚用到的 PID Namespace，Linux 操作系统还提供了 Mount、UTS、IPC、Network 和 User 这些 Namespace，用来对各种不同的进程上下文进行 “障眼法” 操作。
这就是 Linux 容器最基本的实现原理。

所以说，容器，其实是一种特殊的进程而已。

鉴于我对容器本质的讲解，你觉得上面这张容器和虚拟机对比图右侧关于容器的部分，怎么画才更精确？
你是否知道最新的 Docker 项目默认会为容器启用哪些 Namespace 吗？

其中 pid、net、ipc、mnt、uts 等命名空间将容器的进程、网络、消息、文件系统和 hostname 隔离开

进程隔离：pid namespace
网络隔离：net namespace
挂载点隔离：mount namespace
进程间通信隔离：ipc namespace
独立的用户、用户组：user namespace
独立的 hostname、domain name：uts namespace

Linux 的 3.12 内核支持 6 种 Namespace：

UTS: 主机名
IPC: 进程间通信
PID: "chroot" 进程树
NS: 挂载点，首次登陆 Linux
NET: 网络访问，包括接口
USER: 将本地的虚拟 user-id 映射到真实的 user-id

## 隔离与限制

容器基础之隔离与限制

Namespace 技术实际上修改了应用进程看待整个计算机“视图”，即它的“视线”被操作系统做了限制，只能“看到”某些指定的内容。

所以说，“敏捷”和“高性能”是容器相较于虚拟机最大的优势，也是它能够在 PaaS 这种更细粒度的资源管理平台上大行其道的重要原因。
不过，有利就有弊，基于 Linux Namespace 的隔离机制相比于虚拟化技术也有很多不足之处，其中最主要的问题就是：隔离得不彻底。

    首先，既然容器只是运行在宿主机上的一种特殊的进程，那么多个容器之间使用的就还是同一个宿主机的操作系统内核。
    其次，在 Linux 内核中，有很多资源和对象是不能被 Namespace 化的，最典型的例子就是：时间。

Linux Cgroups 就是 Linux 内核中用来为进程设置资源限制的一个重要功能。

Linux Cgroups 的全称是 Linux Control Group。它最主要的作用，就是限制一个进程组能够使用的资源上限，包括 CPU、内存、磁盘、网络带宽等等。

除 CPU 子系统外，Cgroups 的每一项子系统都有其独有的资源限制能力，比如：blkio，为​​​块​​​设​​​备​​​设​​​定​​​I/O 限​​​制，一般用于磁盘等设备；cpuset，为进程分配单独的 CPU 核和对应的内存节点；memory，为进程设定内存使用的限制。

Linux Cgroups 的设计还是比较易用的，简单粗暴地理解呢，它就是一个子系统目录加上一组资源限制文件的组合。

另外，跟 Namespace 的情况类似，Cgroups 对资源的限制能力也有很多不完善的地方，被提及最多的自然是 /proc 文件系统的问题。

你是否知道如何修复容器中的 top 指令以及 /proc 文件系统中的信息呢？（提示：lxcfs）
在从虚拟机向容器环境迁移应用的过程中，你还遇到哪些容器与虚拟机的不一致问题？

## 深入理解容器镜像

Mount Namespace 修改的，是容器进程对文件系统“挂载点”的认知。

这就是 Mount Namespace 跟其他 Namespace 的使用略有不同的地方：它对容器进程视图的改变，一定是伴随着挂载操作（mount）才能生效。

实际上，Mount Namespace 正是基于对 chroot 的不断改良才被发明出来的，它也是 Linux 操作系统里的第一个 Namespace。

而这个挂载在容器根目录上、用来为容器进程提供隔离后执行环境的文件系统，就是所谓的“容器镜像”。
它还有一个更为专业的名字，叫作：rootfs（根文件系统）。

现在，你应该可以理解，对 Docker 项目来说，它最核心的原理实际上就是为待创建的用户进程：
启用 Linux Namespace 配置；
设置指定的 Cgroups 参数；
切换进程的根目录（Change Root）

需要明确的是，rootfs 只是一个操作系统所包含的文件、配置和目录，并不包括操作系统内核。
在 Linux 操作系统中，这两部分是分开存放的，操作系统只有在开机启动时才会加载指定版本的内核镜像。

由于 rootfs 里打包的不只是应用，而是整个操作系统的文件和目录，也就意味着，应用以及它运行所需要的所有依赖，都被封装在了一起。

对一个应用来说，操作系统本身才是它运行所需要的最完整的“依赖库”。

Docker 在镜像的设计中，引入了层（layer）的概念。也就是说，用户制作镜像的每一步操作，都会生成一个层，也就是一个增量 rootfs。

当然，这个想法不是凭空臆造出来的，而是用到了一种叫作联合文件系统（Union File System）的能力。
Union File System 也叫 UnionFS，最主要的功能是将多个不同位置的目录联合挂载（union mount）到同一个目录下。

它是对 Linux 原生 UnionFS 的重写和改进；它的作者怨气好像很大。
我猜是 Linus Torvalds（Linux 之父）一直不让 AuFS 进入 Linux 内核主干的缘故，所以我们只能在 Ubuntu 和 Debian 这些发行版上使用它。

而且，从这个结构可以看出来，这个容器的 rootfs 由如下图所示的三部分组成：

第一部分，只读层。
第二部分，可读写层。
第三部分，Init 层。

通过结合使用 Mount Namespace 和 rootfs，容器就能够为进程构建出一个完善的文件系统隔离环境。

通过结合使用 Mount Namespace 和 rootfs，容器就能够为进程构建出一个完善的文件系统隔离环境。
当然，这个功能的实现还必须感谢 chroot 和 pivot_root 这两个系统调用切换进程根目录的能力。

容器镜像将会成为未来软件的主流发布方式。

既然容器的 rootfs（比如，Ubuntu 镜像），是以只读方式挂载的，那么又如何在容器里修改 Ubuntu 镜像的内容呢？（提示：Copy-on-Write）
除了 AuFS，你知道 Docker 项目还支持哪些 UnionFS 实现吗？你能说出不同宿主机环境下推荐使用哪种实现吗？

1. 上面的读写层通常也称为容器层，下面的只读层称为镜像层，所有的增删查改操作都只会作用在容器层，相同的文件上层会覆盖掉下层。知道这一点，就不难理解镜像文件的修改，比如修改一个文件的时候，首先会从上到下查找有没有这个文件，找到，就复制到容器层中，修改，修改的结果就会作用到下层的文件，这种方式也被称为copy-on-write。

2. 查了一下，包括但不限于以下这几种：aufs, device mapper, btrfs, overlayfs, vfs, zfs。aufs是ubuntu 常用的，device mapper 是 centos，btrfs 是 SUSE，overlayfs ubuntu 和 centos 都会使用，现在最新的 docker 版本中默认两个系统都是使用的 overlayfs，vfs 和 zfs 常用在 solaris 系统。

## 重新认识 Docker 容器

一、docker 镜像如何制作的两种方式是什么？
二、容器既然是一个封闭的进程，那么外接程序是如何进入容器这个进程的呢？
三、docker commit 对挂载点 volume 内容修改的影响是什么？
四、容器与宿主机如何进行文件读写？或 volume 是为了解决什么题？
五、Docker 的 copyData 功能是什么？解决了什么问题？
六、bind mount 机制是什么？
七、cgroup Namespace 的作用是什么？

在前面的三次分享中，我分别从 Linux Namespace 的隔离能力、Linux Cgroups 的限制能力，以及基于 rootfs 的文件系统三个角度，为你剖析了一个 Linux 容器的核心实现原理。

Dockerfile

```
# 使用官方提供的Python开发镜像作为基础镜像
FROM python:2.7-slim

# 将工作目录切换为/app
WORKDIR /app

# 将当前目录下的所有内容复制到/app下
ADD . /app

# 使用pip命令安装这个应用所需要的依赖
RUN pip install --trusted-host pypi.python.org -r requirements.txt

# 允许外界访问容器的80端口
EXPOSE 80

# 设置环境变量
ENV NAME World

# 设置容器进程为：python app.py，即：这个Python应用的启动命令
CMD ["python", "app.py"]
```

通过这个文件的内容，你可以看到 Dockerfile 的设计思想，是使用一些标准的原语（即大写高亮的词语），
描述我们所要构建的 Docker 镜像。并且这些原语，都是按顺序处理的。

这也就意味着：一个进程，可以选择加入到某个进程已有的 Namespace 当中，从而达到“进入”这个进程所在容器的目的，这正是 docker exec 的实现原理。


容器里进程新建的文件，怎么才能让宿主机获取到？
宿主机上的文件和目录，怎么才能让容器里的进程访问到？

这正是 Docker Volume 要解决的问题：Volume 机制，允许你将宿主机上指定的目录或者文件，挂载到容器里面进行读取和修改操作。

而这里要使用到的挂载技术，就是 Linux 的绑定挂载（bind mount）机制。它的主要作用就是，
允许你将一个目录或者文件，而不是整个设备，挂载到一个指定的目录上。
并且，这时你在该挂载点上进行的任何操作，只是发生在被挂载的目录或者文件上，
而原挂载点的内容则会被隐藏起来且不受影响。

所以，在一个正确的时机，进行一次绑定挂载，Docker 就可以成功地将一个宿主机上的目录或文件，不动声色地挂载到容器中。

你在查看 Docker 容器的 Namespace 时，是否注意到有一个叫 cgroup 的 Namespace？

它是 Linux 4.6 之后新增加的一个 Namespace，你知道它的作用吗？如果你执行 docker run -v /home:/test 的时候，容器镜像里的 /test 目录下本来就有内容的话，你会发现，在宿主机的 /home 目录下，也会出现这些内容。这是怎么回事？为什么它们没有被绑定挂载隐藏起来呢？（提示：Docker 的“copyData”功能）

请尝试给这个 Python 应用加上 CPU 和 Memory 限制，然后启动它。根据我们前面介绍的 Cgroups 的知识，请你查看一下这个容器的 Cgroups 文件系统的设置，是不是跟我前面的讲解一致。


## 谈谈 Kubernetes 的本质

一个“容器”，实际上是一个由 Linux Namespace、Linux Cgroups 和 rootfs 三种技术构建出来的进程的隔离环境。

CI/CD、监控、安全、网络、存储

在 Kubernetes 项目中，kubelet 主要负责同容器运行时（比如 Docker 项目）打交道。

此外，kubelet 还通过 gRPC 协议同一个叫作 Device Plugin 的插件进行交互。

而 kubelet 的另一个重要功能，则是调用网络插件和存储插件为容器配置网络和持久化存储。

这两个插件与 kubelet 进行交互的接口，分别是 CNI（Container Networking Interface）和 CSI（Container Storage Interface）。

运行在大规模集群中的各种任务之间，实际上存在着各种各样的关系。这些关系的处理，才是作业编排和管理系统最困难的地方。

Kubernetes 项目最主要的设计思想是，从更宏观的角度，以统一的方式来定义任务之间的各种关系，并且为将来支持更多种类的关系留有余地。

可是，我们知道，对于一个容器来说，它的 IP 地址等信息不是固定的，那么 Web 应用又怎么找到数据库容器的 Pod 呢？

所以，Kubernetes 项目的做法是给 Pod 绑定一个 Service 服务，
而 Service 服务声明的 IP 地址等信息是“终生不变”的。
这个 Service 服务的主要作用，就是作为 Pod 的代理入口（Portal），从而代替 Pod 对外暴露一个固定的网络地址。

除了应用与应用之间的关系外，应用运行的形态是影响“如何容器化这个应用”的第二个重要因素。

为此，Kubernetes 定义了新的、基于 Pod 改进后的对象。
比如 Job，用来描述一次性运行的 Pod（比如，大数据任务）；
再比如 DaemonSet，用来描述每个宿主机上必须且只能运行一个副本的守护进程服务；
又比如 CronJob，则用于描述定时任务等等。

相比之下，在 Kubernetes 项目中，我们所推崇的使用方法是：
首先，通过一个“编排对象”，比如 Pod、Job、CronJob 等，来描述你试图管理的应用；
然后，再为它定义一些“服务对象”，比如 Service、Secret、Horizontal Pod Autoscaler（自动水平扩展器）等。这些对象，会负责具体的平台级功能。

这种使用方法，就是所谓的“声明式 API”。这种 API 对应的“编排对象”和“服务对象”，都是 Kubernetes 项目中的 API 对象（API Object）。

Kubernetes 项目如何启动一个容器化任务呢？

然后，我重点介绍了 Kubernetes 项目的架构，详细讲解了它如何使用“声明式 API”来描述容器化业务和容器间关系的设计思想。
实际上，过去很多的集群管理项目（比如 Yarn、Mesos，以及 Swarm）所擅长的，都是把一个容器，按照某种规则，放置在某个最佳节点上运行起来。这种功能，我们称为“调度”。

而 Kubernetes 项目所擅长的，是按照用户的意愿和整个系统的规则，完全自动化地处理好容器之间的各种关系。这种功能，就是我们经常听到的一个概念：编排。

这今天的分享中，我介绍了 Kubernetes 项目的架构。
你是否了解了 Docker Swarm（SwarmKit 项目）跟 Kubernetes 在架构上和使用方法上的异同呢？
在 Kubernetes 之前，很多项目都没办法管理“有状态”的容器，即，不能从一台宿主机“迁移”到另一台宿主机上的容器。你是否能列举出，阻止这种“迁移”的原因都有哪些呢？

# Kubernetes 搭建集群

## Kubernetes 一键部署利器：kubeadm

SaltStack、Ansible 等运维工具自动化安装 Kubernetes 集群。
把 kubelet 直接运行在宿主机上，然后使用容器部署其他的 Kubernetes 组件。

kubeadm init 的工作流程

当你执行 kubeadm init 指令后，kubeadm 首先要做的，是一系列的检查工作，以确定这台机器可以用来部署 Kubernetes。
这一步检查，我们称为“Preflight Checks”，它可以为你省掉很多后续的麻烦。

其实，Preflight Checks 包括了很多方面，比如：
Linux 内核的版本必须是否是 3.10 以上？
Linux Cgroups 模块是否可用？
机器的 hostname 是否标准？
在 Kubernetes 项目里，机器的名字以及一切存储在 Etcd 中的 API 对象，都必须使用标准的 DNS 命名（RFC 1123）。
用户安装的 kubeadm 和 kubelet 的版本是否匹配？
机器上是不是已经安装了 Kubernetes 的二进制文件？
Kubernetes 的工作端口 10250/10251/10252 端口是不是已经被占用？
ip、mount 等 Linux 指令是否存在？
Docker 是否已经安装？

在通过了 Preflight Checks 之后，kubeadm 要为你做的，是生成 Kubernetes 对外提供服务所需的各种证书和对应的目录。
证书生成后，kubeadm 接下来会为其他组件生成访问 kube-apiserver 所需的配置文件。
接下来，kubeadm 会为 Master 组件生成 Pod 配置文件。
然后，kubeadm 就会为集群生成一个 bootstrap token。
在 token 生成之后，kubeadm 会将 ca.crt 等 Master 节点的重要信息，通过 ConfigMap 的方式保存在 Etcd 当中，供后续部署 Node 节点使用。
kubeadm init 的最后一步，就是安装默认插件。

kubeadm join 的工作流程

这个流程其实非常简单，kubeadm init 生成 bootstrap token 之后，
你就可以在任意一台安装了 kubelet 和 kubeadm 的机器上执行 kubeadm join 了。

如果你有部署规模化生产环境的需求，我推荐使用 kops 或者 SaltStack 这样更复杂的部署工具。

在 Linux 上为一个类似 kube-apiserver 的 Web Server 制作证书，你知道可以用哪些工具实现吗？
回忆一下我在前面文章中分享的 Kubernetes 架构，你能够说出 Kubernetes 各个功能组件之间（包含 Etcd），
都有哪些建立连接或者调用的方式吗？（比如：HTTP/HTTPS，远程调用等等）



# 容器编排与 Kubernetes 作业管理

## 为什么需要 Pod

“Namespace 做隔离，Cgroups 做限制，rootfs 做文件系统”这样的“三句箴言”可以朗朗上口了，为什么 Kubernetes 项目又突然搞出一个 Pod 来呢？

Pod 在 Kubernetes 项目里还有更重要的意义，那就是：容器设计模式。

首先，关于 Pod 最重要的一个事实是：它只是一个逻辑概念。

Pod 里的所有容器，共享的是同一个 Network Namespace，并且可以声明共享同一个 Volume。

所以，在 Kubernetes 项目里，Pod 的实现需要使用一个中间容器，这个容器叫作 Infra 容器。

它们可以直接使用 localhost 进行通信；
它们看到的网络设备跟 Infra 容器看到的完全一样；
一个 Pod 只有一个 IP 地址，也就是这个 Pod 的 Network Namespace 对应的 IP 地址；
当然，其他的所有网络资源，都是一个 Pod 一份，并且被该 Pod 中的所有容器共享；
Pod 的生命周期只跟 Infra 容器一致，而与容器 A 和 B 无关。

第一个最典型的例子是：WAR 包与 Web 服务器。

像这样，我们就用一种“组合”方式，解决了 WAR 包与 Tomcat 容器之间耦合关系的问题。
实际上，这个所谓的“组合”操作，正是容器设计模式里最常用的一种模式，它的名字叫：sidecar。

第二个例子，则是容器的日志收集。

但不要忘记，Pod 的另一个重要特性是，它的所有容器都共享同一个 Network Namespace。
这就使得很多与 Pod 网络相关的配置和管理，也都可以交给 sidecar 完成，而完全无须干涉用户容器。
这里最典型的例子莫过于 Istio 这个微服务治理项目了。

但实际上，无论是从具体的实现原理，还是从使用方法、特性、功能等方面，容器与虚拟机几乎没有任何相似的地方；
也不存在一种普遍的方法，能够把虚拟机里的应用无缝迁移到容器中。
因为，容器的性能优势，必然伴随着相应缺陷，即：它不能像虚拟机那样，完全模拟本地物理机环境中的部署方法。

除了 Network Namespace 外，Pod 里的容器还可以共享哪些 Namespace 呢？你能说出共享这些 Namesapce 的具体应用场景吗？

## 深入解析 Pod 对象（一）：基本概念

凡是调度、网络、存储，以及安全相关的属性，基本上是 Pod 级别的。

NodeSelector：是一个供用户将 Pod 与 Node 进行绑定的字段
NodeName：一旦 Pod 的这个字段被赋值，Kubernetes 项目就会被认为这个 Pod 已经经过了调度，调度的结果就是赋值的节点名字。
HostAliases：定义了 Pod 的 hosts 文件（比如 /etc/hosts）里的内容
凡是 Pod 中的容器要共享宿主机的 Namespace，也一定是 Pod 级别的定义

Kubernetes 项目中对 Container 的定义，和 Docker 相比并没有什么太大区别。我在前面的容器技术概念入门系列文章中，和你分享的 Image（镜像）、Command（启动命令）、workingDir（容器的工作目录）、Ports（容器要开发的端口），以及 volumeMounts（容器要挂载的 Volume）都是构成 Kubernetes 项目中 Container 的主要字段。不过在这里，还有这么几个属性值得你额外关注。

首先，是 ImagePullPolicy 字段
其次，是 Lifecycle 字段。
Pod 生命周期的变化，主要体现在 Pod API 对象的 Status 部分，这是它除了 Metadata 和 Spec 之外的第三个重要字段。其中，pod.status.phase，就是 Pod 的当前状态，它有如下几种可能的情况：

Pending。这个状态意味着，Pod 的 YAML 文件已经提交给了 Kubernetes，API 对象已经被创建并保存在 Etcd 当中。但是，这个 Pod 里有些容器因为某种原因而不能被顺利创建。比如，调度不成功。
Running。这个状态下，Pod 已经调度成功，跟一个具体的节点绑定。它包含的容器都已经创建成功，并且至少有一个正在运行中。
Succeeded。这个状态意味着，Pod 里的所有容器都正常运行完毕，并且已经退出了。这种情况在运行一次性任务时最为常见。
Failed。这个状态下，Pod 里至少有一个容器以不正常的状态（非 0 的返回码）退出。这个状态的出现，意味着你得想办法 Debug 这个容器的应用，比如查看 Pod 的 Events 和日志。
Unknown。这是一个异常状态，意味着 Pod 的状态不能持续地被 kubelet 汇报给 kube-apiserver，这很有可能是主从节点（Master 和 Kubelet）间的通信出现了问题。

你能否举出一些 Pod（即容器）的状态是 Running，但是应用其实已经停止服务的例子？相信 Java Web 开发者的亲身体会会比较多吧。

1. 程序本身有 bug，本来应该返回 200，但因为代码问题，返回的是500；
2. 程序因为内存问题，已经僵死，但进程还在，但无响应；
3. Dockerfile 写的不规范，应用程序不是主进程，那么主进程出了什么问题都无法发现；
4. 程序出现死循环。

## 深入解析 Pod 对象（二）：使用进阶

这种特殊的 Volume，叫作 Projected Volume，你可以把它翻译为“投射数据卷”。
备注：Projected Volume 是 Kubernetes v1.11 之后的新特性这是什么意思呢？

到目前为止，Kubernetes 支持的 Projected Volume 一共有四种：
Secret；
ConfigMap；
Downward API；
ServiceAccountToken。

需要注意的是，这个更新可能会有一定的延时。所以在编写应用程序时，在发起数据库连接的代码处写好重试和超时的逻辑，绝对是个好习惯。

不过，需要注意的是，Downward API 能够获取到的信息，一定是 Pod 里的容器进程启动之前就能够确定下来的信息。
而如果你想要获取 Pod 容器运行后才会出现的信息，比如，容器进程的 PID，
那就肯定不能使用 Downward API 了，而应该考虑在 Pod 里定义一个 sidecar 容器。

其实，Secret、ConfigMap，以及 Downward API 这三种 Projected Volume 定义的信息，
大多还可以通过环境变量的方式出现在容器里。但是，通过环境变量获取这些信息的方式，
不具备自动更新的能力。所以，一般情况下，我都建议你使用 Volume 文件的方式获取这些信息。

在明白了 Secret 之后，我再为你讲解 Pod 中一个与它密切相关的概念：Service Account。

这种把 Kubernetes 客户端以容器的方式运行在集群里，
然后使用 default Service Account 自动授权的方式，被称作“InClusterConfig”，
也是我最推荐的进行 Kubernetes API 编程的授权方式。

接下来，我们再来看 Pod 的另一个重要的配置：容器健康检查和恢复机制。

在 Kubernetes 中，你可以为 Pod 里的容器定义一个健康检查“探针”（Probe）。
这样，kubelet 就会根据这个 Probe 的返回值决定这个容器的状态，
而不是直接以容器进行是否运行（来自 Docker 返回的信息）作为依据。
这种机制，是生产环境中保证应用健康存活的重要手段。

这个功能就是 Kubernetes 里的 Pod 恢复机制，也叫 restartPolicy。
它是 Pod 的 Spec 部分的一个标准字段（pod.spec.restartPolicy），
默认值是 Always，即：任何时候这个容器发生了异常，它一定会被重新创建。

而作为用户，你还可以通过设置 restartPolicy，改变 Pod 的恢复策略。
除了 Always，
它还有 OnFailure 和 Never 两种情况：
Always：在任何情况下，只要容器不在运行状态，就自动重启容器；
OnFailure: 只在容器 异常时才自动重启容器；
Never: 从来不重启容器。

只要 Pod 的 restartPolicy 指定的策略允许重启异常的容器（比如：Always），那么这个 Pod 就会保持 Running 状态，并进行容器重启。否则，Pod 就会进入 Failed 状态 。对于包含多个容器的 Pod，只有它里面所有的容器都进入异常状态后，Pod 才会进入 Failed 状态。在此之前，Pod 都是 Running 状态。此时，Pod 的 READY 字段会显示正常容器的个数，比如：

需要说明的是，PodPreset 里定义的内容，只会在 Pod API 对象被创建之前追加在这个对象本身上，而不会影响任何 Pod 的控制器的定义。

在没有 Kubernetes 的时候，你是通过什么方法进行应用的健康检查的？Kubernetes 的 livenessProbe 和 readinessProbe 提供的几种探测机制，是否能满足你的需求？

## 编排其实很简单：谈谈“控制器”模型

实际上，你可能已经有所感悟：Pod 这个看似复杂的 API 对象，实际上就是对容器的进一步抽象和封装而已。

比如，现在有一种待编排的对象 X，它有一个对应的控制器。那么，我就可以用一段 Go 语言风格的伪代码，为你描述这个控制循环：

在具体实现中，实际状态往往来自于 Kubernetes 集群本身。

而期望状态，一般来自于用户提交的 YAML 文件。

比如，增加 Pod，删除已有的 Pod，或者更新 Pod 的某个字段。这也是 Kubernetes 项目“面向 API 对象编程”的一个直观体现。

如上图所示，类似 Deployment 这样的一个控制器，实际上都是由上半部分的控制器定义（包括期望状态），加上下半部分的被控制对象的模板组成的。

你能否说出，Kubernetes 使用的这个“控制器模式”，跟我们平常所说的“事件驱动”，有什么区别和联系吗？

## 经典PaaS的记忆：作业副本与水平扩展

在上一篇文章中，我为你详细介绍了 Kubernetes 项目中第一个重要的设计思想：控制器模式。
而在今天这篇文章中，我就来为你详细讲解一下，Kubernetes 里第一个控制器模式的完整实现：Deployment。

Deployment 看似简单，但实际上，它实现了 Kubernetes 项目中一个非常重要的功能：
Pod 的“水平扩展 / 收缩”（horizontal scaling out/in）。这个功能，是从 PaaS 时代开始，一个平台级项目就必须具备的编排能力。

从这个 YAML 文件中，我们可以看到，一个 ReplicaSet 对象，其实就是由副本数目的定义和一个 Pod 模板组成的。不难发现，它的定义其实是 Deployment 的一个子集。更重要的是，Deployment 控制器实际操纵的，正是这样的 ReplicaSet 对象，而不是 Pod 对象。

首先，我需要使用 kubectl rollout history 命令，查看每次 Deployment 变更对应的版本。而由于我们在创建这个 Deployment 的时候，指定了–record 参数，所以我们创建这些版本时执行的 kubectl 命令，都会被记录下来。

然后，我们就可以在 kubectl rollout undo 命令行最后，加上要回滚到的指定版本的版本号，就可以回滚到指定版本了。

通过这些讲解，你应该了解到：Deployment 实际上是一个两层控制器。
首先，它通过 ReplicaSet 的个数来描述应用的版本；
然后，它再通过 ReplicaSet 的属性（比如 replicas 的值），来保证 Pod 的副本数量。

备注：Deployment 控制 ReplicaSet（版本），ReplicaSet 控制 Pod（副本数）。这个两层控制关系一定要牢记。

可是，在实际使用场景中，应用发布的流程往往千差万别，也可能有很多的定制化需求。
比如，我的应用可能有会话黏连（session sticky），这就意味着“滚动更新”的时候，哪个 Pod 能下线，是不能随便选择的。
这种场景，光靠 Deployment 自己就很难应对了。对于这种需求，我在专栏后续文章中重点介绍的“自定义控制器”，就可以帮我们实现一个功能更加强大的 Deployment Controller。
当然，Kubernetes 项目本身，也提供了另外一种抽象方式，帮我们应对其他一些用 Deployment 无法处理的应用编排场景。这个设计，就是对有状态应用的管理。

你听说过金丝雀发布（Canary Deployment）和蓝绿发布（Blue-Green Deployment）吗？你能说出它们是什么意思吗？实际上，有了 Deployment 的能力之后，你可以非常轻松地用它来实现金丝雀发布、蓝绿发布，以及 A/B 测试等很多应用发布模式。

金丝雀部署：优先发布一台或少量机器升级，等验证无误后再更新其他机器。优点是用户影响范围小，不足之处是要额外控制如何做自动更新。
蓝绿部署：2组机器，蓝代表当前的V1版本，绿代表已经升级完成的V2版本。通过LB将流量全部导入V2完成升级部署。优点是切换快速，缺点是影响全部用户。
本文学习的滚动更新，我觉得就是一个自动化更新的金丝雀发布。

实践地址
https://github.com/ContainerSolutions/k8s-deployment-strategies/tree/master/canary

## 深入理解StatefulSet（一）：拓扑状态

但是，在实际的场景中，并不是所有的应用都可以满足这样的要求。
尤其是分布式应用，它的多个实例之间，往往有依赖关系，比如：主从关系、主备关系。

还有就是数据存储类应用，它的多个实例，往往都会在本地磁盘上保存一份数据。
而这些实例一旦被杀掉，即便重建出来，实例与数据之间的对应关系也已经丢失，从而导致应用失败。

得益于“控制器模式”的设计思想，Kubernetes 项目很早就在 Deployment 的基础上，扩展出了对“有状态应用”的初步支持。
这个编排功能，就是：StatefulSet。

StatefulSet 的设计其实非常容易理解。它把真实世界里的应用状态，抽象为了两种情况：

拓扑状态。这种情况意味着，应用的多个实例之间不是完全对等的关系。这些应用实例，必须按照某些顺序启动，比如应用的主节点 A 要先于从节点 B 启动。而如果你把 A 和 B 两个 Pod 删除掉，它们再次被创建出来时也必须严格按照这个顺序才行。并且，新创建出来的 Pod，必须和原来 Pod 的网络标识一样，这样原先的访问者才能使用同样的方法，访问到这个新 Pod。
存储状态。这种情况意味着，应用的多个实例分别绑定了不同的存储数据。对于这些应用实例来说，Pod A 第一次读取到的数据，和隔了十分钟之后再次读取到的数据，应该是同一份，哪怕在此期间 Pod A 被重新创建过。这种情况最典型的例子，就是一个数据库应用的多个存储实例。

所以，StatefulSet 的核心功能，就是通过某种方式记录这些状态，然后在 Pod 被重新创建时，能够为新 Pod 恢复这些状态。

那么，这个 Service 又是如何被访问的呢？

第一种方式，是以 Service 的 VIP（Virtual IP，即：虚拟 IP）方式。比如：当我访问 10.0.23.1 这个 Service 的 IP 地址时，10.0.23.1 其实就是一个 VIP，它会把请求转发到该 Service 所代理的某一个 Pod 上。这里的具体原理，我会在后续的 Service 章节中进行详细介绍。

第二种方式，就是以 Service 的 DNS 方式。比如：这时候，只要我访问“my-svc.my-namespace.svc.cluster.local”这条 DNS 记录，就可以访问到名叫 my-svc 的 Service 所代理的某一个 Pod。

而在第二种 Service DNS 的方式下，具体还可以分为两种处理方法：第一种处理方法，是 Normal Service。这种情况下，你访问“my-svc.my-namespace.svc.cluster.local”解析到的，正是 my-svc 这个 Service 的 VIP，后面的流程就跟 VIP 方式一致了。
而第二种处理方法，正是 Headless Service。这种情况下，你访问“my-svc.my-namespace.svc.cluster.local”解析到的，直接就是 my-svc 代理的某一个 Pod 的 IP 地址。可以看到，这里的区别在于，Headless Service 不需要分配一个 VIP，而是可以直接以 DNS 记录的方式解析出被代理 Pod 的 IP 地址。

编写一个 StatefulSet 的 YAML 文件

apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: web
spec:
  serviceName: "nginx"
  replicas: 2
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.9.1
        ports:
        - containerPort: 80
          name: web

这个 YAML 文件，和我们在前面文章中用到的 nginx-deployment 的唯一区别，就是多了一个 serviceName=nginx 字段。

通过这种方法，Kubernetes 就成功地将 Pod 的拓扑状态（比如：哪个节点先启动，哪个节点后启动），按照 Pod 的“名字 + 编号”的方式固定了下来。此外，Kubernetes 还为每一个 Pod 提供了一个固定并且唯一的访问入口，即：这个 Pod 对应的 DNS 记录。

不过，相信你也已经注意到了，尽管 web-0.nginx 这条记录本身不会变，但它解析到的 Pod 的 IP 地址，并不是固定的。这就意味着，对于“有状态应用”实例的访问，你必须使用 DNS 记录或者 hostname 的方式，而绝不应该直接访问这些 Pod 的 IP 地址。

StatefulSet 这个控制器的主要作用之一，就是使用 Pod 模板创建 Pod 的时候，对它们进行编号，
并且按照编号顺序逐一完成创建工作。而当 StatefulSet 的“控制循环”发现 Pod 的“实际状态”与“期望状态”不一致，
需要新建或者删除 Pod 进行“调谐”的时候，它会严格按照这些 Pod 编号的顺序，逐一完成这些操作。


你曾经运维过哪些有拓扑状态的应用呢（比如：主从、主主、主备、一主多从等结构）？
你觉得这些应用实例之间的拓扑关系，能否借助这种为 Pod 实例编号的方式表达出来呢？
如果不能，你觉得 Kubernetes 还应该为你提供哪些支持来管理这个拓扑状态呢？


## 深入理解StatefulSet（二）：存储状态

如果你并不知道有哪些 Volume 类型可以用，要怎么办呢？

Kubernetes 项目引入了一组叫作 Persistent Volume Claim（PVC）和 Persistent Volume（PV）的 API 对象，
大大降低了用户声明和使用持久化 Volume 的门槛。

举个例子，有了 PVC 之后，一个开发人员想要使用一个 Volume，只需要简单的两步即可。

第一步：定义一个 PVC，声明想要的 Volume 的属性：

第二步：在应用的 Pod 中，声明使用这个 PVC：

所以，Kubernetes 中 PVC 和 PV 的设计，实际上类似于“接口”和“实现”的思想。
开发者只要知道并会使用“接口”，即：PVC；而运维人员则负责给“接口”绑定具体的实现，即：PV。

通过这种方式，Kubernetes 的 StatefulSet 就实现了对应用存储状态的管理。

首先，StatefulSet 的控制器直接管理的是 Pod。这是因为，StatefulSet 里的不同 Pod 实例，不再像 ReplicaSet 中那样都是完全一样的，而是有了细微区别的。比如，每个 Pod 的 hostname、名字等都是不同的、携带了编号的。而 StatefulSet 区分这些实例的方式，就是通过在 Pod 的名字里加上事先约定好的编号。

其次，Kubernetes 通过 Headless Service，为这些有编号的 Pod，在 DNS 服务器中生成带有同样编号的 DNS 记录。只要 StatefulSet 能够保证这些 Pod 名字里的编号不变，那么 Service 里类似于 web-0.nginx.default.svc.cluster.local 这样的 DNS 记录也就不会变，而这条记录解析出来的 Pod 的 IP 地址，则会随着后端 Pod 的删除和再创建而自动更新。这当然是 Service 机制本身的能力，不需要 StatefulSet 操心。

最后，StatefulSet 还为每一个 Pod 分配并创建一个同样编号的 PVC。这样，Kubernetes 就可以通过 Persistent Volume 机制为这个 PVC 绑定上对应的 PV，从而保证了每一个 Pod 都拥有一个独立的 Volume。在这种情况下，即使 Pod 被删除，它所对应的 PVC 和 PV 依然会保留下来。所以当这个 Pod 被重新创建出来之后，Kubernetes 会为它找到同样编号的 PVC，挂载这个 PVC 对应的 Volume，从而获取到以前保存在 Volume 里的数据。

这么一看，原本非常复杂的 StatefulSet，是不是也很容易理解了呢？

从这些讲述中，我们不难看出 StatefulSet 的设计思想：StatefulSet 其实就是一种特殊的 Deployment，而其独特之处在于，它的每个 Pod 都被编号了。而且，这个编号会体现在 Pod 的名字和 hostname 等标识信息上，这不仅代表了 Pod 的创建顺序，也是 Pod 的重要网络标识（即：在整个集群里唯一的、可被的访问身份）。有了这个编号后，StatefulSet 就使用 Kubernetes 里的两个标准功能：Headless Service 和 PV/PVC，实现了对 Pod 的拓扑状态和存储状态的维护。


在实际场景中，有一些分布式应用的集群是这么工作的：当一个新节点加入到集群时，或者老节点被迁移后重建时，这个节点可以从主节点或者其他从节点那里同步到自己所需要的数据。在这种情况下，你认为是否还有必要将这个节点 Pod 与它的 PV 进行一对一绑定呢？（提示：这个问题的答案根据不同的项目是不同的。关键在于，重建后的节点进行数据恢复和同步的时候，是不是一定需要原先它写在本地磁盘里的数据）

## 深入理解StatefulSet（三）：有状态应用实践


部署的“有状态应用”

是一个“主从复制”（Maser-Slave Replication）的 MySQL 集群；
有 1 个主节点（Master）；有多个从节点（Slave）；
从节点需要能水平扩展；
所有的写操作，只能在主节点上执行；
读操作可以在所有节点上执行。
这就是一个非常典型的主从模式的 MySQL 集群

所以，在安装好 MySQL 的 Master 节点之后，你需要做的第一步工作，就是通过 XtraBackup 将 Master 节点的数据备份到指定目录。

第二步：配置 Slave 节点。

第三步，启动 Slave 节点。

第四步，在这个集群中添加更多的 Slave 节点。

通过上面的叙述，我们不难看到，将部署 MySQL 集群的流程迁移到 Kubernetes 项目上，需要能够“容器化”地解决下面的“三座大山”：
Master 节点和 Slave 节点需要有不同的配置文件（即：不同的 my.cnf）；
Master 节点和 Slave 节点需要能够传输备份信息文件；
在 Slave 节点第一次启动之前，需要执行一些初始化 SQL 操作；

而由于 MySQL 本身同时拥有拓扑状态（主从节点的区别）和存储状态（MySQL 保存在本地的数据），我们自然要通过 StatefulSet 来解决这“三座大山”的问题。

其中，“第一座大山：Master 节点和 Slave 节点需要有不同的配置文件”，很容易处理：我们只需要给主从节点分别准备两份不同的 MySQL 配置文件，然后根据 Pod 的序号（Index）挂载进去即可。

接下来，我们再一起解决“第二座大山：Master 节点和 Slave 节点需要能够传输备份文件”的问题。翻越这座大山的思路，我比较推荐的做法是：先搭建框架，再完善细节。其中，Pod 部分如何定义，是完善细节时的重点。

然后，我们来重点设计一下这个 StatefulSet 的 Pod 模板，也就是 template 字段。

第一步：从 ConfigMap 中，获取 MySQL 的 Pod 对应的配置文件。

第二步：在 Slave Pod 启动前，从 Master 或者其他 Slave Pod 里拷贝数据库数据到自己的目录下。

这就是我们需要解决的“第三座大山”的问题，即：如何在 Slave 节点的 MySQL 容器第一次启动之前，执行初始化 SQL。


如果我们现在的需求是：所有的读请求，只由 Slave 节点处理；所有的写请求，只由 Master 节点处理。那么，你需要在今天这篇文章的基础上再做哪些改动呢？


## 容器化守护进程的意义：DaemonSet

DaemonSet

这个 Pod 运行在 Kubernetes 集群里的每一个节点（Node）上；每个节点上只有一个这样的 Pod 实例；当有新的节点加入 Kubernetes 集群后，该 Pod 会自动地在新节点上被创建出来；而当旧节点被删除后，它上面的 Pod 也相应地会被回收掉。

这个机制听起来很简单，但 Daemon Pod 的意义确实是非常重要的。
我随便给你列举几个例子：各种网络插件的 Agent 组件，都必须运行在每一个节点上，用来处理这个节点上的容器网络；
各种存储插件的 Agent 组件，也必须运行在每一个节点上，用来在这个节点上挂载远程存储目录，操作容器的 Volume 目录；
各种监控组件和日志组件，也必须运行在每一个节点上，负责这个节点上的监控信息和日志搜集。

那么，DaemonSet 又是如何保证每个 Node 上有且只有一个被管理的 Pod 呢？

没有这种 Pod，那么就意味着要在这个 Node 上创建这样一个 Pod；
有这种 Pod，但是数量大于 1，那就说明要把多余的 Pod 从这个 Node 上删除掉；
正好只有一个这种 Pod，那说明这个节点是正常的。

但是，如何在指定的 Node 上创建新 Pod 呢？

所以，我们的 DaemonSet Controller 会在创建 Pod 的时候，自动在这个 Pod 的 API 对象里，
加上这样一个 nodeAffinity 定义。其中，需要绑定的节点名字，正是当前正在遍历的这个 Node。

而通过这样一个 Toleration，调度器在调度这个 Pod 的时候，就会忽略当前节点上的“污点”，
从而成功地将网络插件的 Agent 组件调度到这台机器上启动起来。

至此，通过上面这些内容，你应该能够明白，DaemonSet 其实是一个非常简单的控制器。
在它的控制循环中，只需要遍历所有节点，然后根据节点上是否有被管理 Pod 的情况，来决定是否要创建或者删除一个 Pod。

相比于 Deployment，DaemonSet 只管理 Pod 对象，然后通过 nodeAffinity 和 Toleration 这两个调度器的小功能，
保证了每个节点上有且只有一个 Pod。这个控制器的实现原理简单易懂，希望你能够快速掌握。
与此同时，DaemonSet 使用 ControllerRevision，来保存和管理自己对应的“版本”。
这种“面向 API 对象”的设计思路，大大简化了控制器本身的逻辑，
也正是 Kubernetes 项目“声明式 API”的优势所在。


我在文中提到，在 Kubernetes v1.11 之前，DaemonSet 所管理的 Pod 的调度过程，
实际上都是由 DaemonSet Controller 自己而不是由调度器完成的。你能说出这其中有哪些原因吗？

## 撬动离线业务：Job与CronJob

实际上，它们主要编排的对象，都是“在线业务”，即：Long Running Task（长作业）。
比如，我在前面举例时常用的 Nginx、Tomcat，以及 MySQL 等等。这些应用一旦运行起来，
除非出错或者停止，它的容器进程会一直保持在 Running 状态。但是，有一类作业显然不满足这样的条件，
这就是“离线业务”，或者叫作 Batch Job（计算业务）。这种业务在计算完成后就直接退出了，
而此时如果你依然用 Deployment 来管理这种业务的话，就会发现 Pod 会在计算结束后退出，
然后被 Deployment Controller 不断地重启；而像“滚动更新”这样的编排功能，更无从谈起了。

接下来，我再和你分享三种常用的、使用 Job 对象的方法。第一种用法，也是最简单粗暴的用法：外部管理器 +Job 模板。

这种模式最典型的应用，就是 TensorFlow 社区的 KubeFlow 项目。

第二种用法：拥有固定任务数目的并行 Job。

第三种用法，也是很常用的一个用法：指定并行度（parallelism），但不设置固定的 completions 的值。

最后，我再来和你分享一个非常有用的 Job 对象，叫作：CronJob。定时任务。

在今天这篇文章中，我主要和你分享了 Job 这个离线业务的编排方法，
讲解了 completions 和 parallelism 字段的含义，以及 Job Controller 的执行原理。
紧接着，我通过实例和你分享了 Job 对象三种常见的使用方法。但是，根据我在社区和生产环境中的经验，
大多数情况下用户还是更倾向于自己控制 Job 对象。所以，相比于这些固定的“模式”，掌握 Job 的 API 对象，
和它各个字段的准确含义会更加重要。最后，我还介绍了一种 Job 的控制器，叫作：CronJob。
这也印证了我在前面的分享中所说的：用一个对象控制另一个对象，是 Kubernetes 编排的精髓所在。

根据 Job 控制器的工作原理，如果你定义的 parallelism 比 completions 还大的话，
比如： parallelism: 4 completions: 2那么，这个 Job 最开始创建的时候，会同时启动几个 Pod 呢？原因是什么？

## 声明式 API 与 Kubernetes 编程范式

像这样的两条命令，就是用 Docker Swarm 启动了两个 Nginx 容器实例。
其中，第一条 create 命令创建了这两个容器，
而第二条 update 命令则把它们“滚动更新”为了一个新的镜像。
对于这种使用方式，我们称为命令式命令行操作。

对于上面这种先 kubectl create，再 replace 的操作，我们称为命令式配置文件操作。

那么，到底什么才是“声明式 API”呢？答案是，kubectl apply 命令。

在上面这个架构图中，我们不难看到 Istio 项目架构的核心所在。
Istio 最根本的组件，是运行在每一个应用 Pod 里的 Envoy 容器。

实际上，Istio 项目使用的，是 Kubernetes 中的一个非常重要的功能，叫作 Dynamic Admission Control。

首先，Istio 会将这个 Envoy 容器本身的定义，以 ConfigMap 的方式保存在 Kubernetes 当中。

接下来，Istio 将一个编写好的 Initializer，作为一个 Pod 部署在 Kubernetes 中。

这也就意味着，当你在 Initializer 里完成了要做的操作后，
一定要记得将这个 metadata.initializers.pending 标志清除掉。
这一点，你在编写 Initializer 代码的时候一定要非常注意。

相信你此时已经明白，Istio 项目的核心，就是由无数个运行在应用 Pod 中的 Envoy 容器组成的服务代理网格。
这也正是 Service Mesh 的含义。

而这个机制得以实现的原理，正是借助了 Kubernetes 能够对 API 对象进行在线更新的能力，
这也正是 Kubernetes“声明式 API”的独特之处：首先，所谓“声明式”，
指的就是我只需要提交一个定义好的 API 对象来“声明”，我所期望的状态是什么样子。
其次，“声明式 API”允许有多个 API 写端，以 PATCH 的方式对 API 对象进行修改，
而无需关心本地原始 YAML 文件的内容。最后，也是最重要的，
有了上述两个能力，Kubernetes 项目才可以基于对 API 对象的增、删、改、查，
在完全无需外界干预的情况下，完成对“实际状态”和“期望状态”的调谐（Reconcile）过程。

你是否对 Envoy 项目做过了解？你觉得为什么它能够击败 Nginx 以及 HAProxy 等竞品，成为 Service Mesh 体系的核心？

## 深入解析声明式API（一）：API对象的奥秘

那么，Kubernetes 是如何对 Resource、Group 和 Version 进行解析，从而在 Kubernetes 项目里找到 CronJob 对象的定义呢？

首先，Kubernetes 会匹配 API 对象的组。

然后，Kubernetes 会进一步匹配到 API 对象的版本号。

最后，Kubernetes 会匹配 API 对象的资源类型。

这时候，APIServer 就可以继续创建这个 CronJob 对象了。为了方便理解，我为你总结了一个如下所示流程图来阐述这个创建过程：

首先，当我们发起了创建 CronJob 的 POST 请求之后，我们编写的 YAML 的信息就被提交给了 APIServer。

而 APIServer 的第一个功能，就是过滤这个请求，并完成一些前置性的工作，比如授权、超时处理、审计等。
然后，请求会进入 MUX 和 Routes 流程。如果你编写过 Web Server 的话就会知道，MUX 和 Routes 是 APIServer 完成 URL 和 Handler 绑定的场所。而 APIServer 的 Handler 要做的事情，就是按照我刚刚介绍的匹配过程，找到对应的 CronJob 类型定义。
接着，APIServer 最重要的职责就来了：根据这个 CronJob 类型定义，使用用户提交的 YAML 文件里的字段，创建一个 CronJob 对象。而在这个过程中，APIServer 会进行一个 Convert 工作，即：把用户提交的 YAML 文件，转换成一个叫作 Super Version 的对象，它正是该 API 资源类型所有版本的字段全集。这样用户提交的不同版本的 YAML 文件，就都可以用这个 Super Version 对象来进行处理了。
接下来，APIServer 会先后进行 Admission() 和 Validation() 操作。比如，我在上一篇文章中提到的 Admission Controller 和 Initializer，就都属于 Admission 的内容。而 Validation，则负责验证这个对象里的各个字段是否合法。这个被验证过的 API 对象，都保存在了 APIServer 里一个叫作 Registry 的数据结构中。也就是说，只要一个 API 对象的定义能在 Registry 里查到，它就是一个有效的 Kubernetes API 对象。
最后，APIServer 会把验证过的 API 对象转换成用户最初提交的版本，进行序列化操作，并调用 Etcd 的 API 把它保存起来。由此可见，声明式 API 对于 Kubernetes 来说非常重要。所以，APIServer 这样一个在其他项目里“平淡无奇”的组件，却成了 Kubernetes 项目的重中之重。它不仅是 Google Borg 设计思想的集中体现，也是 Kubernetes 项目里唯一一个被 Google 公司和 RedHat

在今天这篇文章中，我为你详细解析了 Kubernetes 声明式 API 的工作原理，讲解了如何遵循声明式 API 的设计，为 Kubernetes 添加一个名叫 Network 的 API 资源类型。从而达到了通过标准的 kubectl create 和 get 操作，来管理自定义 API 对象的目的。不过，创建出这样一个自定义 API 对象，我们只是完成了 Kubernetes 声明式 API 的一半工作。接下来的另一半工作是：为这个 API 对象编写一个自定义控制器（Custom Controller）。这样， Kubernetes 才能根据 Network API 对象的“增、删、改”操作，在真实环境中做出相应的响应。比如，“创建、删除、修改”真正的 Neutron 网络。而这，正是 Network 这个 API 对象所关注的“业务逻辑”。这个业务逻辑的实现过程，以及它所使用的 Kubernetes API 编程库的工作原理，就是我要在下一篇文章中讲解的主要内容。


在了解了 CRD 的定义方法之后，你是否已经在考虑使用 CRD（或者已经使用了 CRD）来描述现实中的某种实体了呢？
能否分享一下你的思路？（举个例子：某技术团队使用 CRD 描述了“宿主机”，然后用 Kubernetes 部署了 Kubernetes）

## 深入解析声明式API（二）：编写自定义控制器

所谓的 Informer，就是一个自带缓存和索引机制，可以触发 Handler 的客户端库。这个本地缓存在 Kubernetes 中一般被称为 Store，索引一般被称为 Index。Informer 使用了 Reflector 包，它是一个可以通过 ListAndWatch 机制获取并监视 API 对象变化的客户端封装。Reflector 和 Informer 之间，用到了一个“增量先进先出队列”进行协同。而 Informer 与你要编写的控制循环之间，则使用了一个工作队列来进行协同。在实际应用中，除了控制循环之外的所有代码，实际上都是 Kubernetes 为你自动生成的，即：pkg/client/{informers, listers, clientset}里的内容。而这些自动生成的代码，就为我们提供了一个可靠而高效地获取 API 对象“期望状态”的编程库。所以，接下来，作为开发者，你就只需要关注如何拿到“实际状态”，然后如何拿它去跟“期望状态”做对比，从而决定接下来要做的业务逻辑即可。以上内容，就是 Kubernetes API 编程范式的核心思想。

请思考一下，为什么 Informer 和你编写的控制循环之间，一定要使用一个工作队列来进行协作呢？

## 基于角色的权限控制：RBAC

我们知道，Kubernetes 中所有的 API 对象，都保存在 Etcd 里。可是，
对这些 API 对象的操作，却一定都是通过访问 kube-apiserver 实现的。
其中一个非常重要的原因，就是你需要 APIServer 来帮助你做授权工作。
而在 Kubernetes 项目中，负责完成授权（Authorization）工作的机制，
就是 RBAC：基于角色的访问控制（Role-Based Access Control）。

而在这里，我只希望你能明确三个最基本的概念。
Role：角色，它其实是一组规则，定义了一组对 Kubernetes API 对象的操作权限。
Subject：被作用者，既可以是“人”，也可以是“机器”，也可以使你在 Kubernetes 里定义的“用户”。
RoleBinding：定义了“被作用者”和“角色”的绑定关系。而这三个概念，其实就是整个 RBAC 体系的核心所在。

接下来，我们会看到一个 roleRef 字段。正是通过这个字段，
RoleBinding 对象就可以直接通过名字，来引用我们前面定义的 Role 对象（example-role），
从而定义了“被作用者（Subject）”和“角色（Role）”之间的绑定关系。

那么，对于非 Namespaced（Non-namespaced）对象（比如：Node），或者，
某一个 Role 想要作用于所有的 Namespace 的时候，我们又该如何去做授权呢？

这个由 Kubernetes 负责管理的“内置用户”，正是我们前面曾经提到过的：ServiceAccount。

在今天这篇文章中，我主要为你讲解了基于角色的访问控制（RBAC）。
其实，你现在已经能够理解，所谓角色（Role），其实就是一组权限规则列表。
而我们分配这些权限的方式，就是通过创建 RoleBinding 对象，将被作用者（subject）和权限列表进行绑定。
另外，与之对应的 ClusterRole 和 ClusterRoleBinding，则是 Kubernetes 集群级别的 Role 和 RoleBinding，
它们的作用范围不受 Namespace 限制。而尽管权限的被作用者可以有很多种（比如，User、Group 等），
但在我们平常的使用中，最普遍的用法还是 ServiceAccount。
所以，Role + RoleBinding + ServiceAccount 的权限分配方式是你要重点掌握的内容。
我们在后面编写和安装各种插件的时候，会经常用到这个组合。


请问，如何为所有 Namespace 下的默认 ServiceAccount（default ServiceAccount），
绑定一个只读权限的 Role 呢？请你提供 ClusterRoleBinding（或者 RoleBinding）的 YAML 文件。

## 聪明的微创新：Operator 工作原理解读

在前面的几篇文章中，我已经和你分享了 Kubernetes 项目中的大部分编排对象
（比如 Deployment、StatefulSet、DaemonSet，以及 Job），
也介绍了“有状态应用”的管理方法，还阐述了为 Kubernetes 
添加自定义 API 对象和编写自定义控制器的原理和流程。

Operator 的工作原理，实际上是利用了 Kubernetes 的自定义 API 资源（CRD），
来描述我们想要部署的“有状态应用”；然后在自定义控制器里，
根据自定义 API 对象的变化，来完成具体的部署和运维工作。

不过，考虑到你可能还不太清楚Etcd 集群的组建方式，我在这里先简单介绍一下这部分知识。
Etcd Operator 部署 Etcd 集群，采用的是静态集群（Static）的方式。

在 Operator 的实现过程中，我们再一次用到了 CRD。
可是，你一定要明白，CRD 并不是万能的，它有很多场景不适用，还有性能瓶颈。
你能列举出一些不适用 CRD 的场景么？你知道造成 CRD 性能瓶颈的原因主要在哪里么？

# Kubernetes 容器持久化存储

## PV、PVC、StorageClass，这些到底在说啥？

不难看出，PVC 和 PV 的设计，其实跟“面向对象”的思想完全一致。
PVC 可以理解为持久化存储的“接口”，它提供了对某种持久化存储的描述，
但不提供具体的实现；而这个持久化存储的实现部分则由 PV 负责完成。

而所谓的“持久化 Volume”，指的就是这个宿主机上的目录，具备“持久性”。

这个准备“持久化”宿主机目录的过程，我们可以形象地称为“两阶段处理”。

这一步为虚拟机挂载远程磁盘的操作，对应的正是“两阶段处理”的第一阶段。在 Kubernetes 中，我们把这个阶段称为 Attach。

这个将磁盘设备格式化并挂载到 Volume 宿主机目录的操作，对应的正是“两阶段处理”的第二个阶段，我们一般称为：Mount。

其实很简单，在具体的 Volume 插件的实现接口上，Kubernetes 分别给这两个阶段提供了两种不同的参数列表：
对于“第一阶段”（Attach），Kubernetes 提供的可用参数是 nodeName，即宿主机的名字。
而对于“第二阶段”（Mount），Kubernetes 提供的可用参数是 dir，即 Volume 的宿主机目录。

所以，在 Kubernetes 中，上述关于 PV 的“两阶段处理”流程，
是靠独立于 kubelet 主控制循环（Kubelet Sync Loop）之外的两个控制循环来实现的。

kubelet 的一个主要设计原则，就是它的主控制循环绝对不可以被 block。

在了解了 Kubernetes 的 Volume 处理机制之后，我再来为你介绍这个体系里最后一个重要概念：StorageClass。

所以，Kubernetes 为我们提供了一套可以自动创建 PV 的机制，即：Dynamic Provisioning。
相比之下，前面人工管理 PV 的方式就叫作 Static Provisioning。
Dynamic Provisioning 机制工作的核心，在于一个名叫 StorageClass 的 API 对象。
而 StorageClass 对象的作用，其实就是创建 PV 的模板。

具体地说，StorageClass 对象会定义如下两个部分内容：
第一，PV 的属性。比如，存储类型、Volume 的大小等等。
第二，创建这种 PV 需要用到的存储插件。比如，Ceph 等等。

需要注意的是，这套容器持久化存储体系，完全是 Kubernetes 项目自己负责管理的，
并不依赖于 docker volume 命令和 Docker 的存储插件。
当然，这套体系本身就比 docker volume 命令的诞生时间还要早得多。

在了解了 PV、PVC 的设计和实现原理之后，你是否依然觉得它有“过度设计”的嫌疑？
或者，你是否有更加简单、足以解决你 90% 需求的 Volume 的用法？

## PV、PVC体系是不是多此一举？从本地持久化卷谈起

不过，首先需要明确的是，Local Persistent Volume 并不适用于所有应用。
事实上，它的适用范围非常固定，比如：高优先级的系统应用，需要在多个不同节点上存储数据，并且对 I/O 较为敏感。
典型的应用包括：分布式数据存储比如 MongoDB、Cassandra 等，分布式文件系统比如 GlusterFS、Ceph 等，
以及需要在本地磁盘上进行大量数据缓存的分布式应用。

可以看到，正是通过 PV 和 PVC，以及 StorageClass 这套存储体系，这个后来新添加的持久化存储方案，
对 Kubernetes 已有用户的影响，几乎可以忽略不计。作为用户，
你的 Pod 的 YAML 和 PVC 的 YAML，并没有任何特殊的改变，这个特性所有的实现只会影响到 PV 的处理，
也就是由运维人员负责的那部分工作。

正是由于需要使用“延迟绑定”这个特性，Local Persistent Volume 目前还不能支持 Dynamic Provisioning。
你是否能说出，为什么“延迟绑定”会跟 Dynamic Provisioning 有冲突呢？

## 编写自己的存储插件：FlexVolume与CSI

在 Kubernetes 中，存储插件的开发有两种方式：FlexVolume 和 CSI。

所以说，当你编写完了 FlexVolume 的实现之后，一定要把它的可执行文件放在每个节点的插件目录下。

不过，像这样的 FlexVolume 实现方式，虽然简单，但局限性却很大。

这也是为什么，我们需要有 Container Storage Interface（CSI）这样更完善、更编程友好的插件方式。

相比之下，CSI 插件体系的设计思想，就是把这个 Provision 阶段，以及 Kubernetes 里的一部分存储管理功能，
从主干代码里剥离出来，做成了几个单独的组件。这些组件会通过 Watch API 
监听 Kubernetes 里与存储相关的事件变化，比如 PVC 的创建，来执行具体的存储管理动作。

其中，Driver Registrar 组件，负责将插件注册到 kubelet 里面（这可以类比为，将可执行文件放在插件目录下）。而在具体实现上，Driver Registrar 需要请求 CSI 插件的 Identity 服务来获取插件信息。而 External Provisioner 组件，负责的正是 Provision 阶段。在具体实现上，External Provisioner 监听（Watch）了 APIServer 里的 PVC 对象。当一个 PVC 被创建时，它就会调用 CSI Controller 的 CreateVolume 方法，为你创建对应 PV。

最后一个 External Attacher 组件，负责的正是“Attach 阶段”。在具体实现上，
它监听了 APIServer 里 VolumeAttachment 对象的变化。VolumeAttachment 对象是 Kubernetes 确认一个 Volume 可以进入“Attach 阶段”的重要标志

可以看到，相比于 FlexVolume，CSI 的设计思想，把插件的职责从“两阶段处理”，
扩展成了 Provision、Attach 和 Mount 三个阶段。其中，Provision 等价于“创建磁盘”，
Attach 等价于“挂载磁盘到虚拟机”，Mount 等价于“将该磁盘格式化后，挂载在 Volume 的宿主机目录上”。

假设现在，你的宿主机是阿里云的一台虚拟机，你要实现的容器持久化存储，是基于阿里云提供的云盘。
你能准确地描述出，在 Provision、Attach 和 Mount 阶段，CSI 插件都需要做哪些操作吗？

## 容器存储实践：CSI插件编写指南

我们这次编写的 CSI 插件的功能，就是：让我们运行在 DigitalOcean 上的 Kubernetes 集群能够使用它的块存储服务，作为容器的持久化存储。

这就需要从 CSI 插件的第一个服务 CSI Identity 说起了。


tree $GOPATH/src/github.com/digitalocean/csi-digitalocean/driver  
$GOPATH/src/github.com/digitalocean/csi-digitalocean/driver 
├── controller.go
├── driver.go
├── identity.go
├── mounter.go
└── node.go

第一，通过 DaemonSet 在每个节点上都启动一个 CSI 插件，来为 kubelet 提供 CSI Node 服务。
第二，通过 StatefulSet 在任意一个节点上再启动一个 CSI 插件，为 External Components 提供 CSI Controller 服务。

请你根据编写 FlexVolume 和 CSI 插件的流程，分析一下什么时候该使用 FlexVolume，什么时候应该使用 CSI？

# Kubernetes 容器网络

## 浅谈容器网络

而所谓“网络栈”，就包括了：网卡（Network Interface）、回环设备（Loopback Device）、
路由表（Routing Table）和 iptables 规则。对于一个进程来说，这些要素，
其实就构成了它发起和响应网络请求的基本环境。

在大多数情况下，我们都希望容器进程能使用自己 Network Namespace 里的网络栈，即：拥有属于自己的 IP 地址和端口。

在 Linux 中，能够起到虚拟交换机作用的网络设备，是网桥（Bridge）。
它是一个工作在数据链路层（Data Link）的设备，
主要功能是根据 MAC 地址学习来将数据包转发到网桥的不同端口（Port）上。

可是，我们又该如何把这些容器“连接”到 docker0 网桥上呢？这时候，我们就需要使用一种名叫 Veth Pair 的虚拟设备了。

熟悉了 docker0 网桥的工作方式，你就可以理解，在默认情况下，被限制在 Network Namespace 里的容器进程，
实际上是通过 Veth Pair 设备 + 宿主机网桥的方式，实现了跟同其他容器的数据交换。

这里的关键在于，容器要想跟外界进行通信，它发出的 IP 包就必须从它的 Network Namespace 里出来，来到宿主机上。

而解决这个问题的方法就是：为容器创建一个一端在容器里充当默认网卡、另一端在宿主机上的 Veth Pair 设备。

尽管容器的 Host Network 模式有一些缺点，但是它性能好、配置简单，并且易于调试，
所以很多团队会直接使用 Host Network。那么，如果要在生产环境中使用容器的 Host Network 模式，
你觉得需要做哪些额外的准备工作呢？

## 深入解析容器跨主机网络

Flannel 项目是 CoreOS 公司主推的容器网络方案。
事实上，Flannel 项目本身只是一个框架，真正为我们提供容器网络功能的，是 Flannel 的后端实现。
目前，Flannel 支持三种后端实现，
分别是：
VXLAN；
host-gw；
UDP。

在 Linux 中，TUN 设备是一种工作在三层（Network Layer）的虚拟网络设备。
TUN 设备的功能非常简单，即：在操作系统内核和用户应用程序之间传递 IP 包。

等一下，flanneld 又是如何知道这个 IP 地址对应的容器，是运行在 Node 2 上的呢？
这里，就用到了 Flannel 项目里一个非常重要的概念：子网（Subnet）。

根据我前面讲解的 TUN 设备的原理，这正是一个从用户态向内核态的流动方向（Flannel 进程向 TUN 设备发送数据包），
所以 Linux 内核网络栈就会负责处理这个 IP 包，具体的处理方法，就是通过本机的路由表来寻找这个 IP 包的下一步流向。

可以看到，Flannel UDP 模式提供的其实是一个三层的 Overlay 网络，即：它首先对发出端的 IP 包进行 UDP 封装，
然后在接收端进行解封装拿到原始的 IP 包，进而把这个 IP 包转发给目标容器。
这就好比，Flannel 在不同宿主机上的两个容器之间打通了一条“隧道”，使得这两个容器可以直接使用 IP 地址进行通信，
而无需关心容器和宿主机的分布情况。

我前面曾经提到，上述 UDP 模式有严重的性能问题，所以已经被废弃了。通过我上面的讲述，你有没有发现性能问题出现在了哪里呢？

实际上，相比于两台宿主机之间的直接通信，基于 Flannel UDP 模式的容器通信多了一个额外的步骤，
即 flanneld 的处理过程。而这个过程，由于使用到了 flannel0 这个 TUN 设备，仅在发出 IP 包的过程中，
就需要经过三次用户态与内核态之间的数据拷贝。

所以说，我们在进行系统级编程的时候，有一个非常重要的优化原则，就是要减少用户态到内核态的切换次数，
并且把核心的处理逻辑都放在内核态进行。这也是为什么，Flannel 后来支持的VXLAN 模式，
逐渐成为了主流的容器网络方案的原因。

Flannel UDP 和 VXLAN 模式


可以看到，Flannel 通过上述的“隧道”机制，实现了容器之间三层网络（IP 地址）的连通性。
但是，根据这个机制的工作原理，你认为 Flannel 能负责保证二层网络（MAC 地址）的连通性吗？为什么呢？

## Kubernetes 网络模型与 CNI 网络插件

Kubernetes 之所以要设置这样一个与 docker0 网桥功能几乎一样的 CNI 网桥，
主要原因包括两个方面：一方面，Kubernetes 项目并没有使用 Docker 的网络模型（CNM），所以它并不希望、也不具备配置 docker0 网桥的能力；
另一方面，这还与 Kubernetes 如何配置 Pod，也就是 Infra 容器的 Network Namespace 密切相关。
我们知道，Kubernetes 创建一个 Pod 的第一步，就是创建并启动一个 Infra 容器，
用来“hold”住这个 Pod 的 Network Namespace（这里，你可以再回顾一下专栏第 13 篇文章《为什么我们需要 Pod？》中的相关内容）。
所以，CNI 的设计思想，就是：Kubernetes 在启动 Infra 容器之后，就可以直接调用 CNI 网络插件，
为这个 Infra 容器的 Network Namespace，配置符合预期的网络栈。

这些 CNI 的基础可执行文件，按照功能可以分为三类：

第一类，叫作 Main 插件，它是用来创建具体网络设备的二进制文件。比如，bridge（网桥设备）、ipvlan、loopback（lo 设备）、macvlan、ptp（Veth Pair 设备），以及 vlan。我在前面提到过的 Flannel、Weave 等项目，都属于“网桥”类型的 CNI 插件。所以在具体的实现中，它们往往会调用 bridge 这个二进制文件。这个流程，我马上就会详细介绍到。第二类，叫作 IPAM（IP Address Management）插件，它是负责分配 IP 地址的二进制文件。比如，dhcp，这个文件会向 DHCP 服务器发起请求；host-local，则会使用预先配置的 IP 地址段来进行分配。第三类，是由 CNI 社区维护的内置 CNI 插件。比如：flannel，就是专门为 Flannel 项目提供的 CNI 插件；tuning，是一个通过 sysctl 调整网络设备参数的二进制文件；portmap，是一个通过 iptables 配置端口映射的二进制文件；bandwidth，是一个使用 Token Bucket Filter (TBF) 来进行限流的二进制文件。从这些二进制文件中，我们可以看到，如果要实现一个给 Kubernetes 用的容器网络方案，其实需要做两部分工作，以 Flannel 项目为例：

在本篇文章中，我为你详细讲解了 Kubernetes 中 CNI 网络的实现原理。根据这个原理，你其实就很容易理解所谓的“Kubernetes 网络模型”了：所有容器都可以直接使用 IP 地址与其他容器通信，而无需使用 NAT。所有宿主机都可以直接使用 IP 地址与所有容器通信，而无需使用 NAT。反之亦然。容器自己“看到”的自己的 IP 地址，和别人（宿主机或者容器）看到的地址是完全一样的。可以看到，这个网络模型，其实可以用一个字总结，那就是“通”。

请你思考一下，为什么 Kubernetes 项目不自己实现容器网络，而是要通过 CNI 做一个如此简单的假设呢？

##  解读 Kubernetes 三层网络方案

在上一篇文章中，我以网桥类型的 Flannel 插件为例，为你讲解了 Kubernetes 里容器网络和 CNI 插件的主要工作原理。
不过，除了这种模式之外，还有一种纯三层（Pure Layer 3）网络方案非常值得你注意。
其中的典型例子，莫过于 Flannel 的 host-gw 模式和 Calico 项目了。

在 Kubernetes v1.7 之后，类似 Flannel、Calico 的 CNI 网络插件都是可以直接连接 Kubernetes 的 APIServer 来访问 Etcd 的，
无需额外部署 Etcd 给它们使用。

而在这种模式下，容器通信的过程就免除了额外的封包和解包带来的性能损耗。
根据实际的测试，host-gw 的性能损失大约在 10% 左右，而其他所有基于 VXLAN“隧道”机制的网络方案，性能损失都在 20%~30% 左右。

而在容器生态中，要说到像 Flannel host-gw 这样的三层网络方案，我们就不得不提到这个领域里的“龙头老大”Calico 项目了。

不过，不同于 Flannel 通过 Etcd 和宿主机上的 flanneld 来维护路由信息的做法，
Calico 项目使用了一个“重型武器”来自动地在整个集群中分发路由信息。这个“重型武器”，就是 BGP。

BGP 的全称是 Border Gateway Protocol，即：边界网关协议。
它是一个 Linux 内核原生就支持的、专门用在大规模数据中心里维护不同的“自治系统”之间路由信息的、无中心的路由协议。

在了解了 BGP 之后，Calico 项目的架构就非常容易理解了。
它由三个部分组成：Calico 的 CNI 插件。这是 Calico 与 Kubernetes 对接的部分。我已经在上一篇文章中，和你详细分享了 CNI 插件的工作原理，这里就不再赘述了。
Felix。它是一个 DaemonSet，负责在宿主机上插入路由规则（即：写入 Linux 内核的 FIB 转发信息库），以及维护 Calico 所需的网络设备等工作。
BIRD。它就是 BGP 的客户端，专门负责在集群里分发路由规则信息。除了对路由信息的维护方式之外，Calico 项目与 Flannel 的 host-gw 模式的另一个不同之处，就是它不会在宿主机上创建任何网桥设备。
这时候，Calico 的工作方式，可以用一幅示意图来描述，如下所示（在接下来的讲述中，我会统一用“BGP 示意图”来指代它）：
Calico 工作原理其中的绿色实线标出的路径，就是一个 IP 包从 Node 1 上的 Container 1，到达 Node 2 上的 Container 4 的完整路径。可以看到，Calico 的 CNI 插件会为每个容器设置一个 Veth Pair 设备，然后把其中的一端放置在宿主机上（它的名字以 cali 前缀开头）。此外，由于 Calico 没有使用 CNI 的网桥模式，Calico 的 CNI 插件还需要在宿主机上为每个容器的 Veth Pair 设备配置一条路由规则，用于接收传入的 IP 包。比如，宿主机 Node 2 上的 Container 4 对应的路由规则，如下所示：10.233.2.3 dev cali5863f3 scope link即：发往 10.233.2.3 的 IP 包，应该进入 cali5863f3 设备。基于上述原因，Calico 项目在宿主机上设置的路由规则，肯定要比 Flannel 项目多得多。不过，Flannel host-gw 模式使用 CNI 网桥的主要原因，其实是为了跟 VXLAN 模式保持一致。否则的话，Flannel 就需要维护两套 CNI 插件了。有了这样的 Veth Pair 设备之后，容器发出的 IP 包就会经过 Veth Pair 设备出现在宿主机上。然后，宿主机网络栈就会根据路由规则的下一跳 IP 地址，把它们转发给正确的网关。

需要注意的是，Calico 维护的网络在默认配置下，是一个被称为“Node-to-Node Mesh”的模式。

所以，Node-to-Node Mesh 模式一般推荐用在少于 100 个节点的集群里。而在更大规模的集群中，你需要用到的是一个叫作 Route Reflector 的模式。

此外，我在前面提到过，Flannel host-gw 模式最主要的限制，就是要求集群宿主机之间是二层连通的。而这个限制对于 Calico 来说，也同样存在。

上面这条规则里的下一跳地址是 192.168.2.2，可是它对应的 Node 2 跟 Node 1 却根本不在一个子网里，
没办法通过二层网络把 IP 包发送到下一跳地址。在这种情况下，你就需要为 Calico 打开 IPIP 模式。

不难看到，当 Calico 使用 IPIP 模式的时候，集群的网络性能会因为额外的封包和解包工作而下降。
在实际测试中，Calico IPIP 模式与 Flannel VXLAN 模式的性能大致相当。
所以，在实际使用时，如非硬性需求，我建议你将所有宿主机节点放在一个子网里，避免使用 IPIP。

而在 Calico 项目中，它已经为你提供了两种将宿主机网关设置成 BGP Peer 的解决方案。

这种方案，是使用一个或多个独立组件负责搜集整个集群里的所有路由信息，然后通过 BGP 协议同步给网关。
而我们前面提到，在大规模集群中，Calico 本身就推荐使用 Route Reflector 节点的方式进行组网。
所以，这里负责跟宿主机网关进行沟通的独立组件，直接由 Route Reflector 兼任即可。

在本篇文章中，我为你详细讲述了 Fannel host-gw 模式和 Calico 这两种纯三层网络方案的工作原理。
需要注意的是，在大规模集群里，三层网络方案在宿主机上的路由规则可能会非常多，这会导致错误排查变得困难。
此外，在系统故障的时候，路由规则出现重叠冲突的概率也会变大。基于上述原因，如果是在公有云上，由于宿主机网络本身比较“直白”，我一般会推荐更加简单的 Flannel host-gw 模式。但不难看到，在私有部署环境里，Calico 项目才能够覆盖更多的场景，并为你提供更加可靠的组网方案和架构思路。

你能否总结一下三层网络方案和“隧道模式”的异同，以及各自的优缺点？

三层和隧道的异同：
相同之处是都实现了跨主机容器的三层互通，而且都是通过对目的 MAC 地址的操作来实现的；
不同之处是三层通过配置下一条主机的路由规则来实现互通，隧道则是通过通过在 IP 包外再封装一层 MAC 包头来实现。
三层的优点：少了封包和解包的过程，性能肯定是更高的。
三层的缺点：需要自己想办法维护路由规则。
隧道的优点：简单，原因是大部分工作都是由 Linux 内核的模块实现了，应用层面工作量较少。
隧道的缺点：主要的问题就是性能低。

## 为什么说 Kubernetes 只有 soft multi-tenancy？

你肯定会问了，Kubernetes 的网络方案对“隔离”到底是如何考虑的呢？难道 Kubernetes 就不管网络“多租户”的需求吗？

NetworkPolicy

而 NetworkPolicy 定义的规则，其实就是“白名单”。

然后，在 ingress 字段里，我定义了 from 和 ports，
即：允许流入的“白名单”和端口。
其中，这个允许流入的“白名单”里，我指定了三种并列的情况，分别是：ipBlock、namespaceSelector 和 podSelector。

可以看到，Kubernetes 网络插件对 Pod 进行隔离，其实是靠在宿主机上生成 NetworkPolicy 对应的 iptable 规则来实现的。

在 CNI 网络插件中，上述需求可以通过设置两组 iptables 规则来实现。

第一组规则，负责“拦截”对被隔离 Pod 的访问请求。

实际上，iptables 只是一个操作 Linux 内核 Netfilter 子系统的“界面”。
顾名思义，Netfilter 子系统的作用，就是 Linux 内核里挡在“网卡”和“用户态进程”之间的一道“防火墙”。

其中，第一条 FORWARD 链“拦截”的是一种特殊情况

第二条 FORWARD 链“拦截”的则是最普遍的情况，即：容器跨主通信。

当然，随着 Kubernetes 社区以及 CNCF 生态的不断发展，Kubernetes 项目也已经开始逐步下探，
“吃”掉了基础设施领域的很多“蛋糕”。这也正是容器生态继续发展的一个必然方向。

请你编写这样一个 NetworkPolicy：它使得指定的 Namespace（比如 my-namespace）里的所有 Pod，
都不能接收任何 Ingress 请求。然后，请你说说，这样的 NetworkPolicy 有什么实际的作用？

job，cronjob 这类计算型 pod

## 找到容器不容易：Service、DNS 与服务发现

而 Kubernetes 之所以需要 Service，一方面是因为 Pod 的 IP 不是固定的，
另一方面则是因为一组 Pod 实例之间总会有负载均衡的需求。

实际上，Service 是由 kube-proxy 组件，加上 iptables 来共同实现的。

可以看到，这一组规则，实际上是一组随机模式（–mode random）的 iptables 链。而随机转发的目的地，分别是 KUBE-SEP-WNBA2IHDGP2BOBGZ、KUBE-SEP-X3P2623AGDH6CDF3 和 KUBE-SEP-57KPRZ3JQVENLNBR。而这三条链指向的最终目的地，其实就是这个 Service 代理的三个 Pod。所以这一组规则，就是 Service 实现负载均衡的位置。

不难想到，当你的宿主机上有大量 Pod 的时候，成百上千条 iptables 规则不断地被刷新，
会大量占用该宿主机的 CPU 资源，甚至会让宿主机“卡”在这个过程中。
所以说，一直以来，基于 iptables 的 Service 实现，都是制约 Kubernetes 项目承载更多量级的 Pod 的主要障碍。

而 IPVS 模式的 Service，就是解决这个问题的一个行之有效的方法。IPVS 模式的工作原理，其实跟 iptables 模式类似。
当我们创建了前面的 Service 之后，kube-proxy 首先会在宿主机上创建一个虚拟网卡（叫作：kube-ipvs0），并为它分配 Service VIP 作为 IP 地址。

而相比于 iptables，IPVS 在内核中的实现其实也是基于 Netfilter 的 NAT 模式，
所以在转发这一层上，理论上 IPVS 并没有显著的性能提升。但是，IPVS 并不需要在宿主机上
为每个 Pod 设置 iptables 规则，而是把对这些“规则”的处理放到了内核态，
从而极大地降低了维护这些规则的代价。这也正印证了我在前面提到过的，
“将重要操作放入内核态”是提高性能的重要手段。

不过需要注意的是，IPVS 模块只负责上述的负载均衡和代理功能。而一个完整的 Service 流程正常工作所需要的包过滤、SNAT 等操作，
还是要靠 iptables 来实现。只不过，这些辅助性的 iptables 规则数量有限，也不会随着 Pod 数量的增加而增加。
所以，在大规模集群里，我非常建议你为 kube-proxy 设置–proxy-mode=ipvs 来开启这个功能。
它为 Kubernetes 集群规模带来的提升，还是非常巨大的。

需要注意的是，在 Kubernetes 里，/etc/hosts 文件是单独挂载的，
这也是为什么 kubelet 能够对 hostname 进行修改并且 Pod 重建后依然有效的原因。这跟 Docker 的 Init 层是一个原理。

在这篇文章里，我为你详细讲解了 Service 的工作原理。实际上，Service 机制，以及 Kubernetes 里的 DNS 插件，都是在帮助你解决同样一个问题，即：如何找到我的某一个容器？这个问题在平台级项目中，往往就被称作服务发现，即：当我的一个服务（Pod）的 IP 地址是不固定的且没办法提前获知时，我该如何通过一个固定的方式访问到这个 Pod 呢？而我在这里讲解的、ClusterIP 模式的 Service 为你提供的，就是一个 Pod 的稳定的 IP 地址，即 VIP。并且，这里 Pod 和 Service 的关系是可以通过 Label 确定的。而 Headless Service 为你提供的，则是一个 Pod 的稳定的 DNS 名字，并且，这个名字是可以通过 Pod 名字和 Service 名字拼接出来的。在实际的场景里，你应该根据自己的具体需求进行合理选择。

请问，Kubernetes 的 Service 的负载均衡策略，在 iptables 和 ipvs 模式下，都有哪几种？具体工作模式是怎样的？

一种是通过<serviceName>.<namespace>.svc.cluster.local访问。对应于clusterIP。
另一种是通过<podName>.<serviceName>.<namesapce>.svc.cluster.local访问,对应于headless service。

## 从外界连通 Service 与 Service 调试“三板斧”

所以，在使用 Kubernetes 的 Service 时，一个必须要面对和解决的问题就是：
如何从外部（Kubernetes 集群之外），访问到 Kubernetes 里创建的 Service？

可是，为什么一定要对流出的包做 SNAT操作呢？

这里最常用的一种方式就是：NodePort。

从外部访问 Service 的第二种方式，适用于公有云上的 Kubernetes 服务。
这时候，你可以指定一个 LoadBalancer 类型的 Service

而第三种方式，是 Kubernetes 在 1.7 之后支持的一个新特性，叫作 ExternalName

实际上，在理解了 Kubernetes Service 机制的工作原理之后，很多与 Service 相关的问题，
其实都可以通过分析 Service 在宿主机上对应的 iptables 规则（或者 IPVS 配置）得到解决。

比如，当你的 Service 没办法通过 DNS 访问到的时候。
你就需要区分到底是 Service 本身的配置问题，还是集群的 DNS 出了问题。
一个行之有效的方法，就是检查 Kubernetes 自己的 Master 节点的 Service DNS 是否正常。

如果上面访问 kubernetes.default 返回的值都有问题，那你就需要检查 kube-dns 的运行状态和日志了。
否则的话，你应该去检查自己的 Service 定义是不是有问题。

而如果你的 Service 没办法通过 ClusterIP 访问到的时候，你首先应该检查的是这个 Service 是否有 Endpoints：

而如果 Endpoints 正常，那么你就需要确认 kube-proxy 是否在正确运行。

如果 kube-proxy 一切正常，你就应该仔细查看宿主机上的 iptables 了。
而一个 iptables 模式的 Service 对应的规则，我在上一篇以及这一篇文章里已经全部介绍到了，
它们包括：KUBE-SERVICES 或者 KUBE-NODEPORTS 规则对应的 Service 的入口链，这个规则应该与 VIP 和 Service 端口一一对应；
KUBE-SEP-(hash) 规则对应的 DNAT 链，这些规则应该与 Endpoints 一一对应；
KUBE-SVC-(hash) 规则对应的负载均衡链，这些规则的数目应该与 Endpoints 数目一致；
如果是 NodePort 模式的话，还有 POSTROUTING 处的 SNAT 链。
通过查看这些链的数量、转发目的地址、端口、过滤条件等信息，你就能很容易发现一些异常的蛛丝马迹。

当然，还有一种典型问题，就是 Pod 没办法通过 Service 访问到自己。
这往往就是因为 kubelet 的 hairpin-mode 没有被正确设置。关于 Hairpin 的原理我在前面已经介绍过，这里就不再赘述了。
你只需要确保将 kubelet 的 hairpin-mode 设置为 hairpin-veth 或者 promiscuous-bridge 即可。

在本篇文章中，我为你详细讲解了从外部访问 Service 的三种方式（NodePort、LoadBalancer 和 External Name）和具体的工作原理。然后，我还为你讲述了当 Service 出现故障的时候，如何根据它的工作原理，按照一定的思路去定位问题的可行之道。通过上述讲解不难看出，所谓 Service，其实就是 Kubernetes 为 Pod 分配的、固定的、基于 iptables（或者 IPVS）的访问入口。而这些访问入口代理的 Pod 信息，则来自于 Etcd，由 kube-proxy 通过控制循环来维护。并且，你可以看到，Kubernetes 里面的 Service 和 DNS 机制，也都不具备强多租户能力。比如，在多租户情况下，每个租户应该拥有一套独立的 Service 规则（Service 只应该看到和代理同一个租户下的 Pod）。再比如 DNS，在多租户情况下，每个租户应该拥有自己的 kube-dns（kube-dns 只应该为同一个租户下的 Service 和 Pod 创建 DNS Entry）。当然，在 Kubernetes 中，kube-proxy 和 kube-dns 其实也是普通的插件而已。你完全可以根据自己的需求，实现符合自己预期的 Service。

为什么 Kubernetes 要求 externalIPs 必须是至少能够路由到一个 Kubernetes 的节点？

## 谈谈 Service 与 Ingress

这种全局的、为了代理不同后端 Service 而设置的负载均衡服务，就是 Kubernetes 里的 Ingress 服务。

所以，Ingress 的功能其实很容易理解：所谓 Ingress，就是 Service 的“Service”。

然后，这个 Ingress Controller 会根据你定义的 Ingress 对象，提供对应的代理能力。
目前，业界常用的各种反向代理项目，比如 Nginx、HAProxy、Envoy、Traefik 等，
都已经为 Kubernetes 专门维护了对应的 Ingress Controller。


接下来，我就以最常用的 Nginx Ingress Controller 为例，
在我们前面用 kubeadm 部署的 Bare-metal 环境中，和你实践一下 Ingress 机制的使用过程。

可以看到，Nginx Ingress Controller 为我们创建的 Nginx 负载均衡器，已经成功地将请求转发给了对应的后端 Service。
以上，就是 Kubernetes 里 Ingress 的设计思想和使用方法了。不过，你可能会有一个疑问，
如果我的请求没有匹配到任何一条 IngressRule，那么会发生什么呢？

如果我的需求是，当访问www.mysite.com和 forums.mysite.com时，
分别访问到不同的 Service（比如：site-svc 和 forums-svc）。
那么，这个 Ingress 该如何定义呢？请你描述出 YAML 文件中的 rules 字段。


# Kubernetes 作业调度与资源管理

## Kubernetes 的资源模型与资源管理

而作为 Kubernetes 的资源管理与调度部分的基础，我们要从它的资源模型开始说起。
在 Kubernetes 中，像 CPU 这样的资源被称作“可压缩资源”（compressible resources）。它的典型特点是，当可压缩资源不足时，Pod 只会“饥饿”，但不会退出。而像内存这样的资源，则被称作“不可压缩资源（incompressible resources）。当不可压缩资源不足时，Pod 就会因为 OOM（Out-Of-Memory）被内核杀掉。

此外，不难看到，Kubernetes 里 Pod 的 CPU 和内存资源，实际上还要分为 limits 和 requests 两种情况。

Kubernetes 这种对 CPU 和内存资源限额的设计，实际上参考了 Borg 论文中对“动态资源边界”的定义，
既：容器化作业在提交时所设置的资源边界，并不一定是调度系统所必须严格遵守的，
这是因为在实际场景中，大多数作业使用到的资源其实远小于它所请求的资源限额。

当 Pod 里的每一个 Container 都同时设置了 requests 和 limits，并且 requests 和 limits 值相等的时候，这个 Pod 就属于 Guaranteed 类别。

而当 Pod 不满足 Guaranteed 的条件，但至少有一个 Container 设置了 requests。那么这个 Pod 就会被划分到 Burstable 类别。

而如果一个 Pod 既没有设置 requests，也没有设置 limits，那么它的 QoS 类别就是 BestEffort。

实际上，QoS 划分的主要应用场景，是当宿主机资源紧张的时候，kubelet 对 Pod 进行 Eviction（即资源回收）时需要用到的。

在这个配置中，你可以看到 Eviction 在 Kubernetes 里其实分为 Soft 和 Hard 两种模式。

而当 Eviction 发生的时候，kubelet 具体会挑选哪些 Pod 进行删除操作，就需要参考这些 Pod 的 QoS 类别了。
首当其冲的，自然是 BestEffort 类别的 Pod。其次，是属于 Burstable 类别、并且发生“饥饿”的资源使用量已经超出了 requests 的 Pod。
最后，才是 Guaranteed 类别。并且，Kubernetes 会保证只有当 Guaranteed 类别的 Pod 的资源使用量超过了其 limits 的限制，
或者宿主机本身正处于 Memory Pressure 状态时，Guaranteed 的 Pod 才可能被选中进行 Eviction 操作。

在理解了 Kubernetes 里的 QoS 类别的设计之后，我再来为你讲解一下Kubernetes 里一个非常有用的特性：cpuset 的设置。

这种情况下，由于操作系统在 CPU 之间进行上下文切换的次数大大减少，容器里应用的性能会得到大幅提升。
事实上，cpuset 方式，是生产环境里部署在线应用类型的 Pod 时，非常常用的一种方式。

可是，这样的需求在 Kubernetes 里又该如何实现呢？其实非常简单。
首先，你的 Pod 必须是 Guaranteed 的 QoS 类型；
然后，你只需要将 Pod 的 CPU 资源的 requests 和 limits 设置为同一个相等的整数值即可。

在本篇文章中，我先为你详细讲解了 Kubernetes 里对资源的定义方式和资源模型的设计。
然后，我为你讲述了 Kubernetes 里对 Pod 进行 Eviction 的具体策略和实践方式。
正是基于上述讲述，在实际的使用中，我强烈建议你将 DaemonSet 的 Pod 都设置为 Guaranteed 的 QoS 类型。
否则，一旦 DaemonSet 的 Pod 被回收，它又会立即在原宿主机上被重建出来，这就使得前面资源回收的动作，完全没有意义了。

为什么宿主机进入 MemoryPressure 或者 DiskPressure 状态后，新的 Pod 就不会被调度到这台宿主机上呢？

这是因为给宿主机打了污点标记。

## 十字路口上的 Kubernetes 默认调度器

从集群所有的节点中，根据调度算法挑选出所有可以运行该 Pod 的节点；从第一步的结果中，再根据调度算法挑选一个最符合条件的节点作为最终结果。

所以在具体的调度流程中，默认调度器会首先调用一组叫作 Predicate 的调度算法，来检查每个 Node。
然后，再调用一组叫作 Priority 的调度算法，来给上一步得到的结果里的每个 Node 打分。最终的调度结果，就是得分最高的那个 Node。

在 Kubernetes 中，上述调度机制的工作原理。
第一个控制循环，我们可以称之为 Informer Path。
第二个控制循环，是调度器负责 Pod 调度的主循环，我们可以称之为 Scheduling Path。

调度算法执行完成后，调度器就需要将 Pod 对象的 nodeName 字段的值，修改为上述 Node 的名字。这个步骤在 Kubernetes 里面被称作 Bind。但是，为了不在关键调度路径里远程访问 APIServer，Kubernetes 的默认调度器在 Bind 阶段，只会更新 Scheduler Cache 里的 Pod 和 Node 的信息。这种基于“乐观”假设的 API 对象更新方式，在 Kubernetes 里被称作 Assume。

除了上述的“Cache 化”和“乐观绑定”，Kubernetes 默认调度器还有一个重要的设计，那就是“无锁化”。

不过，随着 Kubernetes 项目发展到今天，它的默认调度器也已经来到了一个关键的十字路口。
事实上，Kubernetes 现今发展的主旋律，是整个开源项目的“民主化”。
也就是说，Kubernetes 下一步发展的方向，是组件的轻量化、接口化和插件化。
所以，我们才有了 CRI、CNI、CSI、CRD、Aggregated APIServer、Initializer、Device Plugin 等各个层级的可扩展能力。
可是，默认调度器，却成了 Kubernetes 项目里最后一个没有对外暴露出良好定义过的、可扩展接口的组件。

需要注意的是，上述这些可插拔式逻辑，都是标准的 Go 语言插件机制（Go plugin 机制），也就是说，你需要在编译的时候选择把哪些插件编译进去。

有了上述设计之后，扩展和自定义 Kubernetes 的默认调度器就变成了一件非常容易实现的事情。这也意味着默认调度器在后面的发展过程中，必然不会在现在的实现上再添加太多的功能，反而还会对现在的实现进行精简，最终成为 Scheduler Framework 的一个最小实现。而调度领域更多的创新和工程工作，就可以交给整个社区来完成了。这个思路，是完全符合我在前面提到的 Kubernetes 的“民主化”设计的。

不过，这样的 Scheduler Framework 也有一个不小的问题，那就是一旦这些插入点的接口设计不合理，
就会导致整个生态没办法很好地把这个插件机制使用起来。而与此同时，这些接口本身的变更又是一个费时费力的过程，
一旦把控不好，就很可能会把社区推向另一个极端，即：Scheduler Framework 没法实际落地，大家只好都再次 fork kube-scheduler。

在本篇文章中，我为你详细讲解了 Kubernetes 里默认调度器的设计与实现，分析了它现在正在经历的重构，以及未来的走向。
不难看到，在 Kubernetes 的整体架构中，kube-scheduler 的责任虽然重大，但其实它却是在社区里最少受到关注的组件之一。
这里的原因也很简单，调度这个事情，在不同的公司和团队里的实际需求一定是大相径庭的，
上游社区不可能提供一个大而全的方案出来。
所以，将默认调度器进一步做轻做薄，并且插件化，才是 kube-scheduler 正确的演进方向。

请问，Kubernetes 默认调度器与 Mesos 的“两级”调度器，有什么异同呢？

messos 二级调度是资源调度和业务调度分开；
优点：插件化调度框架（用户可以自定义自己调度器然后注册到messos资源调度框架即可），灵活可扩展性高。
缺点：资源和业务调度分开无法获取资源使用情况，进而无法做更细粒度的调度。
k8s调度是统一调度也就是业务和资源调度进行统一调度，可以进行更细粒度的调度；缺点其调度器扩展性差。

## Kubernetes 默认调度器调度策略解析

在上一篇文章中，我主要为你讲解了 Kubernetes 默认调度器的设计原理和架构。
在今天这篇文章中，我们就专注在调度过程中 Predicates 和 Priorities 这两个调度策略主要发生作用的阶段。

首先，我们一起看看 Predicates。Predicates 在调度过程中的作用，可以理解为 Filter，
即：它按照调度策略，从当前集群的所有节点中，“过滤”出一系列符合条件的节点。
这些节点，都是可以运行待调度 Pod 的宿主机。

第一种类型，叫作 GeneralPredicates。
顾名思义，这一组过滤规则，负责的是最基础的调度策略。比如，PodFitsResources 计算的就是宿主机的 CPU 和内存资源等是否够用。

第二种类型，是与 Volume 相关的过滤规则。
这一组过滤规则，负责的是跟容器持久化 Volume 相关的调度策略。

第三种类型，是宿主机相关的过滤规则。
这一组规则，主要考察待调度 Pod 是否满足 Node 本身的某些条件。

第四种类型，是 Pod 相关的过滤规则。
这一组规则，跟 GeneralPredicates 大多数是重合的。
而比较特殊的，是 PodAffinityPredicate。
这个规则的作用，是检查待调度 Pod 与 Node 上的已有 Pod 之间的亲密（affinity）和反亲密（anti-affinity）关系。

在具体执行的时候， 当开始调度一个 Pod 时，Kubernetes 调度器会同时启动 16 个 Goroutine，
来并发地为集群里的所有 Node 计算 Predicates，最后返回可以运行这个 Pod 的宿主机列表。

NodeAffinityPriority、TaintTolerationPriority 和 InterPodAffinityPriority 这三种 Priority
PodMatchNodeSelector、PodToleratesNodeTaints 和 PodAffinityPredicate 这三个 Predicate 的含义和计算方法是类似的。
但是作为 Priority，一个 Node 满足上述规则的字段数目越多，它的得分就会越高。

在默认 Priorities 里，还有一个叫作 ImageLocalityPriority 的策略。
它是在 Kubernetes v1.12 里新开启的调度规则，即：如果待调度 Pod 需要使用的镜像很大，
并且已经存在于某些 Node 上，那么这些 Node 的得分就会比较高。

当然，为了避免这个算法引发调度堆叠，调度器在计算得分的时候还会根据镜像的分布进行优化，
即：如果大镜像分布的节点数目很少，那么这些节点的权重就会被调低，从而“对冲”掉引起调度堆叠的风险。

此外，对于比较复杂的调度算法来说，比如 PodAffinityPredicate，它们在计算的时候不只关注待调度 Pod 和待考察 Node，还需要关注整个集群的信息，比如，遍历所有节点，读取它们的 Labels。这时候，Kubernetes 调度器会在为每个待调度 Pod 执行该调度算法之前，先将算法需要的集群信息初步计算一遍，然后缓存起来。这样，在真正执行该算法的时候，调度器只需要读取缓存信息进行计算即可，从而避免了为每个 Node 计算 Predicates 的时候反复获取和计算整个集群的信息。

在本篇文章中，我为你讲述了 Kubernetes 默认调度器里的主要调度算法。
需要注意的是，除了本篇讲述的这些规则，Kubernetes 调度器里其实还有一些默认不会开启的策略。
你可以通过为 kube-scheduler 指定一个配置文件或者创建一个 ConfigMap ，来配置哪些规则需要开启、哪些规则需要关闭。
并且，你可以通过为 Priorities 设置权重，来控制调度器的调度行为。

请问，如何能够让 Kubernetes 的调度器尽可能地将 Pod 分布在不同机器上，避免“堆叠”呢？请简单描述下你的算法。

## Kubernetes 默认调度器的优先级与抢占机制

优先级（Priority ）和抢占（Preemption）机制。

Kubernetes 规定，优先级是一个 32 bit 的整数，最大值不超过 1000000000（10 亿，1 billion），
并且值越大代表优先级越高。而超出 10 亿的值，其实是被 Kubernetes 保留下来分配给系统 Pod 使用的。
显然，这样做的目的，就是保证系统 Pod 不会被用户抢占掉。

而 Kubernetes 调度器实现抢占算法的一个最重要的设计，就是在调度队列的实现里，使用了两个不同的队列。

第一个队列，叫作 activeQ。
凡是在 activeQ 里的 Pod，都是下一个调度周期需要调度的对象。
所以，当你在 Kubernetes 集群里新创建一个 Pod 的时候，调度器会将这个 Pod 入队到 activeQ 里面。
而我在前面提到过的、调度器不断从队列里出队（Pop）一个 Pod 进行调度，实际上都是从 activeQ 里出队的。

第二个队列，叫作 unschedulableQ，专门用来存放调度失败的 Pod。
而这里的一个关键点就在于，当一个 unschedulableQ 里的 Pod 被更新之后，
调度器会自动把这个 Pod 移动到 activeQ 里，从而给这些调度失败的 Pod “重新做人”的机会。

第一步，调度器会检查牺牲者列表，清理这些 Pod 所携带的 nominatedNodeName 字段。
第二步，调度器会把抢占者的 nominatedNodeName，设置为被抢占的 Node 的名字。
第三步，调度器会开启一个 Goroutine，同步地删除牺牲者。

具体来说，在为某一对 Pod 和 Node 执行 Predicates 算法的时候，如果待检查的 Node 是一个即将被抢占的节点，即：调度队列里有 nominatedNodeName 字段值是该 Node 名字的 Pod 存在（可以称之为：“潜在的抢占者”）。那么，调度器就会对这个 Node ，将同样的 Predicates 算法运行两遍。第一遍， 调度器会假设上述“潜在的抢占者”已经运行在这个节点上，然后执行 Predicates 算法；第二遍， 调度器会正常执行 Predicates 算法，即：不考虑任何“潜在的抢占者”。而只有这两遍 Predicates 算法都能通过时，这个 Pod 和 Node 才会被认为是可以绑定（bind）的。

当整个集群发生可能会影响调度结果的变化（比如，添加或者更新 Node，添加和更新 PV、Service 等）时，
调度器会执行一个被称为 MoveAllToActiveQueue 的操作，把所调度失败的 Pod 从 unscheduelableQ 移动到 activeQ 里面。请问这是为什么？

一个相似的问题是，当一个已经调度成功的 Pod 被更新时，调度器则会将 unschedulableQ 
里所有跟这个 Pod 有 Affinity/Anti-affinity 关系的 Pod，移动到 activeQ 里面。请问这又是为什么呢？

## Kubernetes GPU 管理与 Device Plugin 机制

需要指出的是，Device Plugin 的设计，长期以来都是以 Google Cloud 的用户需求为主导的，
所以，它的整套工作机制和流程上，实际上跟学术界和工业界的真实场景还有着不小的差异。

请你结合自己的需求谈一谈，你希望如何对当前的 Device Plugin 进行改进呢？或者说，你觉得当前的设计已经完全够用了吗？

# Kubernetes 容器运行时

## 幕后英雄：SIG-Node与CRI

而 kubelet 这个组件本身，也是 Kubernetes 里面第二个不可被替代的组件（第一个不可被替代的组件当然是 kube-apiserver）。
也就是说，无论如何，我都不太建议你对 kubelet 的代码进行大量的改动。
保持 kubelet 跟上游基本一致的重要性，就跟保持 kube-apiserver 跟上游一致是一个道理。

可以看到，kubelet 的工作核心，就是一个控制循环，即：SyncLoop（图中的大圆圈）。而驱动这个控制循环运行的事件，包括四种：Pod 更新事件；Pod 生命周期变化；kubelet 本身设置的执行周期；定时的清理事件。

所以，跟其他控制器类似，kubelet 启动的时候，要做的第一件事情，就是设置 Listers，也就是注册它所关心的各种事件的 Informer。这些 Informer，就是 SyncLoop 需要处理的数据的来源。此外，kubelet 还负责维护着很多很多其他的子控制循环（也就是图中的小圆圈）。这些控制循环的名字，一般被称作某某 Manager，比如 Volume Manager、Image Manager、Node Status Manager 等等。

在这里需要注意的是，kubelet 调用下层容器运行时的执行过程，并不会直接调用 Docker 的 API，
而是通过一组叫作 CRI（Container Runtime Interface，容器运行时接口）的 gRPC 接口来间接执行的。

不难看到，在这个过程中，kubelet 的 SyncLoop 和 CRI 的设计，是其中最重要的两个关键点。
也正是基于以上设计，SyncLoop 本身就要求这个控制循环是绝对不可以被阻塞的。
所以，凡是在 kubelet 里有可能会耗费大量时间的操作，比如准备 Pod 的 Volume、拉取镜像等，
SyncLoop 都会开启单独的 Goroutine 来进行操作。

请问，在你的项目中，你是如何部署 kubelet 这个组件的？为什么要这么做呢？

## 解读 CRI 与 容器运行时

所以说，这里的 CRI shim，就是容器项目的维护者们自由发挥的“场地”了。
而除了 dockershim 之外，其他容器运行时的 CRI shim，都是需要额外部署在宿主机上的。

举个例子。CNCF 里的 containerd 项目，就可以提供一个典型的 CRI shim 的能力，即：将 Kubernetes 发出的 CRI 请求，
转换成对 containerd 的调用，然后创建出 runC 容器。而 runC 项目，
才是负责执行我们前面讲解过的设置容器 Namespace、Cgroups 和 chroot 等基础操作的组件。


具体地说，我们可以把 CRI 分为两组：第一组，是 RuntimeService。它提供的接口，主要是跟容器相关的操作。比如，创建和启动容器、删除容器、执行 exec 命令等等。而第二组，则是 ImageService。它提供的接口，主要是容器镜像相关的操作，比如拉取镜像、删除镜像等等。

在这一部分，CRI 设计的一个重要原则，就是确保这个接口本身，只关注容器，不关注 Pod。
这样做的原因，也很容易理解。
第一，Pod 是 Kubernetes 的编排概念，而不是容器运行时的概念。所以，我们就不能假设所有下层容器项目，都能够暴露出可以直接映射为 Pod 的 API。
第二，如果 CRI 里引入了关于 Pod 的概念，那么接下来只要 Pod API 对象的字段发生变化，那么 CRI 就很有可能需要变更。
而在 Kubernetes 开发的前期，Pod 对象的变化还是比较频繁的，但对于 CRI 这样的标准接口来说，这个变更频率就有点麻烦了。

关于像 Kata Containers 或者 gVisor 这种所谓的安全容器项目。

除了上述对容器生命周期的实现之外，CRI shim 还有一个重要的工作，就是如何实现 exec、logs 等接口。
这些接口跟前面的操作有一个很大的不同，就是这些 gRPC 接口调用期间，
kubelet 需要跟容器项目维护一个长连接来传输数据。这种 API，我们就称之为 Streaming API。

所以说，当你对容器这一层有特殊的需求时，我一定优先建议你考虑实现一个自己的 CRI shim ，
而不是修改 kubelet 甚至容器项目的代码。这样通过插件的方式定制 Kubernetes 的做法，
也是整个 Kubernetes 社区最鼓励和推崇的一个最佳实践。
这也正是为什么像 Kata Containers、gVisor 甚至虚拟机这样的“非典型”容器，都可以无缝接入到 Kubernetes 项目里的重要原因。

请你思考一下，我前面讲解过的 Device Plugin 为容器分配的 GPU 信息，是通过 CRI 的哪个接口传递给 dockershim，最后交给 Docker API 的呢？

## 绝不仅仅是安全：Kata Containers 与 gVisor

这两种容器实现的本质，都是给进程分配了一个独立的操作系统内核，从而避免了让容器共享宿主机的内核。
这样，容器进程能够看到的攻击面，就从整个宿主机内核变成了一个极小的、独立的、以容器为单位的内核，
从而有效解决了容器进程发生“逃逸”或者夺取整个宿主机的控制权的问题。



# Kubernetes 容器监控与日志

## Prometheus、Metrics Server 与 Kubernetes 监控体系

第一种 Metrics，是宿主机的监控数据。

而 Node Exporter 可以暴露给 Prometheus 采集的 Metrics 数据， 也不单单是节点的负载（Load）、CPU 、内存、磁盘以及网络这样的常规信息，它的 Metrics 指标可以说是“包罗万象”

第二种 Metrics，是来自于 Kubernetes 的 API Server、kubelet 等组件的 /metrics API。
除了常规的 CPU、内存的信息外，这部分信息还主要包括了各个组件的核心监控指标。
比如，对于 API Server 来说，它就会在 /metrics API 里，
暴露出各个 Controller 的工作队列（Work Queue）的长度、请求的 QPS 和延迟数据等等。
这些信息，是检查 Kubernetes 本身工作情况的主要依据。

第三种 Metrics，是 Kubernetes 相关的监控数据。
这部分数据，一般叫作 Kubernetes 核心监控数据（core metrics）。
这其中包括了 Pod、Node、容器、Service 等主要 Kubernetes 核心概念的 Metrics。

需要注意的是，这里提到的 Kubernetes 核心监控数据，其实使用的是 Kubernetes 的一个非常重要的扩展能力，叫作 Metrics Server。

而且，在这个机制下，你还可以添加更多的后端给这个 kube-aggregator。
所以 kube-aggregator 其实就是一个根据 URL 选择具体的 API 后端的代理服务器。
通过这种方式，我们就可以很方便地扩展 Kubernetes 的 API 了。

最后，在具体的监控指标规划上，我建议你遵循业界通用的 USE 原则和 RED 原则。

其中，USE 原则指的是，按照如下三个维度来规划资源监控指标：

利用率（Utilization），资源被有效利用起来提供服务的平均时间占比；
饱和度（Saturation），资源拥挤的程度，比如工作队列的长度；
错误率（Errors），错误的数量。

而 RED 原则指的是，按照如下三个维度来规划服务监控指标：
每秒请求数量（Rate）；
每秒错误数量（Errors）；
服务响应时间（Duration）。

不难发现， USE 原则主要关注的是“资源”，比如节点和容器的资源使用情况，
而 RED 原则主要关注的是“服务”，比如 kube-apiserver 或者某个应用的工作情况。
这两种指标，在我今天为你讲解的 Kubernetes + Prometheus 组成的监控体系中，都是可以完全覆盖到的。

在监控体系中，对于数据的采集，其实既有 Prometheus 这种 Pull 模式，也有 Push 模式。请问，你如何看待这两种模式的异同和优缺点呢？


## Custom Metrics: 让 Auto Scaling 不再“食之无味”


接下来，我通过一个具体的实例，来为你讲解一下 Custom Metrics 具体的使用方式。

首先，我们当然是先部署 Prometheus 项目。这一步，我当然会使用 Prometheus Operator 来完成

第二步，我们需要把 Custom Metrics APIServer 部署起来

第三步，我们需要为 Custom Metrics APIServer 创建对应的 ClusterRoleBinding，以便能够使用 curl 来直接访问 Custom Metrics 的 API

第四步，我们就可以把待监控的应用和 HPA 部署起来

当然，对于一个多实例应用来说，通过 Service 来采集 Pod 的 Custom Metrics 其实才是合理的做法


在本篇文章中，我为你详细讲解了 Kubernetes 里对自定义监控指标，即 Custom Metrics 的设计与实现机制。这套机制的可扩展性非常强，也终于使得 Auto Scaling 在 Kubernetes 里面不再是一个“食之无味”的鸡肋功能了。另外可以看到，Kubernetes 的 Aggregator APIServer，是一个非常行之有效的 API 扩展机制。而且，Kubernetes 社区已经为你提供了一套叫作 KubeBuilder 的工具库，帮助你生成一个 API Server 的完整代码框架，你只需要在里面添加自定义 API，以及对应的业务逻辑即可。

在你的业务场景中，你希望使用什么样的指标作为 Custom Metrics ，以便对 Pod 进行 Auto Scaling 呢？怎么获取到这个指标呢？

## 让日志无处可逃：容器日志收集与管理

首先需要明确的是，Kubernetes 里面对容器日志的处理方式，都叫作 cluster-level-logging，
即：这个日志处理系统，与容器、Pod 以及 Node 的生命周期都是完全无关的。
这种设计当然是为了保证，无论是容器挂了、Pod 被删除，甚至节点宕机的时候，应用的日志依然可以被正常获取到。

主要为你推荐了三种日志方案

第一种，在 Node 上部署 logging agent，将日志文件转发到后端存储里保存起来。

Kubernetes 容器日志方案的第二种，就是对这种特殊情况的一个处理，
即：当容器的日志只能输出到某些文件里的时候，
我们可以通过一个 sidecar 容器把这些日志文件重新输出到 sidecar 的 stdout 和 stderr 上，
这样就能够继续使用第一种方案

第三种方案，就是通过一个 sidecar 容器，直接把应用的日志文件发送到远程存储里面去。

在本篇文章中，我为你详细讲解了 Kubernetes 项目对容器应用日志的收集方式。
综合对比以上三种方案，我比较建议你将应用日志输出到 stdout 和 stderr，
然后通过在宿主机上部署 logging-agent 的方式来集中处理日志。
这种方案不仅管理简单，kubectl logs 也可以用，而且可靠性高，并且宿主机本身，
很可能就自带了 rsyslogd 等非常成熟的日志收集组件来供你使用。

请问，当日志量很大的时候，直接将日志输出到容器 stdout 和 stderr 上，有没有什么隐患呢？
有没有解决办法呢？你还有哪些容器收集的方案，是否可以分享一下？

# 开源与社区

## 谈谈 Kubernetes 开源社区和未来走向

# 特别放送 | 2019 年，容器技术生态会发生些什么？

即：Serverless = FaaS + BaaS。而究其本质，“高可扩展性”、“工作流驱动”和“按使用计费”，可以认为是 Serverless 最主要的三个特征。

并且，无论是 Function、传统应用、容器、存储服务、网络服务，都会开始尝试以不同的方式和形态嵌入到“高可扩展性”、“工作流驱动”和“按使用计费”这三个特征当中。

Serverless 三个特征背后所体现的，乃是云端应用开发过程向“用户友好”和“低心智负担”方向演进的最直接途径。而这种“简单、经济、可信赖”的朴实诉求，正是云计算诞生的最初期许和永恒的发展方向。

而在这种上层应用服务能力向 Serverless 迁移的演进过程中，不断被优化的 Auto-scaling 能力和细粒度的资源隔离技术，将会成为确保 Serverless 能为用户带来价值的最有力保障。

## Kubernetes：赢开发者赢天下

## 基于 Kubernetes 的云原生应用管理，到底应该怎么做？



