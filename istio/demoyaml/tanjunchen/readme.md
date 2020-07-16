kubectl apply -f svc/
kubectl apply -f deploy/

kubectl apply -f  <(istioctl kube-inject -f deploy/*.yaml)

kubectl apply -f vs/

kubectl exec -it client-746f788c84-gx2mv -n tanjunchen -- sh

wget -q -O - http://web-svc:8080


wget -q -O - http://web-svc:8080 --header="end-user":"tanjunchen"

