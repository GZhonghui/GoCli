package main

import (
	gomod "github.com/GZhonghui/GoMod"
	gomodv2 "github.com/GZhonghui/GoMod/v2"
	"local.dev/GoCli/tools"
)

func main() {
	gomod.PrintAboutMessage()
	gomodv2.PrintAboutMessage()
	tools.Hi()
}
