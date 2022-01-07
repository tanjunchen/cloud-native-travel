install:
	kubectl create ns thrift-demo
	kubectl -n thrift-demo apply -f ns-deploy-svc.yaml


apply-envoyfilter:
	kubectl -n thrift-demo apply -f thrift-envoyfilter.yaml

delete-envoyfilter:
	kubectl -n thrift-demo delete -f thrift-envoyfilter.yaml

clean:
	kubectl -n thrift-demo delete -f thrift-envoyfilter.yaml
	kubectl -n thrift-demo delete -f ns-deploy-svc.yaml
	kubectl ns thrift-demo
