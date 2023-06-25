# eBPF 跟踪数据包

## 测试环境

```
root@instance-00qqerhq:~# docker ps
CONTAINER ID   IMAGE     COMMAND                  CREATED              STATUS              PORTS     NAMES
e2114dbf01fe   nginx     "/docker-entrypoint.…"   About a minute ago   Up About a minute   80/tcp    infallible_archimedes
root@instance-00qqerhq:~# uname -a
Linux instance-00qqerhq 5.8.0-050800-generic #202008022230 SMP Sun Aug 2 22:33:21 UTC 2020 x86_64 x86_64 x86_64 GNU/Linux
root@instance-00qqerhq:~# docker inspect e2114dbf01fe | grep IP
            "LinkLocalIPv6Address": "",
            "LinkLocalIPv6PrefixLen": 0,
            "SecondaryIPAddresses": null,
            "SecondaryIPv6Addresses": null,
            "GlobalIPv6Address": "",
            "GlobalIPv6PrefixLen": 0,
            "IPAddress": "172.17.0.2",
            "IPPrefixLen": 16,
            "IPv6Gateway": "",
                    "IPAMConfig": null,
                    "IPAddress": "172.17.0.2",
                    "IPPrefixLen": 16,
                    "IPv6Gateway": "",
                    "GlobalIPv6Address": "",
                    "GlobalIPv6PrefixLen": 0,
```

## 安装 perf 

源码或者操作系统包管理器安装都行，在此不赘述。

sudo apt install linux-tools-generic

## 安装 bcc

sudo apt-get install bpfcc-tools linux-headers-$(uname -r)

## 测试

perf trace 是 perf 子命令，能够跟踪 packet 路径，跟踪 ping 向 172.17.0.2 容器的包，这里我们只关心 net 事件，忽略系统调用信息：
```
sudo perf trace --no-syscalls --event 'net:*' ping 172.17.0.2  > /dev/null
0.000 ping/48922 net:net_dev_queue:dev=docker0 skbaddr=0xffff893aa3223400 len=98
     0.010 ping/48922 net:net_dev_start_xmit:dev=docker0 queue_mapping=0 skbaddr=0xffff893aa3223400 vlan_tagged=0 vlan_proto=0x0000 vlan_tci=0x0000 protocol=0x0800 ip_summed=0 len=98 data_len=0 network_offset=14 transport_offset_valid=1 transport_offset=34 tx_flags=0 gso_size=0 gso_segs=0 gso_type=0
     0.028 ping/48922 net:net_dev_queue:dev=veth1127065 skbaddr=0xffff893aa3223400 len=98
     0.030 ping/48922 net:net_dev_start_xmit:dev=veth1127065 queue_mapping=0 skbaddr=0xffff893aa3223400 vlan_tagged=0 vlan_proto=0x0000 vlan_tci=0x0000 protocol=0x0800 ip_summed=0 len=98 data_len=0 network_offset=14 transport_offset_valid=1 transport_offset=34 tx_flags=0 gso_size=0 gso_segs=0 gso_type=0
     0.033 ping/48922 net:netif_rx_entry:dev=eth0 napi_id=0x2 queue_mapping=0 skbaddr=0xffff893aa3223400 vlan_tagged=0 vlan_proto=0x0000 vlan_tci=0x0000 protocol=0x0800 ip_summed=0 hash=0x00000000 l4_hash=0 len=84 data_len=0 truesize=768 mac_header_valid=1 mac_header=-14 nr_frags=0 gso_size=0 gso_type=0
......
```

```
root@instance-00qqerhq:~# sudo perf list 'net:*'

List of pre-defined events (to be used in -e):

  net:napi_gro_frags_entry                           [Tracepoint event]
  net:napi_gro_frags_exit                            [Tracepoint event]
  net:napi_gro_receive_entry                         [Tracepoint event]
  net:napi_gro_receive_exit                          [Tracepoint event]
  net:net_dev_queue                                  [Tracepoint event]
  net:net_dev_start_xmit                             [Tracepoint event]
  net:net_dev_xmit                                   [Tracepoint event]
  net:net_dev_xmit_timeout                           [Tracepoint event]
  net:netif_receive_skb                              [Tracepoint event]
  net:netif_receive_skb_entry                        [Tracepoint event]
  net:netif_receive_skb_exit                         [Tracepoint event]
  net:netif_receive_skb_list_entry                   [Tracepoint event]
  net:netif_receive_skb_list_exit                    [Tracepoint event]
  net:netif_rx                                       [Tracepoint event]
  net:netif_rx_entry                                 [Tracepoint event]
  net:netif_rx_exit                                  [Tracepoint event]
  net:netif_rx_ni_entry                              [Tracepoint event]
  net:netif_rx_ni_exit                               [Tracepoint event]


Metric Groups:
```

使用以下命令跟踪事件类型：

```
root@instance-00qqerhq:~# sudo perf trace --no-syscalls               --event 'net:net_dev_queue'               --event 'net:netif_receive_skb_entry'     --event 'net:netif_rx'                    --event 'net:napi_gro_receive_entry'      ping 172.17.0.2  > /dev/null
     0.000 ping/49861 net:net_dev_queue:dev=docker0 skbaddr=0xffff893aa4a03500 len=98
     0.012 ping/49861 net:net_dev_queue:dev=veth1127065 skbaddr=0xffff893aa4a03500 len=98
     0.016 ping/49861 net:netif_rx:dev=eth0 skbaddr=0xffff893aa4a03500 len=84
     0.031 ping/49861 net:net_dev_queue:dev=eth0 skbaddr=0xffff893aa4a03c00 len=98
     0.033 ping/49861 net:netif_rx:dev=veth1127065 skbaddr=0xffff893aa4a03c00 len=84
     0.040 ping/49861 net:netif_receive_skb_entry:dev=docker0 napi_id=0x1 queue_mapping=0 skbaddr=0xffff893aa4a03c00 vlan_tagged=0 vlan_proto=0x0000 vlan_tci=0x0000 protocol=0x0800 ip_summed=2 hash=0x00000000 l4_hash=0 len=84 data_len=0 truesize=768 mac_header_valid=1 mac_header=-14 nr_frags=0 gso_size=0 gso_type=0
```

## 参考

1、https://github.com/yadutaf/tracepkt

2、https://arthurchiao.art/blog/trace-packet-with-tracepoint-perf-ebpf-zh/

## 测试代码

```python
from bcc import BPF

# Hello BPF Program
bpf_text = """
#include <net/inet_sock.h>
#include <bcc/proto.h>

// 1. Attach kprobe to "inet_listen"
int kprobe__inet_listen(struct pt_regs *ctx, struct socket *sock, int backlog)
{
    bpf_trace_printk("Hello World!\\n");
    return 0;
};
"""

# 2. Build and Inject program
b = BPF(text=bpf_text)

# 3. Print debug output
while True:
    print b.trace_readline()
```
