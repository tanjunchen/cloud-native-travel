# nginx 测试

k8s yaml 如下所示：
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx
  labels:
    app: nginx
spec:
  replicas: 1
  minReadySeconds: 0
  strategy:
    type: RollingUpdate # 策略：滚动更新
    rollingUpdate:
      maxSurge: 25%
      maxUnavailable﻿: 25%
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      restartPolicy: Always
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
        - name: nginx
          image: docker.io/tanjunchen/nginx:1.14.2
          imagePullPolicy: Always
          ports:
            - containerPort: 80
          command:
            - /bin/sh
            - -c
            - "cd /usr/share/nginx/html/ && dd if=/dev/zero of=1k bs=1k count=1 && dd if=/dev/zero of=100k bs=1k count=100 && nginx -g \"daemon off;\""
```
