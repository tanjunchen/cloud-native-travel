# 核心概念（13％）
# 多容器 Pod（10％）
# Pod 设计（20％）
# 状态持久性（8％）
# 配置（18％）
# 可观察性（18％）
# 服务和网络（13％）

# 了解 Kubernetes API 原语，创建和配置基本 Pod

1. 列出集群中的所有命名空间

kubectl get namespaces
kubectl get ns

2. 列出所有命名空间中的所有 Pod

kubectl get po --all-namespaces

3. 列出特定命名空间中的所有 Pod

kubectl get pod -n kube-system
kubectl get pod -n 命名空间名称

4. 列出特定命名空间中的所有 Service

kubectl get svc --all-namespaces
kubectl get svc -n 命名空间名称
kubectl get svc -n default

5. 用 json 路径表达式列出所有显示名称和命名空间的 Pod

kubectl get pods -o=jsonpath="{. items[*]['metadata. name','metadata. namespace']}" --all-namespaces --sort-by=metadata. name

kubectl get pods -o=jsonpath="{. items[*]['metadata. name','metadata. namespace']}" --all-namespaces

6. 在默认命名空间中创建一个 Nginx Pod，并验证 Pod 是否正在运行

kubectl run nginx --image=nginx
以上命令
kubectl run --generator=deployment/apps. v1 is DEPRECATED and will be removed in a future version.  Use kubectl run --generator=run-pod/v1 or kubectl create instead

kubectl run nginx --image=nginx --restart=Never
会直接产生 nginx  后面不会产生随机数

7. 使用 yaml 文件创建相同的 Nginx Pod

kubectl run nginx --image=nginx --restart=Never --dry-run -o yaml > nginx. yaml

nginx. yaml 为

```
apiVersion: v1
kind: Pod
metadata:
  creationTimestamp: null
  labels:
    run: nginx
  name: nginx
spec:
  containers:
  - image: nginx
    name: nginx
    resources: {}
  dnsPolicy: ClusterFirst
  restartPolicy: Never
status: {}
```

8. 输出刚创建的 Pod 的 yaml 文件

kubectl get po nginx -o yaml

9. 输出刚创建的 Pod 的 yaml 文件，并且其中不包含特定于集群的信息

kubectl get po nginx -o yaml --export

```
Flag --export has been deprecated, This flag is deprecated and will be removed in future. 
apiVersion: v1
kind: Pod
metadata:
  annotations:
    cni. projectcalico. org/podIP: 10. 244. 1. 4/32
    cni. projectcalico. org/podIPs: 10. 244. 1. 4/32
  creationTimestamp: null
  labels:
    run: nginx
  name: nginx
  selfLink: /api/v1/namespaces/default/pods/nginx
spec:
  containers:
  - image: nginx
    imagePullPolicy: Always
    name: nginx
    resources: {}
    terminationMessagePath: /dev/termination-log
    terminationMessagePolicy: File
    volumeMounts:
    - mountPath: /var/run/secrets/kubernetes. io/serviceaccount
      name: default-token-6t2vf
      readOnly: true
  dnsPolicy: ClusterFirst
  enableServiceLinks: true
  nodeName: k8s-node01
  priority: 0
  restartPolicy: Never
  schedulerName: default-scheduler
  securityContext: {}
  serviceAccount: default
  serviceAccountName: default
  terminationGracePeriodSeconds: 30
  tolerations:
  - effect: NoExecute
    key: node. kubernetes. io/not-ready
    operator: Exists
    tolerationSeconds: 300
  - effect: NoExecute
    key: node. kubernetes. io/unreachable
    operator: Exists
    tolerationSeconds: 300
  volumes:
  - name: default-token-6t2vf
    secret:
      defaultMode: 420
      secretName: default-token-6t2vf
status:
  phase: Pending
  qosClass: BestEffort
```

10. 获取刚刚创建的 Pod 的完整详细信息

kubectl describe pod nginx

```
Name:         nginx
Namespace:    default
Priority:     0
Node:         k8s-node01/192. 168. 17. 151
Start Time:   Sun, 05 Jan 2020 18:57:52 -0800
Labels:       run=nginx
Annotations:  cni. projectcalico. org/podIP: 10. 244. 1. 4/32
              cni. projectcalico. org/podIPs: 10. 244. 1. 4/32
Status:       Running
IP:           10. 244. 1. 4
IPs:
  IP:  10. 244. 1. 4
Containers:
  nginx:
    Container ID:   docker://3d9b8b7aba7c10b3f1fbdd470dee702bc5c0e70a46157674574b3b20bd629335
    Image:          nginx
    Image ID:       docker-pullable://nginx@sha256:b2d89d0a210398b4d1120b3e3a7672c16a4ba09c2c4a0395f18b9f7999b768f2
    Port:           <none>
    Host Port:      <none>
    State:          Running
      Started:      Sun, 05 Jan 2020 18:57:57 -0800
    Ready:          True
    Restart Count:  0
    Environment:    <none>
    Mounts:
      /var/run/secrets/kubernetes. io/serviceaccount from default-token-6t2vf (ro)
Conditions:
  Type              Status
  Initialized       True 
  Ready             True 
  ContainersReady   True 
  PodScheduled      True 
Volumes:
  default-token-6t2vf:
    Type:        Secret (a volume populated by a Secret)
    SecretName:  default-token-6t2vf
    Optional:    false
QoS Class:       BestEffort
Node-Selectors:  <none>
Tolerations:     node. kubernetes. io/not-ready:NoExecute for 300s
                 node. kubernetes. io/unreachable:NoExecute for 300s
Events:
  Type    Reason     Age   From                 Message
  ----    ------     ----  ----                 -------
  Normal  Scheduled  14m   default-scheduler    Successfully assigned default/nginx to k8s-node01
  Normal  Pulling    14m   kubelet, k8s-node01  Pulling image "nginx"
  Normal  Pulled     14m   kubelet, k8s-node01  Successfully pulled image "nginx"
  Normal  Created    14m   kubelet, k8s-node01  Created container nginx
  Normal  Started    14m   kubelet, k8s-node01  Started container nginx
```

