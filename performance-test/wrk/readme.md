# wrk 测试

k8s yaml 如下所示：
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: wrk
spec:
  selector:
    matchLabels:
      run: wrk
  replicas: 1
  template:
    metadata:
      labels:
        run: wrk
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
      - name: wrk
        image: docker.io/tanjunchen/wrk:4.2.0
        ports:
        - containerPort: 80
```
