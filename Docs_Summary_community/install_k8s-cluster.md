sudo kubeadm init --image-repository registry.aliyuncs.com/google_containers --pod-network-cidr=10.244.0.0/16

kubectl run --generator=run-pod/v1 -i --tty load-generator --image=busybox /bin/sh

kubectl 是管理 Kubernetes Cluster 的命令行工具，前面我们已经在所有的节点安装了 kubectl。Master 初始化完成后需要做一些配置工作，然后 kubectl 就能使用了。
依照 kubeadm init 输出的最后提示，推荐用 Linux 普通用户执行 kubectl。

创建普通用户centos
#创建普通用户并设置密码123456
useradd centos && echo "centos:123456" | chpasswd centos

#追加sudo权限,并配置sudo免密
sed -i '/^root/a\centos  ALL=(ALL)       NOPASSWD:ALL' /etc/sudoers

#保存集群安全配置文件到当前用户.kube目录
su - centos
mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config

#启用 kubectl 命令自动补全功能（注销重新登录生效）
echo "source <(kubectl completion bash)" >> ~/.bashrc

需要这些配置命令的原因是：Kubernetes 集群默认需要加密方式访问。所以，这几条命令，就是将刚刚部署生成的 Kubernetes 集群的安全配置文件，保存到当前用户的.kube 目录下，kubectl 默认会使用这个目录下的授权信息访问 Kubernetes 集群。
如果不这么做的话，我们每次都需要通过 export KUBECONFIG 环境变量告诉 kubectl 这个安全配置文件的位置。
配置完成后centos用户就可以使用 kubectl 命令管理集群了。


查看集群状态：
k8s-release 1.16 开始 
kubectl get cs 显示输出已经变化

请参考https://segmentfault.com/a/1190000020912684

通过下面这条命令显示k8s 1.16 以前的格式
kubectl get cs -o=go-template='{{printf "|NAME|STATUS|MESSAGE|\n"}}{{range .items}}{{$name := .metadata.name}}{{range .conditions}}{{printf "|%s|%s|%s|\n" $name .status .message}}{{end}}{{end}}'


查看节点状态
kubectl get nodes 

使用 kubectl describe 命令来查看这个节点（Node）对象的详细信息、状态和事件（Event）：
kubectl describe node k8s-master 

kubectl get pod -n kube-system -o wide
部署网络插件
要让 Kubernetes Cluster 能够工作，必须安装 Pod 网络，否则 Pod 之间无法通信。
Kubernetes 支持多种网络方案，这里我们使用 Calico

https://docs.projectcalico.org/v3.10/getting-started/kubernetes/

kubectl apply -f https://docs.projectcalico.org/v3.10/manifests/calico.yaml

Kubernetes 的 Worker 节点跟 Master 节点几乎是相同的，它们运行着的都是一个 kubelet 组件。唯一的区别在于，在 kubeadm init 的过程中，kubelet 启动后，Master 节点上还会自动运行 kube-apiserver、kube-scheduler、kube-controller-manger 这三个系统 Pod。


kubeadm join 192.168.92.56:6443 --token 67kq55.8hxoga556caxty7s --discovery-token-ca-cert-hash sha256:7d50e704bbfe69661e37c5f3ad13b1b88032b6b2b703ebd4899e259477b5be69


注意重新执行kubeadm token create --print-join-command

然后根据提示，我们可以通过 kubectl get nodes 查看节点的状态：

如果pod状态为Pending、ContainerCreating、ImagePullBackOff 都表明 Pod 没有就绪，Running 才是就绪状态。
如果有pod提示Init:ImagePullBackOff，说明这个pod的镜像在对应节点上拉取失败，我们可以通过 kubectl describe pod 查看 Pod 具体情况，以确认拉取失败的镜像：

docker images

kubeadm config images list --kubernetes-version v1.16.0 # 看下该版本下的镜像名

在master主机内保存镜像为文件：

docker save -o /opt/kube-pause.tar k8s.gcr.io/pause:3.1
docker save -o /opt/kube-proxy.tar k8s.gcr.io/kube-proxy:v1.13.0
docker save -o /opt/kube-flannel1.tar quay.io/coreos/flannel:v0.9.1
docker save -o /opt/kube-flannel2.tar quay.io/coreos/flannel:v0.10.0-amd64
docker save -o /opt/kube-calico1.tar quay.io/calico/cni:v3.3.2
docker save -o /opt/kube-calico2.tar quay.io/calico/node:v3.3.2
拷贝文件到node计算机