11. 删除刚创建的 Pod

kubectl delete pod nginx
kubectl delete -f yaml 文件

12. 强制删除刚创建的 Pod

kubectl delete po nginx --grace-period=0 --force

warning: Immediate deletion does not wait for confirmation that the running resource has been terminated.  The resource may continue to run on the cluster indefinitely. 
pod "nginx" force deleted

13. 创建版本为 1. 17. 4 的 Nginx Pod，并将其暴露在端口 80 上

kubectl run nginx --image=nginx:1. 17. 4 --restart=Never --port=80

kubectl run nginx --image=nginx:1. 17. 4 --restart=Never --port=80 --dry-run -o yaml

```
apiVersion: v1
kind: Pod
metadata:
  creationTimestamp: null
  labels:
    run: nginx
  name: nginx
spec:
  containers:
  - image: nginx:1. 17. 4
    name: nginx
    ports:
    - containerPort: 80
    resources: {}
  dnsPolicy: ClusterFirst
  restartPolicy: Never
status: {}
```

14. 将刚创建的容器的镜像更改为 1. 15-alpine，并验证该镜像是否已更新

kubectl set image pod/nginx nginx=nginx:1. 15-alpine

kubectl edit po nginx

kubectl get po nginx -w

15. 对于刚刚更新的 Pod，将镜像版本改回 1. 17. 1，并观察变化

kubectl set image pod/nginx nginx=nginx:1. 17. 1

kubectl edit po nginx

kubectl get po nginx -w

16. 在不用 describe 命令的情况下检查镜像版本

kubectl get po nginx -o=jsonpath='{. spec. containers[]. image}{"\n"}'

17. 创建 Nginx Pod 并在 Pod 上执行简单的 shell

kubectl run nginx --image=nginx --restart=Never
kubectl exec -it nginx /bin/bash

18. 获取刚刚创建的 Pod 的 IP 地址

kubectl get po nginx -o wide

19. 创建一个 busybox Pod，在创建它时运行命令 ls 并检查日志

kubectl run busybox --image=busybox --restart=Never -- ls

kubectl logs busybox

```
bin
dev
etc
home
proc
root
sys
tmp
usr
var
```

20. 如果 Pod 崩溃了，请检查 Pod 的先前日志

kubectl logs busybox -p

21. 用命令 sleep 3600 创建一个 busybox Pod

kubectl run busybox --image=busybox --restart=Never -- /bin/sh -c "sleep 3600"

22. 检查 busybox Pod 中 Nginx Pod 的连接

kubectl get pod nginx -o wide

kubectl exec -it busybox -- wget -o- 上述命令列出的 ip 地址

23. 创建一个能回显消息“How are you”的 busybox Pod，并手动将其删除

kubectl run busybox --image=busybox --restart=Never -it -- echo "How are you"
kubectl delete po busybox

24. 创建一个 Nginx Pod 并列出具有不同复杂度（verbosity）的 Pod

kubectl run nginx --image=nginx --restart=Never --port=80
kubectl get po nginx --v=7
```
I0105 20:54:25.780863  118442 loader.go:375] Config loaded from file:  /home/k8s-master/.kube/config
I0105 20:54:25.790538  118442 round_trippers.go:420] GET https://192.168.17.150:6443/api/v1/namespaces/default/pods/nginx
I0105 20:54:25.790575  118442 round_trippers.go:427] Request Headers:
I0105 20:54:25.790584  118442 round_trippers.go:431]     Accept: application/json;as=Table;v=v1beta1;g=meta.k8s.io, application/json
I0105 20:54:25.790591  118442 round_trippers.go:431]     User-Agent: kubectl/v1.16.3 (linux/amd64) kubernetes/b3cbbae
I0105 20:54:25.797321  118442 round_trippers.go:446] Response Status: 200 OK in 6 milliseconds
NAME    READY   STATUS    RESTARTS   AGE
nginx   1/1     Running   3          91m

```

