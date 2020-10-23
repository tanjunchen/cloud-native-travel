服务端

[TanjunchenEchoServer](https://github.com/tanjunchen/TanjunchenEchoServer)

客户端

[TanjunchenEchoClient](https://github.com/tanjunchen/TanjunchenEchoClient)

先启动服务端：

kubectl apply -f istio/example/tcp-demo/src/java/java-tcp-server.yaml

获取服务端暴露的地址信息

kubectl create ns tcp

export TCPIP=$(kubectl get po -l app=tcp-java-echo-server -n tcp -o jsonpath='{.items[0].status.hostIP}')

export TCPPORT=$(kubectl -n test get service tcp-java-echo-server  -o jsonpath='{.spec.ports[?(@.name=="tcp")].nodePort}')

注意替换 client 端中的文件地址

```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: tcp-java-echo-client
spec:
  replicas: 1
  selector:
    matchLabels:
      app: tcp-java-echo-client
  template:
    metadata:
      labels:
        app: tcp-java-echo-client
    spec:
      containers:
        - name: tcp-java-echo-client
          image: registry.cn-hangzhou.aliyuncs.com/tanjunchen/tcp-java-client:v1.0
          imagePullPolicy: IfNotPresent
          command: ["java", "-jar","/client.jar"]
          args: ["1",TCPIP,TCPPORT]  ## 注意替换 TCP TCPPORT 中的值
```

服务端的日志：

```
[root@test-10 test-watch]# kubectl logs -f tcp-java-echo-client-d9784c864-f5b9v -n test
10.20.11.116 == 31888 == 1
客户端： 服务端响应: Server-C
客户端： 服务端响应: Server-A
客户端： 服务端响应: Server-B
客户端： 服务端响应: Server-A
客户端： 服务端响应: Server-A
客户端： 服务端响应: Server-C
客户端： 服务端响应: Server-A
......
```

客户端的日志：

```
[root@test-10 test-watch]# kubectl logs -f tcp-java-echo-server-5cf55fbbd8-blx2m -n test
 ......启动服务端...... 
客户端数量: 1 服务端收到: Client-1
客户端数量: 2 服务端收到: Client-2
客户端数量: 3 服务端收到: Client-3
客户端数量: 4 服务端收到: Client-2
客户端数量: 5 服务端收到: Client-1
客户端数量: 6 服务端收到: Client-3
客户端数量: 7 服务端收到: Client-2
客户端数量: 8 服务端收到: Client-1
......
```

