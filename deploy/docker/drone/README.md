# Drone

**使用参考**

- [使用Drone进行持续构建与发布](https://jimmysong.io/kubernetes-handbook/practice/drone-ci-cd.html)
- [Docker — 从入门到实践-Drone](https://www.cntofu.com/book/139/cases/ci/drone.md)

**Compose启动**
```bash
docker-compose up -d
```

**常见问题**

- 创建`admin`权限用户，否则在做缓存需要`volume cache`时出错
    - `DRONE_USER_CREATE=username:{GITHUB_USER_NAME},machine:false,admin:true,token:xxx` 


**Pipeline示例**
`.drone.yml`以`gateway`模块为例，完成Go编译、Docker打包发布、K8S部署过程，不是一个API网关

```
curl -v -HHost:api.starter-kit.com -X GET 'http://{host}:{port}/'
```

## Drone Secret
- [Drone configure secrets](https://docs.drone.io/configure/secrets/)

**依赖的Secret**

> 使用drone的cli工具，在web的user settings可以查看使用方式

```shell script
drone secret add --repository=micro-in-cn/starter-kit --name=docker_username --data=
drone secret add --repository=micro-in-cn/starter-kit --name=docker_password --data=
drone secret add --repository=micro-in-cn/starter-kit --name=k8s_server --data=
drone secret add --repository=micro-in-cn/starter-kit --name=k8s_ca --data=
drone secret add --repository=micro-in-cn/starter-kit --name=k8s_token --data=
```

**Docker Hub秘钥**

- `docker_username`
- `docker_password`

**K8S config**

1.取一个Pod的Secret或者单独创建一个account
```
kubectl describe po {pod-name} | grep SecretName
```

2.取Secret的ca和token
```
kubectl get secret {secret-name} -o yaml | egrep 'ca.crt:|token:'
```

3.drone secret配置
- `k8s_server`
    - `cat ~/.kube/config`
- `k8s_ca`
    - ca直接使用
- `k8s_token`
    - token做base64解码后使用

**`kubectl`部署需要对应的account有授权**

> 测试可以将clusterrole=cluster-admin授权给serviceaccount
    
```
# Error from server (Forbidden): services is forbidden: User "system:serviceaccount:default:default" cannot list resource "services" in API group "" in the namespace "default"
kubectl create clusterrolebinding cluster-admin-role-bound --clusterrole=cluster-admin --serviceaccount={namespace}:{service-account}
```

## Ref
- [Drone configure secrets](https://docs.drone.io/configure/secrets/)
- [Drone docker plugin](http://plugins.drone.io/drone-plugins/drone-docker/)
- [使用Drone进行持续构建与发布](https://jimmysong.io/kubernetes-handbook/practice/drone-ci-cd.html)
- [Docker — 从入门到实践-Drone](https://www.cntofu.com/book/139/cases/ci/drone.md)
- [基于drone的CI/CD，对接kubernetes实践教程](https://www.kubernetes.org.cn/4687.html)