scp /opt/*.tar root@192.168.232.203:/opt/
在node节点执行docker导入

docker load -i /opt/kube-flannel1.tar
docker load -i /opt/kube-flannel2.tar
docker load -i /opt/kube-proxy.tar
docker load -i /opt/kube-pause.tar
docker load -i /opt/kube-calico1.tar
docker load -i /opt/kube-calico2.tar

token过期后重新生成
# 生成新的token
kubeadm token create
# 生成新的token hash码
openssl x509 -pubkey -in /etc/kubernetes/pki/ca.crt | openssl rsa -pubin -outform der 2>/dev/null | openssl dgst -sha256 -hex | sed 's/^.* //'
# 利用新的token和hash码加入节点
# master地址，token，hash请自行更换
kubeadm join 192.168.232.204:6443 --token m87q91.gbcqhfx9ansvaf3o --discovery-token-ca-cert-hash sha256:fdd34ef6c801e382f3fb5b87bc9912a120bf82029893db121b9c8eae29e91c62

sudo kubeadm  init --image-repository=registry.aliyuncs.com/google_containers --kubernetes-version=v1.16.3 --pod-network-cidr=192.168.0.0/16

mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config

kubectl apply -f https://docs.projectcalico.org/v3.10/manifests/calico.yaml

watch kubectl get pods --all-namespaces

kubectl get nodes -o wide

1) [preflight] kubeadm 执行初始化前的检查。
2) [kubelet-start] 生成kubelet的配置文件”/var/lib/kubelet/config.yaml”
3) [certificates] 生成相关的各种token和证书
4) [kubeconfig] 生成 KubeConfig 文件，kubelet 需要这个文件与 Master 通信
5) [control-plane] 安装 Master 组件，会从指定的 Registry 下载组件的 Docker 镜像。
6) [bootstraptoken] 生成token记录下来，后边使用kubeadm join往集群中添加节点时会用到
7) [addons] 安装附加组件 kube-proxy 和 kube-dns。
8) Kubernetes Master 初始化成功，提示如何配置常规用户使用kubectl访问集群。
9) 提示如何安装 Pod 网络。
10) 提示如何注册其他节点到 Cluster


# 重新启动电脑，使用free -m查看分区状态

Pod



CentOS:

vim /etc/sysconfig/network-scripts/ifcfg-ens33

IPADDR=192.168.17.132
GATEWAY=192.168.17.2
DNS=8.8.8.8
BOOTPROTO="static"

service network restart

关闭防火墙

禁用 SELINUX

关闭SWAP


sudo mv /etc/yum.repos.d/CentOS-Base.repo /etc/yum.repos.d/CentOS-Base.repo.backup
sudo wget -O /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-7.repo
yum clean all
yum makecache


docker
sudo yum remove docker \
                  docker-client \
                  docker-client-latest \
                  docker-common \
                  docker-latest \
                  docker-latest-logrotate \
                  docker-logrotate \
                  docker-engine



sudo yum install -y vim wget epel-release

sudo yum install -y yum-utils \
  device-mapper-persistent-data \
  lvm2

sudo yum-config-manager --add-repo http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo

sudo yum install -y docker-ce docker-ce-cli containerd.io

yum list docker-ce --showduplicates | sort -r

sudo yum install -y docker-ce-<VERSION_STRING> docker-ce-cli-<VERSION_STRING> containerd.io
例如：
yum install -y docker-ce-19.03.5-3.el7 docker-ce-cli-19.03.5-3.el7 containerd.io
sudo yum -y install docker-ce-18.09.0 docker-ce-cli-18.09.0

sudo systemctl start docker

sudo docker run hello-world



kubernetes:

su root 用户

sudo cat <<EOF > /etc/yum.repos.d/kubernetes.repo
[kubernetes]
name=Kubernetes
baseurl=https://mirrors.aliyun.com/kubernetes/yum/repos/kubernetes-el7-x86_64/
enabled=1
gpgcheck=1
repo_gpgcheck=1
gpgkey=https://mirrors.aliyun.com/kubernetes/yum/doc/yum-key.gpg https://mirrors.aliyun.com/kubernetes/yum/doc/rpm-package-key.gpg
EOF

sudo yum install -y kubelet-1.16.3-0 kubeadm-1.16.3-0 kubectl-1.16.3-0

sudo systemctl enable kubelet && systemctl start kubelet

kubeadm init --kubernetes-version=1.16.3 \
--apiserver-advertise-address=192.168.31.150 \
--image-repository registry.aliyuncs.com/google_containers \
--service-cidr=10.1.0.0/16 \
--pod-network-cidr=10.244.0.0/16

sudo kubeadm init --kubernetes-version=1.16.3 --image-repository registry.aliyuncs.com/google_containers  --pod-network-cidr=10.244.0.0/16 --service-cidr=10.1.0.0/16

–kubernetes-version: 用于指定k8s版本；
–apiserver-advertise-address：用于指定kube-apiserver监听的ip地址,就是 master本机IP地址。
–pod-network-cidr：用于指定Pod的网络范围； 10.244.0.0/16
–service-cidr：用于指定SVC的网络范围；
–image-repository: 指定阿里云镜像仓库地址

root 用户
cat <<EOF >  /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
EOF
sudo  sysctl --system

（仅master）
yum install -y bash-completion

source <(kubectl completion bash)
echo "source <(kubectl completion bash)" >> ~/.bashrc
source  ~/.bashrc


kubectl create deployment nginx --image=nginx

kubectl scale deployment nginx --replicas=2

kubectl get pods -l app=nginx -o wide

kubectl expose deployment nginx --port=80 --type=NodePort

kubectl get services nginx

kubectl edit cm kube-proxy -n kube-system

删除 每个节点的 pod
kubectl get pod -n kube-system | grep kube-proxy | awk '{system("kubectl delete pod "$1" -n kube-system")}'


Pod
Pod的定义
Pod的生命周期

Pending：Pod定义正确，提交到Master，但其所包含的容器镜像还未完全创建。通常，Master对Pod进行调度需要一些时间，Node进行容器镜像的下载也需要一些时间，启动容器也需要一定时间。（写数据到etcd，调度，pull镜像，启动容器）
Running：Pod已经被分配到某个Node上，并且所有的容器都被创建完毕，至少有一个容器正在运行中，或者有容器正在启动或重启中。
Succeeded：Pod中所有的容器都成功运行结束，并且不会被重启。这是Pod的一种最终状态
Failed：Pod中所有的容器都运行结束了，其中至少有一个容器是非正常结束的（exit code不是0）。这也是Pod的一种最终状态。
Unknown：无法获得Pod的状态，通常是由于无法和Pod所在的Node进行通信。


controller:
  Job
  ReplicationController  ReplicaSet  Deployment
  DaemonSet

Pod Disruption Budgets(Pod 中断)

自愿和非自愿的中断

控制 kubernetes 的速率
  集群需要多少个副本
  优雅地中止一个实例需要多长时间
  一个新的实例启动需要多长时间
  控制器的类型
  集群资源容量

Pod Disruption Budget

PDB能够限制同时中断的pod的数量,以保证集群的高可用性.

最常见的要保护的对象是是以下kubernetes内置的controller创建的应用对象

Deployment
ReplicationController
ReplicaSet
StatefulSet

集群如何响应中断

无状态的前端
单实例有状态应用
多实例有状态应用,例如zookeeper,etcd,consul等

minAvailable  minAvailable

PDB对象不能被更新,你只能够删除它然后重新创建.

Deployment

定义Deployment来创建Pod和ReplicaSet
滚动升级和回滚应用
扩容和缩容
暂停和继续Deployment

kubectl create -f nginx-deployment.yaml

kubectl scale deployment nginx-deployment --replicas 10

kubectl get pods -o wide

升级
kubectl set image deployment/nginx-deployment nginx=nginx:1.9.1

回滚
kubectl rollout undo deployment/nginx-deployment


kubeadm token create
kubeadm token list
openssl x509 -pubkey -in /etc/kubernetes/pki/ca.crt | openssl rsa -pubin -outform der 2>/dev/null | openssl dgst -sha256 -hex | sed 's/^.* //'

[root@walker-4 kubernetes]# kubeadm join --token aa78f6.8b4cafc8ed26c34f --discovery-token-ca-cert-hash sha256:0fd95a9bc67a7bf0ef42da968a0d55d92e52898ec37c971bd77ee501d845b538  172.16.6.79:6443

kubeadm token create --print-join-command

kubectl get pods | grep Terminating | awk '{print $1}' | xargs kubectl delete pod

回退Deployment

kubectl rollout status deployments nginx-deployment

kubectl rollout history deployment/nginx-deployment

kubectl rollout undo deployment/nginx-deployment

清理Policy

比例扩容

Deployment 状态
Progressing Deployment

Complete Deployment

Failed Deployment
  无效的引用
  不可读的 probe failure
  镜像拉取错误
  权限不够
  范围限制
  程序运行时配置错误

Deployment 相对 ReplicaSet 优势

Service

Service可以看作是一组提供相同服务的Pod对外的访问接口。借助Service，应用可以方便地实现服务发现和负载均衡。

Service的类型

ClusterIP
NodePort
LoadBalance
ExternalName

虚拟IP 和服务代理
代理  用户空间的代理模式  Iptables的代理模式

最终结果 任何到Service Cluster Ip 和port的流量都会指向合适的Pod

Job

Job是对ReplicaSet、ReplicationController等持久性控制器的补充。
Job中的restart policy必需是"Never"或者"OnFailure"，这个很好理解，因为pod要运行到结束，而不是反复重新启动。
Job不需要选择器，其中的pod也不需要标签，系统在创建Job时会自动添加相关内容。当然用户也可以出于资源组织的目的添加标签，但这个与Job本身的实现没有关系。
Job新增加两个字段：.spec.completions、.spec.parallelism。详细用法在示例中说明
backoffLimit字段：示例中说明。


Coredns CrashLoopBackOff 导致无法成功添加工作节点的问题
添加工作节点时提示token过期
kubectl 执行命令报“The connection to the server localhost:8080 was refused”
网络组件flannel无法完成初始化
部分节点无法启动pod

kubectl log -f coredns-5c98db65d4-8wt9z -n kube-system

细并发Job
粗并发Job
非并发Job

Service、Deployment、RS、RC和Pod之间的关系

使用kubectl rolling-update更新

Deployment的rolling-update

kubectl apply -f nginx-demo-dm.yml --record

kubectl describe deployment nginx-demo

kubectl rollout status deployment/nginx-demo

kubectl rollout history deployment nginx-demo

kubectl rollout history deployment hello-deployment --revision=2

                           +------------+
                           | deployment |
                           +-----+------+
                                 |
                                 |
                                 |
                                 |
       +--------------------------------------------------+
       |                         |                        |
       |                         |                        |
       |                         |                        |
       |                         |                        |
       |                         |                        |
       |                         |                        |
+------v------+           +------v------+          +------v------+
|replicaset:v1|           |replicaset:v2|          |replicaset:v3|
+-------------+           +------+------+          +-------------+
                                 |
                                 |
                        +--------+---------+
                        |                  |
                        |                  |
                    +---v---+          +---v---+
                    |pod:v2 |          |pod:v2 |
                    +-------+          +-------+

apiVersion: apps/v1
kind: Deployment
metadata:
  name: image-deployment
spec:
  replicas: 1
  selector:
    matchLabels:
      app: image-update
  template:
    metadata:
      labels:
        app: image-update
    spec:
      containers:
      - name: nginx
        image: registry.cn-beijing.aliyuncs.com/mrvolleyball/nginx:v1
        imagePullPolicy: Always

kubectl get deploy

kubectl get rs

kubectl get pod

deployment:
kubectl describe deploy name

replicaset:
kubectl describe rs name

pod

先增加一个pod，镜像版本为新版本
pod可用之后，删除一个老版本pod
循环第1、2步，直到老版本pod全部删除，新版本的pod全部可用

这个过程就是replicaset的作用


kubectl apply -f roll_update.yaml

kubectl describe deploy name

kubectl describe rs update-deployment-7db77f7cc6

kubectl get pod image-deployment-f69875fff-smdxg

kubectl patch deployment image-deployment --patch '{"spec": {"template": {"spec": {"containers": [{"name": "nginx","image":"registry.cn-beijing.aliyuncs.com/mrvolleyball/nginx:v2"}]}}}}' && kubectl rollout pause deployment image-deployment

kubectl rollout history deploy image-deployment

本文详细探索deployment在滚动更新时候的行为

livenessProbe：存活性探测。判断pod是否已经停止
readinessProbe：就绪性探测。判断pod是否能够提供正常服务
maxSurge：在滚动更新过程中最多可以存在的pod数
maxUnavailable：在滚动更新过程中最多不可用的pod数


Kubernetes对象之ReplicaSet

deployment 控制副本数 不推荐直接创建

kubernetes对象之cronjob

定时任务示例：

apiVersion: batch/v1beta1
kind: CronJob
metadata:
  name: hello
spec:
  schedule: "*/1 * * * *"
  jobTemplate:
    spec:
      template:
        spec:
          containers:
          - name: hello
            image: busybox
            args:
            - /bin/sh
            - -c
            - date; echo Hello from the Kubernetes cluster
          restartPolicy: OnFailure

kubectl create -f ./cronjob.yaml

或者

kubectl run hello --schedule="*/1 * * * *" --restart=OnFailure --image=busybox -- /bin/sh -c "date; echo Hello from the Kubernetes cluster"

kubectl get cronjob hello

kubectl delete cronjob hello

kubernetes  对象之 secrets

echo -n 'admin' > ./username.txt
echo -n 'a' > ./password.txt

kubectl create secret generic my-secret --from-file=Secret.yml

apiVersion: v1
kind: Secret
metadata:
  name: mysecret
type: Opaque
data:
  username: YWRtaW4=
  password: MWYyZDFlMmU2N2Rm

echo -n 'admin' | base64
YWRtaW4=
echo -n '1f2d1e2e67df' | base64
MWYyZDFlMmU2N2Rm

kubectl get secret mysecret -o yaml


Pod 使用 Secret 的两种方法：volume 与环境变量

创建或者使用已存Secret，同一Secret可被多个pod引用。
修改pod定义，在.spec.volumes[] 下增加新volume，名称随意。在相应的.spec.volumes[].secret.secretName指定Secret名称。
为使用Secret的容器添加.spec.containers[].volumeMounts[]，同时指定.spec.containers[].volumeMounts[].readOnly = true。指定.spec.containers[].volumeMounts[].mountPath为未使用的目录名称。
在容器的image中，通过指定的目录与Secret中的key访问敏感内容。


通过环境使用Secret流程：
创建或者使用已存在 Secrets，一个 Secret 可以被多个 pod 使用。
修改pod声明中使用 Secret 的容器配置，为其添加环境变量 env[].valueFrom.secretKeyRef，每条 key 对应一个环境变量。
在容器的 image 中通过引用环境变量使用敏感数据。


Secrets限制条件

创建Pod时需要对其使用的Secret进行有效性检查，因此Secrets要先于pod创建。
Secrets寄居于namespace之下，只有处于同一namespace下的pod可以引用
单个Secret的size不能超过1M，防止Secrets占用太多内存引起apiServer性能恶化。但过多的Secrets仍然会占用大量的内在，关于限制所有Secrets占用内存的特性正在计划中。
kubelet在不经过apiServer、控制器创建pod，如–manifest-url、–config时不能使用Secrets。
如果在创建pod时Secret不存在或者key不存在，pod不能启动。通过环境变量引用Secret时，如果envFrom中指定的key的名称不合法，pod仍然能启动但会触发相应错误事件。

Kubernetes之网络策略(Network Policy)

Network Policy对象Spec说明

apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: test-network-policy
  namespace: default
spec:
  podSelector:
    matchLabels:
      role: db
  policyTypes:
  - Ingress
  - Egress
  ingress:
  - from:
    - ipBlock:
        cidr: 172.17.0.0/16
        except:
        - 172.17.1.0/24
    - namespaceSelector:
        matchLabels:
          project: myproject
    - podSelector:
        matchLabels:
          role: frontend
    ports:
    - protocol: TCP
      port: 6379
  egress:
  - to:
    - ipBlock:
        cidr: 10.0.0.0/24
    ports:
    - protocol: TCP
      port: 5978

  
默认规则


默认禁止所有入pod流量(Default deny all ingress traffic)
默认允许所有入pod流量(Default allow all ingress traffic)
默认禁止所有出pod流量(Default deny all egress traffic)
默认允许所有出pod流量(Default allow all egress traffic)
默认禁止所有入出pod流量(Default deny all ingress and all egress traffic)

default-deny   allow-all  


kubectl create -f nginx-policy.yaml

kind: NetworkPolicy
apiVersion: networking.k8s.io/v1
metadata:
  name: access-nginx
spec:
  podSelector:
    matchLabels:
      run: nginx
  ingress:
  - from:
    - podSelector:
        matchLabels:
          access: "true"


kubernetes对象之Volume

Volume类型
awsElasticBlockStore
azureDisk
azureFile
**cephfs**
**configMap**
csi
downwardAPI
emptyDir
fc (fibre channel)
flocker
gcePersistentDisk
gitRepo (deprecated)
**glusterfs**
hostPath
iscsi
local
nfs
persistentVolumeClaim
**projected**
portworxVolume
quobyte
rbd
scaleIO
**secret**
storageos
vsphereVolume

kubernetes对象之Ingress

LoadBalancer   NodePort    Ingress

