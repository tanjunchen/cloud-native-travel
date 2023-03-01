# istio mcp-over-xds

mcp-over-xds demo

参考 https://github.com/antoineco/istio-mcp-sample

## Usage

To build and publish the container image and deploy the server to Kubernetes.

```
kubectl apply -f deploy.yaml
```

Then, configure the server as an `istiod` [configSource][istio-cfgsrc]:

```yaml
apiVersion: install.istio.io/v1alpha1
kind: IstioOperator
spec:
  profile: minimal
  meshConfig:
    configSources:
      - address: xds://mcp-sample.default.svc.cluster.local:15010
      - address: k8s://
# ...
```

see the configz from istiod, there are many service entry example.com

```
kubectl -n istio-system exec -it istiod-5dcbbcf9b4-m9j6n -- curl localhost:8080/debug/configz | grep example.com
```

## refer

[mcp-design]: https://docs.google.com/document/d/1lHjUzDY-4hxElWN7g6pz-_Ws7yIPt62tmX3iGs_uLyI/

[istio-cfgsrc]: https://istio.io/latest/docs/reference/config/istio.mesh.v1alpha1/#ConfigSource

[mcpoverxds]: https://github.com/zirain/mcpoverxds
