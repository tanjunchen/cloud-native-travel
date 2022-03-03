# 动态存储管理实战 - 搭建 GlusterFS 集群

在 Node 上安装 GlusterFS 客户端

sudo yum install glusterfs glusterfs-fuse


GlusterFS 管理服务容器需要一特权模式运行，在 kube-apiserver 的启动参数中增加：

我用 kubeadm 搭建的 v1.16.3 k8s 集群目前默认添加了这个参数

--allow-privileged=true

/etc/kubernetes/manifests/kube-apiserver.yaml 

```
apiVersion: v1
kind: Pod
metadata:
  creationTimestamp: null
  labels:
    component: kube-apiserver
    tier: control-plane
  name: kube-apiserver
  namespace: kube-system
spec:
  containers:
  - command:
    - kube-apiserver
    - --advertise-address=192.168.17.130
    - --allow-privileged=true
    - --authorization-mode=Node,RBAC
    - --client-ca-file=/etc/kubernetes/pki/ca.crt
    - --enable-admission-plugins=NodeRestriction
    - --enable-bootstrap-token-auth=true
    - --etcd-cafile=/etc/kubernetes/pki/etcd/ca.crt
    - --etcd-certfile=/etc/kubernetes/pki/apiserver-etcd-client.crt
    - --etcd-keyfile=/etc/kubernetes/pki/apiserver-etcd-client.key
    - --etcd-servers=https://127.0.0.1:2379
    - --insecure-port=0
    - --kubelet-client-certificate=/etc/kubernetes/pki/apiserver-kubelet-client.crt
    - --kubelet-client-key=/etc/kubernetes/pki/apiserver-kubelet-client.key
    - --kubelet-preferred-address-types=InternalIP,ExternalIP,Hostname
    - --proxy-client-cert-file=/etc/kubernetes/pki/front-proxy-client.crt
    - --proxy-client-key-file=/etc/kubernetes/pki/front-proxy-client.key
    - --requestheader-allowed-names=front-proxy-client
    - --requestheader-client-ca-file=/etc/kubernetes/pki/front-proxy-ca.crt
    - --requestheader-extra-headers-prefix=X-Remote-Extra-
    - --requestheader-group-headers=X-Remote-Group
    - --requestheader-username-headers=X-Remote-User
    - --secure-port=6443
    - --service-account-key-file=/etc/kubernetes/pki/sa.pub
    - --service-cluster-ip-range=10.96.0.0/12
    - --tls-cert-file=/etc/kubernetes/pki/apiserver.crt
    - --tls-private-key-file=/etc/kubernetes/pki/apiserver.key
    image: registry.aliyuncs.com/google_containers/kube-apiserver:v1.16.3
    imagePullPolicy: IfNotPresent
    livenessProbe:
      failureThreshold: 8
      httpGet:
        host: 192.168.17.130
        path: /healthz
        port: 6443
        scheme: HTTPS
      initialDelaySeconds: 15
      timeoutSeconds: 15
    name: kube-apiserver
    resources:
      requests:
        cpu: 250m
    volumeMounts:
    - mountPath: /etc/ssl/certs
      name: ca-certs
      readOnly: true
    - mountPath: /etc/pki
      name: etc-pki
      readOnly: true
    - mountPath: /etc/kubernetes/pki
      name: k8s-certs
      readOnly: true
  hostNetwork: true
  priorityClassName: system-cluster-critical
  volumes:
  - hostPath:
      path: /etc/ssl/certs
      type: DirectoryOrCreate
    name: ca-certs
  - hostPath:
      path: /etc/pki
      type: DirectoryOrCreate
    name: etc-pki
  - hostPath:
      path: /etc/kubernetes/pki
      type: DirectoryOrCreate
    name: k8s-certs
status: {}
```

给要部署 GlusterFS 管理服务的节点打上 "storagenode=glusterfs" 的标签，是为了将 GlusterFS 容器定向部署到安装了 GlusterFS 的 Node 上：