Ingress Controller   创建 Ingress 对象


Kuebernetes之DaemonSet

用途
运行集群存储守护进程，如glusterd、ceph。
运行集群日志收集守护进程，如fluentd、logstash。
运行节点监控守护进程，如Prometheus Node Exporter, collectd, Datadog agent, New Relic agent, or Ganglia gmond。

总之，可以通过Taint、Toleration、Affinity、node label控制DaemonSet部署pod的节点范围。

DaemonSet自动添加的Toleration

与DaemonSet中pod通信的几种模式

Push：收集数据并向其它服务发送，如将收集到的统计信息发送给统计类型数据库。
NodeIP and Known Port：DaemonSet中的pod可以被设置使用主机网络的一个port，而客户端可以很方便的知道节点IP列表，因此可以通过节点IP地址与port访问DaemonSet pod。
DNS：创建无头服务并且让它的选择器匹配所有DaemonSet的pod，这样DaemonSet中的pod就会成为无头服务的endpoints。类似于StatefulSet。
Service：让Service选中DaemonSet，为访问DaemonSet中的pod提供统一入口与负载均衡。

kubernetes之初始容器(init container)


初始容器能做什么

它们可以包含并且运行一些出于安全考虑不适合和应用放在一块的小工具.

它们可以一些小工具和自定义代码来做些初始化工作,这样就不需要在普通应用容器里使用sed,awk,python或者dig来做初始化工作了

应用构建者和发布者可以独立工作,而不必再联合起来处理同一个pod

它们使用linux namespaces因此它们和普通应用pod拥有不同的文件系统视图.因此他们可以被赋予普通应用容器获取不到的secrets

它们在应用容器启动前运行,因此它们可以阻止或者延缓普通应用容器的初始化直到需要的条件满足

myservice.yaml
kind: Service
apiVersion: v1
metadata:
  name: myservice
spec:
  ports:
  - protocol: TCP
    port: 80
    targetPort: 9376

mydb.yaml
kind: Service
apiVersion: v1
metadata:
  name: mydb
spec:
  ports:
  - protocol: TCP
    port: 80
    targetPort: 9377

myapp.yaml
apiVersion: v1
kind: Pod
metadata:
  name: myapp-pod
  labels:
    app: myapp
spec:
  containers:
  - name: myapp-container
    image: busybox
    command: ['sh', '-c', 'echo The app is running! && sleep 3600']
  initContainers:
  - name: init-myservice
    image: busybox
    command: ['sh', '-c', 'until nslookup myservice; do echo waiting for myservice; sleep 2; done;']
  - name: init-mydb
    image: busybox
    command: ['sh', '-c', 'until nslookup mydb; do echo waiting for mydb; sleep 2; done;']


行为细节

资源

pod重启原因

kubernetes之多容器pod以及通信

什么是pod

kubernetes为什么使用pod作为最小单元,而不是container

kubernetes为什么允许一个pod里有多个容器

多容器pod的使用场景举例

Sidecar containers

代理,桥接和适配器

同一pod间的容器通信

apiVersion: v1
kind: Pod
metadata:
  name: mc1
spec:
  volumes:
  - name: html
    emptyDir: {}
  containers:
  - name: 1st
    image: nginx
    volumeMounts:
    - name: html
      mountPath: /usr/share/nginx/html
  - name: 2nd
    image: debian
    volumeMounts:
    - name: html
      mountPath: /html
    command: ["/bin/sh", "-c"]
    args:
      - while true; do
          date >> /html/index.html;
          sleep 1;
        done

 kubectl exec mc1 -c 1st -- /bin/cat /usr/share/nginx/html/index.html

 kubectl exec mc1 -c 2nd -- /bin/cat /html/index.html

 进程间通信(IPC)

apiVersion: v1
kind: Pod
metadata:
  name: producer-comsumer
spec:
  containers:
  - name: producer
    image: allingeek/ch6_ipc
    command: ["./ipc", "-producer"]
  - name: consumer
    image: allingeek/ch6_ipc
    command: ["./ipc", "-consumer"]
  restartPolicy: Never

kubectl get pods

kubectl logs producer-comsumer -c producer

kubectl logs producer-comsumer -c consumer

容器的依赖关系和启动顺序

**kubernetes No route to host iptable**

清楚防火墙规则
iptables -F

Kubernetes 基本概念之 Label


Label的定义

常见的Label

relase: stable
release: canary
environment: dev
environemnt: qa
environment: production
tier: frontend
tier: backend
tier: middleware

Label Selector


kubectl rollout回滚和autoscale自动扩容

kubernetes 滚动升级

Kubernetes 中采用ReplicaSet（简称RS）来管理Pod。如果当前集群中的Pod实例数少于目标值，RS 会拉起新的Pod，反之，则根据策略删除多余的Pod。Deployment正是利用了这样的特性，通过控制两个RS里面的Pod，从而实现升级。
滚动升级是一种平滑过渡式的升级，在升级过程中，服务仍然可用。


创建 deployment
kubectl create deploy nginx-test --image=nginx:1.14

scale 副本数量
kubectl scale deployment nginx-test --replicas 10

如果集群支持 horizontal pod autoscaling 的话，还可以为Deployment设置自动扩展：
kubectl autoscale deployment nginx-test --min=10 --max=15 --cpu-percent=80

更新 deployment

回滚到上一个版本：

kubectl rollout undo deployment/nginx-test
也可以使用 --revision参数指定某个历史版本：

kubectl rollout undo deployment/nginx-test --to-revision=2
历史记录
kubectl rollout history deployment/nginx-test

验证发布
kubectl rollout status deploy/nginx-test

回滚发布
kubectl rollout undo deployments/nginx-test

kubectl rollout undo deployment/nginx-test --to-revision=<版次>

滚动升级中的重要参数：
maxSurge    maxUnavailable


DESIRED 最终期望处于READY状态的副本数

CURRENT 当前的副本总数

UP-TO-DATE 当前完成更新的副本数

AVAILABLE 当前可用的副本数

使用 autoscaler 自动设置在kubernetes集群中运行的pod数量（水平自动伸缩）。


kubernetes调度之nodeName与NodeSelector

NodeName:
Pod.spec.nodeName用于强制约束将Pod调度到指定的Node节点上，这里说是“调度”，
但其实指定了nodeName的Pod会直接跳过Scheduler的调度逻辑，直接写入PodList列表，该匹配规则是强制匹配。

NodeSelector:
Pod.spec.nodeSelector是通过kubernetes的label-selector机制进行节点选择，
由scheduler调度策略MatchNodeSelector进行label匹配，调度pod到目标节点，该匹配规则是强制约束。


kubernetes里的各种port

apiVersion: extensions/v1beta1
kind: Deployment
metadata:
name: tomcat-deployment
spec:
replicas: 3
template:
metadata:
labels:
app: tomcat
tier: frontend
spec:
containers:
    name: tomcat
    image: docker.cinyi.com:443/tomcat
    ports:
    containerPort: 80   #这里containerPort是容器内部的port

apiVersion: v1
kind: Service
metadata:
name: tomcat-server
spec:
type: NodePort
ports:
    port: 11111  #service暴露在cluster ip上的端口，通过<cluster ip>:port访问服务,通过此端口集群内的服务可以相互访问
    targetPort: 8080  #Pod的外部访问端口，port和nodePort的数据通过这个端口进入到Pod内部，Pod里面的containers的端口映射到这个端口，提供服务
    nodePort: 30001 #Node节点的端口，<nodeIP>:nodePort 是提供给集群外部客户访问service的入口
selector:
tier: frontend

kubernetes调度之 PriorityClass

现在版本支持Pod优先级抢占，通过PriorityClass来实现同一个Node节点内部的Pod对象抢占。
根据 Pod 中运行的作业类型判定各个 Pod 的优先级，对于高优先级的 Pod 可以抢占低优先级 Pod 的资源。
Pod priority指的是Pod的优先级，高优先级的Pod会优先被调度，
或者在资源不足低情况牺牲低优先级的Pod，以便于重要的Pod能够得到资源部署


Kubernetes调度之亲和与反亲和

nodeSelector(节点选择器) 

pod-nginx.yaml

apiVersion: v1
kind: Pod
metadata:
  name: nginx
  labels:
    env: test
spec:
  containers:
  - name: nginx
    image: nginx
    imagePullPolicy: IfNotPresent
  nodeSelector:
    disktype: ssd

亲和与反亲和（Affinity and anti-affinity）

部署pod时，大多数情况下kubernetes的调度程序能将pod调度到集群中合适的节点上。
但有些情况下用户需要对pod调度到哪个节点上施加更多控制，
比如将特定pod部署到拥有SSD存储节点、将同一个服务的多个后端部署在不同的机器上提高安全性、
将通信频繁的服务部署在同一个可用区域降低通信链路长度。
用户对pod部署的节点施加控制都与"label selector"有关。

节点亲和

内部pod亲和与反亲和

示例：
部署2个 redis 工作节点，通过反亲和将2个 redis 副本分别部署在2个不同的节点上。
apiVersion: apps/v1
kind: Deployment
metadata:
  name: redis-cache
spec:
  selector:
    matchLabels:
      app: store
  replicas: 2
  template:
    metadata:
      labels:
        app: store
    spec:
      affinity:
        podAntiAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
          - labelSelector:
              matchExpressions:
              - key: app
                operator: In
                values:
                - store
            topologyKey: "kubernetes.io/hostname"
      containers:
      - name: redis-server
        image: redis:3.2-alpine
2 个 nginx web 前端，要求 2 个副本不对分别部署在不同的节点上，通过与上列相似的反亲和实现。
同时需要将 2 个 web 前端部署在其上已经部署 redis 的节点上，降低通信成本，通过亲和实现，配置如下：
apiVersion: apps/v1
kind: Deployment
metadata:
  name: web-server
spec:
  selector:
    matchLabels:
      app: web-store
  replicas: 3
  template:
    metadata:
      labels:
        app: web-store
    spec:
      affinity:
        podAntiAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
          - labelSelector:
              matchExpressions:
              - key: app
                operator: In
                values:
                - web-store
            topologyKey: "kubernetes.io/hostname"
        podAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
          - labelSelector:
              matchExpressions:
              - key: app
                operator: In
                values:
                - store
            topologyKey: "kubernetes.io/hostname"
      containers:
      - name: web-app
        image: nginx:1.12-alpine

kubernetes调度之污点(taint)和容忍(toleration)

kubectl create -f pod-redis-two.yaml
kubectl create -f  pod-nginx-two.yaml

以上测试的是 pod 的亲和力与反亲和力

可以使用kubectl taint为一个节点(node)添加污点(taint),例如:

kubectl taint nodes node1 key=value:NoSchedule

kubectl taint nodes node1 key:NoSchedule-

基于taint的驱离策略
有条件节点taint

所谓驱离是指pod被从此节点上移除,调度到其它节点上

