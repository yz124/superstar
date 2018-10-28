package main

import (
	"github.com/yz124/superstar/bootstrap"
	"github.com/yz124/superstar/web/middleware/identity"
	"github.com/yz124/superstar/web/routes"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("Superstar database", "一凡Sir")
	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen(":8080")
}