# 给 Node 打标签
kubectl label node k8s-master storagenode=glusterfs
kubectl label node node01 storagenode=glusterfs
kubectl label node node02 storagenode=glusterfs



安装有 glusterfs 的 node 上加载 dm_thin_pool 核心

sudo modprobe dm_thin_pool

lsmod  | grep thin

glusterfs-daemonset.yaml

```
---
kind: DaemonSet
apiVersion: apps/v1
metadata:
  name: glusterfs
  labels:
    glusterfs: daemonset
  annotations:
    description: GlusterFS DaemonSet
    tags: glusterfs
spec:
  selector:
    matchLabels:
      name: glusterfs
  template:
    metadata:
      labels:
        name: glusterfs
    spec:
      nodeSelector:
        storagenode: glusterfs
      hostNetwork: true
      containers:
      - image: gluster/gluster-centos:latest
        imagePullPolicy: IfNotPresent
        name: glusterfs
        env:
        - name: HOST_DEV_DIR
          value: "/mnt/host-dev"
        - name: GLUSTER_BLOCKD_STATUS_PROBE_ENABLE
          value: "1"
        - name: GB_GLFS_LRU_COUNT
          value: "15"
        - name: TCMU_LOGDIR
          value: "/var/log/glusterfs/gluster-block"
        resources:
          requests:
            memory: 100Mi
            cpu: 100m
        volumeMounts:
        - name: glusterfs-heketi
          mountPath: "/var/lib/heketi"
        - name: glusterfs-run
          mountPath: "/run"
        - name: glusterfs-lvm
          mountPath: "/run/lvm"
        - name: glusterfs-etc
          mountPath: "/etc/glusterfs"
        - name: glusterfs-logs
          mountPath: "/var/log/glusterfs"
        - name: glusterfs-config
          mountPath: "/var/lib/glusterd"
        - name: glusterfs-host-dev
          mountPath: "/mnt/host-dev"
        - name: glusterfs-misc
          mountPath: "/var/lib/misc/glusterfsd"
        - name: glusterfs-block-sys-class
          mountPath: "/sys/class"
        - name: glusterfs-block-sys-module
          mountPath: "/sys/module"
        - name: glusterfs-cgroup
          mountPath: "/sys/fs/cgroup"
          readOnly: true
        - name: glusterfs-ssl
          mountPath: "/etc/ssl"
          readOnly: true
        - name: kernel-modules
          mountPath: "/lib/modules"
          readOnly: true
        securityContext:
          capabilities: {}
          privileged: true
        readinessProbe:
          timeoutSeconds: 3
          initialDelaySeconds: 40
          exec:
            command:
            - "/bin/bash"
            - "-c"
            - "if command -v /usr/local/bin/status-probe.sh; then /usr/local/bin/status-probe.sh readiness; else systemctl status glusterd.service; fi"
          periodSeconds: 25
          successThreshold: 1
          failureThreshold: 50
        livenessProbe:
          timeoutSeconds: 3
          initialDelaySeconds: 40
          exec:
            command:
            - "/bin/bash"
            - "-c"
            - "if command -v /usr/local/bin/status-probe.sh; then /usr/local/bin/status-probe.sh liveness; else systemctl status glusterd.service; fi"
          periodSeconds: 25
          successThreshold: 1
          failureThreshold: 50
      volumes:
      - name: glusterfs-heketi
        hostPath:
          path: "/var/lib/heketi"
      - name: glusterfs-run
      - name: glusterfs-lvm
        hostPath:
          path: "/run/lvm"
      - name: glusterfs-etc
        hostPath:
          path: "/etc/glusterfs"
      - name: glusterfs-logs
        hostPath:
          path: "/var/log/glusterfs"
      - name: glusterfs-config
        hostPath:
          path: "/var/lib/glusterd"
      - name: glusterfs-host-dev
        hostPath:
          path: "/dev"
      - name: glusterfs-misc
        hostPath:
          path: "/var/lib/misc/glusterfsd"
      - name: glusterfs-block-sys-class
        hostPath:
          path: "/sys/class"
      - name: glusterfs-block-sys-module
        hostPath:
          path: "/sys/module"
      - name: glusterfs-cgroup
        hostPath:
          path: "/sys/fs/cgroup"
      - name: glusterfs-ssl
        hostPath:
          path: "/etc/ssl"
      - name: kernel-modules
        hostPath:
          path: "/lib/modules"
```