pod的亲和性是以pod为中心的,而节点的污点则是以节点为中心.想要使pod被调度到指定节点,需要亲和属性,想要节点排斥非专用pod,则需要使用taint,同时使用亲和性和污点可以保证专用节点被特定pod专用,特定pod仅使用专用节点

基于taint的驱离策略

不容忍此taint的pod会被马上驱离
容忍此taint但是没有指定tolerationSeconds的pod将会永远运行在此节点
容忍此taint但是包含tolerationSeconds属性的节点将会在此节点上留存指定时间(虽然容忍,但是是有条件的,仅在一段时间内容忍)

DaemonSet类型的pod创建时自动为以下两种类型的taint添加NoExecute效果类型并且没有tolerationSeconds

有条件节点taint

node.kubernetes.io/unreachable
node.kubernetes.io/not-ready
node.kubernetes.io/memory-pressure
node.kubernetes.io/disk-pressure
node.kubernetes.io/out-of-disk (only for critical pods)
node.kubernetes.io/unschedulable (1.10 or later)
node.kubernetes.io/network-unavailable (host network only)


kubernetes 之 StatefulSet 详解
RC、Deployment、DaemonSet都是面向无状态的服务，它们所管理的Pod的IP、名字，启停顺序等都是随机的，而StatefulSet是什么？顾名思义，有状态的集合，管理所有有状态的服务，比如MySQL、MongoDB集群等。
StatefulSet本质上是Deployment的一种变体，在v1.9版本中已成为GA版本，它为了解决有状态服务的问题，它所管理的Pod拥有固定的Pod名称，启停顺序，在StatefulSet中，Pod名字称为网络标识(hostname)，还必须要用到共享存储。
在Deployment中，与之对应的服务是service，而在StatefulSet中与之对应的headless service，headless service，即无头服务，与service的区别就是它没有Cluster IP，解析它的名称时将返回该Headless Service对应的全部Pod的Endpoint列表。
除此之外，StatefulSet在Headless Service的基础上又为StatefulSet控制的每个Pod副本创建了一个DNS域名，这个域名的格式为：

$(podname).(headless server name)
FQDN： $(podname).(headless server name).namespace.svc.cluster.local

apiVersion: v1
kind: Service
metadata:
  name: nginx
  labels:
    app: nginx
spec:
  ports:
  - port: 80
    name: web
  clusterIP: None
  selector:
    app: nginx
---
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: web
spec:
  selector:
    matchLabels:
      app: nginx # has to match .spec.template.metadata.labels
  serviceName: "nginx"  #声明它属于哪个Headless Service.
  replicas: 3 # by default is 1
  template:
    metadata:
      labels:
        app: nginx # has to match .spec.selector.matchLabels
    spec:
      terminationGracePeriodSeconds: 10
      containers:
      - name: nginx
        image: k8s.gcr.io/nginx-slim:0.8
        ports:
        - containerPort: 80
          name: web
        volumeMounts:
        - name: www
          mountPath: /usr/share/nginx/html
  volumeClaimTemplates:   #可看作pvc的模板
  - metadata:
      name: www
    spec:
      accessModes: [ "ReadWriteOnce" ]
      storageClassName: "gluster-heketi"  #存储类名，改为集群中已存在的
      resources:
        requests:
          storage: 1Gi

Headless Service：名为nginx，用来定义Pod网络标识( DNS domain)。
StatefulSet：定义具体应用，名为Nginx，有三个Pod副本，并为每个Pod定义了一个域名。
volumeClaimTemplates： 存储卷申请模板，创建PVC，指定pvc名称大小，将自动创建pvc，且pvc必须由存储类供应。

为什么需要 headless service 无头服务？
在用Deployment时，每一个Pod名称是没有顺序的，是随机字符串，因此是Pod名称是无序的，但是在statefulset中要求必须是有序 ，每一个pod不能被随意取代，pod重建后pod名称还是一样的。而pod IP是变化的，所以是以Pod名称来识别。pod名称是pod唯一性的标识符，必须持久稳定有效。这时候要用到无头服务，它可以给每个Pod一个唯一的名称。

为什么需要volumeClaimTemplate？ 对于有状态的副本集都会用到持久存储，对于分布式系统来讲，它的最大特点是数据是不一样的，所以各个节点不能使用同一存储卷，每个节点有自已的专用存储，但是如果在Deployment中的Pod template里定义的存储卷，是所有副本集共用一个存储卷，数据是相同的，因为是基于模板来的 ，而statefulset中每个Pod都要自已的专有存储卷，所以statefulset的存储卷就不能再用Pod模板来创建了，于是statefulSet使用volumeClaimTemplate，称为卷申请模板，它会为每个Pod生成不同的pvc，并绑定pv， 从而实现各pod有专用存储。这就是为什么要用volumeClaimTemplate的原因。


如果集群中没有StorageClass的动态供应PVC的机制，也可以提前手动创建多个PV、PVC，
手动创建的PVC名称必须符合之后创建的StatefulSet命名规则：(volumeClaimTemplates.name)-(pod_name)

匹配Pod name(网络标识)的模式为：$(statefulset名称)-$(序号)，比如上面的示例：web-0，web-1，web-2。

StatefulSet为每个Pod副本创建了一个DNS域名，这个域名的格式为： $(podname).(headless server name)，也就意味着服务间是通过Pod域名来通信而非Pod IP，因为当Pod所在Node发生故障时，Pod会被飘移到其它Node上，Pod IP会发生变化，但是Pod域名不会有变化。

StatefulSet使用Headless服务来控制Pod的域名，这个域名的FQDN为：$(service name).$(namespace).svc.cluster.local，其中，“cluster.local”指的是集群的域名。

根据volumeClaimTemplates，为每个Pod创建一个pvc，pvc的命名规则匹配模式：(volumeClaimTemplates.name)-(pod_name)，比如上面的volumeMounts.name=www， Pod name=web-[0-2]，因此创建出来的PVC是www-web-0、www-web-1、www-web-2。

删除Pod不会删除其pvc，手动删除pvc将自动释放pv。
关于Cluster Domain、headless service名称、StatefulSet 名称如何影响StatefulSet的Pod的DNS域名的示例：


Cluster Domain	    Service (ns/name)	    StatefulSet (ns/name)	    StatefulSet Domain	    Pod DNS	    Pod Hostname
cluster.local	default/nginx	default/web	nginx.default.svc.cluster.local	web-{0..N-1}.nginx.default.svc.cluster.local	web-{0..N-1}
cluster.local	foo/nginx	foo/web	nginx.foo.svc.cluster.local	web-{0..N-1}.nginx.foo.svc.cluster.local	web-{0..N-1}
kube.local	foo/nginx	foo/web	nginx.foo.svc.kube.local	web-{0..N-1}.nginx.foo.svc.kube.local	web-{0..N-1}

Statefulset的启停顺序：
有序部署：部署StatefulSet时，如果有多个Pod副本，它们会被顺序地创建（从0到N-1）并且，在下一个Pod运行之前所有之前的Pod必须都是Running和Ready状态。
有序删除：当Pod被删除时，它们被终止的顺序是从N-1到0。
有序扩展：当对Pod执行扩展操作时，与部署一样，它前面的Pod必须都处于Running和Ready状态

在v1.7以后，通过允许修改Pod排序策略，同时通过.spec.podManagementPolicy字段确保其身份的唯一性。
OrderedReady：上述的启停顺序，默认设置。
Parallel：告诉StatefulSet控制器并行启动或终止所有Pod，并且在启动或终止另一个Pod之前不等待前一个Pod变为Running and Ready或完全终止。

StatefulSet 使用场景
稳定的持久化存储，即Pod重新调度后还是能访问到相同的持久化数据，基于PVC来实现。
稳定的网络标识符，即Pod重新调度后其PodName和HostName不变。
有序部署，有序扩展，基于 init containers 来实现。
有序收缩。

更新策略：
OnDelete   RollingUpdate   Partitions


kubernetes的Service Account和secret

Service Account概念的引入是基于这样的使用场景：
运行在pod里的进程需要调用Kubernetes API以及非Kubernetes API的其它服务。
Service Account它并不是给kubernetes集群的用户使用的，而是给pod里面的进程使用的，它为pod提供必要的身份认证。

kubectl get sa --all-namespaces

kubectl get sa  default  -o yaml

kubectl get secret default-token-value -o yaml

Secret

Kubernetes提供了Secret来处理敏感信息，目前Secret的类型有3种:
Opaque(default): 任意字符串
kubernetes.io/service-account-token: 作用于ServiceAccount，就是上面说的。
kubernetes.io/dockercfg: 作用于Docker registry，用户下载docker镜像认证使用。

secret-opaque.yaml
apiVersion: v1
kind: Secret
metadata:
  name: mysecret
type: Opaque
data:
  username: YWRtaW4=
  password: MWYyZDFlMmU2N2Rm


imagePullSecrets  当在需要安全验证的环境中拉取镜像的时候，需要通过用户名和密码


kubernetes之计算机资源管理

当你编排一个pod的时候,你也可以可选地指定每个容器需要多少CPU和多少内存(RAM).
当容器请求特定的资源时,调度器可以更好地根据资源请求来确定把pod调度到哪个节点上.
当容器请求限制特定资源时,特定节点会以指定方式对容器的资源进行限制.

pod和容器的资源请求与资源限制

spec.containers[].resources.limits.cpu
spec.containers[].resources.limits.memory
spec.containers[].resources.requests.cpu
spec.containers[].resources.requests.memory

CPU  内存  pod的资源请求/限制是pod里的容器资源请求/限制的和

apiVersion: v1
kind: Pod
metadata:
  name: frontend
spec:
  containers:
  - name: db
    image: mysql
    env:
    - name: MYSQL_ROOT_PASSWORD
      value: "password"
    resources:
      requests:
        memory: "64Mi"
        cpu: "250m"
      limits:
        memory: "128Mi"
        cpu: "500m"
  - name: wp
    image: wordpress
    resources:
      requests:
        memory: "64Mi"
        cpu: "250m"
      limits:
        memory: "128Mi"
        cpu: "500m"

pod挂起 event消息是:failedScheduling

集群中添加更多节点
终止一些非必须进程来为挂起的pod腾出资源
检测确保pod小于node,比如节点的容量是cpu:1,如果pod请求的是cpu:1.1将永远不会被调度.

查看节点 node 的资源

kubectl describe node node的名称

Local ephemeral storage(本地暂存容量)

对Local ephemeral storage的请求/限制

带有Local ephemeral storage的pod如何运行

对于容器级别的隔离,如果容器的可写层( writable layer)和日志(log)超出了容量限制,容器所在的pod将会被驱离;
对于pod级别的隔离,如果pod里所有容器使用的总Local ephemeral storage和pod的emptydir存储卷超过限制,pod将会被驱离.

apiVersion: v1
kind: Pod
metadata:
  name: teststorage
  labels:
    app: teststorage
spec:
  containers:
  - name: busybox
    image:  busybox
    command: ["bash", "-c", "while true; do dd if=/dev/zero of=$(date '+%s').out count=1 bs=10MB; sleep 1; done"] # 持续写入文件到容器的rootfs中
    resources:
      limits:
        ephemeral-storage: 100Mi #定义存储的限制为100M
      requests:
        ephemeral-storage: 100Mi


