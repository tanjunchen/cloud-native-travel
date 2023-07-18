#!/bin/bash
set -x

# 卸载 bpf_redir 程序
sudo bpftool prog detach pinned /sys/fs/bpf/bpf_redir msg_verdict pinned /sys/fs/bpf/sock_ops_map
sudo rm /sys/fs/bpf/bpf_redir

# 卸载 bpf_sockops_v4 程序
sudo bpftool cgroup detach /sys/fs/cgroup/unified/ sock_ops pinned /sys/fs/bpf/bpf_sockops
sudo rm /sys/fs/bpf/bpf_sockops

# 删除 map
sudo rm /sys/fs/bpf/sock_ops_map