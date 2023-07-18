# Simple sleep service

## Note(注意)

Istio 使用的是 governmentpaas/curl-ssl 镜像，该镜像 Dockerfile 是：

```
FROM alpine:3.12

ENV PACKAGES "jq gettext curl openssl ca-certificates"5

RUN apk add --no-cache $PACKAGES
```

本地重新打包镜像并且上传到阿里云，如下所示：

registry.cn-hangzhou.aliyuncs.com/tanjunchen/sleep:v1.1

To use it:

1. Install Istio by following the [istio install instructions](https://istio.io/docs/setup/).

1. Start the sleep service:

    If you have [automatic sidecar injection](https://istio.io/docs/setup/additional-setup/sidecar-injection/#automatic-sidecar-injection) enabled:

    ```bash
    kubectl apply -f sleep.yaml
    ```

    Otherwise manually inject the sidecars before applying:

    ```bash
    kubectl apply -f <(istioctl kube-inject -f sleep.yaml)
    ```

1. Start some other services, for example, the [Bookinfo sample](https://istio.io/docs/examples/bookinfo/).

    Now you can `kubectl exec` into the sleep service to experiment with Istio networking.
    For example, the following commands can be used to call the Bookinfo `ratings` service:

    ```bash
    export SLEEP_POD=$(kubectl get pod -l app=sleep -o jsonpath={.items..metadata.name})
    kubectl exec -it $SLEEP_POD -c sleep curl www.baidu.com}
    ```