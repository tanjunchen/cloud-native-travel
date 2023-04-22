// 如果宏没有定义，则定义AF_INET
#ifndef AF_INET
#define AF_INET 2
#endif

// 如果宏没有定义，则定义AF_INET
#ifndef NULL
#define NULL ((void*)0)
#endif

// 如果宏定义 16 进制
#define INBOUND_ENVOY_IP 0x600007f
// 宏定义 SOCKOPS_MAP_SIZE 最大值
#define SOCKOPS_MAP_SIZE 65535

#include <bpf/bpf_endian.h>

struct addr_2_tuple {
    uint32_t ip4;
    uint32_t port;
};

struct socket_4_tuple {
    struct addr_2_tuple local;
    struct addr_2_tuple remote;
};

/* when active establish, record local addr as key and remote addr as value
|--------------------------------------------------------------------|
|   key(local ip, local port)   |     Val(remote ip, remoteport)     |
|--------------------------------------------------------------------|
|        A-ip,A-app-port        |    B-cluster-ip,B-cluster-port     |
|--------------------------------------------------------------------|
|       A-ip,A-envoy-port       |              B-ip,B-port           |
|--------------------------------------------------------------------|
*/
struct {
        __uint(type, BPF_MAP_TYPE_HASH);
        __uint(max_entries, SOCKOPS_MAP_SIZE);
        __type(key, struct addr_2_tuple);
        __type(value, struct addr_2_tuple);
        __uint(pinning, LIBBPF_PIN_BY_NAME);
} map_active_estab SEC(".maps");

/* This is a proxy map to store current socket 4-tuple and other side socket 4-tuple
|-------------------------------------------------------------------------------------------|
|          key(current socket 4-tuple)        |        Val(other side socket 4-tuple)       |
|-------------------------------------------------------------------------------------------|
| A-ip,A-app-port,B-cluster-ip,B-cluster-port |    127.0.0.1,A-outbound,A-ip:A-app-port     |
|-------------------------------------------------------------------------------------------|
|   127.0.0.1,A-outbound,A-ip:A-app-port      | A-ip:A-app-port,B-cluster-ip,B-cluster-port |
|-------------------------------------------------------------------------------------------|
*/

struct {
        __uint(type, BPF_MAP_TYPE_HASH);
        __uint(max_entries, SOCKOPS_MAP_SIZE);
        __type(key, struct socket_4_tuple);
        __type(value, struct socket_4_tuple);
        __uint(pinning, LIBBPF_PIN_BY_NAME);
} map_proxy SEC(".maps");

/* This is a sockhash map for sk_msg redirect
|------------------------------------------------------------------------|
|  key(local_ip:local_port, remote_ip:remote_port) |     Val(skops)      |
|------------------------------------------------------------------------|
|   A-ip:A-app-port, B-cluster-ip,B-cluster-port   |     A-app-skops     |    <--- A-app active_estab CB
|------------------------------------------------------------------------|
|          A-ip:A-envoy-port, B-ip:B-port          |    A-envoy-skops    |    <--- A-envoy active_estab CB
|------------------------------------------------------------------------|
|       127.0.0.1:A-outbound, A-ip:A-app-port      |   A-outbound-skops  |    <--- A-outbound passive_estab CB
|------------------------------------------------------------------------|
|        B-ip:B-inbound, A-ip:A-envoy-port         |   B-inbound-skops   |    <--- B-inbound passive_estab CB
|------------------------------------------------------------------------|
*/
struct {
        __uint(type, BPF_MAP_TYPE_SOCKHASH);
        __uint(max_entries, SOCKOPS_MAP_SIZE);
        __uint(key_size, sizeof(struct socket_4_tuple));
        __uint(value_size, sizeof(uint32_t));
        __uint(pinning, LIBBPF_PIN_BY_NAME);
} map_redir SEC(".maps");

/* This a array map for debug configuration and record bypassed packet number
|-----------|------------------------------------|
|     0     |   0/1 (disable/enable debug info)  |
|-----------|------------------------------------|
|     1     |       bypassed packets number      |
|------------------------------------------------|
*/
struct {
        __uint(type, BPF_MAP_TYPE_ARRAY);
        __uint(max_entries, 2);
        __type(key, uint32_t);
        __type(value, uint32_t);
        __uint(pinning, LIBBPF_PIN_BY_NAME);
} debug_map SEC(".maps");

// 函数内联
// sk_ops_extract4_key 从 struct bpf_sock_ops *ops（socket metadata）中提取 key
static __inline__ void sk_ops_extract4_key(struct bpf_sock_ops *ops,
                struct socket_4_tuple *key)
{
    key->local.ip4 = ops->local_ip4;
    key->local.port = ops->local_port;
    key->remote.ip4 = ops->remote_ip4;
    key->remote.port = bpf_ntohl(ops->remote_port);
}

// 函数内联
// sk_msg_extract4_keys 从 struct sk_msg_md *ops（socket metadata）中提取 key
static __inline__ void sk_msg_extract4_keys(struct sk_msg_md *msg,
                struct socket_4_tuple *proxy_key, struct socket_4_tuple *key)
{
    proxy_key->local.ip4 = msg->local_ip4;
    proxy_key->local.port = msg->local_port;
    proxy_key->remote.ip4 = msg->remote_ip4;
    proxy_key->remote.port = bpf_ntohl(msg->remote_port);
    key->local.ip4 = msg->remote_ip4;
    key->local.port = bpf_ntohl(msg->remote_port);
    key->remote.ip4 = msg->local_ip4;
    key->remote.port = msg->local_port;
}
