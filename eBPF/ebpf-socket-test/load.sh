#!/bin/bash

set -x
set -e

# 挂载 bpf 文件系统
sudo mount -t bpf bpf /sys/fs/bpf/

# 编译 bpf_sockops 程序
clang -O2 -g -target bpf -c bpf_sockops.c -o bpf_sockops.o

# 加载并附加 bpf_sockops 程序
sudo bpftool prog load bpf_sockops.o /sys/fs/bpf/bpf_sockops
sudo bpftool cgroup attach /sys/fs/cgroup/unified/ sock_ops pinned /sys/fs/bpf/bpf_sockops

# 提取 bpf_sockops 程序使用的 sockhash map 的 id
# 然后这个 map id 被固定到 bpf 虚拟文件系统
MAP_ID=$(sudo bpftool prog show pinned /sys/fs/bpf/bpf_sockops | grep -o -E 'map_ids [0-9]+' | cut -d ' ' -f2-)
sudo bpftool map pin id $MAP_ID /sys/fs/bpf/sock_ops_map

# 加载 bpf_redir 程序并将其附加到 sock_ops_map
clang -O2 -g -Wall -target bpf -c bpf_redir.c -o bpf_redir.o

sudo bpftool prog load bpf_redir.o /sys/fs/bpf/bpf_redir map name sock_ops_map pinned /sys/fs/bpf/sock_ops_map
sudo bpftool prog attach pinned /sys/fs/bpf/bpf_redir msg_verdict pinned /sys/fs/bpf/sock_ops_map