# K8s Demon 项目

## 1、项目前置

### 1.1、项目打包

#### 1.1.1、前端

1、前端项目打包：

```shell
# 打开终端，进入你的 vue 项目根目录，执行以下命令：（生产构建，推荐，包含类型检查）
# 构建完成的目标文件夹名: dist
npm run build:pro
```

> ```json
>   "scripts": {
>     "dev": "vite",
>     "build": "run-p type-check \"build-only {@}\" --",
>     "preview": "vite preview",
>     "build-only": "vite build",
>     "type-check": "vue-tsc --build",
>     "build:pro": "vue-tsc && vite build --mode production"
>   },
> ```



#### 1.1.2、后端

2、后端项目打包（交叉编译，在 Windows 上编译，Linux 上运行）：

```shell
# 打开终端，进入你的 Go 项目根目录，执行以下命令：
# 下面为 powershell 命令
$env:CGO_ENABLED=0; $env:GOOS="linux"; $env:GOARCH="amd64"; go build -o <你的项目名> main.go
```

> 构建参数详解：
>
> - `CGO_ENABLED=0`：这是关键一步！它表示禁用 CGO（C语言调用）。设置为 `0` 可以进行**静态编译**，将程序依赖的所有库都打包进二进制文件里，生成的文件不依赖外部的 C 库，可以直接在“干净”的 Linux 系统上运行，避免了很多依赖问题。
> - `GOOS=linux`：指定目标操作系统为 Linux。
> - `GOARCH=amd64`：指定目标系统架构。这是目前最常见的 64 位 x86 架构服务器。如果你的目标机器是 ARM 架构（如树莓派），则需设置为 `arm64`。
> - `-o your-app-name`：指定编译后输出的文件名，方便识别。如果你不指定 `-o`，生成的文件名默认就是你的项目名或 `main`。



### 1.2、项目测试

#### 1.2.1、安装 Nginx

1、在 CentOS 7 上安装 Nginx ：

> - **准备安装环境**
>
>   ```shell
>   # 避免 Nginx 与 Httpd 发生端口冲突
>   (rpm -qa | grep -P "^httpd-([0-9].)+") && rpm -e --nodeps httpd || echo "httpd未安装"
>
>   # 下载相关依赖
>   yum install -y gcc pcre-devel zlib-devel openssl-devel gd gd-devel
>
>   # 创建用户和组
>   groupadd nginx
>   useradd nginx -g nginx -M -s /sbin/nologin
>
>   # 下载 Nginx 源码包
>   which wget || yum -y install wget
>   wget https://nginx.org/download/nginx-1.24.0.tar.gz
>
>   # 解压
>   tar -axf nginx-1.24.0.tar.gz
>   cd nginx-1.24.0
>
>   # 预编译（基本）
>   ./configure --prefix=/usr/local/nginx --user=nginx --group=nginx
>   # 或这个预编译 （进阶：安装多个功能模块）
>   ./configure --prefix=/usr/local/nginx --user=nginx  --group=nginx --sbin-path=/usr/local/nginx/nginx --conf-path=/usr/local/nginx/conf/nginx.conf --error-log-path=/var/log/nginx/nginx.log --http-log-path=/var/log/nginx/access.log --modules-path=/usr/local/nginx/modules --with-select_module --with-poll_module --with-threads  --with-http_ssl_module --with-http_v2_module  --with-http_realip_module --with-http_image_filter_module --with-http_sub_module --with-http_flv_module --with-http_gunzip_module --with-http_gzip_static_module  --with-http_stub_status_module --with-stream
>   # 检查 （0 : 代表成功）
>   echo $?
>
>   # 编译 （0 : 代表成功）
>   make
>   [root@midden nginx-1.24.0]# echo $?
>   0
>
>   # 安装 （0 : 代表成功）
>   make install
>   [root@midden nginx-1.24.0]# echo $?
>   0
>
>   # 检查
>   nginx -V
>   ```
>
> - **配置环境变量**
>
>   - **方法一**
>
>   ```shell
>   # 配置环境变量方法一
>   cat > /etc/profile.d/nginx.sh<<EOF
>   export PATH="/usr/local/nginx:\$PATH"
>   EOF
>   ```
>
>   ```shell
>   # 使配置的环境变量生效
>   source /etc/profile
>   ```
>
>   ```shell
>   # 查看帮助
>   [root@midden nginx]# nginx -h
>   ```
>
> - **启动 Nginx**
>
>   ```shell
>   # 设置开机自启动
>   echo "/usr/local/nginx/nginx" >> /etc/rc.d/rc.local
>                             
>   # 赋予可执行权限（不然开机自启不生效）
>   chmod +x /etc/rc.d/rc.local
>                             
>   # 启动 Nginx 服务
>   nginx  或  nginx -c /usr/local/nginx/conf/nginx.conf
>   ```
>
>   ```shell
>   # 停止 Nginx 服务
>   nginx -s stop
>                             
>   ```
>   
>   ```shell
>   # 在浏览器输入虚拟机的IP（注意是http，走80端口,不是https,走443端口）
>   http://<ip>/
>   ```



