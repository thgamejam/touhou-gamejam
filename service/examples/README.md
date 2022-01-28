# SpringBoot 项目样例

## 前提条件
1. 确保安装 `Docker`
2. 运行 `./service/testing-environment.ps1` 脚本，安装依赖容器


## 创建项目
- `名称`应与服务实际内容相关
- `位置`应在 `service` 目录下
- `组`统一使用 `cc.thjam`

![图1-1创建项目](https://github.com/thgamejam/touhou-gamejam/raw/main/service/examples/img/图1-1创建项目.png)

先择基础依赖如下

![图1-1创建项目](https://github.com/thgamejam/touhou-gamejam/raw/main/service/examples/img/图1-2创建项目.png)


## 环境与配置

### 使用 Maven 切换开发环境
分别使用三种环境, 即: 
- 两种测试环境: 本地测试环境 `dev` 与 线上测试环境 `test`
- 生成环境 `prod`

向 `pom.xml` 中的节点 `project` 内添加以下内容，并重新加载 `Mavne` 项目

```xml
<profiles>
    <profile>
        <!-- 本地开发环境 -->
        <id>dev</id>
        <properties>
            <profiles.active>dev</profiles.active>
            <spring.active>dev</spring.active>
        </properties>
        <!-- 默认是dev环境 -->
        <activation>
            <activeByDefault>true</activeByDefault>
        </activation>
    </profile>
    <profile>
        <!-- 测试环境 -->
        <id>test</id>
        <properties>
            <profiles.active>test</profiles.active>
            <spring.active>test</spring.active>
        </properties>
    </profile>
    <profile>
        <!-- 生产环境 -->
        <id>prod</id>
        <properties>
            <profiles.active>prod</profiles.active>
            <spring.active>prod</spring.active>
        </properties>
    </profile>
</profiles>
```

勾选 `dev` 本地测试环境

![图2-1切换开发环境](https://github.com/thgamejam/touhou-gamejam/raw/main/service/examples/img/图2-1切换开发环境.png)


### 项目配置

配置文件定义:
- `application.yml` 基础配置信息。如 `Consul`、`Actuator` 等配置
- `application-dev.yml` 本地调试配置信息。如 `MariaDB`、`Redis` 等配置
- `application-test.yml` 线上调试配置信息
- `application-prod.yml` 生产环境部署配置信息

注: `application-dev.yml`、`application-test.yml`、`application-prod.yml` 将覆盖 `application.yml` 与 `Consul` 配置中心的配置内容。

创建 `application.yml` 配置文件, 设置 `UTF-8` 编码, 并复制如下内容: 

```yaml
spring:
  profiles:
    active: @spring.active@   # 控制 Maven 环境切换
  application:
    name: examples
  config:
    # 必要！
    # https://docs.spring.io/spring-cloud-consul/docs/current/reference/html/#config-data-import
    import: optional:consul:localhost:8500
  cloud:
    consul:
      host: localhost
      port: 8500
      config:
        enabled: true
        # 以下内容组合为 Consul KV 路径 "prefixes/name/data-key"
        # 即: "service/examples/dev"
        prefixes:
          - service
        name: ${spring.application.name}
        data-key: ${spring.profiles.active}
        format: YAML # Value 类型
      discovery:
        enabled: true
        register: true    # 注册自身到 Consul
        deregister: true  # 服务停止时取消注册
        service-name: ${spring.application.name}  # 注册到 Consul 的服务名
        health-check-path: /actuator/health   # 健康检查的接口 由 Spring Boot Actuator 提供
        heartbeat:
          enabled: true   # 开启 Consul 心跳机制
```

同样的, 创建 `application-dev.yml` 配置文件, 设置 `UTF-8` 编码, 并复制如下内容:

```yaml
server:
  port: 3000  # 修改 Spring Boot Web 端口
spring:
  datasource:
    driver-class-name: org.mariadb.jdbc.Driver
    username: root
    password: "123456"
    url: jdbc:mysql://localhost:3306/touhou_gamejam?serverTimezone=GMT%2B8&useUnicode=true&characterEncoding=utf-8&useSSL=false
  jpa:
    hibernate:
      ddl-auto: update
```


### (非必要) 测试环境下的 Consul 配置分发

- 确保 `Consul Docker` 正确运行
- 确保 `application.yml` [正确配置](#项目配置) `Consul`
- 注意 `application-dev.yml` 将覆盖 `Consul` 配置中心的配置内容, 必要时请注释 `application-dev.yml` 中的内容

打开 `Consul Web UI` 默认情况下, 浏览器访问 <http://localhost:8500/ui> , 点击左侧导航栏, 进入 `Key/Value`, 点击 `Create` 进行配置文件的创建

- `Key or folder` 中填写 `service/examples/dev`
- `Value` 选择 `YAML`

填入以下配置内容:

```yaml
server:
  port: 5000
```

启动服务, 命令行中显示如下内容, 即成功下发配置并从 `5000` 端口启动 `Spring Boot Web` 容器

```sh
Tomcat initialized with port(s): 5000 (http)
```

