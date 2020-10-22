# TCP Echo Service

## Usage

To run the TCP Echo Service sample:

1. Install Istio by following the [istio install instructions](https://istio.io/docs/setup/kubernetes/quick-start.html).

1. Start the `tcp-echo-server` service inside the Istio service mesh:

    ```console
    $ kubectl apply -f <(istioctl kube-inject -f tcp-echo.yaml)
    service/tcp-echo created
    deployment.apps/tcp-echo-v1 created
    deployment.apps/tcp-echo-v2 created
    ```

1. Test by running the `nc` command from a `busybox` container from within the cluster.

    ```console
    $ kubectl run -i --rm --restart=Never dummy --image=busybox -- sh -c "echo world | nc tcp-echo 9000"
    istio-version-v1 world
    pod "dummy" deleted
    ```

    As you observe, sending _world_ on a TCP connection to the server results in
    the server prepending _istio-version-v1_ and echoing back with _istio-version-v1 world_.

1. To clean up, execute the following command:

    ```console
    $ kubectl delete -f tcp-echo.yaml
    service "tcp-echo" deleted
    deployment.apps "tcp-echo-v1" deleted
    deployment.apps "tcp-echo-v2" deleted
    ```

## testing TCP traffic shifting

1. kubectl create namespace istio-io-tcp-traffic-shifting

1. kubectl label namespace istio-io-tcp-traffic-shifting istio-injection=enabled

1. kubectl apply -f istio/example/sleep/sleep.yaml -n istio-io-tcp-traffic-shifting

1. kubectl apply -f istio/example/tcp-demo/src/tcp/tcp-echo.yaml -n istio-io-tcp-traffic-shifting

1. kubectl apply -f istio/example/tcp-demo/src/tcp/tcp-echo-all-v1.yaml -n istio-io-tcp-traffic-shifting

1. 流量全部打到 v1 版本
    
   ```
   export TCP_INGRESS_PORT = $(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="tcp")].nodePort}')
   export INGRESS_HOST = $(kubectl get po -l istio=ingressgateway -n istio-system -o jsonpath='{.items[0].status.hostIP}')
   
   for i in {1..30}; do \
   kubectl exec "$(kubectl get pod -l app=sleep -n istio-io-tcp-traffic-shifting -o jsonpath={.items..metadata.name})" \
   -c sleep -n istio-io-tcp-traffic-shifting -- sh -c "(date; sleep 1) | nc $INGRESS_HOST $TCP_INGRESS_PORT"; \
   done
   ```

1. 切换 30% 流量到 v2 版本 

    kubectl apply -f istio/example/tcp-demo/src/tcp/tcp-echo-30-v2.yaml -n istio-io-tcp-traffic-shifting

1. 查看流量效果

   ```
   export TCP_INGRESS_PORT = $(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="tcp")].nodePort}')
   export INGRESS_HOST = $(kubectl get po -l istio=ingressgateway -n istio-system -o jsonpath='{.items[0].status.hostIP}')
   
   for i in {1..30}; do \
   kubectl exec "$(kubectl get pod -l app=sleep -n istio-io-tcp-traffic-shifting -o jsonpath={.items..metadata.name})" \
   -c sleep -n istio-io-tcp-traffic-shifting -- sh -c "(date; sleep 1) | nc $INGRESS_HOST $TCP_INGRESS_PORT"; \
   done
   ```

1. 删除测试案例

   kubectl delete -f istio/example/tcp-demo/src/tcp/tcp-echo.yaml -n istio-io-tcp-traffic-shifting
   kubectl delete -f istio/example/tcp-demo/src/tcp/tcp-echo-30-v2.yaml -n istio-io-tcp-traffic-shifting
   kubectl delete -f istio/example/tcp-demo/src/tcp/tcp-echo-all-v1.yaml -n istio-io-tcp-traffic-shifting
   kubectl delete namespace istio-io-tcp-traffic-shifting