#### 1.2.2、配置 Nginx

1、将前端打包好的文件 `dist` 文件夹的内容放到下面的目录下 ：

```shell
[root@test01 html]# tree /usr/local/nginx/html
/usr/local/nginx/html
├── 50x.html
├── assets
├── favicon.ico
└── index.html
```

2、配置 Nginx 的配置文件

```shell
cp /usr/local/nginx/conf/nginx.conf /usr/local/nginx/conf/nginx.conf.bak
vim /usr/local/nginx/conf/nginx.conf
```

```shell
worker_processes  1;
events {
    worker_connections  1024;
}
http {
    include       mime.types;
    default_type  application/octet-stream;
    sendfile        on;
    keepalive_timeout  65;
    server {
        listen       80;
        server_name  localhost;
        location / {
            root   html;
            index  index.html index.htm;
            # 添加这行，对 SPA 应用很重要
            try_files $uri $uri/ /index.html;
        }
        # 后端接口代理
        location /api/ {
            proxy_pass http://192.168.15.30:8080/;
        }
        
        error_page   500 502 503 504  /50x.html;
        location = /50x.html {
            root   html;
        }
    }
}
```

> 1、SPA 的路由问题（**SPA** = Single Page Application（单页应用）），单页应用（如 Vue/React/Angular）使用**前端路由**：
>
> - 用户访问 `http://example.com/about`
> - 但实际上服务器上没有 `/about` 这个文件
> - 所有路由都应该由前端的 JavaScript 处理
>
> 2、没用 `try_files` 的情况：
>
> - 访问 `/` → ✅ 正常显示 index.html
> - 访问 `/about` → ❌ 404 Not Found（服务器找不到 about 文件
>
> 简单来说，`try_files $uri $uri/ /index.html;` 就是告诉 Nginx："**先找真实文件，找不到就交给前端路由处理**"，这是 SPA 应用能正常刷新的关键！

3、重新加载 Nginx 的配置，然后就可以访问了：

```shell
nginx -s reload
```

---

<img src="./imgs/web.png" style="zoom: 33%;" />

---



#### 1.2.3、配置后端项目

1、将后端打包的项目放在以下目录：

```shell
cd /opt
#[root@test01 opt]# tree /opt
#/opt
#└── k8sdemon
```

```shell
# 添加可执行权限
chmod +x /opt/k8sdemon
#在 Go 项目中，配置文件通常不会被编译进二进制文件
vim config.yaml
```

```shell
# 配置文件
# 程序配置信息
system:
  port: 8080
  host: "0.0.0.0"
# 数据库连接信息
mysql:
  port: 3306
  host: "192.168.15.36"
  username: "root"
  password: "QQlpx1314."
  dbname: "k8sdemon"
```



2、配置使用 `systemctl` 接管后端项目：

```shell
vim /etc/systemd/system/k8sdemon.service
```

```ini
[Unit]
Description=My Go Application Service
After=network.target
Wants=network.target

[Service]
Type=simple
# 如果你的应用不需要 root 权限，建议使用普通用户运行
User=root
Group=root

# 工作目录 - 如果你的应用有读取本地文件（如配置文件、静态资源），这个很重要
WorkingDirectory=/opt

# 启动命令 - 写绝对路径最保险
ExecStart=/opt/k8sdemon

# 重启策略
Restart=always
RestartSec=5

# 日志输出
StandardOutput=journal
StandardError=journal

[Install]
WantedBy=multi-user.target
```



3、重新加载  `systemctl` 服务的配置：

```shell
# 1. 重载配置文件，让 systemd 识别新的服务
sudo systemctl daemon-reload

# 2. 启动服务
sudo systemctl start k8sdemon

# 3. 设置开机自启
sudo systemctl enable k8sdemon

# 4. 查看运行状态
sudo systemctl status k8sdemon
```

