package main

import (
	_ "gf-playground-1/internal/packed"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/os/gctx"

	"gf-playground-1/internal/cmd"
)

func main() {
	err := g.I18n().SetPath("manifest/i18n")
	if err != nil {
		panic(err)
	}
	g.I18n().SetLanguage("zh-CN")

	cmd.Main.Run(gctx.GetInitCtx())
}
