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
- Vscode launch config are included : 
  - Run the golang app natively with go delve
  - Run the app within docker
  - Send a curl request to retrieve the player-data

# Installation of the application with argocd

The docker image of the app is published on GHCR
```
ghcr.io/pierre-zachary/opusmajor:main
```

- How to render the kustomization 
```
kubectl kustomize k8s/player-data
```
- Install the app in your cluster: 
```
kubectl apply -f https://raw.githubusercontent.com/Pierre-ZACHARY/opusmajor/refs/heads/main/argocd/player-data-application.yaml
```

- Check sync/health:
```
kubectl -n argocd get application player-data
kubectl -n player-data get deploy,pods,svc,hpa
```

## Ingress access with playerdata.minikube

Enable the default Minikube NGINX ingress addon, then Kubernetes will route requests for `playerdata.minikube` to this app through the Ingress resource in this repo. On macOS with Docker driver, keep `minikube tunnel` running in a separate terminal so the ingress address is reachable.

Reference doc:
https://v1-33.docs.kubernetes.io/docs/tasks/access-application-cluster/ingress-minikube/

Enable ingress:
```
minikube addons enable ingress
# Verify 
kubectl get pods -n ingress-nginx
```

/!\ Important : keep minikube tunnel open to expose the ingress from minikube on your localhost
```
minikube tunnel
```

Add this line to your dns hostnames : 
```
127.0.0.1 playerdata.minikube
```

Test ingress routing:
```
curl http://playerdata.minikube/player-data
```

### Change the root URL suffix

The Ingress host is built as `playerdata.<ROOT_URL>`.

1. Edit `ROOT_URL` in `k8s/player-data/kustomization.yaml`:
```
configMapGenerator:
  - name: player-data-config
    literals:
      - ROOT_URL=minikube
```

2. Replace `minikube` with your own suffix, for example `local`.

3. Re-apply manifests:
```
kubectl apply -k k8s/player-data
```

4. Add/update your local hosts entry to match the new hostname, for example:
```
127.0.0.1 playerdata.local
```

5. Test with the new URL:
```
curl http://playerdata.local/player-data
```


# API Documentation (Swagger / OpenAPI)

OpenAPI spec file:
```
docs/swagger.json
```

## Render Option 1: Swagger Editor (quickest)

1. Open:
```
https://editor.swagger.io/
```
2. Click File -> Import File and select:
```
docs/swagger.json
```

## Render Option 2: Swagger UI with Docker (local)

Run:
```
docker run --rm -p 8081:8080 -e SWAGGER_JSON=/app/swagger.json -v "$(pwd)/docs/swagger.json:/app/swagger.json" swaggerapi/swagger-ui
```

Then open:
```
http://localhost:8081
```