25. 使用自定义列 PODNAME 和 PODSTATUS 列出 Nginx Pod

kubectl get po nginx -o=custom-columns="POD_NAME:.metadata.name,POD_STATUS:.status.containerStatuses"

```
nginx      [map[containerID:docker://9ff65fc8e8771aa9bc069735a7d9bffa915737cdc8da2f491e330c9c11082ebb image:nginx:latest imageID:docker-pullable://nginx@sha256:b2d89d0a210398b4d1120b3e3a7672c16a4ba09c2c4a0395f18b9f7999b768f2 lastState:map[terminated:map[containerID:docker://ade08d0814b2073079bbad9ebf51b3a3fa2af2df4a411f99b255a78d3b669eeb exitCode:0 finishedAt:2020-01-06T03:34:31Z reason:Completed startedAt:2020-01-06T03:31:59Z]] name:nginx ready:true restartCount:3 started:true state:map[running:map[startedAt:2020-01-06T03:34:31Z]]]]
```

26. 列出所有按名称排序的 Pod

kubectl get po --sort-by=.metadata.name

27. 列出所有按创建时间排序的 Pod

kubectl get po --sort-by=.metadata.creationTimestamp

28. 用“ls; sleep 3600;”“echo Hello World; sleep 3600;”及“echo this is the third container; sleep 3600”三个命令创建一个包含三个 busybox 容器的 Pod，并观察其状态

sudo vim multi-containers.yaml

apiVersion: v1
kind: Pod
metadata:
  creationTimestamp: null
  labels:
    run: busybox
  name: busybox
spec:
  containers:
  - args:
    - bin/sh
    - -c
    - ls; sleep 3600
    image: busybox
    name: busybox1
    resources: {}
  - args:
    - bin/sh
    - -c
    - echo Hello World;sleep 3600
    image: busybox
    name: busybox2
    resources: {}
  - args:
    - bin/sh
    - -c
    - echo this is third containers;sleep 3600
    image: busybox
    name: busybox3
    resources: {}
  dnsPolicy: ClusterFirst
  restartPolicy: Never
status: {}

29. 检查刚创建的每个容器的日志

kubectl get po

kubectl logs busybox -c busybox1
kubectl logs busybox -c busybox2
kubectl logs busybox -c busybox3

30. 检查第二个容器 busybox2 的先前日志（如果有）

kubectl logs busybox -c busybox2 --previous

31. 在上述容器的第三个容器 busybox3 中运行命令 ls

kubectl exec busybox -c busybox3 -- ls

32. 显示以上容器的 metrics，将其放入 file.log 中并进行验证

kubectl top pod busybox --containers > file.log && cat file.log

33. 用主容器 busybox 创建一个 Pod，并执行“while true; do echo ‘Hi I am from Main container’ >> /var/log/index.html; sleep 5; done”，并带有暴露在端口 80 上的 Nginx 镜像的 sidecar 容器。用 emptyDir Volume 将该卷安装在 /var/log 路径（用于 busybox）和 /usr/share/nginx/html 路径（用于nginx容器）。验证两个容器都在运行。

kubectl run muilti-containers-pod --image=busybox --restart=Never --dry-run -o yaml > nulti-containers-pod.yaml

```
apiVersion: v1
kind: Pod
metadata:
  creationTimestamp: null
  labels:
    run: muilti-containers-pod
  name: muilti-containers-pod
spec:
  volumes: 
  - name: var-logs
    emptyDir: {}
  containers:
  - image: busybox
    command: ["/bin/sh"]
    args: ["-c", "while true; do echo 'Hi I am from Main container' >> /var/log/index.html; sleep 5; done"]
    name: main-containers
    resources: {}
    volumeMounts: 
    - name: var-logs
      mountPath: /var/log
  - image: nginx
    name: sidercar-container
    resources: {}
    ports: 
      - containerPort: 80
    volumeMounts: 
    - name: var-logs
      mountPath: /usr/share/nginx/html
  dnsPolicy: ClusterFirst
  restartPolicy: Never
status: {}
```
kubectl create -f multi-containers-pod.yaml

34. 进入两个容器并验证 main.txt 是否存在，并用 curl localhost 从 sidecar 容器中查询 main. txt

kubectl exec -it muilti-containers-pod -c main-container -- sh
cat var/log/index.html
kubectl exec -it muilti-containers-pod -c sidercar-container -- sh
cat /usr/share/nginx/html/index.html

35. 获取带有标签信息的 Pod

kubectl get po --show-labels

36. 创建 5 个 Nginx Pod，其中两个标签为 env = prod，另外三个标签为 env = dev

kubectl run nginx-dev1 --image=nginx --restart=Never --labels=env=dev
kubectl run nginx-dev2 --image=nginx --restart=Never --labels=env=dev
kubectl run nginx-dev3 --image=nginx --restart=Never --labels=env=dev

kubectl run nginx-pro1 --image=nginx --restart=Never --labels=env=pro
kubectl run nginx-pro2 --image=nginx --restart=Never --labels=env=pro

