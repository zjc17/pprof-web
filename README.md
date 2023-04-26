# Pprof Web Visualizer
[![CodeQL](https://github.com/zjc17/pprof-web/actions/workflows/github-code-scanning/codeql/badge.svg)](https://github.com/zjc17/pprof-web/actions/workflows/github-code-scanning/codeql)
[![Security Scan](https://github.com/zjc17/pprof-web/actions/workflows/scan.yml/badge.svg)](https://github.com/zjc17/pprof-web/actions/workflows/scan.yml)
[![Release](https://github.com/zjc17/pprof-web/actions/workflows/release.yml/badge.svg)](https://github.com/zjc17/pprof-web/actions/workflows/release.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/zjc17/pprof-web)](https://goreportcard.com/report/github.com/zjc17/pprof-web)
[![Docker Image](https://img.shields.io/docker/pulls/lovecho/pprof-web.svg)](https://hub.docker.com/r/lovecho/pprof-web)

Golang Pprof Web Visualizer is a web application that visualizes the output of the pprof tool in a web browser.

The Size of the binary is < 10MB.

Support on WebUI, x86, ARM, Linux and MacOS.

![preview 01](.github/preview-01.png)

![preview 02](.github/preview-02.png)

## Download

Download the binaries for your system and architecture from the [release page](https://github.com/zjc17/pprof-web/releases).

If you prefer docker, you can use the following command (DockerHub):

```bash
docker pull lovecho/pprof-web:latest
docker pull lovecho/pprof-web:latest
```

## Usage

Use default parameters to format all configuration files in the current directory:

```bash
./pprof-web
```

This will start a web server on port 8080.

Then you can access it at [localhost:8080](http://localhost:8080).

### Docker Usage

There is no difference between using parameters in Docker and the above, for example, we start a Web UI formatting tool service in Docker:

```bash
docker run --rm -it -p 8080:8080 lovecho/pprof-web:latest
```

## Credit

Web Components:

- Gin is a HTTP web framework written in Go (Golang), under [MIT license].
  - https://github.com/gin-gonic/gin
- Crayons - A UI Kit comprising of web components for building Freshworks Apps! - [License not specified yet]
  - https://github.com/freshworks/crayons