kubectl apply -f glusterfs-daemonset.yaml 

kubectl describe pod glusterfs-96dz2

```
Name:         glusterfs-96dz2
Namespace:    default
Priority:     0
Node:         k8s-master/192.168.17.130
Start Time:   Wed, 04 Mar 2020 06:20:38 -0800
Labels:       controller-revision-hash=5588b77db7
              name=glusterfs
              pod-template-generation=1
Annotations:  <none>
Status:       Running
IP:           192.168.17.130
IPs:
  IP:           192.168.17.130
Controlled By:  DaemonSet/glusterfs
Containers:
  glusterfs:
    Container ID:   docker://b2efa225f1fff85e73652a6d807d175092a0239165a2f2e0337045194c070b33
    Image:          gluster/gluster-centos:latest
    Image ID:       docker-pullable://gluster/gluster-centos@sha256:8167034b9abf2d16581f3f4571507ce7d716fb58b927d7627ef72264f802e908
    Port:           <none>
    Host Port:      <none>
    State:          Running
      Started:      Wed, 04 Mar 2020 06:20:39 -0800
    Ready:          True
    Restart Count:  0
    Requests:
      cpu:      100m
      memory:   100Mi
    Liveness:   exec [/bin/bash -c if command -v /usr/local/bin/status-probe.sh; then /usr/local/bin/status-probe.sh liveness; else systemctl status glusterd.service; fi] delay=40s timeout=3s period=25s #success=1 #failure=50
    Readiness:  exec [/bin/bash -c if command -v /usr/local/bin/status-probe.sh; then /usr/local/bin/status-probe.sh readiness; else systemctl status glusterd.service; fi] delay=40s timeout=3s period=25s #success=1 #failure=50
    Environment:
      HOST_DEV_DIR:                        /mnt/host-dev
      GLUSTER_BLOCKD_STATUS_PROBE_ENABLE:  1
      GB_GLFS_LRU_COUNT:                   15
      TCMU_LOGDIR:                         /var/log/glusterfs/gluster-block
    Mounts:
      /etc/glusterfs from glusterfs-etc (rw)
      /etc/ssl from glusterfs-ssl (ro)
      /lib/modules from kernel-modules (ro)
      /mnt/host-dev from glusterfs-host-dev (rw)
      /run from glusterfs-run (rw)
      /run/lvm from glusterfs-lvm (rw)
      /sys/class from glusterfs-block-sys-class (rw)
      /sys/fs/cgroup from glusterfs-cgroup (ro)
      /sys/module from glusterfs-block-sys-module (rw)
      /var/lib/glusterd from glusterfs-config (rw)
      /var/lib/heketi from glusterfs-heketi (rw)
      /var/lib/misc/glusterfsd from glusterfs-misc (rw)
      /var/log/glusterfs from glusterfs-logs (rw)
      /var/run/secrets/kubernetes.io/serviceaccount from default-token-hmmcv (ro)
Conditions:
  Type              Status
  Initialized       True 
  Ready             True 
  ContainersReady   True 
  PodScheduled      True 
Volumes:
  glusterfs-heketi:
    Type:          HostPath (bare host directory volume)
    Path:          /var/lib/heketi
    HostPathType:  
  glusterfs-run:
    Type:       EmptyDir (a temporary directory that shares a pod's lifetime)
    Medium:     
    SizeLimit:  <unset>
  glusterfs-lvm:
    Type:          HostPath (bare host directory volume)
    Path:          /run/lvm
    HostPathType:  
  glusterfs-etc:
    Type:          HostPath (bare host directory volume)
    Path:          /etc/glusterfs
    HostPathType:  
  glusterfs-logs:
    Type:          HostPath (bare host directory volume)
    Path:          /var/log/glusterfs
    HostPathType:  
  glusterfs-config:
    Type:          HostPath (bare host directory volume)
    Path:          /var/lib/glusterd
    HostPathType:  
  glusterfs-host-dev:
    Type:          HostPath (bare host directory volume)
    Path:          /dev
    HostPathType:  
  glusterfs-misc:
    Type:          HostPath (bare host directory volume)
    Path:          /var/lib/misc/glusterfsd
    HostPathType:  
  glusterfs-block-sys-class:
    Type:          HostPath (bare host directory volume)
    Path:          /sys/class
    HostPathType:  
  glusterfs-block-sys-module:
    Type:          HostPath (bare host directory volume)
    Path:          /sys/module
    HostPathType:  
  glusterfs-cgroup:
    Type:          HostPath (bare host directory volume)
    Path:          /sys/fs/cgroup
    HostPathType:  
  glusterfs-ssl:
    Type:          HostPath (bare host directory volume)
    Path:          /etc/ssl
    HostPathType:  
  kernel-modules:
    Type:          HostPath (bare host directory volume)
    Path:          /lib/modules
    HostPathType:  
  default-token-hmmcv:
    Type:        Secret (a volume populated by a Secret)
    SecretName:  default-token-hmmcv
    Optional:    false
QoS Class:       Burstable
Node-Selectors:  storagenode=glusterfs
Tolerations:     node.kubernetes.io/disk-pressure:NoSchedule
                 node.kubernetes.io/memory-pressure:NoSchedule
                 node.kubernetes.io/network-unavailable:NoSchedule
                 node.kubernetes.io/not-ready:NoExecute
                 node.kubernetes.io/pid-pressure:NoSchedule
                 node.kubernetes.io/unreachable:NoExecute
                 node.kubernetes.io/unschedulable:NoSchedule
Events:
  Type    Reason     Age   From                 Message
  ----    ------     ----  ----                 -------
  Normal  Scheduled  3m7s  default-scheduler    Successfully assigned default/glusterfs-96dz2 to k8s-master
  Normal  Pulled     3m6s  kubelet, k8s-master  Container image "gluster/gluster-centos:latest" already present on machine
  Normal  Created    3m6s  kubelet, k8s-master  Created container glusterfs
  Normal  Started    3m6s  kubelet, k8s-master  Started container glusterfs
```

