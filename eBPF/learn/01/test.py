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