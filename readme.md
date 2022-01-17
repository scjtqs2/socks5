# 说明

### cmd/daemon 的入口 是用于daemon方式运行的

```shell
go build ./cmd/daemon/main.go
# win使用方法，安装为服务
./socks5.exe install
#linux/mac 安装为服务
./socks5 install
#卸载方法
./socks5 uninstall
```

### 正常使用

```shell
go build
#linux/mac
./socks5
#win
./socks5.exe
```

### 配置文件就不做独立导入了

在main.go里面修改`UserName`和`Password`
