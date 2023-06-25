# 使用 ebpf 监听 tcp_sendmsg 事件案例

## 环境信息

参考 readme 中准备 ebpf 调试环境

## 测试过程

安装 Python3 bcc 辅助工具：
```
sudo apt install python3-bpfcc
```

执行 python 测试代码：
```
pytho3 test.py
```

## 测试结果

```
root@ebpf-test:~# python3 test.py
tcp_sendmsg: 1
tcp_sendmsg: 2
tcp_sendmsg: 3
tcp_sendmsg: 4
tcp_sendmsg: 20
tcp_sendmsg: 21
tcp_sendmsg: 22
```

## 代码

```python
#!/usr/bin/python3

from bcc import BPF
from time import sleep

# 定义 eBPF 程序
bpf_text = """
#include <uapi/linux/ptrace.h>

BPF_HASH(stats, u32);

int count(struct pt_regs *ctx) {
    u32 key = 0;
    u64 *val, zero=0;
    val = stats.lookup_or_init(&key, &zero);
    (*val)++;
    return 0;
}
"""

# 编译 eBPF 程序
b = BPF(text=bpf_text, cflags=["-Wno-macro-redefined"])

# 加载 eBPF 程序
b.attach_kprobe(event="tcp_sendmsg", fn_name="count")

name = {
  0: "tcp_sendmsg"
}
# 输出统计结果
while True:
    try:
        #print("Total packets: %d" % b["stats"][0].value)
        for k, v in b["stats"].items():
           print("{}: {}".format(name[k.value], v.value))
        sleep(1)
    except KeyboardInterrupt:
        exit()
```
