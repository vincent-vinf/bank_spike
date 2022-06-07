# golang秒杀系统后端

## 架构设计

### 微服务功能划分

![bank_spike.drawio](assets/bank_spike.drawio.png)

* 用户(user)
  * 登录
  * 注册
  * 刷新token
  
* 秒杀(spike)
  * 核心的秒杀逻辑
  * 将订单减扣的消息发送到消息队列
  
* 准入(access)
  * 对指定的秒杀，根据灵活的规则配置来判断用户是否有资格

* 管理(admin)
  * 管理秒杀活动(CRUD)

* 订单(order)
  * 查询订单
  * 付款
  * 接收消息队列的消息，生成订单插入数据库




## 构建&安装

#### CI

项目提供了 `Dockerfile`和`Makefile`

```bash
# 构建docker镜像
make all_image
# 推送镜像到远端仓库
make push
# 执行构建和推送
make build_push
# 打包helm chart 供上传到k8s master安装
make tar_chart
```

#### 安装

> 前提：
>
> * k8s集群
> * 支持ingress
> * helm v3

```bash
helm install spike-backend spike-chart-latest.tar.gz
```

运行`kubectl get all`，结果如下

```
NAME                                        READY   STATUS    RESTARTS   AGE
pod/spike-spike-service-66bf5fc76f-c9k4z    1/1     Running   0          15m
pod/spike-user-service-85c94ff74f-jgdzx     1/1     Running   0          15m
pod/spike-access-service-774c7bcb9f-g2mg4   1/1     Running   0          15m
pod/spike-order-service-7cc84c5bf-k6c9r     1/1     Running   0          15m
pod/spike-admin-service-5db98f8b87-pht2x    1/1     Running   0          15m

NAME                           TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)          AGE
service/kubernetes             ClusterIP   10.43.0.1      <none>        443/TCP          156m
service/spike-user-service     NodePort    10.43.81.176   <none>        8080:30085/TCP   15m
service/spike-admin-service    NodePort    10.43.192.49   <none>        8080:30066/TCP   15m
service/spike-order-service    NodePort    10.43.30.117   <none>        8080:30621/TCP   15m
service/spike-access-service   NodePort    10.43.226.2    <none>        8081:31816/TCP   15m
service/spike-spike-service    NodePort    10.43.48.208   <none>        8080:31830/TCP   15m

NAME                                   READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/spike-spike-service    1/1     1            1           15m
deployment.apps/spike-user-service     1/1     1            1           15m
deployment.apps/spike-access-service   1/1     1            1           15m
deployment.apps/spike-order-service    1/1     1            1           15m
deployment.apps/spike-admin-service    1/1     1            1           15m

NAME                                              DESIRED   CURRENT   READY   AGE
replicaset.apps/spike-spike-service-66bf5fc76f    1         1         1       15m
replicaset.apps/spike-user-service-85c94ff74f     1         1         1       15m
replicaset.apps/spike-access-service-774c7bcb9f   1         1         1       15m
replicaset.apps/spike-order-service-7cc84c5bf     1         1         1       15m
replicaset.apps/spike-admin-service-5db98f8b87    1         1         1       15m
```

#### 中间件部署

##### mysql

```bash
docker run -d --restart=always --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=<passwd> mysql
```

##### redis

写入/root/redis.conf

```
bind 0.0.0.0
daemonize NO
protected-mode no
requirepass <passwd>
```

```bash
docker run -d --restart=always --name redis -p 6379:6379 -v /root/redis.conf:/etc/redis/redis.conf -d redis /etc/redis/redis.conf
```

##### rabbitmq

```bash
docker run -d --restart=always --name rabbitmq -p 15672:15672 -p 5672:5672 -e RABBITMQ_DEFAULT_USER=<username> -e RABBITMQ_DEFAULT_PASS=<passwd> rabbitmq:management
```

  

  