37. 确认所有 Pod 都使用正确的标签创建

kubectl get po --show-labels

```
nginx-dev1                  1/1     Running   0          2m2s    env=dev
nginx-dev2                  1/1     Running   0          63s     env=dev
nginx-dev3                  1/1     Running   0          57s     env=dev
nginx-pro1                  1/1     Running   0          52s     env=pro
nginx-pro2                  1/1     Running   0          46s     env=pro
```

38. 获得带有标签 env=dev 的 Pod

kubectl get pods -l env=dev

```
nginx-dev1   1/1     Running   0          4m55s
nginx-dev2   1/1     Running   0          3m56s
nginx-dev3   1/1     Running   0          3m50s
```

39. 获得带标签 env=dev 的 Pod 并输出标签

kubectl get pods -l env=dev --show-labels

40. 获得带有标签 env=pro 的 Pod

kubectl get pods -l env=pro

```
NAME         READY   STATUS    RESTARTS   AGE
nginx-pro1   1/1     Running   0          8m25s
nginx-pro2   1/1     Running   0          8m19s
```

41. 获得带标签 env=prod 的 Pod 并输出标签

kubectl get pods -l env=pro --show-labels

42. 获取带有标签 env 的 Pod

kubectl get po -L env

43. 获得带标签 env=dev、env=pro 的 Pod

kubectl get po -l 'env in (dev,pro)'

44. 获取带有标签 env=dev 和 env=pro 的 Pod 并输出标签

kubectl get po -l 'env in (dev,pro)' --show-labels

45. 将其中一个容器的标签更改为 env=uat 并列出所有要验证的容器

kubectl label pod/nginx-pro1 env=aa --overwrite
kubectl get pods --show-labels

46. 删除刚才创建的 Pod 标签，并确认所有标签均已删除

kubectl label pod nginx-dev{1..3} env-
kubectl label pod nginx-pro{1..2} env-
kubectl get pods --show-labels

47. 为所有 Pod 添加标签 app = nginx 并验证

kubectl label pod nginx-dev{1..3} app=nginx
kubectl label pod nginx-pro{1..2} app=nginx
kubectl get pods --show-labels

48. 获取所有带有标签的节点（如果使用 minikube，则只会获得主节点）

kubectl get nodes --show-labels

```
NAME         STATUS   ROLES    AGE   VERSION   LABELS
k8s-master   Ready    master   13d   v1.16.3   beta.kubernetes.io/arch=amd64,beta.kubernetes.io/os=linux,kubernetes.io/arch=amd64,kubernetes.io/hostname=k8s-master,kubernetes.io/os=linux,node-role.kubernetes.io/master=
k8s-node01   Ready    <none>   13d   v1.16.3   beta.kubernetes.io/arch=amd64,beta.kubernetes.io/os=linux,kubernetes.io/arch=amd64,kubernetes.io/hostname=k8s-node01,kubernetes.io/os=linux
k8s-node02   Ready    <none>   13d   v1.16.3   beta.kubernetes.io/arch=amd64,beta.kubernetes.io/os=linux,kubernetes.io/arch=amd64,kubernetes.io/hostname=k8s-node02,kubernetes.io/os=linux
```

49. 标记节点（如果正在使用，则为 minikube）nodeName = nginxnode

kubectl label node minikube nodeName=nginxnode

50. 建一个标签为 nginx=dev 的 Pod 并将其部署在此节点上

kubectl label node k8s-node01  nginx=dev
kubectl label node k8s-node02  nginx=pro


参考文献
medium.com/bb-tutorials-and-thoughts/practice-enough-with-these-questions-for-the-ckad-exam

# Pod 设计、状态持久性


52. 使用节点选择器验证已调度的 Pod

pod.yaml
```
apiVersion: v1
kind: Pod
metadata:
  creationTimestamp: null
  labels:
    run: nginx
  name: nginx
spec:
  containers:
  - image: nginx
    name: nginx
    resources: {}
  dnsPolicy: ClusterFirst
  restartPolicy: Never
  nodeSelector: 
  nginx: dev
status: {}
```

在 50 基础上

kubectl describe po nginx | grep Node-Selectors

```
Node-Selectors:  nginx=dev
```

  
53. 验证我们刚刚创建的 Pod Nginx 是否具有 nginx=dev 这个标签

kubectl describe po nginx | grep Labels

```
Labels:       run=nginx
```

54. 用 name=webapp 注释 Pod nginx-dev.*、nginx-pro.*

kubectl annotate po nginx-dev{1..3} name=webapp
kubectl annotate po nginx-pro{1..2} name=webapp

55. 验证已正确注释的 Pod

kubectl describe po nginx-dev{1..3} | grep -i annotations
kubectl describe po nginx-pro{1..2} | grep -i annotations

56. 删除 Pod 上的注释并验证

kubectl annotate po nginx-dev{1..3} name-
kubectl annotate po nginx-pro{1..2} name-
kubectl describe po nginx-dev{1..3} | grep -i annotations
kubectl describe po nginx-pro{1..2} | grep -i annotations