kubectl describe ds glusterfs

```
Name:           glusterfs
Selector:       name=glusterfs
Node-Selector:  storagenode=glusterfs
Labels:         glusterfs=daemonset
Annotations:    deprecated.daemonset.template.generation: 1
                description: GlusterFS DaemonSet
                kubectl.kubernetes.io/last-applied-configuration:
                  {"apiVersion":"apps/v1","kind":"DaemonSet","metadata":{"annotations":{"description":"GlusterFS DaemonSet","tags":"glusterfs"},"labels":{"g...
                tags: glusterfs
Desired Number of Nodes Scheduled: 3
Current Number of Nodes Scheduled: 3
Number of Nodes Scheduled with Up-to-date Pods: 3
Number of Nodes Scheduled with Available Pods: 3
Number of Nodes Misscheduled: 0
Pods Status:  3 Running / 0 Waiting / 0 Succeeded / 0 Failed
Pod Template:
  Labels:  name=glusterfs
  Containers:
   glusterfs:
    Image:      gluster/gluster-centos:latest
    Port:       <none>
    Host Port:  <none>
    Requests:
      cpu:      100m
      memory:   100Mi
    Liveness:   exec [/bin/bash -c if command -v /usr/local/bin/status-probe.sh; then /usr/local/bin/status-probe.sh liveness; else systemctl status glusterd.service; fi] delay=40s timeout=3s period=25s #success=1 #failure=50
    Readiness:  exec [/bin/bash -c if command -v /usr/local/bin/status-probe.sh; then /usr/local/bin/status-probe.sh readiness; else systemctl status glusterd.service; fi] delay=40s timeout=3s period=25s #success=1 #failure=50
    Environment:
      HOST_DEV_DIR:                        /mnt/host-dev
      GLUSTER_BLOCKD_STATUS_PROBE_ENABLE:  1
      GB_GLFS_LRU_COUNT:                   15
      TCMU_LOGDIR:                         /var/log/glusterfs/gluster-block
    Mounts:
      /etc/glusterfs from glusterfs-etc (rw)
      /etc/ssl from glusterfs-ssl (ro)
      /lib/modules from kernel-modules (ro)
      /mnt/host-dev from glusterfs-host-dev (rw)
      /run from glusterfs-run (rw)
      /run/lvm from glusterfs-lvm (rw)
      /sys/class from glusterfs-block-sys-class (rw)
      /sys/fs/cgroup from glusterfs-cgroup (ro)
      /sys/module from glusterfs-block-sys-module (rw)
      /var/lib/glusterd from glusterfs-config (rw)
      /var/lib/heketi from glusterfs-heketi (rw)
      /var/lib/misc/glusterfsd from glusterfs-misc (rw)
      /var/log/glusterfs from glusterfs-logs (rw)
  Volumes:
   glusterfs-heketi:
    Type:          HostPath (bare host directory volume)
    Path:          /var/lib/heketi
    HostPathType:  
   glusterfs-run:
    Type:       EmptyDir (a temporary directory that shares a pod's lifetime)
    Medium:     
    SizeLimit:  <unset>
   glusterfs-lvm:
    Type:          HostPath (bare host directory volume)
    Path:          /run/lvm
    HostPathType:  
   glusterfs-etc:
    Type:          HostPath (bare host directory volume)
    Path:          /etc/glusterfs
    HostPathType:  
   glusterfs-logs:
    Type:          HostPath (bare host directory volume)
    Path:          /var/log/glusterfs
    HostPathType:  
   glusterfs-config:
    Type:          HostPath (bare host directory volume)
    Path:          /var/lib/glusterd
    HostPathType:  
   glusterfs-host-dev:
    Type:          HostPath (bare host directory volume)
    Path:          /dev
    HostPathType:  
   glusterfs-misc:
    Type:          HostPath (bare host directory volume)
    Path:          /var/lib/misc/glusterfsd
    HostPathType:  
   glusterfs-block-sys-class:
    Type:          HostPath (bare host directory volume)
    Path:          /sys/class
    HostPathType:  
   glusterfs-block-sys-module:
    Type:          HostPath (bare host directory volume)
    Path:          /sys/module
    HostPathType:  
   glusterfs-cgroup:
    Type:          HostPath (bare host directory volume)
    Path:          /sys/fs/cgroup
    HostPathType:  
   glusterfs-ssl:
    Type:          HostPath (bare host directory volume)
    Path:          /etc/ssl
    HostPathType:  
   kernel-modules:
    Type:          HostPath (bare host directory volume)
    Path:          /lib/modules
    HostPathType:  
Events:
  Type    Reason            Age    From                  Message
  ----    ------            ----   ----                  -------
  Normal  SuccessfulCreate  6m10s  daemonset-controller  Created pod: glusterfs-w8mfq
  Normal  SuccessfulCreate  6m10s  daemonset-controller  Created pod: glusterfs-96dz2
  Normal  SuccessfulCreate  6m10s  daemonset-controller  Created pod: glusterfs-r6gk6
```

