# Docker Compose

## Network

```bash
docker network create micro-net
```

## 基础服务

- Etcd
- Redis
- Postgres

```bash
docker-compose -f compose-basic.yml -p starter-kit up 
```

## Gateway

- micro api
- micro web

```bash
docker-compose -f compose-gateway.yml -p starter-kit up
```