kubernetes调度之资源配额

资源配额,通过ResourceQuota定义,提供了对某一名称空间使用资源的总体约束.
它即可以限制这个名称空间下有多少个对象可以被创建,
也可以限制对计算机资源使用量的限制(前面说到过,计算机资源包括cpu,内存,磁盘空间等资源)


计算机资源配额
cpu memory 
扩展资源的资源配额  存储资源配额

Terminating,NotTerminating和NotBestEffort范围限制配额追踪以下资源:

cpu

limits.cpu

limits.memory

memory

pods

requests.cpu

requests.memory

每一个 PriorityClass 的资源配额

你在使用 PriorityClass 的配额,需要启用 ResourceQuotaScopeSelectors

集群中的 pod 有以下三个优先级类之一:low,medium,high

每个优先级类都创建了一个资源配额

apiVersion: v1
kind: List
items:
- apiVersion: v1
  kind: ResourceQuota
  metadata:
    name: pods-high
  spec:
    hard:
      cpu: "100"
      memory: 200Mi
      pods: "10"
    scopeSelector:
      matchExpressions:
      - operator : In
        scopeName: PriorityClass
        values: ["high"]
- apiVersion: v1
  kind: ResourceQuota
  metadata:
    name: pods-medium
  spec:
    hard:
      cpu: "10"
      memory: 20Mi
      pods: "10"
    scopeSelector:
      matchExpressions:
      - operator : In
        scopeName: PriorityClass
        values: ["medium"]
- apiVersion: v1
  kind: ResourceQuota
  metadata:
    name: pods-low
  spec:
    hard:
      cpu: "5"
      memory: 10Mi
      pods: "10"
    scopeSelector:
      matchExpressions:
      - operator : In
        scopeName: PriorityClass
        values: ["low"]

配额资源的申请与限制

cat <<EOF > compute-resources.yaml
apiVersion: v1
kind: ResourceQuota
metadata:
  name: compute-resources
spec:
  hard:
    pods: "4"
    requests.cpu: "1"
    requests.memory: 1Gi
    limits.cpu: "2"
    limits.memory: 2Gi
    requests.nvidia.com/gpu: 4
EOF

kubectl create namespace myspace

kubectl create -f compute-resources.yaml --namespace=myspace

cat <<EOF > object-counts.yaml
apiVersion: v1
kind: ResourceQuota
metadata:
  name: object-counts
spec:
  hard:
    configmaps: "10"
    persistentvolumeclaims: "4"
    replicationcontrollers: "20"
    secrets: "10"
    services: "10"
    services.loadbalancers: "2"
EOF

kubectl create -f object-counts.yaml --namespace=myspace

kubectl get quota --namespace=myspace

kubectl describe quota compute-resources --namespace=myspace

kubectl describe quota object-counts --namespace=myspace


kubectl通过count/<resource>.<group>语法形式支持标准名称空间对象数量配额

kubectl create namespace myspace
kubectl create quota test --hard=count/deployments.extensions=2,count/replicasets.extensions=4,count/pods=3,count/secrets=4 --namespace=myspace
kubectl run nginx --image=nginx --replicas=2 --namespace=myspace
kubectl describe quota --namespace=myspace


配额和集群容量
ResourceQuotas独立于集群的容量,它们通过绝对的单位表示.因此,如果你向集群添加了节点,这并不会给集群中的每个名称空间赋予消费更多资源的能力.

把集群中所有的资源按照比例分配给不同团队
允许每个租户根据需求增加资源使用,但是有一个总体的限制以防资源被耗尽
检测名称空间的需求,添加节点,增加配额

kubernetes调度之pod优先级和资源抢占

Pod可以拥有优先级.优先意味着相对于其它pod某个pod更为重要.如果重要的pod不能被调度,
则kubernetes调度器会优先于(驱离)低优先级的pod来让处于pending状态的高优先级pod被调度.

怎样使用优先级和抢占
怎样禁用抢占
抢占
抢占限制
跨节点抢占

调度pod优先级和抢占

pod优先级和资源抢占可能导致的问题

pod被非必要地抢占
pod被抢占,但是抢占的pod并没有被调度
高优先级的pod比优先级更低的pod被抢占

pod优先级和QoS是两个正交的功能,几乎没有交叉.调度器的抢占逻辑当选择抢占对象时不会考虑QoS.
抢占会考虑pod的优先级并选择出一系列低优先级抢占目标.
只有当即便移除低优先级pod也不足以运行挂起的pod或者低优先级的pod被PodDisruptionBudget保护时,
调度器才会选择更高优先级的pod作为抢占对象.

kubernetes调度之资源耗尽处理配置

kubelet会积极的监视并阻止可用计算机资源耗尽.这种情况下,kubelet会终止一个或者多个pod来重新取回耗尽的资源,
当kubelet终止一个pod时,它将会终止pod的所有容器并把PodPhase设置为Failed

kubelet支持基于下面列表中描述的驱离信号的驱离策略

Eviction Signal	Description
memory.available	memory.available := node.status.capacity[memory] - node.stats.memory.workingSet
nodefs.available	nodefs.available := node.stats.fs.available
nodefs.inodesFree	nodefs.inodesFree := node.stats.fs.inodesFree
imagefs.available	imagefs.available := node.stats.runtime.imagefs.available
imagefs.inodesFree	imagefs.inodesFree := node.stats.runtime.imagefs.inodesFree


驱离阈值

kubelet支持指定驱离阈值来来触发kubelet回收资源

每一个阈值都是以下形式的:

[eviction-signal][operator][quantity]
eviction-signal 是上面表中定义的一个信号token

operator是一种期望的操作符,比如<(小于号)

quantity 驱离阈值的量,比如1Gi,驱离阈值也可以是由%百分号表示的百分比值

比如说一个节点有10Gi总内存值,并且如果可用内存的值如果低于1Gi的时候你想要触发驱离,你可以以如下两种方式中的任一来定义驱离阈值

memory.available<10%或memory.available<1Gi但是你不能两者同时使用

软驱离阈值

软驱离阈值和一个包含管理员指定的优雅时间的驱离阈值成对出现.驱离信号发出后,
在优雅时间没有超出之前,kubelet不会回收资源.如果不指定优雅时间,kubelet会在一开始就返回错误

硬驱离阈值

硬驱离阈值没有优雅时段,kubelet会立马对相关的资源采取动作.如果硬驱离阈值被满足,kubelet会立马杀死pod,而没有优雅终止时段.

以下标识可以被用于配置硬驱离阈值

驱离监视时间间隔

通过 housekeeping-interval 节点状态  kubelet把一个或多个驱离信号映射到的对应的节点状态上

回收节点级别的资源

驱离pod

kubernetes资源调度之LimitRange

LimitRange从字面意义上来看就是对范围进行限制,实际上是对cpu和内存资源使用范围的限制

前面我们讲到过资源配额,资源配额是对整个名称空间的资源的总限制,
是从整体上来限制的,而LimitRange则是对pod和container级别来做限制的

kubectl create namespace default-mem-example

apiVersion: v1
kind: LimitRange
metadata:
  name: mem-limit-range
spec:
  limits:
  - default:
      memory: 512Mi
    defaultRequest:
      memory: 256Mi
    type: Container

kubectl apply -f memory-defaults.yaml --namespace=default-mem-example

apiVersion: v1
kind: Pod
metadata:
  name: default-mem-demo
spec:
  containers:
  - name: default-mem-demo-ctr
    image: nginx

kubectl apply -f memory-defaults-pod.yaml --namespace=default-mem-example

kubectl get pod default-mem-demo --output=yaml --namespace=default-mem-example

仅指定限制,没指定申请

可以看到容器的内存申请值和限制值是一样的.需要注意它并不是LimitRange里的默认值256M

仅声明了申请,没有声明限制

输出信息显示容器的申请值被设置为声明的值.而限制值被设置成了512M,这是LimitRange的默认设置

设置申请和限制值的动机:
在命名空间运行的每一个容器必须有它自己的内存限额（CPU限额）。
在命名空间中所有的容器使用的内存总量（CPU总量）不能超出指定的限额。

注意默认请求值即为创建pod的时候不指定resource申请时默认赋予的值,默认值即为默认限制的上限.
即不指定的时候默认赋予的值.min和max是可以指定的最大值和最小值.并且需要注意的是以上都是Pod级别的.

apiVersion: v1
kind: LimitRange
metadata:
  name: mem-limit-range
  namespace: example
spec:
  limits:
  - default:  # default limit
      memory: 512Mi
      cpu: 2
    defaultRequest:  # default request
      memory: 256Mi
      cpu: 0.5
    max:  # max limit
      memory: 800Mi
      cpu: 3
    min:  # min request
      memory: 100Mi
      cpu: 0.3
    maxLimitRequestRatio:  # max value for limit / request
      memory: 2
      cpu: 2
    type: Container # limit type, support: Container / Pod / PersistentVolumeClaim


kubernetes之创建基于命名空间的内存和cpu限额示例

创建命名空间
kubectl create namespace quota-mem-cpu-example

创建 ResourceQuota
kubectl apply -f quota-mem-cpu.yaml --namespace=quota-mem-cpu-example

资源配额对名称空间quota-mem-cpu-example增加了以下限制:
每一个pod都必须内存申请/限制,cpu申请/限制

这里是针对上面的示例来说的,因为示例中配额同时配置了这4个选项,因此pod必须声明这四个选项


创建一个 pod

kubectl apply -f quota-mem-cpu-pod.yaml --namespace=quota-mem-cpu-example

kubectl get resourcequota mem-cpu-demo --namespace=quota-mem-cpu-example --output=yaml

再创建一个 Pod

apiVersion: v1
kind: Pod
metadata:
  name: quota-mem-cpu-demo-2
spec:
  containers:
  - name: quota-mem-cpu-demo-2-ctr
    image: redis
    resources:
      limits:
        memory: "1Gi"
        cpu: "800m"      
      requests:
        memory: "700Mi"
        cpu: "400m"

kubectl apply -f quota-mem-cpu-pod-2.yaml --namespace=quota-mem-cpu-example

最后输出：
Error from server (Forbidden): error when creating "quota-mem-cpu-pod-2.yaml": pods "quota-mem-cpu-demo-2" is forbidden: exceeded quota: mem-cpu-demo, requested: requests.memory=700Mi, used: requests.memory=600Mi, limited: requests.memory=1Gi


如果想要对pod的资源进行限制,则可以使用LimitRange,使用了LimitRange后,
超过LimitRange限制资源的pod将不会创建,并且如果容器没有指定申请或者限制,
会被赋以LimitRange的默认值.


kubernetes容器探针检测

kubernetes提供了livenessProbe(可用性探针)和readinessProbe(就绪性探针)对容器的健康性进行检测,
当然这仅仅简单的关于可用性方面的探测,实际上我们不仅仅要对容器进行健康检测,
还要对容器内布置的应用进行健康性检测,这不在本篇讨论之列,后面会有专门篇幅来讨论结合APM工具,
grafana和prometheus的应用检测预警机制.