# 部署 Heketi 服务

Hekeit 是一个提供 RESTful API 管理 GlusterFS 卷的框架，并能够在 OpenStack、Kubernetes、OpenShift 等云平台上实现动态存储资源供应，支持 GlusterFS 多集群管理，便于管理员对 GlusterFS 进行操作。

heketi-service-account.yaml

```
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: heketi-service-account
```  

heketi-deploy-svc.yaml

```
---
apiVersion: v1
kind: Service
metadata:
  name: deploy-heketi
  labels:
    glusterfs: heketi-service
    deploy-heketi: support
  annotations:
    description: Exposes Heketi Service
spec:
  selector:
    name: deploy-heketi
  ports:
  - name: deploy-heketi
    port: 8080
    targetPort: 8080
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: deploy-heketi
  labels:
    glusterfs: heketi-deployment
    deploy-heketi: heket-deployment
  annotations:
    description: Defines how to deploy Heketi
spec:
  selector: 
    matchLabels: 
      name: deploy-heketi
  replicas: 1
  template:
    metadata:
      name: deploy-heketi
      labels:
        glusterfs: heketi-pod
        name: deploy-heketi
    spec:
      serviceAccountName: heketi-service-account
      containers:
      - image: heketi/heketi
        imagePullPolicy: IfNotPresent
        name: deploy-heketi
        env:
        - name: HEKETI_EXECUTOR
          value: kubernetes
        - name: HEKETI_FSTAB
          value: "/var/lib/heketi/fstab"
        - name: HEKETI_SNAPSHOT_LIMIT
          value: '14'
        - name: HEKETI_KUBE_GLUSTER_DAEMONSET
          value: "y"
        ports:
        - containerPort: 8080
        volumeMounts:
        - name: db
          mountPath: "/var/lib/heketi"
        readinessProbe:
          timeoutSeconds: 3
          initialDelaySeconds: 3
          httpGet:
            path: "/hello"
            port: 8080
        livenessProbe:
          timeoutSeconds: 3
          initialDelaySeconds: 30
          httpGet:
            path: "/hello"
            port: 8080
      volumes:
      - name: db
        hostPath:
          path: "/heketi-data"
```

