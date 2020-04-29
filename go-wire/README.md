# GO依赖注入wire

* 手动编写启动
```
cmd/main.go 
```

* wire启动
```
1.安装wire
go get github.com/google/wire/cmd/wire

2.生成依赖注入代
码wire ./internal/di/

3.使用wire启动代码
cmd/main_wire.go
```