57. 删除到目前为止我们创建的所有 Pod

kubectl delete pod --all

58. 创建一个名为 webapp 的 Deployment，它带有 5 个副本的镜像 Nginx

kubectl create deployment  webapp --image=nginx --dry-run -o yaml > webapp-deployment.yaml
更改 webapp-deployment 的 replicas 为 5

```
apiVersion: apps/v1
kind: Deployment
metadata:
  creationTimestamp: null
  labels:
    app: webapp
  name: webapp
spec:
  replicas: 5
  selector:
    matchLabels:
      app: webapp
  strategy: {}
  template:
    metadata:
      creationTimestamp: null
      labels:
        app: webapp
    spec:
      containers:
      - image: nginx
        name: nginx
        resources: {}
status: {}
```


59. 用标签获取我们刚刚创建的 Deployment

kubectl get deploy webapp --show-labels

```
NAME     READY   UP-TO-DATE   AVAILABLE   AGE     LABELS
webapp   5/5     5            5           3m59s   app=webapp
```

60. 导出该 Deployment 的 yaml 文件

kubectl get deploy webapp -o yaml

```
apiVersion: apps/v1
kind: Deployment
metadata:
  annotations:
    deployment.kubernetes.io/revision: "1"
    kubectl.kubernetes.io/last-applied-configuration: |
      {"apiVersion":"apps/v1","kind":"Deployment","metadata":{"annotations":{},"creationTimestamp":null,"labels":{"app":"webapp"},"name":"webapp","namespace":"default"},"spec":{"replicas":5,"selector":{"matchLabels":{"app":"webapp"}},"strategy":{},"template":{"metadata":{"creationTimestamp":null,"labels":{"app":"webapp"}},"spec":{"containers":[{"image":"nginx","name":"nginx","resources":{}}]}}},"status":{}}
  creationTimestamp: "2020-01-08T05:52:11Z"
  generation: 1
  labels:
    app: webapp
  name: webapp
  namespace: default
  resourceVersion: "30357"
  selfLink: /apis/apps/v1/namespaces/default/deployments/webapp
  uid: defb62c1-1b1d-43db-ab43-5cfe90f67d68
spec:
  progressDeadlineSeconds: 600
  replicas: 5
  revisionHistoryLimit: 10
  selector:
    matchLabels:
      app: webapp
  strategy:
    rollingUpdate:
      maxSurge: 25%
      maxUnavailable: 25%
    type: RollingUpdate
  template:
    metadata:
      creationTimestamp: null
      labels:
        app: webapp
    spec:
      containers:
      - image: nginx
        imagePullPolicy: Always
        name: nginx
        resources: {}
        terminationMessagePath: /dev/termination-log
        terminationMessagePolicy: File
      dnsPolicy: ClusterFirst
      restartPolicy: Always
      schedulerName: default-scheduler
      securityContext: {}
      terminationGracePeriodSeconds: 30
status:
  availableReplicas: 5
  conditions:
  - lastTransitionTime: "2020-01-08T05:52:26Z"
    lastUpdateTime: "2020-01-08T05:52:26Z"
    message: Deployment has minimum availability.
    reason: MinimumReplicasAvailable
    status: "True"
    type: Available
  - lastTransitionTime: "2020-01-08T05:52:11Z"
    lastUpdateTime: "2020-01-08T05:52:29Z"
    message: ReplicaSet "webapp-58867d7bbb" has successfully progressed.
    reason: NewReplicaSetAvailable
    status: "True"
    type: Progressing
  observedGeneration: 1
  readyReplicas: 5
  replicas: 5
  updatedReplicas: 5
```

61. 获取该 Deployment 的 Pod

kubectl get deploy --show-labels

```
NAME               READY   UP-TO-DATE   AVAILABLE   AGE    LABELS
nginx-deployment   3/3     3            3           19h    app=nginx
webapp             5/5     5            5           7m1s   app=webapp
```
kubectl get po -l app=webapp

```
NAME                      READY   STATUS    RESTARTS   AGE
webapp-58867d7bbb-4q5nn   1/1     Running   0          7m29s
webapp-58867d7bbb-9jxc7   1/1     Running   0          7m29s
webapp-58867d7bbb-hxfv6   1/1     Running   0          7m29s
webapp-58867d7bbb-qj2nk   1/1     Running   0          7m29s
webapp-58867d7bbb-wmms4   1/1     Running   0          7m29s
```

62. 将该 Deployment 从 5 个副本扩展到 20 个副本并验证

kubectl scale deploy webapp --replicas=20

尝试 --replicas=1000 机器配置过低 会发生问题

kubectl get po -l app=webapp

63. 获取该 Deployment 的 rollout 状态

kubectl rollout status deploy webapp

64. 获取使用该 Deployment 创建的副本集

kubectl get rs -l app=webapp

65. 获取该 Deployment 的副本集和 Pod 的 yaml

kubectl get rs -l app=webapp -o yaml
kubectl get po -l app=webapp -o yaml

66. 删除刚创建的 Deployment，并查看所有 Pod 是否已被删除

