package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	//app.Use(logger.New())

	htmlEngine := iris.HTML("./", ".html")
	//htmlEngine.Reload(false)
	app.RegisterView(htmlEngine)

	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString("Hello world! -- from iris.")
	})

	app.Get("/hello", func(ctx iris.Context) {
		ctx.ViewData("Title", "测试页面")
		ctx.ViewData("Content", "Hello world! -- from template")
		ctx.View("hello.html")
	})

	app.Run(iris.Addr(":8080"), iris.WithCharset("UTF-8"))
}
