# wrk2 测试

## 构建与编译

```bash
docker build --platform=linux/amd64 -t tanjunchen/wrk2:latest .
docker push docker.io/tanjunchen/wrk2:latest 
```

k8s yaml 如下所示：
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: wrk2
spec:
  selector:
    matchLabels:
      app:wrk2
  replicas: 1
  template:
    metadata:
      labels:
        app:wrk2
    spec:
      initContainers:
      - name: setsysctl
        image: docker.io/tanjunchen/busybox:latest
        securityContext:
          privileged: true
        command:
        - sh
        - -c
        - |
          sysctl -w net.core.somaxconn=65535
          sysctl -w net.ipv4.ip_local_port_range="1024 65535"
          sysctl -w net.ipv4.tcp_tw_reuse=1
          sysctl -w fs.file-max=1048576
      containers:
      - name: wrk2
        #image: docker.io/tanjunchen/wrk2:latest
        image: docker.io/tanjunchen/haydenjeune-wrk2:latest
        ports:
        - containerPort: 80
```

wrk2 使用的 yaml 文件如下所示：
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: wrk2
spec:
  selector:
    matchLabels:
      app:wrk2
  replicas: 1
  template:
    metadata:
      labels:
        app:wrk2
    spec:
      initContainers:
      - name: setsysctl
        image: docker.io/tanjunchen/busybox:latest
        securityContext:
          privileged: true
        command:
        - sh
        - -c
        - |
          sysctl -w net.core.somaxconn=65535
          sysctl -w net.ipv4.ip_local_port_range="1024 65535"
          sysctl -w net.ipv4.tcp_tw_reuse=1
          sysctl -w fs.file-max=1048576
      containers:
      - name: wrk2
        image: docker.io/tanjunchen/cylab-wrk2:latest
        ports:
        - containerPort: 80
```

## 构建与编译

```bash
docker build --platform=linux/amd64 -t tanjunchen/wrk2:latest .
docker push docker.io/tanjunchen/wrk2:latest 
```

wrk 与 wrk2 同在的 yaml 文件如下所示：
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: wrk-wrk2
spec:
  selector:
    matchLabels:
      app:wrk-wrk2
  replicas: 1
  template:
    metadata:
      labels:
        app:wrk-wrk2
    spec:
      initContainers:
      - name: setsysctl
        image: docker.io/tanjunchen/busybox:latest
        securityContext:
          privileged: true
        command:
        - sh
        - -c
        - |
          sysctl -w net.core.somaxconn=65535
          sysctl -w net.ipv4.ip_local_port_range="1024 65535"
          sysctl -w net.ipv4.tcp_tw_reuse=1
          sysctl -w fs.file-max=1048576
      containers:
      - name: wrk-wrk2
        image: docker.io/tanjunchen/wrk-wrk2:latest
        ports:
        - containerPort: 80
```

## 使用 wrk2 压测

```bash
# threads:4, connections: 100, test duration: 2 minute, CPU cores #4-7
docker run --rm --cpuset-cpus 2-5 -ti docker.io/tanjunchen/wrk-wrk2:latest -t4 -c100 -d120s -R10000 --u_latency http://127.0.0.1:8080/
```

## 参考

1. https://github.com/haydenjeune/wrk2-docker
