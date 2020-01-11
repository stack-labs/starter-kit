# Helm

> 未测试，且`image`配置仅是样例

- charts
    - `gateway`适合`micro api`和`micro web`网关
    - `service`适合`api`、`web`聚合，以及`srv`

**yaml输出**
```bash
helm template micro ./starter-kit --namespace ns-micro > starter-kit.yaml
```