---

![](./imgs/run.png)

---



## 2、制作镜像

### 2.1、前端镜像

1、必要文件准备：

```shell
# 准备一个目录
# 将前端打包的文件也放在该目录下
mkdir frontend_docker
cd frontend_docker
# Nginx 部分配置
vim server.conf
```

```shell
server {
    listen       80;
    listen  [::]:80;
    server_name  localhost;
    location / {
        # 写绝对路径
        root   /usr/share/nginx/html;
        index  index.html index.htm;
        # 添加这行，对 SPA 应用很重要
        try_files $uri $uri/ /index.html;
    }
    # 后端接口代理
    location /api/ {
        proxy_pass http://192.168.15.30:8080/;
    }
        
    error_page   500 502 503 504  /50x.html;
    location = /50x.html {
        root   html;
    }
}
```

2、编写 Dockerfile 文件：

```shell
vim Dockerfile
```

```dockerfile
FROM nginx:1.28
# 复制 Nginx 配置文件
COPY server.conf /etc/nginx/conf.d/
# 复制前端文件
COPY dist/ /usr/share/nginx/html/
# 移除默认配置
RUN rm /etc/nginx/conf.d/default.conf
# 暴露 80 端口
EXPOSE 80
# 健康检查
HEALTHCHECK --interval=30s --timeout=3s --retries=3 --start-period=0s \
    CMD curl -f http://localhost/ || exit 1
# 运行命令
CMD ["nginx", "-g", "daemon off;"]
```

3、使用 Dockerfile 文件构建镜像：

```shell
# 目录层级
frontend_docker/
├── dist
│   ├── assets
│   ├── favicon.ico
│   └── index.html
├── Dockerfile
└── server.conf
```

```shell
# 开始构建
docker build -t frontend:v1.0 .
```

```shell
# 测试并访问
docker run -it -d -p 9000:80 --name myfront frontend:v1.0
```



### 2.2、后端镜像

1、必要文件准备：

```shell
# 创建目录并把后端打包的文件放在该目录下
mkdir backend_docker
# 后端的配置文件
vim config.yaml
```

```yaml
# 配置文件
# 程序配置信息
system:
  port: 8080
  host: "0.0.0.0"
# 数据库连接信息
mysql:
  port: 3306
  host: "192.168.15.36"
  username: "root"
  password: "QQlpx1314."
  dbname: "k8sdemon"
```

2、编写 Dockerflie 文件：

```shell
vim Dockerfile
```

```dockerfile
FROM alpine:3.22
# 设置工作目录
WORKDIR /opt/k8sdemon
# 复制必要的文件
COPY k8sdemon .
COPY config.yaml .
# 赋予可执行权限
RUN chmod +x k8sdemon
# 暴露端口
EXPOSE 8080
# 运行命令
CMD ["./k8sdemon"]
```

3、使用 Dockerfile 文件构建镜像：

```shell
# 目录结构
/backend_docker/
├── config.yaml
├── Dockerfile
└── k8sdemon
```

```shell
# 开始构建
docker build -t backend:v1.0 .
```

```shell
# 测试
docker run -it -d -p 8080:8080 --name mybackend backend:v1.0
```



### 2.3、数据库服务

1、安装 MySQL 5.7 数据库，登录后执行以下命令：

```mysql
# 允许 root 用户远程登录
GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' IDENTIFIED BY '你的root密码' WITH GRANT OPTION;
# 创建 k8sdemon 数据库
CREATE DATABASE IF NOT EXISTS k8sdemon
CHARACTER SET utf8mb4
COLLATE utf8mb4_general_ci;
# 使用数据库
USE k8sdemon;
# 导入表结果和测试数据
source /root/backup.sql;
```

- backup.sql

```sql
-- MySQL dump 10.13  Distrib 5.7.44, for Linux (x86_64)
--
-- Host: localhost    Database: k8sdemon
-- ------------------------------------------------------
-- Server version	5.7.44

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `users` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'张三','123456'),(2,'李四','abcdef'),(3,'王五','888888'),(18,'小廉','123');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping events for database 'k8sdemon'
--

--
-- Dumping routines for database 'k8sdemon'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2026-03-28 23:38:13
```





## 3、部署在 K8s 上

在 Kubernetes 中，前端 Pod 调用后端接口时，地址应该写成后端 Service 的域名。

