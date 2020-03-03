# 升级 master 节点的 kubeadm

切换到 root

kubeadm config images list

apt-mark unhold kubeadm

apt-get update && apt-get install -y kubeadm=1.17.0-00

apt-mark hold kubeadm

$CP_NODE 为当前升级的节点名称
kubectl drain $CP_NODE --ignore-daemonsets

kubeadm upgrade apply v1.17.0 --ignore-preflight-errors=CoreDNSUnsupportedPlugins

$CP_NODE 为当前升级的节点名称
kubectl uncordon $CP_NODE

# 升级 master 节点的 kubelet 与 kubectl

切换到 root

apt-mark unhold kubelet kubectl
apt-get update && apt-get install -y kubelet=1.17.0-00 kubectl=1.17.0-00
apt-mark hold kubelet kubectl
systemctl restart kubelet

# 升级 worker 节点的 kubeadm 以及 kubelet 与 kubectl

切换到 root

apt-mark unhold kubeadm
apt-get update && apt-get install -y kubeadm=1.17.0-00
apt-mark hold kubeadm

在 master 上 drain 需要升级的 worker 节点

    NODE=k8s-node
    kubectl drain $NODE --ignore-daemonsets
    在 worker 节点上执行升级命令

    获取所有的需要升级的 node 
    // kubectl get nodes -o=jsonpath='{range .items[*]}{.metadata.name}{" "}{end}'

    kubeadm upgrade node

    在 worker 节点上升级 kubelet 与 kubectl

    apt-mark unhold kubelet kubectl && \
    apt-get update && apt-get install -y kubelet=1.17.0-00 kubectl=1.17.0-00 && \
    apt-mark hold kubelet kubectl && \
    systemctl restart kubelet

在 master 上解除 worker 节点的放空

kubectl uncordon $NODE

查看所有的节点是否升级完毕