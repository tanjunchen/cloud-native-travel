# golang-ebpf 测试

k8s yaml 如下所示：
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: goapp-ebpf
  labels:
    app: goapp-ebpf
spec:
  replicas: 1
  selector:
    matchLabels:
      app: goapp-ebpf
  template:
    metadata:
      labels:
        app: goapp-ebpf
    spec:
      containers:
        - name: goapp-ebpf
          image: docker.io/tanjunchen/goapp-ebpf:latest
          resources:
            requests:
              cpu: "100m"
          imagePullPolicy: Always
          ports:
            - containerPort: 8080
```

## 构建与编译

```bash
docker build --platform=linux/amd64 -t tanjunchen/goapp-ebpf:latest .
docker push docker.io/tanjunchen/goapp-ebpf:latest
```

## docker 运行

docker run -it -d --name goapp-ebpf -p 8080:8080 docker.io/tanjunchen/goapp-ebpf:latest

