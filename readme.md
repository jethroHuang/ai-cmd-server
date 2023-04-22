# ai-cmd-server

本项目为 ai-cmd 的服务端，用于接收客户端的请求并返回 ai 生成的命令行

### ⚡ 语言切换

- [中文](./readme.md)
- [English](./readme_en.md)

## 编译并运行
```bash
$ go build
$ ./ai-cmd-server <your openai apikey>
```

## docker 编译并运行
```bash
$ docker build -t ai-cmd-server .
$ docker run -d -p 1323:1323 --name ai-cmd-container ai-cmd-server <your openai apikey>
```

## 使用
```bash
$ curl "http://localhost:1323/help?content=查看用户目录所有后缀为log的文件名称&sys=windows"
$ {"status":0,"data":"dir %userprofile%\\*.log","msg":"ok"}
```

## 请求与响应
### 请求
此 api 服务只有一个接口，接口为 `/help`

请求方式为 `GET`，请求参数为 `content` 和 `sys`

`content` 为用户想要执行的命令，例如：**查看用户目录所有后缀为log的文件**

`sys` 为用户的操作系统，目前支持 `windows`,`linux`,`macos`

### 响应
响应为 json 格式，包含三个字段
status: 状态码，0 为成功，其他为失败
data: 响应数据，成功时为 ai 生成的命令行，失败时为空字符串
msg: 响应信息，成功时为 ok，失败时为错误信息