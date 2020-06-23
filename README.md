# betPawa

## Setup

Get

- k3d [v3.0.0-rc.6](https://github.com/rancher/k3d/releases/tag/v3.0.0-rc.6)

- kubectl

- skaffold

## Run

Run local registry

```
docker run --rm -p 5000:5000 --restart=always --name registry registry:2
echo '127.0.0.1 registry.local' | sudo tee -a /etc/hosts
```

then

```bash
k3d create cluster --api-port 6550 -p 8081:80@loadbalancer --workers 2 --k3s-server-arg '--no-deploy=traefik'
export KUBECONFIG=$(k3d get kubeconfig k3s-default)
skaffold dev --default-repo registry.local:5000
```

and test

```bash
curl localhost:8081/ping
OK
```

## References

- 1 https://github.com/kubernetes/ingress-nginx/issues/5401#issuecomment-617006161

- 2 https://rancher.com/docs/k3s/latest/en/networking/

- 3 https://k3d.io/usage/guides/exposing_services/
