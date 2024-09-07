# tiktok-live-assistant
### 注意：本项目仅支持Chrome浏览器，且仅用于学习使用
### 需要自行在项目目录下创建 .env文件，该文件内容如下：
```text
TIKTOK_URL=https://www.douyin.com/
TIKTOK_LIVE_URL=https://live.douyin.com/
ROOM_ID=换成直播间ID
ROOM_XPATH=换成评论区的评论item（在浏览器使用“检查”复制元素的class）
```

## 使用说明

### 1.环境准备
先安装并配置Go环境，版本为1.20.6  

保姆级教程链接：https://blog.csdn.net/ohmygodes/article/details/122646716

使用 `win` + `R` 组合键打开CMD终端

在CMD终端使用 `go version` 命令验证是否安装成功



### 2.拉取项目
将项目拉取到本地，在CMD终端cd到项目的目录中

例如：

`cd/d D:\tiktok-live-assistant`

### 3.运行项目
先执行初始化依赖项命令： `go mod init`

然后再使用启动程序命令： `go run ./main.go`

最后按照提示完成步骤
### 项目结构

```text
tiktok-live-assistant/
├── main.go    # 主程序入口
├── configs/   # 存放配置 
├── models/    # 封装数据结构 
├── public/    # 公共资源 
│   └── cookies/     
├── services/    # 服务 由main直接调用，或service之间互相调用
│   └── browsers-service.go
├── handlers/    # 操作 由services层调用，或handler之间互相调用。负责更具体的读、写和判断等操作 
│   └── ctrl-cookies.go/  
└── go.mod    // 依赖
```

### 基础功能项

1.启动浏览器，跳转登录页面（已完成）

2.登录抖音并保存cookies（已完成）

3.完成登录跳转直播间（已完成）

4.自动获取直播间信息（已完成）

5.算法判断问句（已完成）

6.自动发送信息（未完成）

### 扩展功能项

1.优化缓存配置，减少空间占用（已完成）

2.使用协程提高效率（已完成）

3.数据实时同步至多维表（未完成）

4.问答数据（未完成）

### 项目后续

1.使用GORM框架操作SQLite（已完成）

2.使用HTTP框架调用大模型（暂定gin框架）

3.实现更多的自动操作，减少手动操作

4.使用GUI框架，构建友好的交互界面

5.打包为可执行程序



