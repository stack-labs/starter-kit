# Helm

- charts
    - `gateway`适合`micro api`和`micro web`网关
    - `service`适合`api`、`web`聚合，以及`srv`
    - `gateway`和`service`的chart单独使用参考[.drone.yml](/.drone.yml)的`pipeline`

**yaml输出**
```bash
helm template micro ./starter-kit --namespace starter-kit > starter-kit.yaml
```

**部署**
```bash
helm template micro ./starter-kit --namespace starter-kit | kubectl apply -f -
```
