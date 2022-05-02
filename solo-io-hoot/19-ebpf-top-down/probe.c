#include "vmlinux.h"
#include "bpf/bpf_helpers.h"
#include "bpf/bpf_core_read.h"
#include "bpf/bpf_tracing.h"
#include "solo_types.h"

// 1. Change the license if necessary
char __license[] SEC("license") = "Dual MIT/GPL";

struct event_t {
        ipv4_addr daddr;
        u32 pid;
} __attribute__((packed));

// This is the definition for the global map which both our
// bpf program and user space program can access.
// More info and map types can be found here: https://www.man7.org/linux/man-pages/man2/bpf.2.html
struct {
        __uint(max_entries, 1 << 24);
        __uint(type, BPF_MAP_TYPE_RINGBUF);
        __type(value, struct event_t);
} events SEC(".maps.print");

SEC("kprobe/tcp_v4_connect")
int BPF_KPROBE(tcp_v4_connect, struct sock *sk, struct sockaddr *uaddr)
{
        // Init event pointer
        struct event_t *event;
        __u32 daddr;

        struct sockaddr_in *usin = (struct sockaddr_in *)uaddr;
        daddr = BPF_CORE_READ(usin, sin_addr.s_addr);

        // Reserve a spot in the ringbuffer for our event
        event = bpf_ringbuf_reserve(&events, sizeof(struct event_t), 0);
        if (!event) {
                return 0;
        }

        // 3. set data for our event,
        event->daddr = daddr;
        event->pid = bpf_get_current_pid_tgid();

        bpf_ringbuf_submit(event, 0);

        return 0;
}
