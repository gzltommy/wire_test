原文地址：https://zhuanlan.zhihu.com/p/338926709

### (1) Installing
```
$ go install github.com/google/wire/cmd/wire@latest
```

### (2)调用 wire 命令生成依赖文件：
```
$ wire
wire: github.com/DrmagicE/wire-examples/quickstart: wrote XXXX\github.com\DrmagicE\wire-examples\quickstart\wire_gen.go

$ ls
main.go  wire.go  wire_gen.go
```