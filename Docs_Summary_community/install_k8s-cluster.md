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

亲和与反亲和（Affinity and anti-affinity）

部署pod时，大多数情况下kubernetes的调度程序能将pod调度到集群中合适的节点上。
但有些情况下用户需要对pod调度到哪个节点上施加更多控制，
比如将特定pod部署到拥有SSD存储节点、将同一个服务的多个后端部署在不同的机器上提高安全性、
将通信频繁的服务部署在同一个可用区域降低通信链路长度。
用户对pod部署的节点施加控制都与"label selector"有关。

节点亲和

内部pod亲和与反亲和

kubernetes调度之污点(taint)和容忍(toleration)

可以使用kubectl taint为一个节点(node)添加污点(taint),例如:

kubectl taint nodes node1 key=value:NoSchedule

kubectl taint nodes node1 key:NoSchedule-

基于taint的驱离策略
有条件节点taint

kubernetes之StatefulSet详解

RC、Deployment、DaemonSet都是面向无状态的服务，它们所管理的Pod的IP、名字，启停顺序等都是随机的，
而StatefulSet是什么？顾名思义，有状态的集合，管理所有有状态的服务，比如MySQL、MongoDB集群等。

