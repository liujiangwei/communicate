package web

import (
	"communicate/wx/controller"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func Start() {
	app := iris.New()
	app.Get("/", func(context context.Context) {
		_, _ = context.HTML("test")
	})

	wx := app.Party("/wx/")
	wx.Get("/", controller.Login)
	wx.Get("/test", func(context context.Context) {
		_,_ = context.HTML("/wx/text")
	})

	_ = app.Run(iris.Addr(":8080"))
}