#include <linux/bpf.h>

#include "bpf_sockops.h"


/*
 * extract the key that identifies the destination socket in the sock_ops_map
 */
static inline
void extract_key4_from_msg(struct sk_msg_md *msg, struct sock_key *key)
{
    key->sip4 = msg->remote_ip4;
    key->dip4 = msg->local_ip4;
    key->family = 1;

    key->dport = (bpf_htonl(msg->local_port) >> 16);
    key->sport = FORCE_READ(msg->remote_port) >> 16;
}

__section("sk_msg")
int bpf_redir(struct sk_msg_md *msg)
{
    struct sock_key key = {};
    extract_key4_from_msg(msg, &key);
    msg_redirect_hash(msg, &sock_ops_map, &key, BPF_F_INGRESS);
    return SK_PASS;
}

char ____license[] __section("license") = "GPL";
