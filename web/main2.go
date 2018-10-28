package main

import (
	"github.com/yz124/superstar/bootstrap"
	"github.com/yz124/superstar/web/middleware/identity"
	"github.com/yz124/superstar/web/routes"
)

func main() {
	app := bootstrap.New("Superstar database", "一凡Sir")
	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure)
	app.Listen(":8081")
}
