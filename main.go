package main

import (
	app "restapi/api"
	config "restapi/api"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
