# 什么是 calico

Calico为容器或虚拟机提供安全的网络连接，它创建一个扁平化的第3层网络，为每个节点分配一个可路由的IP地址。网络中的节点不需要NAT或IP隧道就可以相互通信，因此性能很好，接近于物理网络，不使用网络策略的情况下，可能引入0.01毫秒级别的延迟，带宽影响不明显）。 在需要Overlay的应用场景下，Calico可以支持IP-in-IP隧道，也可以与其它Overlay网络（例如Flannel）配合，IP-in-IP隧道会带来较小的性能下降。

Calico支持动态的网络安全策略（NetPolicy），你可以细粒度的控制容器、虚拟机、物理机端点之间的网络通信。

Calico在每个节点运行一个虚拟路由器（vRouter），vRouter利用Linux内核自带的IP转发功能，工作负载依赖于此路由器和外部通信。节点上的代理组件Felix负责根据分配到节点上的工作负载的IP地址信息为vRouter提供L3转发规则。vRouter基于BIRD实现了边界网关协议（BGP）。通过vRouter，工作负载直接基于物理网络进行通信，甚至可以被分配外网IP并直接暴露到互联网上。

# 集成 K8S

## 安装 Calico

## 网络策略

## 静态 IP

## 高可用

## 外部连接

### 出站连接

### 入站连接

## 使用 RR

## 配置

## calicoctl

参考资料
https://blog.gmem.cc/cni-provided-by-calico

地址链接

https://docs.projectcalico.org/v3.11/getting-started/kubernetes/installation/flannel#installing-with-the-kubernetes-api-datastore-recommended

vxlan(virtual Extensible LAN)虚拟可扩展局域网，是一种overlay的网络技术，使用MAC in UDP的方法进行封装，共50字节的封装报文头。

IPIP 是一种将各 Node 的路由之间做一个 tunnel，再把两个网络连接起来的模式，启用 IPIP 模式时，Calico 将在各 Node 上创建一个名为 "tunl0" 的虚拟网络接口。

centos 7:
uname -a

Linux node01 3.10.0-1062.4.3.el7.x86_64 #1 SMP Wed Nov 13 23:58:53 UTC 2019 x86_64 x86_64 x86_64 GNU/Linux

node01 node02
sudo yum install NetworkManager

sudo systemctl start NetworkManager.service

sudo systemctl enable NetworkManager.service



编辑文件，添加一下内容
/etc/NetworkManager/conf.d/calico.conf

[keyfile]
unmanaged-devices=interface-name:cali*;interface-name:tunl*

初始化操作：
sudo kubeadm init --kubernetes-version=1.16.3 --image-repository registry.aliyuncs.com/google_containers  --pod-network-cidr=10.244.0.0/16 --service-cidr=10.1.0.0/16  

master:

    mkdir -p $HOME/.kube
    sudo cp -i /etc/kubernetes/admin.conf $HOME/.ikube/config
    sudo chown $(id -u):$(id -g) $HOME/.kube/config

node:
    sudo kubeadm join 192.168.17.130:6443 --token 8e6j89.15v4pv57l995gnos \
    --discovery-token-ca-cert-hash sha256:be9a815e6d8d5cf17cd5523265f301913db4d7f3f4108f59518594ec9e66b80b 


--cluster-cidr=<your-pod-cidr> and --allocate-node-cidrs=true
On kubeadm, you can pass --pod-network-cidr=<your-pod-cidr> to kubeadm to set both Kubernetes controller flags 

curl https://docs.projectcalico.org/v3.11/manifests/canal.yaml -O

kubectl get --namespace=kube-system daemonset canal

kubectl get --namespace=kube-system pod -o wide|grep canal

calicoctl get ippool default-ipv4-pool -o yaml

sudo yum list telnet*              列出telnet相关的安装包
sudo yum install telnet-server -y          安装telnet服务
sudo yum install telnet.* -y        安装telnet客户端


# 安装 calicoctl 客户端

wget -O /usr/local/bin/calicoctl https://github.com/projectcalico/calicoctl/releases/download/v3.11.1/calicoctl

现在只安装到主节点：
scp -r ./calicoctl node01@192.168.1.131:/home/node01/ 
scp -r ./calicoctl node02@192.168.1.132:/home/node02/

sudo chmod +x /usr/local/bin/calicoctl

创建 calicoctl.cfg 配置文件

stat $HOME/.kube/config: no such file or directory



calicoctl get nodes
ifconfig
kubectl get node
kubectl get pod -n kube-system
ip a
brctl show
brctl show virbr0
brctl -h
brctl show virbr0
kubectl get pod
kubectl get pod --all-namespaces
docker images | grep nginx
kubectl run nginx --image registry.cn-beijing.aliyuncs.com/mrvolleyball/nginx
kubectl run nginx --image registry.cn-beijing.aliyuncs.com/mrvolleyball/nginx:v1 --replicas 3
kubectl get deployment
kubectl delete deployment nginx
kubectl run nginx --image registry.cn-beijing.aliyuncs.com/mrvolleyball/nginx:v1 --replicas 3
kubectl get pod -o wide
ping 10.244.2.4
ping 10.244.1.3
sudo -s iptables -t nat -L -n
ip r
kubectl get pod -n kube-system
kubectl get pod -n kube-system -o wide
kubectl -n kube-system logs -f canal-2r44f
kubectl -n kube-system logs -f canal-2r44f -c kube-flannel
ip a
sudo -s ip l s dev flannel.1 up