这是 Kubernetes 服务发现的标准做法，也是官方推荐的最佳实践 。千万不要写具体 Pod 的 IP 地址，因为 Pod 重建后 IP 会变，写了就“写死”了。地址具体怎么写，取决于你的前端 Pod 和后端 Service 是否在同一个命名空间（Namespace）。

无论是否在同一命名空间，下面这个格式都通用：

- **地址格式**：`http://<后端Service名>.<后端所在命名空间>.svc.cluster.local:<端口号>`
- **例子**：
  - 假设后端 Service 叫 `backend-service`，它在 `production` 命名空间，端口是 `8080`。
  - 那么，在任何地方（任何命名空间），前端都应该调用：

```shell
http://backend-service.production.svc.cluster.local:8080
```



### 3.1、编写 Deployment

#### 3.1.1、后端 Deployment

```shell
vim backend.yaml
```

```yaml
# 定义项目的命名空间
apiVersion: v1
kind: Namespace
metadata:
  name: k8sdemon
  labels:
    environment: production
    team: frontend
  annotations:
    description: "k8sdemon服务空间"
---
# 定义后端 Service
apiVersion: v1
kind: Service
metadata:
  name: backend-service
  namespace: k8sdemon
  labels:
    app: backend
  annotations:
    description: "后端的 Service"
spec:
  type: ClusterIP
  selector:
    app: backend
  ports:
  - name: http
    port: 8080
    targetPort: 8080
    protocol: TCP
---
# 定义后端
apiVersion: apps/v1
kind: Deployment
metadata:
  name: backend-deployment
  namespace: k8sdemon
spec:
  selector:
    matchLabels:
      app: backend
  replicas: 2
  template:
    metadata:
      labels:
        app: backend
    spec:
      containers:
      - name: backend
        image: registry.cn-hangzhou.aliyuncs.com/pork-registry/backend:v1.0
        ports:
        - containerPort: 8080
```



#### 3.1.2、重新编译前端镜像

1、需要将前端文件里，请求后端的接口地址改成后端 Service 的名字，还要修改镜像里 Nginx 的代理后端地址，如下：

```shell
http://backend-service.k8sdemon.svc.cluster.local:8080
```

```shell
# Nginx 部分配置
vim server.conf
```

```ini
server {
    listen       80;
    listen  [::]:80;
    server_name  localhost;
    location / {
        # 写绝对路径
        root   /usr/share/nginx/html;
        index  index.html index.htm;
        # 添加这行，对 SPA 应用很重要
        try_files $uri $uri/ /index.html;
    }
    # 后端接口代理改成 backend-svice
    location /api/ {
        proxy_pass http://backend-service.k8sdemon.svc.cluster.local:8080/;
    }
        
    error_page   500 502 503 504  /50x.html;
    location = /50x.html {
        root   html;
    }
}
```

```shell
# 编写 Dockerfile
vim Dockerfile
```

```dockerfile
FROM nginx:1.28
# 复制 Nginx 配置文件
COPY server.conf /etc/nginx/conf.d/
# 复制前端文件
COPY dist/ /usr/share/nginx/html/
# 移除默认配置
RUN rm /etc/nginx/conf.d/default.conf
# 暴露 80 端口
EXPOSE 80
# 健康检查
HEALTHCHECK --interval=30s --timeout=3s --retries=3 --start-period=0s \
    CMD curl -f http://localhost/ || exit 1
# 运行命令
CMD ["nginx", "-g", "daemon off;"]
```

2、使用 Dockerfile 文件构建镜像：

```shell
# 目录层级
frontend_docker_k8s/
├── dist
│   ├── assets
│   ├── favicon.ico
│   └── index.html
├── Dockerfile
└── server.conf
```

```shell
# 开始构建
docker build -t frontend:v2.0 .
```



#### 3.1.3、前端 Deployment

```shell
vim frontend.yaml
```

```yaml
# 定义前端的 Service
apiVersion: v1
kind: Service
metadata:
  name: frontend-service
  namespace: k8sdemon
  labels:
    app: frontend
  annotations:
    description: "Nginx Web Service"
spec:
  type: NodePort
  selector:
    app: frontend
  ports:
  - name: http
    port: 80
    targetPort: 80
    protocol: TCP
    nodePort: 30001
---
# 定义前端
apiVersion: apps/v1
kind: Deployment
metadata:
  name: frontend-deployment
  namespace: k8sdemon
spec:
  selector:
    matchLabels:
      app: frontend
  replicas: 2
  template:
    metadata:
      labels:
        app: frontend
    spec:
      containers:
      - name: frontend
        image: registry.cn-hangzhou.aliyuncs.com/pork-registry/frontend:v2.0
        ports:
        - containerPort: 80
```