kubectl delete deploy webapp
kubectl get po -l app=webapp -w


67. 使用镜像 nginx：1.17.1 和容器端口 80 创建 webapp Deployment，并验证镜像版本

kubectl create deploy webapp --image=nginx:1.17.1 --dry-run -o yaml > webapp.yaml

```
apiVersion: apps/v1
kind: Deployment
metadata:
  creationTimestamp: null
  labels:
    app: webapp
  name: webapp
spec:
  replicas: 1
  selector:
    matchLabels:
      app: webapp
  strategy: {}
  template:
    metadata:
      creationTimestamp: null
      labels:
        app: webapp
    spec:
      containers:
      - image: nginx:1.17.1
        name: nginx
        ports: 
        - containerPort: 80
        resources: {}
status: {}
```

68. 使用镜像版本 1.17.4 更新 Deployment 并验证

kubectl set image deploy/webapp nginx=nginx:1.17.4
kubectl get deploy  -o wide

```
NAME               READY   UP-TO-DATE   AVAILABLE   AGE    CONTAINERS   IMAGES         SELECTOR
nginx-deployment   3/3     3            3           20h    nginx        nginx:1.7.9    app=nginx
webapp             1/1     1            1           5m7s   nginx        nginx:1.17.4   app=webapp
```

kubectl describe deploy webapp | grep Image

```
  Image:        nginx:1.17.4
```

69. 检查 rollout 历史记录，并确保更新后一切正常

kubectl rollout history deploy webapp

```
deployment.apps/webapp 
REVISION  CHANGE-CAUSE
1         <none>
2         <none>
```
kubectl get deploy webapp --show-labels
kubectl get rs -l app=webapp
kubectl get po -l app=webapp

70. 撤消之前使用版本 1.17.1 的 Deployment，并验证镜像是否还有老版本

kubectl rollout undo deploy webapp
kubectl rollout history deploy webapp
```
deployment.apps/webapp 
REVISION  CHANGE-CAUSE
2         <none>
3         <none>
```
kubectl describe deploy webapp | grep Image

71. 使用镜像版本 1.16.1 更新 Deployment，并验证镜像、检查 rollout 历史记录

kubectl set image deploy/webapp nginx=nginx:1.16.1
kubectl rollout status deploy webapp
```
Waiting for deployment "webapp" rollout to finish: 1 old replicas are pending termination...
Waiting for deployment "webapp" rollout to finish: 1 old replicas are pending termination...
deployment "webapp" successfully rolled out
```
kubectl rollout history deploy webapp

```
deployment.apps/webapp 
REVISION  CHANGE-CAUSE
2         <none>
3         <none>
4         <none>
```
```
kubectl describe deploy webapp | grep Image
  Image:        nginx:1.16.1
```

72. 将 Deployment 更新到镜像 1.17.1 并确认一切正常

kubectl rollout undo deploy webapp --to-revision=3
kubectl rollout history deploy webapp
kubectl describe deploy webapp | grep Image

73. 使用错误的镜像版本 1.100 更新 Deployment，并验证有问题

kubectl set image deploy/webapp nginx=nginx:1.10000
kubectl rollout history deploy webapp
kubectl get pods
kubectl describe po pod名称
Warning  Failed     2s (x2 over 52s)   kubelet, k8s-node01  Error: ImagePullBackOff

74. 撤消使用先前版本的 Deployment，并确认一切正常

kubectl rollout undo deploy webapp
kubectl rollout status deploy webapp
kubectl get pods

75. 检查该 Deployment 的特定修订版本的历史记录

kubectl rollout history deploy webapp --revision=特定版本记录号


76. 暂停 Deployment rollout

kubectl rollout pause deploy  webapp


77. 用最新版本的镜像更新 Deployment，并检查历史记录

kubectl set image deploy/webapp nginx=nginx:lastest
kubectl rollout history deploy webapp

78. 恢复 Deployment rollout

kubectl rollout resume deploy  webapp

79. 检查 rollout 历史记录，确保是最新版本

kubectl rollout history deploy webapp
kubectl rollout history deploy webapp --revision=9

80. 将自动伸缩应用到该 Deployment 中，最少副本数为 10，最大副本数为 20，
目标 CPU 利用率 85%，并验证 hpa 已创建，将副本数从 1 个增加到 10 个

kubectl autoscale deploy webapp --min=10 --max=20 --cpu-percent=85
```
kubectl get hpa
NAME     REFERENCE           TARGETS         MINPODS   MAXPODS   REPLICAS   AGE
webapp   Deployment/webapp   <unknown>/85%   10        20        0          9s
```
kubectl get po -l app=webapp
```
NAME                      READY   STATUS    RESTARTS   AGE
webapp-7668577c8f-98ltc   1/1     Running   0          29s
webapp-7668577c8f-bjmhz   1/1     Running   0          29s
webapp-7668577c8f-cbl2k   1/1     Running   0          29s
webapp-7668577c8f-dnb5z   1/1     Running   0          2m54s
webapp-7668577c8f-dvfrt   1/1     Running   0          29s
webapp-7668577c8f-gcqgt   1/1     Running   0          29s
webapp-7668577c8f-ht8pg   1/1     Running   0          29s
webapp-7668577c8f-nhjzs   1/1     Running   0          29s
webapp-7668577c8f-nsmqd   1/1     Running   0          29s
webapp-7668577c8f-rjg58   1/1     Running   0          29s
```

