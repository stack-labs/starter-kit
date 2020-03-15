# Kubernetes

**目录**

- [环境](#环境)
- [网关部署](#网关部署)

## 环境

- Kubernetes
- Options
	- Ingress
		- [nginx ingress](https://github.com/nginxinc/kubernetes-ingress)

## 网关部署

> deprecated，使用[helm](/deploy/k8s/helm)

<details>
  <summary> 网关部署 </summary>

- [网关Docker镜像](/gateway#Docker)
- [部署yaml](/deploy/k8s/gateway)

```bash
$ cd deploy/k8s

# 手动部署
# 略

# 脚本部署
./run.sh start micro
```

### 部署状态
```bash
$ kubectl config set-context --current --namespace=ns-micro
$ kubectl get pod
  NAME                         READY   STATUS    RESTARTS   AGE
  micro-api-7ff59495bb-xznt6   1/1     Running   0          21m
  micro-web-65d4f-tcxd5        1/1     Running   0          21m
  
$ kubectl get service
  NAME        TYPE           CLUSTER-IP      EXTERNAL-IP   PORT(S)        AGE
  micro-api   LoadBalancer   10.107.85.217   <pending>     80:30000/TCP   21m
  micro-web   LoadBalancer   10.111.68.118   <pending>     80:30002/TCP   21m
  
$ kubectl get ingress
  NAME                HOSTS                 ADDRESS   PORTS   AGE
  ingress-micro-api   api.starter-kit.com             80      20m
  ingress-micro-web   www.starter-kit.com             80      20m
```

### 测试API及Web网关
```bash
$ minikube ip
192.168.39.147
```

#### Ingress模式
```bash
# nginx-ingress port 80:32134/TCP,443:31089/TCP
$ curl -HHost:api.starter-kit.com 'http://192.168.39.147:32134'
  {"version": "1.14.0"}
$ curl -HHost:www.starter-kit.com 'http://192.168.39.147:32134'
  <html>
  	……
  </html>
```

#### NodePort模式
```bash
$ curl 'http://192.168.39.147:30000'
  {"version": "1.14.0"}
$ curl 'http://192.168.39.147:30002'
  <html>
  	……
  </html>
```

</details>


