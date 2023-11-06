# 常见 net 网络调试工具

## Kubernetes 集群登录到某个节点调试工具

使用的 yaml 文件如下所示：
```yaml
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: node-shell-debug
spec:
  selector:
    matchLabels:
      app: node-shell-debug
  template:
    metadata:
      labels:
        app: node-shell-debug
    spec:
      tolerations:
      - effect: NoSchedule
        key: node-role.kubernetes.io/master
        operator: Exists
      containers:
      - args:
        - -t
        - "1"
        - -m
        - -u
        - -i
        - -n
        - sleep
        - "140000000"
        command:
        - nsenter
        image: docker.io/tanjunchen/node-shell:dev
        imagePullPolicy: Always
        name: shell
        securityContext:
          privileged: true
      hostIPC: true
      hostNetwork: true
      hostPID: true
```

原始 dockerfile 文件如下所示：
```dockerfile
FROM ubuntu:latest

# 安装必要的软件包
RUN apt-get update && apt-get install -y \
    curl \
    dnsutils \
    iputils-ping \
    net-tools \
    util-linux \
    && rm -rf /var/lib/apt/lists/*

# 设置工作目录
WORKDIR /

# 启动容器时执行的命令
CMD ["/bin/bash"]
```

## 常见网络调试工具

net 网络调试 yaml 文件如下所示：
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: network
  labels:
    app: network
spec:
  selector:
    matchLabels:
      app: network
  template:
    metadata:
      labels:
        app: network
    spec:
      containers:
      - name: network
        imagePullPolicy: Always
        image: registry.cn-hangzhou.aliyuncs.com/tanjunchen/network-multitool:v1
        env:
        - name: HTTP_PORT
          value: "1180"
        - name: HTTPS_PORT
          value: "11443"
        ports:
        - containerPort: 1180
          name: http-port
        - containerPort: 11443
          name: https-port
        resources:
          requests:
            cpu: "1m"
            memory: "20Mi"
          limits:
            cpu: "10m"
            memory: "20Mi"
        securityContext:
          runAsUser: 0
          capabilities:
            add: ["NET_ADMIN"]
```
更多的详情可参考 [Network-MultiTool](https://github.com/Praqma/Network-MultiTool)。

使用说明见以下文件：
```bash
# Usage - on Docker:
# ------------------
# docker run --rm -it praqma/network-multitool /bin/bash 
# OR
# docker run -d  praqma/network-multitool
# OR
# docker run -p 80:80 -p 443:443 -d  praqma/network-multitool
# OR
# docker run -e HTTP_PORT=1180 -e HTTPS_PORT=11443 -p 1180:1180 -p 11443:11443 -d  praqma/network-multitool


# Usage - on Kubernetes:
# ---------------------
# kubectl run multitool --image=praqma/network-multitool
```