#### 3.1.4、开始部署

1、在 kubernetes 集群的 master 节点上：

---

![](./imgs/node.png)

---

```shell
# 部署到 k8s 集群
kubectl apply -f backend.yaml
kubectl apply -f frontend.yaml
```

2、查看 Service 的情况：

---

![](./imgs/svc.png)

---

3、查看 Deployment 的情况：

---

![](./imgs/deploy.png)

---

4、查看 Pod 的情况：

---

![](./imgs/pod.png)

---

5、验证：

```shell
# master 节点
http://192.168.15.31:30001/user/info
# node 节点
http://192.168.15.32:30001/user/info
http://192.168.15.33:30001/user/info
```

---

![](./imgs/succeed.png)

---



## 4、Gateway API 方式访问 （需要K8s高版本）

Kubernetes 项目推荐使用 Gateway 而不是 Ingress。 Ingress API 已经被冻结。

这意味着：

- Ingress API 是正式发布的，并且遵循正式发布 API 的稳定性保证。 Kubernetes 项目没有计划从 Kubernetes 中移除 Ingress。
- Ingress API 不再进行开发，也不会对其进行进一步的更改或更新。

Gateway API 是 Ingress API 的后继者。 

```shell
# 官网地址
https://gateway-api.sigs.k8s.io/guides/getting-started/
```

对于新项目来说，直接采用 Gateway API 是最正确的选择。这让你从一开始就能站在 Kubernetes 网络演进的前沿，避免未来不必要的迁移成本。

---

<img src="./imgs/gateway.png" style="zoom: 80%;" />

----

```shell
# 项目地址
https://github.com/envoyproxy/gateway
https://github.com/alibaba/higress
```



### 4.1、Higress 启动

1、下面为常用的一些官方地址：

```shell
https://github.com/helm/helm
https://helm.sh/
https://artifacthub.io/
```

2、安装 Helm ：

> ⚠ 注意：Helm 版本需要 >= 3.10 。

```shell
# 下载 helm
wget https://get.helm.sh/helm-v3.20.0-linux-amd64.tar.gz
# 解压安装
tar -zxvf helm-v3.20.0-linux-amd64.tar.gz
mv linux-amd64/helm /usr/local/bin/helm
# 测试
helm version
```

---

![](./imgs/helm.png)

---

3、初始化（将 Bitnami 的 Helm 仓库添加到你的本地 Helm 环境中）：

```shell
helm repo add bitnami https://charts.bitnami.com/bitnami
cd /root/.cache/helm/
```

4、安装 Higress ：

```shell
helm repo add higress.io https://higress.io/helm-charts
helm install higress -n higress-system higress.io/higress --create-namespace --render-subchart-notes --set global.local=true --set global.o11y.enabled=false
```

---

<img src="./imgs/higress.png" style="zoom:80%;" />

---

5、通过更改 Higress 的 Service 的类型为 `NodePort` 实现访问Higress 控制界面：

```shell
# 备份防止改错，可省略
kubectl get service -n higress-system higress-console -o yaml > higress-console-svc.yaml
# 编辑
# 将 type: LoadBalancer 改为 type: NodePort
kubectl edit svc higress-console -n higress-system
```

---

<img src="./imgs/edit.png" style="zoom:80%;" />

---

```shell
# 访问 Higress 控制界面
http://192.168.15.31:31183/
```

---

<img src="./imgs/console.png" style="zoom: 80%;" />

---

<img src="./imgs/webui.png"  />

---



### 4.2、安装 Gateway API CRD

在开始配置之前，需要确保容器集群中已经安装了 Gateway API CRD 。

```shell
# 官方项目
https://github.com/kubernetes-sigs/gateway-api
```

1、安装标准通道：

> standard-install.yaml 文件里只定义了 Kubernetes 的 API 资源（主要是 CRD），该文件主要由多个 `CustomResourceDefinition` (CRD) 对象组成。CRD 是 Kubernetes 的一种内置资源类型，用于扩展 Kubernetes API。
>
> 命令的 `--server-side` 参数可以更好地处理 CRD 的冲突和更新，是官方推荐的做法。

