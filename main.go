package main

import (
	// 无法使用 main 包
	//testmain "github.com/oceanweave/testgomod"
	// 即使引用时，也要指定 go 包好路径
	"github.com/oceanweave/testgomod/pkg/demo1"
	"github.com/oceanweave/testgomod/tools"
)

func main() {
	// 此处可以正常调用
	demo1.Hello1()
	// 报错 import "github.com/oceanweave/testgomod" is a program, not an importable package
	// 无法引用 main 包的函数
	//testmain.SayHello()

	// 打印出 this sayHello from testgomod Repo‘s tools go package
	// 符合预期
	tools.SayHello()
}