Pod 的生命周期
Pending：表示集群系统正在创建Pod，但是Pod中的container还没有全部被创建，这其中也包含集群为container创建网络，或者下载镜像的时间；
Running：表示pod已经运行在一个节点商量，并且所有的container都已经被创建。但是并不代表所有的container都运行，
它仅仅代表至少有一个container是处于运行的状态或者进程出于启动中或者重启中；
Succeeded：所有Pod中的container都已经终止成功，并且没有处于重启的container；
Failed：所有的Pod中的container都已经终止了，但是至少还有一个container没有被正常的终止(其终止时的退出码不为0)

对于liveness probes的结果也有几个固定的可选项值：

Success：表示通过检测
Failure：表示没有通过检测
Unknown：表示检测没有正常进行

Liveness Probe的种类：

ExecAction：在container中执行指定的命令。当其执行成功时，将其退出码设置为0；
TCPSocketAction：执行一个TCP检查使用container的IP地址和指定的端口作为socket。如果端口处于打开状态视为成功；
HTTPGetAcction：执行一个HTTP默认请求使用container的IP地址和指定的端口以及请求的路径作为url，
用户可以通过host参数设置请求的地址，通过scheme参数设置协议类型(HTTP、HTTPS)如果其响应代码在200~400之间，设为成功。


当前kubelet拥有两个检测器，他们分别对应不通的触发器(根据触发器的结构执行进一步的动作)：

Liveness Probe：表示container是否处于live状态。如果 LivenessProbe失败，
LivenessProbe将会通知kubelet对应的container不健康了。随后kubelet将kill掉 container，
并根据RestarPolicy进行进一步的操作。默认情况下LivenessProbe在第一次检测之前初始化值为 Success，
如果container没有提供LivenessProbe，则也认为是Success；

ReadinessProbe：表示container是否以及处于可接受service请求的状态了。如果ReadinessProbe失败，
endpoints controller将会从service所匹配到的endpoint列表中移除关于这个container的IP地址。
因此对于Service匹配到的 endpoint的维护其核心是ReadinessProbe。
默认Readiness的初始值是Failure，如果一个container没有提供 Readiness则被认为是Success。

对于LivenessProbe和ReadinessProbe用法都一样，拥有相同的参数和相同的监测方式。
initialDelaySeconds：用来表示初始化延迟的时间，也就是告诉监测从多久之后开始运行，单位是秒
periodSeconds:检测的间隔时间,kubernetes每隔一段时间会检测一次,默认为10秒,最小为1秒
timeoutSeconds: 用来表示监测的超时时间，如果超过这个时长后，则认为监测失败
当前对每一个Container都可以设置不同的restartpolicy，有三种值可以设置：

Always: 只要container退出就重新启动
OnFailure: 当container非正常退出后重新启动
Never: 从不进行重新启动

apiVersion: v1
kind: Pod
metadata:
  name: probe-exec
  namespace: coocla
spec:
  containers:
  - name: nginx
    image: nginx
    livenessProbe:
      exec:
        command:
        - cat
        - /tmp/health
      initialDelaySeconds: 5
      timeoutSeconds: 1
---
apiVersion: v1
kind: Pod
metadata:
  name: probe-http
  namespace: coocla
spec:
  containers:
  - name: nginx
    image: nginx
    livenessProbe:
      httpGet:
        path: /
        port: 80
        host: www.baidu.com
        scheme: HTTPS
      initialDelaySeconds: 5
      timeoutSeconds: 1
---
apiVersion: v1
kind: Pod
metadata:
  name: probe-tcp
  namespace: coocla
spec:
  containers:
  - name: nginx
    image: nginx
    livenessProbe:
      initialDelaySeconds: 5
      timeoutSeconds: 1
      tcpSocket:
        port: 80

  
检测方式
exec-命令   TCPSocket   HTTPGet

xec-命令
在用户容器内执行一次命令，如果命令执行的退出码为0，则认为应用程序正常运行，其他任务应用程序运行不正常。

……
  livenessProbe:
    exec:
      command:
      - cat
      - /home/laizy/test/hostpath/healthy
……
TCPSocket
将会尝试打开一个用户容器的Socket连接（就是IP地址：端口）。如果能够建立这条连接，则认为应用程序正常运行，否则认为应用程序运行不正常。

HTTPGet

调用容器内Web应用的web hook，如果返回的HTTP状态码在200和399之间，则认为应用程序正常运行，
否则认为应用程序运行不正常。每进行一次HTTP健康检查都会访问一次指定的URL。

……
  httpGet: #通过httpget检查健康，返回200-399之间，则认为容器正常
    path: / #URI地址
    port: 80 #端口号
    #host: 127.0.0.1 #主机地址
    scheme: HTTP #支持的协议，http或者https
  httpHeaders：’’ #自定义请求的header
……


linessprobe.yaml

apiVersion: v1
kind: ReplicationController
metadata: 
  name: linessprobe
  labels: 
    project: lykops
    app: linessprobe
    version: v1
spec: 
  replicas: 3
  selecttor: 
    project: lykops
    app: linessprobe
    version: v1
    name: linessprobe
  template: 
    metadata: 
      labels: 
        project: lykops
        app: linessprobe
        version: v1
        name: linessprobe
    spec: 
      restartPolicy: Always
      containers: 
      - name: linessprobe
        image: web:apache
        imagePullPolicy: Never
        command: ['sh',"/etc/run.sh" ]
        ports:
        - containerPort: 80
          name: httpd
          protocol: TCP
        readinessProbe:
          httpGet:
            path: /
            port: 80
            scheme: HTTP
          initialDelaySeconds: 120 
          periodSeconds: 15 
          timeoutSeconds: 5
        livenessProbe: 
          httpGet: 
            path: /
            port: 80
            scheme: HTTP
          initialDelaySeconds: 180 
          timeoutSeconds: 5 
          periodSeconds: 15 

linessprobe-svc.yaml

apiVersion: v1
kind: Service
metadata:
  name: linessprobe
  labels:
    project: lykops
    app: linessprobe
    version: v1
spec:
  selector:
    project: lykops
    app: linessprobe
    version: v1
  ports:
  - name: http
    port: 80
    protocol: TCP


kubectl create -f linessprobe-svc.yaml
kubectl create -f linessprobe.yaml

initialDelaySeconds：容器启动后第一次执行探测是需要等待多少秒。
periodSeconds：执行探测的频率。默认是10秒，最小1秒。
timeoutSeconds：探测超时时间。默认1秒，最小1秒。
successThreshold：探测失败后，最少连续探测成功多少次才被认定为成功。默认是1。对于liveness必须是1。最小值是1。
failureThreshold：探测成功后，最少连续探测失败多少次才被认定为失败。默认是3。最小值是1。

查看健康检测的一些信息：
kubectl explain pods.spec.containers.livenessProbe

KIND:     Pod
VERSION:  v1

RESOURCE: livenessProbe <Object>

DESCRIPTION:
     Periodic probe of container liveness. Container will be restarted if the
     probe fails. Cannot be updated. More info:
     https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes

     Probe describes a health check to be performed against a container to
     determine whether it is alive or ready to receive traffic.

FIELDS:
   exec	<Object>
     One and only one of the following should be specified. Exec specifies the
     action to take.

   failureThreshold	<integer>
     Minimum consecutive failures for the probe to be considered failed after
     having succeeded. Defaults to 3. Minimum value is 1.

   httpGet	<Object>
     HTTPGet specifies the http request to perform.

   initialDelaySeconds	<integer> 初始化探测，指定初始化时间
     Number of seconds after the container has started before liveness probes
     are initiated. More info:
     https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes

   periodSeconds	<integer>
     How often (in seconds) to perform the probe. Default to 10 seconds. Minimum
     value is 1.

   successThreshold	<integer>
     Minimum consecutive successes for the probe to be considered successful
     after having failed. Defaults to 1. Must be 1 for liveness. Minimum value
     is 1.

   tcpSocket	<Object>
     TCPSocket specifies an action involving a TCP port. TCP hooks not yet
     supported

   timeoutSeconds	<integer>
     Number of seconds after which the probe times out. Defaults to 1 second.
     Minimum value is 1. More info:
     https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes


1 livenessProbe存活性探测
  1.1 exec探针
  1.2 httpGet探针
2 readinessProbe就绪性探测

第二个案例：

liveness-exec.yaml

apiVersion: v1
kind: Pod
metadata: 
  name: liveness-exec-pod
  namespace: default
spec:
  containers:
  - name: liveness-exec-container
    image: busybox:latest
    imagePullPolicy: IfNotPresent
    command: ["/bin/sh","-c","touch /tmp/healthy; sleep 30; rm -f /tmp/healthy; sleep 3600"]
    livenessProbe:
      exec:
        command: ["test","-e","/tmp/healthy"]
      initialDelaySeconds: 2
      periodSeconds: 3
  
kubectl create -f liveness-exec.yaml
kubectl get pods liveness-exec-pod

liveness-httpget.yaml 

apiVersion: v1
kind: Pod
metadata: 
  name: liveness-httpget-pod
  namespace: default
spec:
  containers:
  - name: liveness-httpget-container
    image: nginx:1.14-alpine
    imagePullPolicy: IfNotPresent
    ports:
    - name: http
      containerPort: 80
    livenessProbe:
      httpGet:
        port: http
        path: /index.html
      initialDelaySeconds: 2
      periodSeconds: 3

periodSeconds：代表每次探测时间间隔
initialDelaySeconds：代表初始化延迟时间，
即在一个容器启动后如果直接开始探测那么很有可能会直接探测失败，需要给一个系统初始化的时间

kubectl create -f liveness-httpget.yaml
kubectl get pods liveness-httpget-pod

Error from server: error dialing backend: dial tcp 192.168.17.132:10250: 
connect: no route to host



kubectl explain pods.spec.containers.readinessProbe

KIND:     Pod
VERSION:  v1

RESOURCE: readinessProbe <Object>

DESCRIPTION:
     Periodic probe of container service readiness. Container will be removed
     from service endpoints if the probe fails. Cannot be updated. More info:
     https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes

     Probe describes a health check to be performed against a container to
     determine whether it is alive or ready to receive traffic.

FIELDS:
   exec	<Object>
     One and only one of the following should be specified. Exec specifies the
     action to take.

   failureThreshold	<integer>
     Minimum consecutive failures for the probe to be considered failed after
     having succeeded. Defaults to 3. Minimum value is 1.

   httpGet	<Object>
     HTTPGet specifies the http request to perform.

   initialDelaySeconds	<integer>
     Number of seconds after the container has started before liveness probes
     are initiated. More info:
     https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes

   periodSeconds	<integer>
     How often (in seconds) to perform the probe. Default to 10 seconds. Minimum
     value is 1.

   successThreshold	<integer>
     Minimum consecutive successes for the probe to be considered successful
     after having failed. Defaults to 1. Must be 1 for liveness and startup.
     Minimum value is 1.

   tcpSocket	<Object>
     TCPSocket specifies an action involving a TCP port. TCP hooks not yet
     supported

   timeoutSeconds	<integer>
     Number of seconds after which the probe times out. Defaults to 1 second.
     Minimum value is 1. More info:
     https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes


