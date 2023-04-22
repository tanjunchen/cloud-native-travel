#include "vmlinux.h"
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_endian.h>
#include "bpf_sockops.h"

static inline void bpf_sock_ops_active_establish_cb(struct bpf_sock_ops *skops) {
    struct socket_4_tuple key = {};

    sk_ops_extract4_key(skops, &key);
    if (key.local.ip4 == INBOUND_ENVOY_IP) {
        // 调用 BPF 辅助函数去更新套接字 map
        bpf_sock_hash_update(skops, &map_redir, &key, BPF_ANY);
        return;
    }
    if (key.local.ip4 == key.remote.ip4) {
        return;
    }

    // 调用 BPF 辅助函数去更新套接字 map map_active_estab
    /* update map_active_estab*/
    bpf_map_update_elem(&map_active_estab, &key.local, &key.remote, BPF_NOEXIST);

    // 调用 BPF 辅助函数去更新套接字 map map_redir
    /* update map_redir */
    bpf_sock_hash_update(skops, &map_redir, &key, BPF_ANY);
}

static inline void bpf_sock_ops_passive_establish_cb(struct bpf_sock_ops *skops) {
    struct socket_4_tuple key = {};
    struct socket_4_tuple proxy_key = {};
    struct socket_4_tuple proxy_val = {};
    struct addr_2_tuple *original_dst;

    sk_ops_extract4_key(skops, &key);
    if (key.remote.ip4 == INBOUND_ENVOY_IP) {
        // 调用 BPF 辅助函数去更新套接字 map map_redir
        bpf_sock_hash_update(skops, &map_redir, &key, BPF_ANY);
    }
    // 调用 BPF 辅助函数从 bpf_map_lookup_elem 去获取原始地址
    original_dst = bpf_map_lookup_elem(&map_active_estab, &key.remote);
    if (original_dst == NULL) {
        return;
    }
    /* update map_proxy */
    proxy_key.local = key.remote;
    proxy_key.remote = *original_dst;
    proxy_val.local = key.local;
    proxy_val.remote = key.remote;
    // 调用 BPF 辅助函数 key
    bpf_map_update_elem(&map_proxy, &proxy_key, &proxy_val, BPF_ANY);
    // 调用 BPF 辅助函数 key
    bpf_map_update_elem(&map_proxy, &proxy_val, &proxy_key, BPF_ANY);

    /* update map_redir */
    // 调用 BPF 辅助函数去更新套接字 map map_redir
    bpf_sock_hash_update(skops, &map_redir, &key, BPF_ANY);

    /* delete element in map_active_estab*/
    // 调用 BPF 辅助函数去更新套接字 map map_active_estab
    bpf_map_delete_elem(&map_active_estab, &key.remote);
}

static inline void bpf_sock_ops_state_cb(struct bpf_sock_ops *skops) {
    struct socket_4_tuple key = {};
    sk_ops_extract4_key(skops, &key);
    /* delete elem in map_proxy */
    bpf_map_delete_elem(&map_proxy, &key);
    /* delete elem in map_active_estab */
    bpf_map_delete_elem(&map_active_estab, &key.local);
}

// 将函数或者变量放在指定段中，可在指定的地方取函数执行
SEC("sockops")
int bpf_sockmap(struct bpf_sock_ops *skops)
{
    if (!(skops->family == AF_INET || skops->remote_ip4)) {
        /* support dual-stack socket */
        return 0;
    }
    bpf_sock_ops_cb_flags_set(skops, BPF_SOCK_OPS_STATE_CB_FLAG);
    switch (skops->op) {
    case BPF_SOCK_OPS_ACTIVE_ESTABLISHED_CB:
        bpf_sock_ops_active_establish_cb(skops);
        break;
    case BPF_SOCK_OPS_PASSIVE_ESTABLISHED_CB:
        bpf_sock_ops_passive_establish_cb(skops);
        break;
    case BPF_SOCK_OPS_STATE_CB:
        if (skops->args[1] == BPF_TCP_CLOSE) {
            bpf_sock_ops_state_cb(skops);
        }
        break;
    default:
        break;
    }
    return 0;
}

char _license[] SEC("license") = "GPL";