81. 通过删除刚刚创建的 Deployment 和 hpa 来清理集群

kubectl delete deploy webapp
kubectl delete hpa webapp

82. 用镜像 node 创建一个 Job，并验证是否有对应的 Pod 创建

kubectl create job nodeversion --image=node -- node v
kubectl get job -w
kubectl get pod

83. 获取刚刚创建的 Job 的日志

kubectl logs pod 名称

84. 用镜像 busybox 输出 Job 的 yaml 文件，并回显“Hello I am from job”

kubectl create job hello-job --image=busybox --dry-run -o yaml -- echo "Hello I am from job"

```
apiVersion: batch/v1
kind: Job
metadata:
  creationTimestamp: null
  name: hello-job
spec:
  template:
    metadata:
      creationTimestamp: null
    spec:
      containers:
      - command:
        - echo
        - Hello I am from job
        image: busybox
        name: hello-job
        resources: {}
      restartPolicy: Never
status: {}
```

85. 将上面的 yaml 文件复制到 hello-job.yaml 文件并创建 Job

kubectl create job hello-job --image=busybox --dry-run -o yaml -- echo "Hello I am from job" > hello-job.yaml

kubectl apply -f hello-job.yaml

86. 验证 Job 并创建关联的容器，检查日志

kubectl get pod
kubectl get po
 kubectl logs po 名称

87. 删除我们刚刚创建的 Job

kubectl delete job hello-job

88. 创建一个相同的 Job，并使它一个接一个地运行 10 次

kubectl create job hello-job --image=busybox --dry-run -o yaml -- echo "Hello I am from job" > 10-job.yaml

在 10-job.yaml 添加 completions: 10
```
apiVersion: batch/v1
kind: Job
metadata:
  creationTimestamp: null
  name: hello-job
spec:
  completions: 10
  template:
    metadata:
      creationTimestamp: null
    spec:
      containers:
      - command:
        - echo
        - Hello I am from Job
        image: busybox
        name: hello-job
        resources: {}
      restartPolicy: Never
status: {}
```
kubectl get job -w
kubectl get po
```
NAME          COMPLETIONS   DURATION   AGE
hello-job     9/10          53s        53s
nodeversion   1/1           3m14s      18m
hello-job     10/10         59s        59s
```


89. 运行 10 次，确认已创建 10 个 Pod，并在完成后删除它们

kubectl delete job hello-job

90. 创建相同的 Job 并使它并行运行 10 次

kubectl create job hello-job --image=busybox --dry-run -o yaml -- echo "Hello I am from job" > 10-parallelism-job.yaml

```
apiVersion: batch/v1
kind: Job
metadata:
  creationTimestamp: null
  name: hello-job
spec:
  parallelism: 10
  template:
    metadata:
      creationTimestamp: null
    spec:
      containers:
      - command:
        - echo
        - Hello I am from Job
        image: busybox
        name: hello-job
        resources: {}
      restartPolicy: Never
status: {}
```

91. 并行运行 10 次，确认已创建 10 个 Pod，并在完成后将其删除

kubectl get job -w
kubectl get po


92. 创建一个带有 busybox 镜像的 Cronjob，每分钟打印一次来自 Kubernetes 集群消息的日期和 hello

kubectl create cronjob date-job --image=busybox --schedule="*/1 * * * *" -- bin/sh -c "date; echo Hello from kubernetes cluster"

kubectl get cronjob

kubectl get po

kubectl logs pod 名称

93. 输出上述 cronjob 的 yaml 文件

kubectl get cj date-job -o yaml

```
apiVersion: batch/v1beta1
kind: CronJob
metadata:
  creationTimestamp: "2020-01-08T08:50:52Z"
  name: date-job
  namespace: default
  resourceVersion: "50543"
  selfLink: /apis/batch/v1beta1/namespaces/default/cronjobs/date-job
  uid: 22d49b8d-58ec-468a-b589-d5f60f0030c0
spec:
  concurrencyPolicy: Allow
  failedJobsHistoryLimit: 1
  jobTemplate:
    metadata:
      creationTimestamp: null
      name: date-job
    spec:
      template:
        metadata:
          creationTimestamp: null
        spec:
          containers:
          - command:
            - bin/sh
            - -c
            - date; echo Hello from kubernetes cluster
            image: busybox
            imagePullPolicy: Always
            name: date-job
            resources: {}
            terminationMessagePath: /dev/termination-log
            terminationMessagePolicy: File
          dnsPolicy: ClusterFirst
          restartPolicy: OnFailure
          schedulerName: default-scheduler
          securityContext: {}
          terminationGracePeriodSeconds: 30
  schedule: '*/1 * * * *'
  successfulJobsHistoryLimit: 3
  suspend: false
status:
  lastScheduleTime: "2020-01-08T08:51:00Z"
```