rediness-httpget.yaml

apiVersion: v1
kind: Pod
metadata: 
  name: readiness-httpget-pod
  namespace: default
spec:
  containers:
  - name: readiness-httpget-container
    image: nginx
    imagePullPolicy: IfNotPresent
    ports:
    - name: http
      containerPort: 80
    readinessProbe:
      httpGet:
        port: http
        path: /index.html
      initialDelaySeconds: 1
      periodSeconds: 3

kubectl get pods readiness-httpget-pod


docker 脚本打包
#!/bin/bash
KUBE_VERSION=v1.14.3
KUBE_PAUSE_VERSION=3.1
ETCD_VERSION=3.3.10
KUBE_DASHBOARD_VERSION=v1.10.1
username=mirrorgooglecontainers
images=(
        kube-proxy-amd64:${KUBE_VERSION}
        kube-scheduler-amd64:${KUBE_VERSION}
        kube-controller-manager-amd64:${KUBE_VERSION}
        kube-apiserver-amd64:${KUBE_VERSION}
        pause:${KUBE_PAUSE_VERSION}
        etcd-amd64:${ETCD_VERSION}
        )

docker pull quay.io/coreos/flannel:v0.10.0-amd64
docker pull coredns/coredns:1.3.1
docker tag coredns/coredns:1.3.1 k8s.gcr.io/coredns:1.3.1
docker rmi coredns/coredns:1.3.1

for image in ${images[@]}
do
	NEW_IMAGE=`echo ${image}|awk '{gsub(/-amd64/,"",$0);print}'`
	echo ${NEW_IMAGE}
	docker pull ${username}/${image}
 	docker tag ${username}/${image} k8s.gcr.io/${NEW_IMAGE}
	docker rmi ${username}/${image} //删除
done

kubernetes 之QoS服务质量管理

在kubernetes中，每个POD都有个QoS标记，通过这个Qos标记来对POD进行服务质量管理。QoS的英文全称为"Quality of Service",中文名为"服务质量"，它取决于用户对服务质量的预期，也就是期望的服务质量。对于POD来说，服务质量体现在两个指标上，一个指标是CPU，另一个指标是内存。在实际运行过程中，当NODE节点上内存资源紧张的时候，kubernetes根据POD具有的不同QoS标记，采取不同的处理策略。


高
^  +------------------------+
|  |                        |
|  |       Guaranteed       |
|  |                        |
|  +------------------------+
|  |                        |
|  |       Burstable        |
|  |                        |
|  |                        |
|  +------------------------+
|  |                        |
|  |       BestEffort       |
+  |                        |
低  +------------------------+


QoS级别
BestEffort
POD中的所有容器都没有指定CPU和内存的requests和limits，那么这个POD的QoS就是BestEffort级别

Burstable
POD中只要有一个容器，这个容器requests和limits的设置同其他容器设置的不一致，那么这个POD的QoS就是Burstable级别

Guaranteed
POD中所有容器都必须统一设置了limits，并且设置参数都一致，如果有一个容器要设置requests，那么所有容器都要设置，并设置参数同limits一致，那么这个POD的QoS就是Guaranteed级别


QoS级别决定了kubernetes处理这些POD的方式，我们以内存资源为例：

当NODE节点上内存资源不够的时候，QoS级别是BestEffort的POD会最先被kill掉；当NODE节点上内存资源充足的时候，QoS级别是BestEffort的POD可以使用NODE节点上剩余的所有内存资源。

当NODE节点上内存资源不够的时候，如果QoS级别是BestEffort的POD已经都被kill掉了，那么会查找QoS级别是Burstable的POD，并且这些POD使用的内存已经超出了requests设置的内存值，这些被找到的POD会被kill掉；当NODE节点上内存资源充足的时候，QoS级别是Burstable的POD会按照requests和limits的设置来使用。

当NODE节点上内存资源不够的时候，如果QoS级别是BestEffort和Burstable的POD都已经被kill掉了，那么会查找QoS级别是Guaranteed的POD，并且这些POD使用的内存已经超出了limits设置的内存值，这些被找到的POD会被kill掉；当NODE节点上内存资源充足的时候，QoS级别是Burstable的POD会按照requests和limits的设置来使用。

从容器的角度出发，为了限制容器使用的CPU和内存，是通过cgroup来实现的，目前kubernetes的QoS只能管理CPU和内存，所以kubernetes现在也是通过对cgroup的配置来实现QoS管理的。


kubenetes之配置pod的QoS

上节提到过,QoS影响pod的调度和驱离,本节讲解如何通过配置pod来使它自动被赋予一个QoS

实际上是pod的配置达到一定标准,则kubernetes会自动为其它添加一个QoS类

Guaranteed
Burstable
BestEffor

创建一个会被赋予Guaranteed类型QoS的pod

满足以下条件的pod将会被赋予GuaranteedQoS类型
pod中每个容器都必须包含内存请求和限制,并且值相等
pod中每个容器都必须包含cpu请求和限制,并且值相等

kubectl create namespace qos-example


我们来查看它的信息
kubectl get pod qos-demo --namespace=qos-example --output=yaml


kubectl get pod qos-demo --namespace=qos-example --output=yaml
apiVersion: v1
kind: Pod
metadata:
  annotations:
    cni.projectcalico.org/podIP: 10.244.140.94/32
    kubectl.kubernetes.io/last-applied-configuration: |
      {"apiVersion":"v1","kind":"Pod","metadata":{"annotations":{},"name":"qos-demo","namespace":"qos-example"},"spec":{"containers":[{"image":"nginx","name":"qos-demo-ctr","resources":{"limits":{"cpu":"700m","memory":"200Mi"},"requests":{"cpu":"700m","memory":"200Mi"}}}]}}
  creationTimestamp: "2019-12-17T07:38:47Z"
  name: qos-demo
  namespace: qos-example
  resourceVersion: "60079"
  selfLink: /api/v1/namespaces/qos-example/pods/qos-demo
  uid: c6b34419-344c-4d40-a0b6-fca98f4b37ed
spec:
  containers:
  - image: nginx
    imagePullPolicy: Always
    name: qos-demo-ctr
    resources:
      limits:
        cpu: 700m
        memory: 200Mi
      requests:
        cpu: 700m
        memory: 200Mi
    terminationMessagePath: /dev/termination-log
    terminationMessagePolicy: File
    volumeMounts:
    - mountPath: /var/run/secrets/kubernetes.io/serviceaccount
      name: default-token-4k4h7
      readOnly: true
  dnsPolicy: ClusterFirst
  enableServiceLinks: true
  nodeName: node02
  priority: 0
  restartPolicy: Always
  schedulerName: default-scheduler
  securityContext: {}
  serviceAccount: default
  serviceAccountName: default
  terminationGracePeriodSeconds: 30
  tolerations:
  - effect: NoExecute
    key: node.kubernetes.io/not-ready
    operator: Exists
    tolerationSeconds: 300
  - effect: NoExecute
    key: node.kubernetes.io/unreachable
    operator: Exists
    tolerationSeconds: 300
  volumes:
  - name: default-token-4k4h7
    secret:
      defaultMode: 420
      secretName: default-token-4k4h7
status:
  conditions:
  - lastProbeTime: null
    lastTransitionTime: "2019-12-17T07:38:47Z"
    status: "True"
    type: Initialized
  - lastProbeTime: null
    lastTransitionTime: "2019-12-17T07:39:02Z"
    status: "True"
    type: Ready
  - lastProbeTime: null
    lastTransitionTime: "2019-12-17T07:39:02Z"
    status: "True"
    type: ContainersReady
  - lastProbeTime: null
    lastTransitionTime: "2019-12-17T07:38:47Z"
    status: "True"
    type: PodScheduled
  containerStatuses:
  - containerID: docker://0fed050593b54695181eaa0718c80f0bfea3c422a2e8bb42a9c4cfee4a724060
    image: nginx:latest
    imageID: docker-pullable://nginx@sha256:50cf965a6e08ec5784009d0fccb380fc479826b6e0e65684d9879170a9df8566
    lastState: {}
    name: qos-demo-ctr
    ready: true
    restartCount: 0
    started: true
    state:
      running:
        startedAt: "2019-12-17T07:39:02Z"
  hostIP: 192.168.17.132
  phase: Running
  podIP: 10.244.140.94
  podIPs:
  - ip: 10.244.140.94
  qosClass: Guaranteed
  startTime: "2019-12-17T07:38:47Z"

输出信息显示kubernetes给它了一个Guaranteed类型的QoS.
同时也验证了它的内存请求与限制相同,cpu请求与限制也相同.

创建一个会被赋予BurstableQoS类型的pod
Pod不符合Guaranteed类型的QoS要求
pod至少设置了内存或者cpu请任一条件

apiVersion: v1
kind: Pod
metadata:
  name: qos-demo-2
  namespace: qos-example
spec:
  containers:
  - name: qos-demo-2-ctr
    image: nginx
    resources:
      limits:
        memory: "200Mi"
      requests:
        memory: "100Mi"

 kubectl apply -f qos-pod-2.yaml --namespace=qos-example

 kubectl get pod qos-demo-2 --namespace=qos-example --output=yaml

 创建一个会被赋予BestEffortQoS类型的pod
一个pod即没有内存限制或请求也没有cpu限制或请求,则会被赋予BestEffort

apiVersion: v1
kind: Pod
metadata:
  name: qos-demo-3
  namespace: qos-example
spec:
  containers:
  - name: qos-demo-3-ctr
    image: nginx

kubectl apply -f qos-pod-3.yaml --namespace=qos-example

kubectl get pod qos-demo-3 --namespace=qos-example --output=yaml

创建一个包含两个容器的pod
以下配置的pod包含两个容器,其中一个声明了内存限制为200M.另一个则没有声明任何请求或者限制

apiVersion: v1
kind: Pod
metadata:
  name: qos-demo-4
  namespace: qos-example
spec:
  containers:

  - name: qos-demo-4-ctr-1
    image: nginx
    resources:
      requests:
        memory: "200Mi"

  - name: qos-demo-4-ctr-2
    image: redis

以下配置的pod包含两个容器,其中一个声明了内存限制为200M.另一个则没有声明任何请求或者限制


kubectl 技巧之查看资源列表,资源版本和资源 schema 配置


在kubernetes里,pod,service,rs,rc,deploy,resource等对象都需要使用yaml文件来创建,很多时候我们都是参照照官方示例或者一些第三方示例来编写yaml文件以创建对象.虽然这些示例很有典型性和代表性,能够满足我们大部分时候的需求,然而这往往还是不够的,根据项目不同,实际配置可能远比官方提供的demo配置复杂的多,这就要求我们除了掌握常用的配置外,还需要对其它配置有所了解.如果有一个文档能够速查某一对象的所有配置,不但方便我们学习不同的配置,也可以做为一个小手册以便我们记不起来某些配置时可以速查.