```shell
kubectl apply --server-side -f https://github.com/kubernetes-sigs/gateway-api/releases/download/v1.4.1/standard-install.yaml
```

2、查看安装

```shell
# Gateway API相关的CRD都位于 gateway.networking.k8s.io 这个API组中
# 查询并显示所有属于该API组的CRD
kubectl get crd | grep gateway.networking.k8s.io
```

---

<img src="./imgs/envoy.png" style="zoom:80%;" />

---



### 4.3、安装 Envoy Gateway

#### 4.3.1、安装

安装 Envoy 会自动安装 CRD。

1、安装 Envoy Gateway：

```shell
wget https://github.com/envoyproxy/gateway/releases/download/v1.7.0/install.yaml
# Kubernetes 对注解的大小有 256KB（262144 字节）的硬性限制
# 所以使用 --server-side 参数
kubectl apply --server-side -f install.yaml
# 查看 CRD
kubectl get crd | grep gateway
```

---

<img src="./imgs/envoy1.png" style="zoom:80%;" />

---



#### 4.3.2、暴露我们的 Demon

1、要想使用 LoadBalancer 类型的 Service，可以安装 MetalLB ：

```shell
# 官网
https://metallb.io/
```

在没有 MetalLB  之前，裸机集群通常只能用`NodePort`或`externalIPs`类型的Service来从外部访问服务，但这两种方式在生产环境都有明显的缺点 。MetalLB的出现，就是为了解决这个问题，让裸机集群上的外部服务也能“即开即用” 。

MetalLB 的核心工作包含两个步骤 ：

1. **地址分配**：你需要预先为MetalLB提供一个或多个**IP地址池**（IPAddressPool）。当有`LoadBalancer`类型的Service创建时，MetalLB的控制器会从这个池中选取一个空闲的IP地址，分配给这个Service 。这个IP地址会一直绑定在这个Service上，即使后端节点故障，IP也不会改变，保证了服务的稳定性 。
2. **对外宣告**：地址分配好后，MetalLB需要让外部网络知道这个新IP“住”在集群里。它通过以下两种模式来实现这一点，你可以根据自身网络环境选择 

```shell
# 如果你的集群 kube-proxy 使用的是 ipvs 模式（1.33 版本很可能默认就是）
# 需要先开启一个开关，让 MetalLB 能够回应 ARP 请求。
kubectl edit configmap -n kube-system kube-proxy

# 找到 ipvs 段，将 strictARP 设置为 true：
apiVersion: kubeproxy.config.k8s.io/v1alpha1
kind: KubeProxyConfiguration
mode: "ipvs"
ipvs:
  strictARP: true  # 将 false 改为 true
  
# 用官方提供的 YAML 文件一键安装所有组件（控制器和 Speaker）。这会创建一个叫 metallb-system 的命名空间
kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.15.3/config/manifests/metallb-native.yaml

# 检查一下 Pod 状态
watch kubectl get pods -n metallb-system

# 创建一个 YAML 文件，例如 metallb-config.yaml：
# 你需要创建两个资源：
# IPAddressPool: 定义 MetalLB 可以使用的 IP 地址范围。
# L2Advertisement: 宣告使用 L2 模式进行广播。
apiVersion: metallb.io/v1beta1
kind: IPAddressPool
metadata:
  name: first-pool
  namespace: metallb-system
spec:
  addresses:
  # --- 请务必修改这行！---
  - 192.168.15.100-192.168.15.200
  # ---------------------
---
apiVersion: metallb.io/v1beta1
kind: L2Advertisement
metadata:
  name: l2-advert
  namespace: metallb-system
spec:
  ipAddressPools:
  - first-pool
  
# 应用该配置
kubectl apply -f metallb-config.yaml
```

2、通过 Envoy Gateway 暴露我们的服务：

> 需要暴露的是 `frontend-service` 这个服务。
>
> 确保 Gateway 和 HTTPRoute 都在 `k8sdemon` 命名空间

```shell
vim expost_prometheus_service_by_gateway.yaml
# 测试格式是否正确
kubectl apply -f expost_prometheus_service_by_gateway.yaml --dry-run=client
kubectl apply -f expost_prometheus_service_by_gateway.yaml
```

