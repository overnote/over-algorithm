# 本源码的使用

必须开启go mod，笔者建议直接安装go1.13以上版本即可，会默认开启 go mod。  

源码使用步骤：
```
cd sources/go
go run main.go          // 在 main.go 中打开想要的测试方法即可
```

# Go语言的一些规范

Go语言中不推荐使用 _ 命名标识符，比如定义常量，推荐如下：
```go
const ListMaxSize int = 100
```