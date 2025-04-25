// 如果没有想要代码复用 这里 Module Path 可以随便写
module local.dev/GoCli

// 语言版本
go 1.24.2

// 导入的第三方代码 一般是 Github 地址
// github url / v2 (if version >= 2) / version (tag)
// 安装（下载）第三方的代码使用 go get
// go get module_path/(v2)@v2.1.0

// 可以同时导入同一个 module 的多个不同的版本
// 条件是它们的「大版本」都不同
// 因为「大版本」不同的话，module path就不同（后面的/v<x>不同）
require (
	github.com/GZhonghui/GoMod v1.0.1
	github.com/GZhonghui/GoMod/v2 v2.1.1
)
