# 开发并运行你的第一个 eBPF 程序

## 学习内容

eBPF 的开发环境搭建方法、运行原理、编程接口，以及各种类型 eBPF 程序的事件触发机制和应用场景。

## 如何选择 eBPF 开发环境？

在前两讲中，我曾提到，虽然 Linux 内核很早就已经支持了 eBPF，但很多新特性都是在 4.x 版本中逐步增加的。
所以，想要稳定运行 eBPF 程序，内核至少需要 4.9 或者更新的版本。而在开发和学习 eBPF 时，为了体验和掌握最新的 eBPF 特性，我推荐使用更新的 5.x 内核。

作为 eBPF 最重大的改进之一，一次编译到处执行（简称 CO-RE）解决了内核数据结构在不同版本差异导致的兼容性问题。不过，在使用 CO-RE 之前，内核需要开启 CONFIG_DEBUG_INFO_BTF=y 和 CONFIG_DEBUG_INFO=y 这两个编译选项。

为了避免你在首次学习 eBPF 时就去重新编译内核，我推荐使用已经默认开启这些编译选项的发行版，
作为你的开发环境，比如：

```
Ubuntu 20.10+
Fedora 31+
RHEL 8.2+
Debian 11+
```

## 如何搭建 eBPF 开发环境？

虚拟机创建好之后，接下来就需要安装 eBPF 开发和运行所需要的开发工具，这包括：
```
将 eBPF 程序编译成字节码的 LLVM；
C 语言程序编译工具 make；
最流行的 eBPF 工具集 BCC 和它依赖的内核头文件；
与内核代码仓库实时同步的 libbpf；
同样是内核代码提供的 eBPF 程序管理工具 bpftool。你可以执行下面的命令，来安装这些必要的开发工具：
```

## 如何改进第一个 eBPF 程序？

BCC 是一个 BPF 编译器集合，包含了用于构建 BPF 程序的编程框架和库，并提供了大量可以直接使用的工具。使用 BCC 的好处是，它把上述的 eBPF 执行过程通过内置框架抽象了起来，并提供了 Python、C++ 等编程语言接口。这样，你就可以直接通过 Python 语言去跟 eBPF 的各种事件和数据进行交互。

第一步：使用 C 开发一个 eBPF 程序

新建一个 hello.c  文件，并输入下面的内容：
```

int hello_world(void *ctx)
{
    bpf_trace_printk("Hello, World!");
    return 0;
}
```

第二步：使用 Python 和 BCC 库开发一个用户态程序

```

#!/usr/bin/env python3
# 1) import bcc library
from bcc import BPF

# 2) load BPF program
b = BPF(src_file="hello.c")
# 3) attach kprobe
b.attach_kprobe(event="do_sys_openat2", fn_name="hello_world")
# 4) read and print /sys/kernel/debug/tracing/trace_pipe
b.trace_print()
```

第三步：执行 eBPF 程序

```
sudo python3 hello.py
```

## 如何改进第一个 eBPF 程序？

改进的源代码

https://github.com/feiskyer/ebpf-apps/blob/main/bcc-apps/python/trace_open.py


## 总结

搭建 eBPF 的开发环境，安装了 eBPF 开发时常用的工具和依赖库。并且，我从最简单的 Hello World 开始，带你借助 BCC 从零开发了一个跟踪 openat()  系统调用的 eBPF 程序。

通常，开发一个 eBPF 程序需要经过开发 C 语言 eBPF 程序、编译为 BPF 字节码、加载 BPF 字节码到内核、内核验证并运行 BPF 字节码，以及用户程序读取 BPF 映射五个步骤。

使用 BCC 的好处是，它把这几个步骤通过内置框架抽象了起来，并提供了简单易用的 Python 接口，这可以帮你大大简化 eBPF 程序的开发。

除此之外，BCC 提供的一系列工具不仅可以直接用在生产环境中，还是你学习和开发新的 eBPF 程序的最佳参考示例。在课程后续的内容中，我还会带你深入 BCC 的详细使用方法。

trace-open.c
```

// 包含头文件
#include <uapi/linux/openat2.h>
#include <linux/sched.h>

// 定义数据结构
struct data_t {
  u32 pid;
  u64 ts;
  char comm[TASK_COMM_LEN];
  char fname[NAME_MAX];
};

// 定义性能事件映射
BPF_PERF_OUTPUT(events);

// 定义kprobe处理函数
int hello_world(struct pt_regs *ctx, int dfd, const char __user * filename, struct open_how *how)
{
  struct data_t data = { };

  // 获取PID和时间
  data.pid = bpf_get_current_pid_tgid();
  data.ts = bpf_ktime_get_ns();

  // 获取进程名
  if (bpf_get_current_comm(&data.comm, sizeof(data.comm)) == 0)
  {
    bpf_probe_read(&data.fname, sizeof(data.fname), (void *)filename);
  }

  // 提交性能事件
  events.perf_submit(ctx, &data, sizeof(data));
  return 0;
}
```

trace-open.py
```

from bcc import BPF

# 1) load BPF program
b = BPF(src_file="trace-open.c")
b.attach_kprobe(event="do_sys_openat2", fn_name="hello_world")

# 2) print header
print("%-18s %-16s %-6s %-16s" % ("TIME(s)", "COMM", "PID", "FILE"))

# 3) define the callback for perf event
start = 0
def print_event(cpu, data, size):
    global start
    event = b["events"].event(data)
    if start == 0:
            start = event.ts
    time_s = (float(event.ts - start)) / 1000000000
    print("%-18.9f %-16s %-6d %-16s" % (time_s, event.comm, event.pid, event.fname))

# 4) loop with callback to print_event
b["events"].open_perf_buffer(print_event)
while 1:
    try:
        b.perf_buffer_poll()
    except KeyboardInterrupt:
        exit()
```