下面我们介绍一些小技巧来快速查看kubernetes api

kubectl api-resources

NAME                              SHORTNAMES   APIGROUP                       NAMESPACED   KIND
bindings                                                                      true         Binding
componentstatuses                 cs                                          false        ComponentStatus
configmaps                        cm                                          true         ConfigMap
endpoints                         ep                                          true         Endpoints
events                            ev                                          true         Event
limitranges                       limits                                      true         LimitRange
namespaces                        ns                                          false        Namespace
nodes                             no                                          false        Node
persistentvolumeclaims            pvc                                         true         PersistentVolumeClaim
persistentvolumes                 pv                                          false        PersistentVolume
pods                              po                                          true         Pod
podtemplates                                                                  true         PodTemplate
replicationcontrollers            rc                                          true         ReplicationController
resourcequotas                    quota                                       true         ResourceQuota
secrets                                                                       true         Secret
serviceaccounts                   sa                                          true         ServiceAccount
services                          svc                                         true         Service
mutatingwebhookconfigurations                  admissionregistration.k8s.io   false        MutatingWebhookConfiguration
validatingwebhookconfigurations                admissionregistration.k8s.io   false        ValidatingWebhookConfiguration
customresourcedefinitions         crd,crds     apiextensions.k8s.io           false        CustomResourceDefinition
apiservices                                    apiregistration.k8s.io         false        APIService
controllerrevisions                            apps                           true         ControllerRevision
daemonsets                        ds           apps                           true         DaemonSet
deployments                       deploy       apps                           true         Deployment
replicasets                       rs           apps                           true         ReplicaSet
statefulsets                      sts          apps                           true         StatefulSet
tokenreviews                                   authentication.k8s.io          false        TokenReview
localsubjectaccessreviews                      authorization.k8s.io           true         LocalSubjectAccessReview
selfsubjectaccessreviews                       authorization.k8s.io           false        SelfSubjectAccessReview
selfsubjectrulesreviews                        authorization.k8s.io           false        SelfSubjectRulesReview
subjectaccessreviews                           authorization.k8s.io           false        SubjectAccessReview
horizontalpodautoscalers          hpa          autoscaling                    true         HorizontalPodAutoscaler
cronjobs                          cj           batch                          true         CronJob
jobs                                           batch                          true         Job
certificatesigningrequests        csr          certificates.k8s.io            false        CertificateSigningRequest
leases                                         coordination.k8s.io            true         Lease
bgpconfigurations                              crd.projectcalico.org          false        BGPConfiguration
bgppeers                                       crd.projectcalico.org          false        BGPPeer
blockaffinities                                crd.projectcalico.org          false        BlockAffinity
clusterinformations                            crd.projectcalico.org          false        ClusterInformation
felixconfigurations                            crd.projectcalico.org          false        FelixConfiguration
globalnetworkpolicies                          crd.projectcalico.org          false        GlobalNetworkPolicy
globalnetworksets                              crd.projectcalico.org          false        GlobalNetworkSet
hostendpoints                                  crd.projectcalico.org          false        HostEndpoint
ipamblocks                                     crd.projectcalico.org          false        IPAMBlock
ipamconfigs                                    crd.projectcalico.org          false        IPAMConfig
ipamhandles                                    crd.projectcalico.org          false        IPAMHandle
ippools                                        crd.projectcalico.org          false        IPPool
networkpolicies                                crd.projectcalico.org          true         NetworkPolicy
networksets                                    crd.projectcalico.org          true         NetworkSet
events                            ev           events.k8s.io                  true         Event
ingresses                         ing          extensions                     true         Ingress
ingresses                         ing          networking.k8s.io              true         Ingress
networkpolicies                   netpol       networking.k8s.io              true         NetworkPolicy
runtimeclasses                                 node.k8s.io                    false        RuntimeClass
poddisruptionbudgets              pdb          policy                         true         PodDisruptionBudget
podsecuritypolicies               psp          policy                         false        PodSecurityPolicy
clusterrolebindings                            rbac.authorization.k8s.io      false        ClusterRoleBinding
clusterroles                                   rbac.authorization.k8s.io      false        ClusterRole
rolebindings                                   rbac.authorization.k8s.io      true         RoleBinding
roles                                          rbac.authorization.k8s.io      true         Role
priorityclasses                   pc           scheduling.k8s.io              false        PriorityClass
csidrivers                                     storage.k8s.io                 false        CSIDriver
csinodes                                       storage.k8s.io                 false        CSINode
storageclasses                    sc           storage.k8s.io                 false        StorageClass
volumeattachments                              storage.k8s.io                 false        VolumeAttachment


除了可以看到资源的对象名称外,还可以看到对象的别名,这时候我们再看到别人的命令如kubectl get no这样费解的命令时就可以知道它实际上代表的是kubectl get nodes命令

查看api的版本,很多yaml配置里都需要指定配置的资源版本,
我们经常看到v1,beta1,beta2这样的配置,到底某个资源的最新版本是什么呢?

可以通过kubectl api-versions来查看api的版本

admissionregistration.k8s.io/v1
admissionregistration.k8s.io/v1beta1
apiextensions.k8s.io/v1
apiextensions.k8s.io/v1beta1
apiregistration.k8s.io/v1
apiregistration.k8s.io/v1beta1
apps/v1
authentication.k8s.io/v1
authentication.k8s.io/v1beta1
authorization.k8s.io/v1
authorization.k8s.io/v1beta1
autoscaling/v1
autoscaling/v2beta1
autoscaling/v2beta2
batch/v1
batch/v1beta1
certificates.k8s.io/v1beta1
coordination.k8s.io/v1
coordination.k8s.io/v1beta1
crd.projectcalico.org/v1
events.k8s.io/v1beta1
extensions/v1beta1
networking.k8s.io/v1
networking.k8s.io/v1beta1
node.k8s.io/v1beta1
policy/v1beta1
rbac.authorization.k8s.io/v1
rbac.authorization.k8s.io/v1beta1
scheduling.k8s.io/v1
scheduling.k8s.io/v1beta1
storage.k8s.io/v1
storage.k8s.io/v1beta1
v1


通过 kubectl explain 查看 api 字段

通过kubectl explain <资源名对象名>查看资源对象拥有的字段

如：

KIND:     Deployment
VERSION:  apps/v1

DESCRIPTION:
     Deployment enables declarative updates for Pods and ReplicaSets.

FIELDS:
   apiVersion	<string>
     APIVersion defines the versioned schema of this representation of an
     object. Servers should convert recognized schemas to the latest internal
     value, and may reject unrecognized values. More info:
     https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources

   kind	<string>
     Kind is a string value representing the REST resource this object
     represents. Servers may infer this from the endpoint the client submits
     requests to. Cannot be updated. In CamelCase. More info:
     https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds

   metadata	<Object>
     Standard object metadata.

   spec	<Object>
     Specification of the desired behavior of the Deployment.

   status	<Object>
     Most recently observed status of the Deployment.

列出所有的信息：

kubectl explain svc --recursive

KIND:     Service
VERSION:  v1

DESCRIPTION:
     Service is a named abstraction of software service (for example, mysql)
     consisting of local port (for example 3306) that the proxy listens on, and
     the selector that determines which pods will answer requests sent through
     the proxy.

FIELDS:
   apiVersion	<string>
   kind	<string>
   metadata	<Object>
      annotations	<map[string]string>
      clusterName	<string>
      creationTimestamp	<string>
      deletionGracePeriodSeconds	<integer>
      deletionTimestamp	<string>
      finalizers	<[]string>
      generateName	<string>
      generation	<integer>
      labels	<map[string]string>
      managedFields	<[]Object>
         apiVersion	<string>
         fieldsType	<string>
         fieldsV1	<map[string]>
         manager	<string>
         operation	<string>
         time	<string>
      name	<string>
      namespace	<string>
      ownerReferences	<[]Object>
         apiVersion	<string>
         blockOwnerDeletion	<boolean>
         controller	<boolean>
         kind	<string>
         name	<string>
         uid	<string>
      resourceVersion	<string>
      selfLink	<string>
      uid	<string>
   spec	<Object>
      clusterIP	<string>
      externalIPs	<[]string>
      externalName	<string>
      externalTrafficPolicy	<string>
      healthCheckNodePort	<integer>
      ipFamily	<string>
      loadBalancerIP	<string>
      loadBalancerSourceRanges	<[]string>
      ports	<[]Object>
         name	<string>
         nodePort	<integer>
         port	<integer>
         protocol	<string>
         targetPort	<string>
      publishNotReadyAddresses	<boolean>
      selector	<map[string]string>
      sessionAffinity	<string>
      sessionAffinityConfig	<Object>
         clientIP	<Object>
            timeoutSeconds	<integer>
      type	<string>
   status	<Object>
      loadBalancer	<Object>
         ingress	<[]Object>
            hostname	<string>
            ip	<string>


以上输出的内容是经过格式化了的,我们可以根据缩进很容易看到某一个字段从属于关系

查看具体字段
通过上面kubectl explain service --recursive可以看到所有的api名称,但是以上仅仅是罗列了所有的api名称,如果想要知道某一个api名称的详细信息,则可以通过kubectl explain <资源对象名称.api名称>的方式来查看,比如以下示例可以查看到service下的spec下的ports字段的信息

kubectl explain svc.spec.ports

KIND:     Service
VERSION:  v1

RESOURCE: ports <[]Object>

DESCRIPTION:
     The list of ports that are exposed by this service. More info:
     https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies

     ServicePort contains information on service's port.

FIELDS:
   name	<string>
     The name of this port within the service. This must be a DNS_LABEL. All
     ports within a ServiceSpec must have unique names. When considering the
     endpoints for a Service, this must match the 'name' field in the
     EndpointPort. Optional if only one ServicePort is defined on this service.

   nodePort	<integer>
     The port on each node on which this service is exposed when type=NodePort
     or LoadBalancer. Usually assigned by the system. If specified, it will be
     allocated to the service if unused or else creation of the service will
     fail. Default is to auto-allocate a port if the ServiceType of this Service
     requires one. More info:
     https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport

   port	<integer> -required-
     The port that will be exposed by this service.

   protocol	<string>
     The IP protocol for this port. Supports "TCP", "UDP", and "SCTP". Default
     is TCP.

   targetPort	<string>
     Number or name of the port to access on the pods targeted by the service.
     Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME. If
     this is a string, it will be looked up as a named port in the target Pod's
     container ports. If this is not specified, the value of the 'port' field is
     used (an identity map). This field is ignored for services with
     clusterIP=None, and should be omitted or set equal to the 'port' field.
     More info:
     https://kubernetes.io/docs/concepts/services-networking/service/#defining-a-service



kubernetes 容器编排之定义环境变量以及通过 downwardapi 把 pod 信息作为环境变量传入容器内

