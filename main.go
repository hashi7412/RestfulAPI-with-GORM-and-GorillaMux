package main

import (
	app "github.com/hashi7412/restfulapi-with-gorm-and-gorillamux/api"
	config "github.com/hashi7412/restfulapi-with-gorm-and-gorillamux/api"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
