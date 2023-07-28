package main

import (
	"github.com/hashi7412/restapi/api/app"
	"github.com/hashi7412/restapi/api/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