## 使用 docker 构建测试环境

docker run -it  ubuntu:21.10 -d --privileged=true

apt-get update && apt-get install sudo -y

sudo apt-get install  make clang llvm libelf-dev libbpf-dev bpfcc-tools libbpfcc-dev  -y

sudo apt-get install linux-headers-generic linux-tools-generic  -y

sudo apt-get install vim -y

问题报错：
```
For Ubuntu20.10+sudo apt-get install -y make clang llvm libelf-dev libbpf-dev bpfcc-tools libbpfcc-dev linux-tools-$(uname -r) linux-headers-$(uname -r) 
```

```
root@3e572e0067ec:/# sudo apt-get install -y make clang llvm libelf-dev libbpf-dev bpfcc-tools libbpfcc-dev linux-tools-$(uname -r) linux-headers-$(uname -r)
Reading package lists... Done
Building dependency tree... Done
Reading state information... Done
E: Unable to locate package linux-tools-5.10.25-linuxkit
E: Couldn't find any package by glob 'linux-tools-5.10.25-linuxkit'
E: Couldn't find any package by regex 'linux-tools-5.10.25-linuxkit'
E: Unable to locate package linux-headers-5.10.25-linuxkit
E: Couldn't find any package by glob 'linux-headers-5.10.25-linuxkit'
E: Couldn't find any package by regex 'linux-headers-5.10.25-linuxkit'
```

```
Unable to find kernel headers. Try rebuilding kernel with CONFIG_IKHEADERS=m (module) or installing the kernel development package for your running kernel version.
chdir(/lib/modules/5.10.25-linuxkit/build): No such file or directory
Traceback (most recent call last):
  File "/tanjunchen/ebpf/trace_open.py", line 8, in <module>
    b = BPF(src_file="trace_open.c")
  File "/usr/lib/python3/dist-packages/bcc/__init__.py", line 364, in __init__
    raise Exception("Failed to compile BPF module %s" % (src_file or "<text>"))
Exception: Failed to compile BPF module b'trace_open.c'
```

##  使用 macOS 上 Docker 构建 eBPF ubuntu 测试镜像

eBPF 及其编译器 bcc 需要访问内核（kernel）的某些部分与头文件（header）才能工作。 
下述过程显示如何使用 Docker Desktop for mac 的 linuxkit 主机 VM 执行构建如下的流程。

### 构建镜像

docker build -t docker.io/tanjunchen/ebpf-for-mac:v1 . 

```
FROM docker/for-desktop-kernel:5.10.25-6594e668feec68f102a58011bb42bd5dc07a7a9b AS ksrc

FROM ubuntu:latest

WORKDIR /
COPY --from=ksrc /kernel-dev.tar /
RUN tar xf kernel-dev.tar && rm kernel-dev.tar

RUN apt-get update
RUN apt install -y kmod python3-bpfcc

COPY hello_world.py /root

WORKDIR /root
CMD mount -t debugfs debugfs /sys/kernel/debug && /bin/bash
```


hello_world.py 内容如下所示：
```
#!/usr/bin/env python3

# Original source https://gist.github.com/lizrice/47ad44a15cce912502f8667a403f5649

from bcc import BPF

prog = """
int hello(void *ctx) {
    bpf_trace_printk("Hello world\\n");
    return 0;
}
"""

b = BPF(text=prog)
#b.attach_kprobe(event="sys_clone", fn_name="hello")
b.attach_kprobe(event="__x64_sys_clone", fn_name="hello")
b.trace_print()

# This prints out a trace line every time the clone system call is called

# If you rename hello() to kprobe__sys_clone() you can delete the
# b.attach_kprobe() line, because bcc can work out what event to attach this to
# from the function name.
```

### 运行镜像

需要以特权身份与访问主机的 PID 命名空间运行镜像。

```
docker run -it --rm --privileged -v /lib/modules:/lib/modules:ro  -v /etc/localtime:/etc/localtime:ro  --pid=host docker.io/tanjunchen/ebpf-for-mac:v1
```

Note: /lib/modules probably doesn't exist on your mac host, so Docker will map the volume in from the linuxkit host VM.

Docker 在 Docker hub 上发布了他们的桌面内核，
可能需要更新 Dockerfile 以获得与你的 linuxkit 主机 VM 匹配的最新内核。
即 FROM docker/for-desktop-kernel:5.10.25-6594e668feec68f102a58011bb42bd5dc07a7a9b AS ksrc 内部部分。

## 参考链接

https://ebpf.io/summit-2020-slides/eBPF_Summit_2020-Lightning-Bryce_Kahle-How_and_When_You_Should_Measure_CPU_Overhead_of_eBPF_Programs.pdf

https://code.visualstudio.com/docs/remote/ssh

https://man7.org/linux/man-pages/man2

http://kerneltravel.net/blog/2020/ebpf_ljr_7/

https://github.com/iovisor/bcc/blob/master/docs/reference_guide.md#2-bpf_perf_output

https://github.com/iovisor/bcc/blob/master/docs/kernel-versions.md

