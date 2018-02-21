# kube-monkey Demo

Requirements:

- [Kubectl - kubernetes cli](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
- [Minikube - to run kubernetes locally](https://github.com/kubernetes/minikube#installation) (which requires [VirtualBox or similar](https://github.com/kubernetes/minikube#requirements))
- [Docker](https://docs.docker.com/engine/installation/)
- [Go](https://golang.org/doc/install)
- [kube-monkey](https://github.com/asobti/kube-monkey)

```sh
# Start a kubernetes cluster running on a local VM
minikube start --kubernetes-version v1.6.4

# Setup docker client to point to the docker instance running inside VM
eval $(minikube docker-env)

# Build the docker image
docker build -t kube-monkey-demo-img:v1 .

# Create the kubernetes resources
kubectl apply -f kubernetes.yaml

# Inspect things
kubectl get deployments
kubectl get pods
kubectl get services

# Inspect the exposed service url
minikube service kube-monkey-demo --url

# Hit the service
curl -v $(minikube service kube-monkey-demo --url)

# Building kube-monkey
go get github.com/asobti/kube-monkey
cd $GOPATH/src/github.com/asobti/kube-monkey
make

# Configuring
// TODO

# Deploying monkeys
// TODO
```
