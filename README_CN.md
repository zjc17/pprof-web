# Pprof Web Visualizer
[![CodeQL](https://github.com/zjc17/pprof-web/actions/workflows/github-code-scanning/codeql/badge.svg)](https://github.com/zjc17/pprof-web/actions/workflows/github-code-scanning/codeql)
[![Security Scan](https://github.com/zjc17/pprof-web/actions/workflows/scan.yml/badge.svg)](https://github.com/zjc17/pprof-web/actions/workflows/scan.yml)
[![Release](https://github.com/zjc17/pprof-web/actions/workflows/release.yml/badge.svg)](https://github.com/zjc17/pprof-web/actions/workflows/release.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/zjc17/pprof-web)](https://goreportcard.com/report/github.com/zjc17/pprof-web)
[![Docker Image](https://img.shields.io/docker/pulls/lovecho/pprof-web.svg)](https://hub.docker.com/r/lovecho/pprof-web)

<p style="text-align: center;">
  <a href="README.md" target="_blank">ENGLISH</a> | <a href="README_CN.md">中文文档</a>
</p>

Golang Pprof Web Visualizer 是一个网络应用程序，可以在网络浏览器中对 pprof 工具的输出进行可视化。

二进制文件小于10MB。

支持 WebUI, x86, ARM, Linux 和 MacOS。

![预览01](.github/preview-01.png)

![预览02](.github/preview-02.png)

## 下载

从[发布页](https://github.com/zjc17/pprof-web/releases)下载适合你的系统和架构的二进制文件。

如果你喜欢docker，你可以使用以下命令（DockerHub）：

``bash
docker pull lovecho/pprof-web:latest
docker pull lovecho/pprof-web:latest
```

## 用法

使用默认参数来格式化当前目录下的所有配置文件：

``bash
./pprof-web
```

这将在8080端口启动一个Web服务器。

然后你可以在[localhost:8080](http://localhost:8080)访问它。

### Docker用法

在Docker中使用参数和上面的方法没有区别，例如，我们在Docker中启动一个Web UI格式化工具服务：

```bash
docker run --rm -it -p 8080:8080 lovecho/pprof-web:latest
```

## 现场演示

你可以在 [pprof.gotool.tech](https://pprof.gotool.tech/) 访问实时演示。

### 最佳实践

在远程机器内通过curl上传pprof的结果，然后从你的本地机器访问 web UI。

```bash
# 用curl上传
curl -F "file=@$FILE_PATH" https://pprof.gotool.tech/submit -v
```

然后你可以检查输出，发现如下内容

```bash
< HTTP/2 307
< date: Thu, 27 Apr 2023 08:14:54 GMT
< location: /pprof/?file_id=XXXXXXXX
< vary: Accept-Encoding
< cf-cache-status: DYNAMIC
```

然后你可以访问 `https://pprof.gotool.tech/pprof/?file_id=XXXXXXXX` 来查看你的pprof结果。


## 信用

网络组件：

- Gin是一个用Go（Golang）编写的HTTP网络框架，在[MIT许可]下。
  - https://github.com/gin-gonic/gin
- Crayons - 一个由Web组件组成的UI工具包，用于构建Freshworks的应用程序! - [还没有指定许可证] 。
  - https://github.com/freshworks/crayons
