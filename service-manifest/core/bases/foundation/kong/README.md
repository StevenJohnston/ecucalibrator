```
kubectl create -f https://bit.ly/k4k8s
export PROXY_IP=$(minikube service -n kong kong-proxy --url | head -1)
```