# 为 Heketi 设置 GlusterFS 集群

在 Heketi 能够管理 GlusterFS 集群之前，首先要为其设置 GlusterFS 集群的信息。可以用一个 topology.json 配置文件来完成各个 GlusterFS 节点和设备的定义。Heketi 要求在一个 GlusterFS 集群中至少有 3 个节点。在 topology.json 配置文件 hostnmaes 字段的 manage 上填写主机名，在 storage 上填写 IP 地址，devices 要求为未创建文件系统的裸设备（可以有多块盘），以供 Heketi 自动完成 PV（Physical Volume）、VG（Volume Group）和 LV（Logical Volume）的创建。topology.json 文件内容如下：

topology.json

```
{
  "clusters": [
    {
      "nodes": [
        {
          "node": {
            "hostnames": {
              "manage": [
                "k8s-master"
              ],
              "storage": [
                "192.168.17.130"
              ]
            },
            "zone": 1
          },
          "devices": [
            "/dev/sdb1"
          ]
        },
        {
          "node": {
            "hostnames": {
              "manage": [
                "node01"
              ],
              "storage": [
                "192.168.17.131"
              ]
            },
            "zone": 1
          },
          "devices": [
            "/dev/sdb1"
          ]
        },
        {
          "node": {
            "hostnames": {
              "manage": [
                "node02"
              ],
              "storage": [
                "192.168.17.132"
              ]
            },
            "zone": 1
          },
          "devices": [
            "/dev/sdb1"
          ]
        }
      ]
    }
  ]
}
```

进入 Heketi 容器，使用命令行工具 heketi-cli 完成 GlusterFS 创建集群之前需要做下面的操作：

1. 建立该服务帐户控制 gluster pod 的能力

kubectl create clusterrolebinding heketi-gluster-admin --clusterrole=edit --serviceaccount=default:heketi-service-account

1. 添加新的硬盘，使其成为原始块设备



1. 