```yaml
# 定义 GatewayClass - 网关的类别和控制器
apiVersion: gateway.networking.k8s.io/v1
kind: GatewayClass
metadata:
  name: envoy-gateway-class
spec:
  # 指定使用 Envoy Gateway 的官方控制器
  controllerName: gateway.envoyproxy.io/gatewayclass-controller
---
# 定义 Gateway - 具体的负载均衡器实例
apiVersion: gateway.networking.k8s.io/v1
kind: Gateway
metadata:
  name: k8sdemon-gateway
  namespace: k8sdemon
spec:
  gatewayClassName: envoy-gateway-class
  # 定义监听列表
  listeners:
    - name: http
      protocol: HTTP
      port: 80
      # 允许的主机名（可选）
      hostname: "*.example.com"
---
# 定义 HTTPRoute - HTTP 流量的路由规则（核心配置）
# 定义 Prometheus HTTPRoute
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: frontend-route
  namespace: k8sdemon
spec:
  parentRefs:
    # 引用上面创建的 Gateway
    - name: k8sdemon-gateway
  hostnames:
    # 访问 frontend 的域名
    - "frontend.example.com"
  rules:
    - matches:
      - path:
          # 匹配所以路径
          type: PathPrefix
          value: /
      backendRefs:
        # frontend 的服务名
        - name: frontend-service
          port: 80
          weight: 1
---
# 定义 Grafana HTTPRoute
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: backend-route
  namespace: k8sdemon
spec:
  parentRefs:
    # 引用上面创建的 Gateway
    - name: k8sdemon-gateway
  hostnames:
    # 访问 backend 的域名
    - "backend.example.com"
  rules:
    - matches:
      - path:
          type: PathPrefix
          value: /
      backendRefs:
        - name: backend-service
          port: 8080
```

3、查看资源：

```shell
# 1. 检查 Gateway 是否 Ready
# 2. 检查路由是否被正确绑定
# 3. 查看 Envoy Gateway 代理 Pod 是否正常运行
kubectl get gateway -n k8sdemon -o wide
kubectl get httproute -n k8sdemon -o wide
kubectl get pods -n k8sdemon
kubectl get svc -n k8sdemon
```

---

<img src="./imgs/env1.png" style="zoom:80%;" />

---

3、浏览器访问：

> 需要更改 Windows 下的 hosts （C:\Windows\System32\drivers\etc）文件：`192.168.15.101 frontend.example.com` 。

```shell
http://backend.example.com/user/info
http://frontend.example.com/user/info
```

---

<img src="./imgs/env3.png" style="zoom:80%;" />

---

<img src="./imgs/env2.png" style="zoom:80%;" />

---





#### 4.3.3、配置说明

 Ingress-nginx 的暴露模式：

```shell
1. 部署 ingress-nginx-controller (控制平面)
   ↓
2. 手动创建/暴露一个 Service (LoadBalancer/NodePort)
   ↓
3. 创建多个 Ingress 资源 (路由规则)
   ↓
4. 所有 Ingress 共用同一个 LoadBalancer IP

```

Gateway API 的暴露模式：

```shell
1. 部署 Envoy Gateway (控制平面，无需暴露)
   ↓
2. 创建 Gateway 资源 (定义监听器)
   ↓
3. 系统自动创建 Envoy Proxy Deployment + Service (LoadBalancer)
   ↓
4. 创建多个 HTTPRoute/GRPCRoute 资源 (路由规则)
   ↓
5. 所有 Route 通过 parentRef 引用同一个 Gateway
   ↓
6. 复用同一个 LoadBalancer IP
```

---

<img src="./imgs/envoy2.png" style="zoom: 67%;" />

---

1、官方提供的快速开始：

```shell
wget https://github.com/envoyproxy/gateway/releases/download/v1.7.1/quickstart.yaml
```

