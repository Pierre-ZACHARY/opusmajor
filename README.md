# Requirements

## Minikube Installation
https://kubernetes.io/fr/docs/tasks/tools/install-minikube/

## ArgoCD Installation 
https://argo-cd.readthedocs.io/en/stable/getting_started/

```
kubectl create namespace argocd
kubectl apply -n argocd --server-side --force-conflicts -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```

## Dev
- Golang
https://go.dev/learn/
- Delve ( go debugger )

# Installation of the application with argocd

The docker image of the app is published on GHCR
```
ghcr.io/pierre-zachary/opusmajor:main
```

## How to render the kustomization 

```
kubectl kustomize k8s/player-data
```