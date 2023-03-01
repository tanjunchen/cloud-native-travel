# 选择性服务发现

## 安装

istioctl install --set profile=default --set spec.components.pilot.k8s.env[0].name="ENABLE_ENHANCED_RESOURCE_SCOPING" --set spec.components.pilot.k8s.env[0].value='"true"' --set hub=gcr.io/istio-testing --set tag=1.16-alpha.558e88a06e4434ed8075798160a58f3d5b7630d2 


## 下载最新的镜像

参考: https://github.com/istio/istio/wiki/Dev-Builds

## 参考

https://github.com/istio/istio/issues/40316

https://github.com/istio/istio/pull/36639

https://docs.google.com/document/d/1y4liRJbQW0NCMeQtqMma46flVqs-izV1/edit