```yaml
# GatewayClass - 定义网关的类别和控制器
apiVersion: gateway.networking.k8s.io/v1
kind: GatewayClass
metadata:
  name: eg                      # GatewayClass的名称，集群中唯一
spec:
  controllerName: gateway.envoyproxy.io/gatewayclass-controller  # 指定使用Envoy Gateway作为控制器
  # controllerName 告诉Kubernetes使用哪个具体的网关控制器来管理这个类的网关实例
  # 格式通常是：<domain>/<controller-name>
  # 这里指定使用 Envoy Gateway 的官方控制器

---
# Gateway - 定义具体的负载均衡器实例
apiVersion: gateway.networking.k8s.io/v1
kind: Gateway
metadata:
  name: eg                      # Gateway的名称，在命名空间内唯一
spec:
  gatewayClassName: eg          # 引用上面定义的GatewayClass，建立关联
  listeners:                     # 定义监听器列表
    - name: http                # 监听器名称，在Gateway内唯一
      protocol: HTTP            # 监听协议：HTTP, HTTPS, TCP, UDP, TLS
      port: 80                  # 监听的端口号

---
# ServiceAccount - 为Pod提供身份标识
apiVersion: v1
kind: ServiceAccount
metadata:
  name: backend                 # ServiceAccount的名称，用于给backend Pod提供身份

---
# Service - 定义网络访问端点
apiVersion: v1
kind: Service
metadata:
  name: backend                 # Service的名称
  labels:                       # 标签，用于选择器和组织资源
    app: backend                # 应用标签
    service: backend            # 服务类型标签
spec:
  ports:                        # 定义Service暴露的端口
    - name: http                # 端口名称，用于区分多个端口
      port: 3000                # Service监听的端口
      targetPort: 3000          # 转发到Pod的容器端口
  selector:                     # 标签选择器，决定哪些Pod属于这个Service
    app: backend                # 选择带有app=backend标签的Pod

---
# Deployment - 定义应用的部署配置
apiVersion: apps/v1
kind: Deployment
metadata:
  name: backend                 # Deployment的名称
spec:
  replicas: 1                   # 期望运行的Pod副本数
  selector:                     # 标签选择器，管理哪些Pod
    matchLabels:                # 匹配的标签
      app: backend              # 选择app=backend的Pod
      version: v1               # 选择version=v1的Pod
  template:                     # Pod模板
    metadata:
      labels:                    # Pod的标签
        app: backend             # 应用标签
        version: v1              # 版本标签
    spec:
      serviceAccountName: backend # 使用的ServiceAccount
      containers:                 # 容器定义
        - image: gcr.io/k8s-staging-gateway-api/echo-basic:v20231214-v1.0.0-140-gf544a46e  # 容器镜像
          imagePullPolicy: IfNotPresent  # 镜像拉取策略：本地不存在才拉取
          name: backend          # 容器名称
          ports:                  # 容器暴露的端口
            - containerPort: 3000 # 容器监听的端口
          env:                    # 环境变量
            - name: POD_NAME      # Pod名称环境变量
              valueFrom:
                fieldRef:
                  fieldPath: metadata.name  # 从Pod metadata获取name字段
            - name: NAMESPACE     # 命名空间环境变量
              valueFrom:
                fieldRef:
                  fieldPath: metadata.namespace  # 从Pod metadata获取namespace字段

---
# HTTPRoute - 定义HTTP流量的路由规则（核心配置）
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: backend                 # HTTPRoute的名称
spec:
  parentRefs:                    # 引用父资源（通常是一个Gateway）
    - name: eg                  # 引用名为eg的Gateway
      # namespace: default      # 如果Gateway在不同namespace，需要指定（默认当前namespace）
      # sectionName: http       # 可以指定Gateway的特定监听器
      # port: 80                # 指定监听的端口
  hostnames:                     # 匹配的域名列表
    - "www.example.com"         # 只有访问此域名的请求才会被此路由处理
  rules:                         # 路由规则列表
    - backendRefs:               # 后端服务引用（流量转发目标）
        - group: ""              # API组，空字符串表示core/v1
          kind: Service          # 资源类型
          name: backend          # 后端Service名称
          port: 3000             # 后端Service的端口
          weight: 1              # 权重（用于负载均衡，金丝雀发布时调整）
      matches:                    # 匹配条件（过滤哪些请求应用此规则）
        - path:                   # 路径匹配
            type: PathPrefix      # 匹配类型：PathPrefix（路径前缀）、Exact（精确匹配）、RegularExpression（正则）
            value: /              # 匹配的值，这里匹配所有以/开头的路径
          # method: GET           # 可选：匹配HTTP方法
          # headers:              # 可选：匹配请求头
          #   - name: "X-Version"
          #     value: "v1"
          # queryParams:          # 可选：匹配查询参数
          #   - name: "test"
          #     value: "true"
```



参考：[Kubernetes Gateway API 与 Envoy Gateway 部署使用指南-CSDN博客](https://blog.csdn.net/ygqygq2/article/details/156461494?ops_request_misc=elastic_search_misc&request_id=2e6911c9792e7b5366d2a79c4842cd54&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~all~sobaiduend~default-1-156461494-null-null.142^v102^pc_search_result_base4&utm_term=envoy Gateway&spm=1018.2226.3001.4187)
