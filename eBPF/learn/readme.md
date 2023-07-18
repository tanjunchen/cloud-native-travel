# 准备基础环境

## 学习环境说明

```
root@ebpf-test:~# uname -a
Linux ebpf-test 5.8.0-050800-generic #202008022230 SMP Sun Aug 2 22:33:21 UTC 2020 x86_64 x86_64 x86_64 GNU/Linux
root@ebpf-test:~# uname -r
5.8.0-050800-generic
```

## 升级 Linux 内核

```
wget https://raw.githubusercontent.com/pimlie/ubuntu-mainline-kernel.sh/master/ubuntu-mainline-kernel.sh
chmod +x ubuntu-mainline-kernel.sh
bash ubuntu-mainline-kernel.sh --no-checksum  -i 5.8.0 
```

## 安装 eBPF 工具

```
sudo apt-get update -y && sudo apt-get upgrade -y

sudo apt-get install -y git cmake make gcc python3 libncurses-dev gawk flex bison openssl libssl-dev dkms libelf-dev libudev-dev libpci-dev libiberty-dev autoconf  

sudo apt-get install -y binutils-dev clang gcc-multilib 
```

安装 bpftool 5.8 版本：
```
git clone -b v5.8 https://github.com/torvalds/linux.git --depth 1
cd /root/linux/tools/bpf/bpftool
make && make install
```