94. 验证 cronJob 为每分钟运行创建一个单独的 Job 和 Pod，并验证 Pod 的日志

kubectl get job
kubectl get po
kubectl logs date-job-<jobid>-<pod>

95. 删除 cronJob，并验证所有关联的 Job 和 Pod 也都被删除

kubectl delete cj date-job
// verify pods and jobs
kubectl get po
kubectl get job

96. 列出集群中的持久卷

kubectl get pv

97. 创建一个名为 task-pv-volume 的 PersistentVolume，其 storgeClassName 为 manual，storage 为 10Gi，accessModes 为 ReadWriteOnce，hostPath 为 /mnt/data

task-pv-volume.yaml

```
apiVersion: v1
kind: PersistentVolume
metadata:
  name: task-pv-volume
  labels: 
    type: local
spec:
  capacity:
    storage: 10Gi
  accessModes:
    - ReadWriteOnce
  storageClassName: manual
  hostPath: 
    path: "/mnt/data"
```
kubectl get pv

```
NAME             CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS      CLAIM   STORAGECLASS   REASON   AGE
task-pv-volume   10Gi       RWO            Retain           Available           manual                  88s
```



98. 创建一个存储至少 3Gi、访问模式为 ReadWriteOnce 的 PersistentVolumeClaim，并确认它的状态是否是绑定的

kubectl create -f task-pv-claim.yaml
kubectl get pvc

```
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: task-pv-claim
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 3Gi
  storageClassName: manual
```

99. 删除我们刚刚创建的持久卷和 PersistentVolumeClaim

kubectl delete pvc task-pv-claim
kubectl delete pv task-pv-volume

100. 使用镜像 Redis 创建 Pod，并配置一个在 Pod 生命周期内可持续使用的卷

redis-storage.yaml
```
apiVersion: v1
kind: Pod
metadata:
  name: redis
spec:
  containers:
  - name: redis
    image: redis
    volumeMounts:
    - name: redis-storage
      mountPath: /data/redis
  volumes:
  - name: redis-storage
    emptyDir: {}
```

101. 在上面的 Pod 中执行操作，并在 /data/redis 路径中创建一个名为 file.txt 的文件，其文本为“This is the file”，然后打开另一个选项卡，再次使用同一 Pod 执行，并验证文件是否在同一路径中

kubectl exec -it redis /bin/sh

cd /data/redis

echo "This is the file" > file.txt

102. 删除上面的 Pod，然后从相同的 yaml 文件再次创建，并验证路径 /data/redis 中是否没有 file.txt

kubectl delete po redis

kubectl apply -f redis-storage.yaml

kubectl exec -it redis /bin/sh

cat /data/redis/file.txt
cat: /data/redis/file.txt: No such file or directory

103. 创建一个名为 task-pv-volume 的 PersistentVolume，其 storgeClassName 为 manual，storage 为 10Gi，accessModes 为 ReadWriteOnce，hostPath 为 /mnt/data；并创建一个存储至少 3Gi、访问模式为 ReadWriteOnce 的 PersistentVolumeClaim，并确认它的状态是否是绑定的

kubectl create -f task-pv-volume.yaml
kubectl create -f task-pv-claim.yaml
kubectl get pv
kubectl get pvc

配置如下所述：
task-pv-volume.yaml
```
apiVersion: v1
kind: PersistentVolume
metadata:
  name: task-pv-volume
  labels: 
    type: local
spec:
  capacity:
    storage: 10Gi
  accessModes:
    - ReadWriteOnce
  storageClassName: manual
  hostPath: 
    path: "/mnt/data"
```
task-pv-claim.yaml
```
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: task-pv-claim
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 3Gi
  storageClassName: manual
```


104. 用容器端口 80 和 PersistentVolumeClaim task-pv-claim 创建一个 Nginx 容器，且具有路径“/usr/share/nginx/html”

task-pv-pod.yaml

apiVersion: v1
kind: Pod
metadata:
  name: task-pv-pod
spec:
  containers:
    - name: task-pv-container
      image: nginx
      ports:
        - containerPort: 80
          name: "http-server"
      volumeMounts:
      - mountPath: "/usr/share/nginx/html"
        name: task-pv-storage
  volumes:
    - name: task-pv-storage
      persistentVolumeClaim:
        claimName: task-pv-claim

kubectl get pv
NAME             CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS   CLAIM                   STORAGECLASS   REASON   AGE
task-pv-volume   10Gi       RWO            Retain           Bound    default/task-pv-claim   manual                  85m

kubectl get pvc
NAME            STATUS   VOLUME           CAPACITY   ACCESS MODES   STORAGECLASS   AGE
task-pv-claim   Bound    task-pv-volume   10Gi       RWO            manual         10m


参考文献
medium.com/bb-tutorials-and-thoughts/practice-enough-with-these-questions-for-the-ckad-exam

