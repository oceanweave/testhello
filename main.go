package main

import (
	"fmt"
	// 无法使用 main 包
	//testmain "github.com/oceanweave/testgomod"
	// 即使引用时，也要指定 go 包好路径
	// v1 版本
	//"github.com/oceanweave/testgomod/pkg/demo1"
	toolsv1 "github.com/oceanweave/testgomod/tools"
	// v2 版本
	"github.com/oceanweave/testgomod/v2/pkg/demo1"
	"github.com/oceanweave/testgomod/v2/tools"
	// v1 与 v2 比较
	// 拉取命令，后面的版本号忽略不计，可以看出 v2 在 testgomod 后面多了个 /v2, 这也是因为原项目 go.mod 文件中 module 名称变化引起的
	// 这样避免了 v1 版本误拉取 v2.0.0，避免了不兼容引起的出错，如 go get -u github.com/oceanweave/testgomod@v2.0.0 拉取不到项目
	// 同时引用时，使用 v2 时，import 要记得在路径中添加 v2
	// go get -u github.com/oceanweave/testgomod@v1.0.1
	// go get -u github.com/oceanweave/testgomod/v2@v2.0.0
)

// v2 版本调用成功 同时 v1 版本也可以使用
//has big change in this demo1 package, totallly different from v1
//this is Hello-1 from demo1
//has big change in this tools package, totallly different from v1
//this sayHello from testgomod Repo‘s tools go package
//--------- below is v1 tools go package ----
//this sayHello from testgomod Repo‘s tools go package

func main() {
	// 此处可以正常调用
	demo1.Hello1()
	// 报错 import "github.com/oceanweave/testgomod" is a program, not an importable package
	// 无法引用 main 包的函数
	//testmain.SayHello()

	// 打印出 this sayHello from testgomod Repo‘s tools go package
	// 符合预期
	// v2 版本重构后删除了此函数，更改为 SayHi
	// tools.SayHello()
	tools.SayHi()

	fmt.Println("--------- below is v1 tools go package ----")
	toolsv1.SayHello()

}
