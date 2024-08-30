# tiktok-live-assistant
### 注意：本项目仅支持Chrome浏览器
## 使用说明

### 1.环境准备
先安装并配置Go环境，版本为1.20.6  
压缩包： https://golang.google.cn/dl/go1.20.6.windows-amd64.zip

保姆式教程链接：https://blog.csdn.net/ohmygodes/article/details/122646716

### 2.拉取项目
将项目拉取到本地，在CMD终端cd到项目的目录中

可以先使用 `go version` 验证是否安装成功

### 3.运行项目
初次使用前先安装依赖： `go mod init`

使用命令启动程序： `go run ./main.go`

### 项目结构

```text
tiktok-live-assistant/
├── main.go    # 主程序入口
├── public/    # 公共资源 
│   └── cookies/     
├── services/    # 交互操作
│   └── browsers.go
└── go.mod    // 依赖
```

### 已完成

1.启动浏览器，跳转登录页面

2.登录抖音并保存cookies

3.完成登录跳转直播间

### 待完成

1.自动获取直播间信息

2.自动发送信息

3.三线程提高效率

4.算法判断问句

5.数据实时同步至多维表

6.问答数据

### 项目后续

1.使用协程提高并发能力

2.使用GORM框架操作数据库

3.使用Docker容器化技术




