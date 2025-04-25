// 在此目录下执行：
// go run .

// 入口函数必须在这个包里面
// 目录名：GoCli
// 包名：main
package main

// import的对象是package
// import path都是从module path开始写
import (
	// 通过使用别名 导入了同一个 module 的两个不同的版本
	// Package名称冲突的时候 就要使用别名
	gomod "github.com/GZhonghui/GoMod" // 为 package 定义别名
	gomodv2 "github.com/GZhonghui/GoMod/v2"
	"local.dev/GoCli/tools" // 这个是本地的 package，写法也是一样的规则
)

// 上面有从Module Path导入 也有从 Package Path导入
// 但是本质上 都是导入 Package
// 也就是说 import 不会递归（不会顺便导入子文件夹下的内容）

func main() {
	gomod.PrintAboutMessage()
	gomodv2.PrintAboutMessage()
	tools.Hi()
}
