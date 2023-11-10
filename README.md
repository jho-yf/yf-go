# yf-go
YF-GO

## Go程序的运行与构建

```sh
# 运行hello.go程序
go run hello.go

# 将hello.go程序构建c
go build hello.go
```
## go mod初始化项目

开启go mod

```shell
go env -w GO111MODULE=on
# 或
export GO111MODULE=on
```

初始化项目

```shell
go mod init github.com/jho-yf/go12-gomod
# 或
go mod init go12-gomod
```

手动安装依赖，例如：

```she
go get github.com/aceld/zinx/ziface
go get github.com/aceld/zinx/znet
```

也可以运行`go run`，会自动安装依赖。

`go.mod`文件，`indirect`表示间接依赖

```go
module github.com/jho-yf/go12-gomod

go 1.17

require (
	github.com/aceld/zinx v1.2.1 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/klauspost/cpuid/v2 v2.1.1 // indirect
	github.com/klauspost/reedsolomon v1.11.8 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/templexxx/cpufeat v0.0.0-20180724012125-cef66df7f161 // indirect
	github.com/templexxx/xor v0.0.0-20191217153810-f85b25db303b // indirect
	github.com/tjfoc/gmsm v1.4.1 // indirect
	github.com/xtaci/kcp-go v5.4.20+incompatible // indirect
	golang.org/x/crypto v0.11.0 // indirect
	golang.org/x/net v0.12.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
)
```

`go.sum`文件，罗列当前项目直接或间接依赖所有模块版本，保证依赖版本不被篡改

> h1:hash 表示整体项目的zip文件打开之后的全部文件的校验和生成的hash
>
> xxx/go.mod h1:hash go.mod文件做的hash

```go
cloud.google.com/go v0.26.0/go.mod h1:aQUYkXzVsufM+DwF1aE+0xfcU+56JwCaLick0ClmMTw=
github.com/BurntSushi/toml v0.3.1/go.mod h1:xHWCNGjB5oqiDr8zfno3MHue2Ht5sIBksp03qcyfWMU=
github.com/aceld/zinx v1.2.1 h1:qLAEkIpG4EpSDuI0auWR9GPVDygRB2UZ4N3L5uKvr+s=
github.com/aceld/zinx v1.2.1/go.mod h1:OHTsFyN1ph7LGOvhNf4ekD9Zkcze3CyT69neU/7EXQk=
github.com/census-instrumentation/opencensus-proto v0.2.1/go.mod h1:f6KPmirojxKA12rnyqOA5BBL4O983OfeGPqjHWSTneU=
github.com/client9/misspell v0.3.4/go.mod h1:qj6jICC3Q7zFZvVWo7KLAzC3yx5G7kyvSDkc90ppPyw=
github.com/cncf/udpa/go v0.0.0-20191209042840-269d4d468f6f/go.mod h1:M8M6+tZqaGXZJjfX53e64911xZQV5JYwmTeXPW+k8Sc=
github.com/davecgh/go-spew v1.1.0/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
...
```

